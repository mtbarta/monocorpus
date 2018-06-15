package service

import (
	"context"
	"os"

	ep "github.com/mtbarta/monocorpus/pkg/notes/types"

	"github.com/go-kit/kit/log"
	"gopkg.in/mgo.v2/bson"
)

var logger = log.NewJSONLogger(os.Stdout)

func (s basicService) CreateNote(ctx *context.Context, note *ep.Note) (*ep.Note, error) {
	_id := bson.NewObjectId().Hex()
	if note.ID == "" {
		note.ID = _id
	}
	err := s.noteCollection.Insert(note)

	if err != nil {
		(*s.logger).Log("stmnt", "failed to insert note", "err", err)
		return nil, NoteError{"failed to insert note"}
	}

	return note, nil
}

// func formatNoteForCreation(note *pb.Note) *pb.Note {
// 	msPtr := reflect.ValueOf(note)
// 	msValue := msPtr.Elem()

// 	for i := 0; i < msValue.NumField(); i++ {
// 		field := msValue.Field(i)
// 		if field.Elem() == ""
// 	}
// }

func (s basicService) DeleteNote(ctx *context.Context, note *ep.Note) (*ep.Note, error) {
	err := s.noteCollection.Remove(&bson.M{
		"_id": note.ID,
	})
	if err != nil {
		(*s.logger).Log("stmnt", "failed to delete note", "error", err)
		return nil, NoteError{"failed to delete note"}
	}
	return note, nil
}

func (s basicService) UpdateNote(ctx *context.Context, note *ep.Note) (*ep.Note, error) {
	tempNote := note.Clone()

	err := s.noteCollection.Update(&bson.M{
		"_id": note.ID,
	}, note)

	if err != nil {
		(*s.logger).Log("stmnt", "failed to update note", "error", err, "id", note.ID)
		return nil, NoteError{"failed to update note"}
	}
	return &tempNote, nil
}

func (s basicService) GetNotes(ctx *context.Context, queryRequest *ep.Query) ([]ep.Note, error) {
	var result []ep.Note

	fmtRequest := formatQueryToMongo(queryRequest)

	if fmtRequest == nil {
		return nil, NoteError{"query was malformed"}
	}

	err := s.noteCollection.Find(fmtRequest).Sort("-dateCreated").All(&result)

	if err != nil {
		(*s.logger).Log("stmnt", "failed to find notes", "err", err)
		return result, NoteError{"failed to find notes"}
	}

	return result, nil
}

func formatQueryToMongo(queryRequest *ep.Query) bson.M {
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
		// logger.Log(queryRequest)
		// // fmt.Println(queryRequest.Authors[1])
		// fmt.Println(len(queryRequest.Authors))
		// for _, s := range queryRequest.Authors {
		// 	fmt.Println(s)
		// }

		// fmtRequest["author"] = bson.M{"$in": queryRequest.Authors}
		if len(queryRequest.Authors) != 2 {
			return nil
		}
		fmtRequest["author"] = string(queryRequest.Authors[1])
	}

	dateConstraints := bson.M{}
	if !queryRequest.Fromdate.IsZero() {
		dateConstraints["$gte"] = queryRequest.Fromdate
	}
	if !queryRequest.Todate.IsZero() {
		dateConstraints["$lte"] = queryRequest.Todate
	}

	if len(dateConstraints) >= 1 {
		fmtRequest["dateCreated"] = dateConstraints
	}

	return fmtRequest
}
