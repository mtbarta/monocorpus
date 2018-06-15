package types

import (
	"time"

	pb "github.com/mtbarta/monocorpus/pkg/notes"

	"github.com/golang/protobuf/ptypes"
)

type Tag struct {
	Text  string `json:"text" bson:"text"`
	Color string `json:"color" bson:"color"`
}

type Note struct {
	ID           string    `json:"id" bson:"_id"`
	Title        string    `json:"title" bson:"title"`
	Author       string    `json:"author" bson:"author"`
	Team         string    `json:"team" bson:"team"`
	Body         string    `json:"body" bson:"body"`
	Type         string    `json:"type" bson:"type"`
	DateCreated  time.Time `json:"dateCreated" bson:"dateCreated"`
	DateModified time.Time `json:"dateModified" bson:"dateModified"`
	Link         string    `json:"link" bson:"link"`
	Image        []byte    `json:"image" bson:"image"`
	Tags         []Tag     `json:"tags" bson:"tags"`
}

type ElasticsearchNote struct {
	ID           string `json:"id" bson:"_id"`
	Title        string `json:"title" bson:"title"`
	Author       string `json:"author" bson:"author"`
	Team         string `json:"team" bson:"team"`
	Body         string `json:"body" bson:"body"`
	Type         string `json:"type" bson:"type"`
	DateCreated  int64  `json:"dateCreated" bson:"dateCreated"`
	DateModified int64  `json:"dateModified" bson:"dateModified"`
	Link         string `json:"link" bson:"link"`
	Tags         []Tag  `json:"tags" bson:"tags"`
}

func (note *Note) ToProto() pb.Note {
	created, _ := ptypes.TimestampProto(note.DateCreated)
	mod, _ := ptypes.TimestampProto(note.DateModified)

	pbTags := make([]*pb.Tag, len(note.Tags))
	for _, tag := range note.Tags {
		pbTags = append(pbTags, &pb.Tag{
			Text:  tag.Text,
			Color: tag.Color,
		})
	}
	return pb.Note{
		Id:           note.ID,
		Title:        note.Title,
		Author:       note.Author,
		Team:         note.Team,
		Body:         note.Body,
		Type:         note.Type,
		DateCreated:  created,
		DateModified: mod,
		Link:         note.Link,
		Image:        note.Image,
		Tags:         pbTags,
	}
}

func (note *Note) Clone() Note {
	return Note{
		ID:           note.ID,
		Title:        note.Title,
		Author:       note.Author,
		Team:         note.Team,
		Body:         note.Body,
		Type:         note.Type,
		DateCreated:  note.DateCreated,
		DateModified: note.DateModified,
		Link:         note.Link,
		Image:        note.Image,
		Tags:         note.Tags,
	}
}

func PbNoteToMongo(note *pb.Note) Note {
	tags := make([]Tag, len(note.Tags))
	for _, tag := range note.Tags {
		tags = append(tags, Tag{
			Text:  tag.Text,
			Color: tag.Color,
		})
	}

	return Note{
		ID:           note.Id,
		Title:        note.Title,
		Author:       note.Author,
		Team:         note.Team,
		Body:         note.Body,
		Type:         note.Type,
		Link:         note.Link,
		DateCreated:  time.Unix(note.DateCreated.GetSeconds(), 0),
		DateModified: time.Unix(note.DateModified.GetSeconds(), 0),
		Image:        note.Image,
		Tags:         tags,
	}
}

func (note *Note) ToElasticSearch() *ElasticsearchNote {
	return &ElasticsearchNote{
		ID:           note.ID,
		Title:        note.Title,
		Author:       note.Author,
		Team:         note.Team,
		Body:         note.Body,
		Type:         note.Type,
		DateCreated:  note.DateCreated.Unix(),
		DateModified: note.DateModified.Unix(),
		Link:         note.Link,
		Tags:         note.Tags,
	}
}

func (note *ElasticsearchNote) ToNote() Note {
	return Note{
		ID:           note.ID,
		Title:        note.Title,
		Author:       note.Author,
		Team:         note.Team,
		Body:         note.Body,
		Type:         note.Type,
		DateCreated:  time.Unix(note.DateCreated, 0),
		DateModified: time.Unix(note.DateModified, 0),
		Link:         note.Link,
		Tags:         note.Tags,
	}
}
