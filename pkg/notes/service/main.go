package service

import (
	"context"

	ep "github.com/mtbarta/monocorpus/pkg/notes/types"

	"github.com/go-kit/kit/log"

	mgo "gopkg.in/mgo.v2"
)

type Service interface {
	UpdateNote(ctx *context.Context, note *ep.Note) (*ep.Note, error)
	CreateNote(ctx *context.Context, note *ep.Note) (*ep.Note, error)
	DeleteNote(ctx *context.Context, note *ep.Note) (*ep.Note, error)
	GetNotes(ctx *context.Context, queryRequest *ep.Query) ([]ep.Note, error)
}

type basicService struct {
	noteCollection *mgo.Collection
	logger         *log.Logger
}

func NewBasicService(logger *log.Logger, db *mgo.Collection) Service {
	return basicService{
		noteCollection: db,
		logger:         logger,
	}
}

// New returns a basic Service with all of the expected middlewares wired in.
func New(logger log.Logger, db *mgo.Collection) Service {
	var svc Service
	{
		svc = NewBasicService(&logger, db)
		// svc = LoggingMiddleware(logger)(svc)
		// svc = InstrumentingMiddleware(ints, chars)(svc)
	}
	return svc
}
