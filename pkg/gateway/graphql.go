package gateway

import (
	"context"
	"strconv"

	npb "github.com/mtbarta/monocorpus/pkg/notes"
	search "github.com/mtbarta/monocorpus/pkg/notes/search"
	"github.com/mtbarta/monocorpus/pkg/notes/types"

	"github.com/go-kit/kit/log"
	"github.com/golang/protobuf/ptypes/timestamp"

	graphql "github.com/graph-gophers/graphql-go"
)

//the graphql schema
var Schema = `
	schema {
		query: Query
		mutation: Mutation
	}
	type Query {
		notes(ids: [ID], title: String, authors: [String], dateCreated: Float, team: String, todate: Float, fromdate: Float): [Note]
		search(query: String, authors: [String], dateCreated: Float, team: String, todate: Float, fromdate: Float): [Note]
	}
	type Mutation {
		createNote(id: ID, title:String, body:String, author:String, team:String, dateCreated:String, dateModified:String, type:String, link:String, image:String, tags:[String]): Note
		updateNote(id: ID, title:String, body:String, author:String, team:String, dateCreated:String, dateModified:String, type:String, link:String, image:String, tags:[String]): Note
		deleteNote(id: ID, title:String, body:String, author:String, team:String, dateCreated:String, dateModified:String, type:String, link:String): Note
	}
	type Note {
		id: ID
		title: String
		body: String
		author: String
		team: String
		dateCreated: Float
		dateModified: Float
		type: String
		link: String,
		image: String,
		tags: [String]
	}
`

type Resolver struct {
	notesConn    npb.NotesClient
	searchClient *search.NoteSearcher
	logger       *log.Logger
}

func NewResolver(notesConn npb.NotesClient, searchClient *search.NoteSearcher, logger *log.Logger) *Resolver {
	return &Resolver{notesConn: notesConn, searchClient: searchClient, logger: logger}
}

func (r *Resolver) Search(ctx context.Context, args *struct {
	Query       *string
	Authors     *[]*string
	Team        *string
	DateCreated *float64
	Todate      *float64
	Fromdate    *float64
}) *[]*NoteResolver {
	if args.Query == nil {
		return nil
	}

	var authors []string
	if args.Authors != nil {
		authors = make([]string, len(*args.Authors))
		for _, author := range *args.Authors {
			authors = append(authors, (*author))
		}
	} else {
		authors = nil
	}

	q := pointerToString(args.Query)
	var notes []types.Note
	if r.searchClient != nil {
		var noteResolvers []*NoteResolver

		notes, _ = r.searchClient.Search(q, ctx)

		for _, n := range notes {
			pbNote := n.ToProto()
			noteResolvers = append(noteResolvers, &NoteResolver{Note: &pbNote})
		}

		return &noteResolvers
	}

	var noResult []*NoteResolver
	return &noResult
}

// Notes queries the Notes store to retrieve... notes
func (r *Resolver) Notes(ctx context.Context, args *struct {
	IDs         *[]*string
	Title       *string
	Authors     *[]*string
	Team        *string
	DateCreated *float64
	Todate      *float64
	Fromdate    *float64
	Image       *string
	Tags        *[]*string
}) *[]*NoteResolver {
	var notes *npb.NoteList

	var ids []string
	if args.IDs != nil {
		ids = make([]string, len(*args.IDs))
		for _, id := range *args.IDs {
			ids = append(ids, (*id))
		}
	} else {
		ids = nil
	}

	var authors []string
	if args.Authors != nil {
		authors = make([]string, len(*args.Authors))
		for _, author := range *args.Authors {
			if author != nil {
				authors = append(authors, (*author))
			}

		}
	} else {
		authors = append(authors, ctx.Value("email").(string))
	}

	title := pointerToString(args.Title)

	from, to := ParseFromAndTo(args.Fromdate, args.Todate)

	query := npb.Query{
		IDs:      ids, // *args.IDs,
		Title:    title,
		Authors:  authors, //*args.authors,
		Todate:   to,
		Fromdate: from,
	}
	notes, err := r.notesConn.GetNotes(ctx, &query)

	if err != nil {
		logger.Log("err", err.Error())
		return nil
	}

	var noteResolvers []*NoteResolver
	for _, n := range notes.Notes {
		noteResolvers = append(noteResolvers, &NoteResolver{Note: n})
	}

	return &noteResolvers
}

func ParseFromAndTo(Fromdate *float64, Todate *float64) (*timestamp.Timestamp, *timestamp.Timestamp) {
	var toTime *timestamp.Timestamp
	var fromTime *timestamp.Timestamp
	if Todate != nil {
		to := int64(*Todate)
		toTime = &timestamp.Timestamp{Seconds: to, Nanos: 0}
	}

	if Fromdate != nil {
		from := int64(*Fromdate)
		fromTime = &timestamp.Timestamp{Seconds: from, Nanos: 0}
	}

	return fromTime, toTime
}

func pointerToString(point *string) string {
	if point == nil {
		return ""
	}
	return *point
}

func pointerToNil(point *string) string {
	if point == nil {
		return ""
	}
	return *point
}

func toTimestamp(el string) (*timestamp.Timestamp, error) {
	if el != "" {
		s, err := strconv.ParseInt(el, 10, 64)
		if err != nil {
			return nil, err
		}
		return &timestamp.Timestamp{Seconds: s}, nil
	}
	return &timestamp.Timestamp{Seconds: 0}, nil
}

type NoteResolver struct {
	Note *npb.Note
}

func (r *NoteResolver) ID() *graphql.ID {
	res := graphql.ID(r.Note.Id)
	return &res
}

func (r *NoteResolver) Title() *string {
	return &r.Note.Title
}

func (r *NoteResolver) Body() *string {
	return &r.Note.Body
}

func (r *NoteResolver) Author() *string {
	return &r.Note.Author
}

func (r *NoteResolver) Team() *string {
	return &r.Note.Team
}

func (r *NoteResolver) DateCreated() *float64 {
	v := r.Note.DateCreated.GetSeconds()
	u := float64(v)
	return &u
}

func (r *NoteResolver) DateModified() *float64 {
	v := r.Note.DateModified.GetSeconds()
	u := float64(v)
	return &u
}

func (r *NoteResolver) Type() *string {
	return &r.Note.Type
}

func (r *NoteResolver) Link() *string {
	return &r.Note.Link
}

func (r *NoteResolver) Image() *string {
	var result = string(r.Note.Image[:])
	return &result
}

func (r *NoteResolver) Tags() *[]*string {
	var result []*string

	return &result
}

func (r *Resolver) CreateNote(ctx context.Context, args *struct {
	ID           *string
	Title        *string
	Body         *string
	Team         *string
	Author       *string
	DateCreated  *float64
	DateModified *float64
	Type         *string
	Link         *string
	Image        *string
	Tags         *[]*string
}) (*NoteResolver, error) {

	title := pointerToNil(args.Title)

	id := pointerToNil(args.ID)
	body := pointerToNil(args.Body)
	author := pointerToNil(args.Author)
	if author == "" {
		author = ctx.Value("email").(string)
	}
	noteType := pointerToNil(args.Type)
	link := pointerToNil(args.Link)
	var created, modified int64
	if args.DateCreated == nil {
		created = 0
	} else {
		created = int64(*args.DateCreated)
	}
	if args.DateModified == nil {
		modified = 0
	} else {
		modified = int64(*args.DateModified)
	}

	createdTime := timestamp.Timestamp{Seconds: created, Nanos: 0}
	modifiedTime := timestamp.Timestamp{Seconds: modified, Nanos: 0}

	note, err := r.notesConn.CreateNote(ctx, &npb.Note{
		Id:           id,
		Title:        title,
		Body:         body,
		Author:       author,
		DateCreated:  &createdTime,
		DateModified: &modifiedTime,
		Type:         noteType,
		Link:         link,
		Image:        []byte(*args.Image),
	})

	if err != nil {
		return &NoteResolver{Note: note}, err
	}
	return &NoteResolver{Note: note}, nil
}

// func (r *Resolver) createNote(args)

func (r *Resolver) DeleteNote(ctx context.Context, args *struct {
	ID           *string
	Title        *string
	Body         *string
	Team         *string
	Author       *string
	DateCreated  *float64
	DateModified *float64
	Type         *string
	Link         *string
}) (*NoteResolver, error) {

	title := pointerToNil(args.Title)

	id := pointerToNil(args.ID)
	body := pointerToNil(args.Body)
	author := pointerToNil(args.Author)
	noteType := pointerToNil(args.Type)
	// link := pointerToNil(args.Link)
	var created, modified int64
	if args.DateCreated == nil {
		created = 0
	} else {
		created = int64(*args.DateCreated)
	}
	if args.DateModified == nil {
		modified = 0
	} else {
		modified = int64(*args.DateModified)
	}

	createdTime := timestamp.Timestamp{Seconds: created, Nanos: 0}
	modifiedTime := timestamp.Timestamp{Seconds: modified, Nanos: 0}

	if r.searchClient != nil {
		r.searchClient.Delete(id, ctx)
	}

	note, err := r.notesConn.DeleteNote(ctx, &npb.Note{
		Id:           id,
		Title:        title,
		Body:         body,
		Author:       author,
		DateCreated:  &createdTime,
		DateModified: &modifiedTime,
		Type:         noteType,
	})

	if err != nil {
		return &NoteResolver{Note: note}, err
	}
	return &NoteResolver{Note: note}, nil
}

func (r *Resolver) UpdateNote(ctx context.Context, args *struct {
	ID           *string
	Title        *string
	Body         *string
	Team         *string
	Author       *string
	DateCreated  *float64
	DateModified *float64
	Type         *string
	Link         *string
	Image        *string
	Tags         *[]*string
}) (*NoteResolver, error) {

	title := pointerToNil(args.Title)
	id := pointerToNil(args.ID)
	body := pointerToNil(args.Body)
	author := pointerToNil(args.Author)
	noteType := pointerToNil(args.Type)
	link := pointerToNil(args.Link)
	var created, modified int64
	if args.DateCreated == nil {
		created = 0
	} else {
		created = int64(*args.DateCreated)
	}
	if args.DateModified == nil {
		modified = 0
	} else {
		modified = int64(*args.DateModified)
	}

	createdTime := timestamp.Timestamp{Seconds: created, Nanos: 0}
	modifiedTime := timestamp.Timestamp{Seconds: modified, Nanos: 0}

	var img []byte
	if args.Image != nil {
		img = []byte(*args.Image)
	} else {
		img = nil
	}

	pbNote := npb.Note{
		Id:           id,
		Title:        title,
		Body:         body,
		Author:       author,
		DateCreated:  &createdTime,
		DateModified: &modifiedTime,
		Type:         noteType,
		Link:         link,
		Image:        img,
	}

	note, err := r.notesConn.UpdateNote(ctx, &pbNote)

	if r.searchClient != nil {
		xformedNote := types.PbNoteToMongo(note)
		r.searchClient.Update(&xformedNote, ctx)
	}

	if err != nil {
		return &NoteResolver{Note: note}, err
	}
	return &NoteResolver{Note: note}, nil
}
