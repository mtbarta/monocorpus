package service

import (
	"context"

	"github.com/mtbarta/monocorpus/pkg/notes"

	"gopkg.in/mgo.v2/bson"
)

func (s NoteService) CreateNote(ctx context.Context, note *notes.Note, resp *notes.Note) error {
	_id := bson.NewObjectId().Hex()
	if note.Id == "" {
		note.Id = _id
	}
	err := s.Collection.Insert(note)

	if err != nil {
		return NoteError{"failed to insert note"}
	}

	resp = note
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
	err := s.Collection.Update(&bson.M{
		"_id": note.Id,
	}, note)

	if err != nil {
		return NoteError{"failed to update note"}
	}
	resp = note
	return nil
}

func (s NoteService) GetNotes(ctx context.Context, queryRequest *notes.Query, resp *notes.NoteList) error {
	var result notes.NoteList

	fmtRequest := formatQueryToMongo(queryRequest)

	if fmtRequest == nil {
		return NoteError{"query was malformed"}
	}

	err := s.Collection.Find(fmtRequest).Sort("-dateCreated").All(&result)

	if err != nil {
		return NoteError{"failed to find notes"}
	}

	resp = &result
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
		dateConstraints["$gte"] = queryRequest.Fromdate.GetSeconds()
	}
	if !(queryRequest.Todate.GetSeconds() == 0) {
		dateConstraints["$lte"] = queryRequest.Todate.GetSeconds()
	}

	if len(dateConstraints) >= 1 {
		fmtRequest["dateCreated"] = dateConstraints
	}

	return fmtRequest
}
