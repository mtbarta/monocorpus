package search

import (
	"context"
	"encoding/json"
	"reflect"
	"time"

	"github.com/golang/protobuf/ptypes"
	"github.com/mtbarta/monocorpus/pkg/logging"
	notes "github.com/mtbarta/monocorpus/pkg/notes"
	"github.com/mtbarta/monocorpus/pkg/notes/service"

	"github.com/microcosm-cc/bluemonday"
	"github.com/olivere/elastic"
	blackfriday "gopkg.in/russross/blackfriday.v2"
)

const mapping = `
{
	"settings":{
		"number_of_shards": 1,
		"number_of_replicas": 0,
		"analysis": {
      "analyzer": {
        "html_analyzer": {
          "type":      "custom",
          "tokenizer": "standard",
          "char_filter": [
            "html_strip"
          ],
          "filter": [
            "lowercase",
            "asciifolding"
          ]
        }
      }
    }
	},
	"mappings":{
		"note":{
			"properties":{
				"id": {
					"type": "text"
				},
				"author":{
					"type":"keyword"
				},
				"body":{
					"type":"text",
					"analyzer": "html_analyzer"
				},
				"title":{
					"type":"text"
				},
				"dateCreated":{
					"type":"date"
				},
				"link":{
					"type":"keyword"
				},
				"type":{
					"type":"keyword"
				}
			}
		}
	}
}`

type NoteSearcher struct {
	Client  *elastic.Client
	url     string
	index   string
	docType string
}

func (s *NoteSearcher) Search(ctx context.Context, noteQuery *SearchQuery, resp *notes.NoteList) error {
	q := elastic.NewMultiMatchQuery(noteQuery.GetQuery(), "title", "body")
	searchResult, err := s.Client.Search().
		Index(s.index).
		Query(q).
		Do(ctx)

	if err != nil {
		return err
	}

	logging.Logger.Debugf("search results", "amount", searchResult.TotalHits())

	if searchResult.TotalHits() == 0 {
		return nil
	}

	var ttyp service.Note

	for _, item := range each(searchResult, reflect.TypeOf(ttyp)) {
		n := item.result.(service.Note)
		created, err := ptypes.TimestampProto(n.DateCreated)
		if err != nil {
			created = nil
		}
		modified, err := ptypes.TimestampProto(n.DateModified)
		if err != nil {
			modified = nil
		}

		pbTags := make([]*notes.Tag, len(n.Tags))
		for _, tag := range n.Tags {
			pbTags = append(pbTags, &notes.Tag{
				Text:  tag.Text,
				Color: tag.Color,
			})
		}

		note := notes.Note{
			Id:           n.ID,
			Title:        n.Title,
			Author:       n.Author,
			Team:         n.Team,
			Body:         n.Body,
			Type:         n.Type,
			DateCreated:  created,
			DateModified: modified,
			Link:         n.Link,
			Image:        n.Image,
			Tags:         pbTags,
			Score:        float32(item.score),
		}
		resp.Notes = append(resp.Notes, &note)
	}
	return nil
}

type hitResult struct {
	result interface{}
	score  float64
}

// Each is a utility function to iterate over all hits. It saves you from
// checking for nil values. Notice that Each will ignore errors in
// serializing JSON and hits with empty/nil _source will get an empty
// value
func each(r *elastic.SearchResult, typ reflect.Type) []hitResult {
	if r.Hits == nil || r.Hits.Hits == nil || len(r.Hits.Hits) == 0 {
		return nil
	}
	var slice []hitResult
	for _, hit := range r.Hits.Hits {
		v := reflect.New(typ).Elem()
		if hit.Source == nil {
			slice = append(slice, hitResult{
				result: v.Interface(),
				score:  *hit.Score,
			})
			continue
		}
		if err := json.Unmarshal(*hit.Source, v.Addr().Interface()); err == nil {
			slice = append(slice, hitResult{
				result: v.Interface(),
				score:  *hit.Score,
			})
		}
	}
	return slice
}

type NoteManager struct {
	Client  *elastic.Client
	url     string
	index   string
	docType string
}

func NewNoteSearcher(host string, port string, index string, docType string) (*NoteSearcher, error) {
	url := "http://" + host + ":" + port
	Client, err := elastic.NewClient(elastic.SetURL(url))

	if err != nil {
		return nil, err
	}
	return &NoteSearcher{
		Client:  Client,
		url:     host + ":" + port,
		index:   index,
		docType: docType,
	}, nil
}

func NewNoteManager(host string, port string, index string, docType string) (*NoteManager, error) {
	url := "http://" + host + ":" + port
	Client, err := elastic.NewClient(elastic.SetURL(url))

	if err != nil {
		return nil, err
	}
	return &NoteManager{
		Client:  Client,
		url:     host + ":" + port,
		index:   index,
		docType: docType,
	}, nil
}

func (s *NoteManager) Ping() (*elastic.PingResult, error) {
	ctx := context.Background()
	info, _, err := s.Client.Ping(s.url).Do(ctx)

	return info, err
}

func (s *NoteManager) ElasticSearchVersion() (string, error) {
	esversion, err := s.Client.ElasticsearchVersion(s.url)

	return esversion, err
}

func (s *NoteManager) EnsureIndex() (bool, error) {
	ctx := context.Background()
	exists, err := s.Client.IndexExists(s.index).Do(ctx)
	if err != nil {
		// Handle error
		panic(err)
	}
	if !exists {
		// Create a new index.
		createIndex, err := s.Client.CreateIndex(s.index).BodyString(mapping).Do(ctx)
		if err != nil {
			// Handle error
			panic(err)
		}
		if !createIndex.Acknowledged {
			// Not acknowledged
		}
	}

	return exists, err
}

func (s *NoteManager) Put(ctx context.Context, note *notes.Note) error {
	logging.Logger.Infof("putting note in search")
	cleanInput := blackfriday.Run([]byte(note.Body))
	cleanInput = bluemonday.UGCPolicy().SanitizeBytes(cleanInput)

	cleanTitle := bluemonday.UGCPolicy().Sanitize(note.Title)
	created, err := ptypes.Timestamp(note.DateCreated)
	if err != nil {
		created = time.Now()
	}
	indexableNote := service.Note{
		Author:      note.Author,
		Title:       cleanTitle,
		Body:        string(cleanInput),
		DateCreated: created,
		Link:        note.Link,
	}
	_, err = s.Client.Index().Index(s.index).
		Type(s.docType).
		Id(note.Id).
		BodyJson(indexableNote).
		Timestamp(created.String()).
		Do(ctx)

	if err != nil {
		logging.Logger.Error("failed to put note in elasticsearch")
		return err
	}
	logging.Logger.Infof("successful")
	return nil
}

func (s *NoteManager) Update(ctx context.Context, note *notes.Note) error {
	cleanInput := blackfriday.Run([]byte(note.Body))
	cleanInput = bluemonday.UGCPolicy().SanitizeBytes(cleanInput)

	cleanTitle := bluemonday.UGCPolicy().Sanitize(note.Title)
	created, err := ptypes.Timestamp(note.DateCreated)
	if err != nil {
		created = time.Now()
	}
	indexableNote := service.Note{
		Author:      note.Author,
		Title:       cleanTitle,
		Body:        string(cleanInput),
		DateCreated: created,
		Link:        note.Link,
	}
	_, err = s.Client.Update().Index(s.index).
		Type(s.docType).
		Id(note.Id).
		Doc(indexableNote).
		DocAsUpsert(true).
		Do(ctx)

	if err != nil {
		return err
	}
	return nil
}

func (s *NoteManager) Delete(ctx context.Context, id string) error {
	_, err := s.Client.Delete().
		Id(id).
		Do(ctx)

	if err != nil {
		return err
	}
	return nil
}
