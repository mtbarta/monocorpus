package search

import (
	"context"
	"reflect"

	"github.com/mtbarta/monocorpus/pkg/notes/util"
	"github.com/mtbarta/monocorpus/pkg/notes/types"

	"github.com/microcosm-cc/bluemonday"
	"github.com/olivere/elastic"
	blackfriday "gopkg.in/russross/blackfriday.v2"
)

var klogger = util.CreateNewLogFmtLogger()
var jsonLogger = util.CreateNewJSONLogger()

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

func (s *NoteSearcher) Ping() (*elastic.PingResult, error) {
	ctx := context.Background()
	info, _, err := s.Client.Ping(s.url).Do(ctx)

	return info, err
}

func (s *NoteSearcher) ElasticSearchVersion() (string, error) {
	esversion, err := s.Client.ElasticsearchVersion(s.url)

	return esversion, err
}

func (s *NoteSearcher) EnsureIndex() (bool, error) {
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
			klogger.Log("error", err.Error)
			panic(err)
		}
		if !createIndex.Acknowledged {
			// Not acknowledged
		}
	}

	return exists, err
}

func (s *NoteSearcher) Put(note *types.Note, ctx context.Context) error {
	cleanInput := blackfriday.Run([]byte(note.Body))
	cleanInput = bluemonday.UGCPolicy().SanitizeBytes(cleanInput)

	cleanTitle := bluemonday.UGCPolicy().Sanitize(note.Title)
	klogger.Log("content", cleanInput)
	indexableNote := types.Note{
		Author:      note.Author,
		Title:       cleanTitle,
		Body:        string(cleanInput),
		DateCreated: note.DateCreated,
		Link:        note.Link,
	}
	_, err := s.Client.Index().Index(s.index).
		Type(s.docType).
		Id(note.ID).
		BodyJson(indexableNote).
		Do(ctx)

	if err != nil {
		klogger.Log("error", err.Error)
		return err
	}
	return nil
}

func (s *NoteSearcher) Search(query string, ctx context.Context) ([]types.Note, error) {
	q := elastic.NewMultiMatchQuery(query, "title", "body")
	searchResult, err := s.Client.Search().
		Index(s.index).
		Query(q).
		Do(ctx)

	if err != nil {
		klogger.Log("error", err.Error)
		return nil, err
	}

	// return searchResult, err
	var ttyp types.ElasticsearchNote
	var results []types.Note
	for _, item := range searchResult.Each(reflect.TypeOf(ttyp)) {
		jsonLogger.Log(item)
		note := item.(types.ElasticsearchNote)
		results = append(results, note.ToNote())
	}
	jsonLogger.Log(results)
	return results, nil
}

func (s *NoteSearcher) Update(note *types.Note, ctx context.Context) error {
	internalNote := note.ToElasticSearch()

	_, err := s.Client.Update().
		Id(note.ID).
		Index(s.index).
		Type(s.docType).
		DocAsUpsert(true).
		Doc(internalNote).
		Do(ctx)

	if err != nil {
		klogger.Log("error", err)
		return err
	}

	return nil
}

func (s *NoteSearcher) Delete(id string, ctx context.Context) error {
	_, err := s.Client.Delete().
		Id(id).
		Do(ctx)

	if err != nil {
		klogger.Log("error", err.Error)
		return err
	}
	return nil
}
