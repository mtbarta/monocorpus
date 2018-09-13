package service

import (
	"context"
	"time"

	"github.com/golang/protobuf/ptypes"
	"github.com/mtbarta/monocorpus/pkg/logging"
	"github.com/mtbarta/monocorpus/pkg/notes"

	"gopkg.in/mgo.v2/bson"
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

func ToMongo(note *notes.Note) Note {
	tags := make([]Tag, len(note.Tags))
	for _, tag := range note.Tags {
		tags = append(tags, Tag{
			Text:  tag.Text,
			Color: tag.Color,
		})
	}

	created, err := ptypes.Timestamp(note.DateCreated)
	if err != nil {
		created = time.Now()
	}
	modified, err := ptypes.Timestamp(note.DateModified)
	if err != nil {
		created = time.Now()
	}

	return Note{
		ID:           note.Id,
		Title:        note.Title,
		Author:       note.Author,
		Team:         note.Team,
		Body:         note.Body,
		Type:         note.Type,
		Link:         note.Link,
		DateCreated:  created,
		DateModified: modified,
		Image:        note.Image,
		Tags:         tags,
	}
}

func ToProto(note Note, resp *notes.Note) {
	created, _ := ptypes.TimestampProto(note.DateCreated)
	mod, _ := ptypes.TimestampProto(note.DateModified)

	pbTags := make([]*notes.Tag, len(note.Tags))
	for _, tag := range note.Tags {
		pbTags = append(pbTags, &notes.Tag{
			Text:  tag.Text,
			Color: tag.Color,
		})
	}

	resp.Id = note.ID
	resp.Title = note.Title
	resp.Author = note.Author
	resp.Team = note.Team
	resp.Body = note.Body
	resp.Type = note.Type
	resp.DateCreated = created
	resp.DateModified = mod
	resp.Link = note.Link
	resp.Image = note.Image
	resp.Tags = pbTags

}

func (s NoteService) CreateNote(ctx context.Context, note *notes.Note, resp *notes.Note) error {
	_id := bson.NewObjectId().Hex()
	if note.Id == "" {
		note.Id = _id
	}

	mongoNote := ToMongo(note)

	err := s.Collection.Insert(mongoNote)

	if err != nil {
		return NoteError{"failed to insert note"}
	}

	ToProto(mongoNote, resp)
	return nil
}

func (s NoteService) DeleteNote(ctx context.Context, note *notes.Note, resp *notes.Note) error {
	err := s.Collection.Remove(&bson.M{
		"_id": note.Id,
	})
	if err != nil {
		return NoteError{"failed to delete note"}
	}

	resp = note
	return nil
}

func (s NoteService) UpdateNote(ctx context.Context, note *notes.Note, resp *notes.Note) error {
	mongoNote := ToMongo(note)

	err := s.Collection.Update(&bson.M{
		"_id": note.Id,
	}, mongoNote)

	if err != nil {
		return NoteError{"failed to update note"}
	}
	ToProto(mongoNote, resp)
	return nil
}

func (s NoteService) GetNotes(ctx context.Context, queryRequest *notes.Query, resp *notes.NoteList) error {
	var result []Note

	fmtRequest := formatQueryToMongo(queryRequest)

	if fmtRequest == nil {
		return NoteError{"query was malformed"}
	}
	logging.Logger.Infof("query", fmtRequest)
	err := s.Collection.Find(fmtRequest).Sort("-dateCreated").All(&result)
	logging.Logger.Info(result)
	logging.Logger.Info(err)
	if err != nil {
		return NoteError{"failed to find notes"}
	}

	for _, note := range result {
		var n notes.Note
		ToProto(note, &n)
		resp.Notes = append(resp.Notes, &n)
	}

	return nil
}

func formatQueryToMongo(queryRequest *notes.Query) bson.M {
	// var fmtRequest bson.M
	fmtRequest := bson.M{}

	if queryRequest.IDs != nil {
		fmtRequest["_id"] = bson.M{"$in": queryRequest.IDs}
	}
	if queryRequest.Title != "" {
		// define query so that user input can't be malicious.
		titles := make([]string, 1)
		titles = append(titles, queryRequest.Title)
		fmtRequest["title"] = bson.M{"$in": titles}
	}
	if queryRequest.Team != "" {
		teams := make([]string, 1)
		teams = append(teams, queryRequest.Team)
		fmtRequest["team"] = bson.M{"$in": teams}
	}
	if queryRequest.Authors != nil {
		//this code smells
		if len(queryRequest.Authors) != 2 {
			return nil
		}
		fmtRequest["author"] = string(queryRequest.Authors[1])
	}

	dateConstraints := bson.M{}
	if !(queryRequest.Fromdate.GetSeconds() == 0) {
		from, err := ptypes.Timestamp(queryRequest.GetFromdate())
		if err == nil {
			dateConstraints["$gte"] = from
		}

	}
	if !(queryRequest.Todate.GetSeconds() == 0) {
		to, err := ptypes.Timestamp(queryRequest.GetTodate())
		if err == nil {
			dateConstraints["$lte"] = to
		}
	}

	if len(dateConstraints) >= 1 {
		fmtRequest["dateCreated"] = dateConstraints
	}

	return fmtRequest
}
