package search

import (
	"context"
	"reflect"

	notes "github.com/mtbarta/monocorpus/pkg/notes"

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
					"type":"date",
					"format": "epoch_second"
				},
				"link":{
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

	var ttyp notes.Note
	var results notes.NoteList
	for _, item := range searchResult.Each(reflect.TypeOf(ttyp)) {
		note := item.(notes.Note)
		results.Notes = append(results.Notes, &note)
	}

	resp = &results
	return nil
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
	cleanInput := blackfriday.Run([]byte(note.Body))
	cleanInput = bluemonday.UGCPolicy().SanitizeBytes(cleanInput)

	cleanTitle := bluemonday.UGCPolicy().Sanitize(note.Title)
	indexableNote := notes.Note{
		Author:      note.Author,
		Title:       cleanTitle,
		Body:        string(cleanInput),
		DateCreated: note.DateCreated,
		Link:        note.Link,
	}
	_, err := s.Client.Index().Index(s.index).
		Type(s.docType).
		Id(note.Id).
		BodyJson(indexableNote).
		Do(ctx)

	if err != nil {
		return err
	}
	return nil
}

func (s *NoteManager) Update(ctx context.Context, note *notes.Note) error {
	_, err := s.Client.Update().
		Id(note.Id).
		Index(s.index).
		Type(s.docType).
		DocAsUpsert(true).
		Doc(note).
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
