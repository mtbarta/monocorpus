package endpoints

import (
	"context"
	"time"

	"golang.org/x/time/rate"

	stdopentracing "github.com/opentracing/opentracing-go"
	"github.com/sony/gobreaker"

	"github.com/go-kit/kit/circuitbreaker"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/metrics"
	"github.com/go-kit/kit/ratelimit"
	"github.com/go-kit/kit/tracing/opentracing"

	pb "github.com/mtbarta/monocorpus/pkg/notes"
	"github.com/mtbarta/monocorpus/pkg/notes/service"
	"github.com/mtbarta/monocorpus/pkg/notes/types"
)

type Set struct {
	CreateNoteEndpoint endpoint.Endpoint
	DeleteNoteEndpoint endpoint.Endpoint
	UpdateNoteEndpoint endpoint.Endpoint
	GetNotesEndpoint   endpoint.Endpoint
}

var RATE = 100

// New returns a Set that wraps the provided server, and wires in all of the
// expected endpoint middlewares via the various parameters.
func New(svc service.Service, logger log.Logger, duration metrics.Histogram, trace stdopentracing.Tracer) Set {
	var CreateNoteEndpoint endpoint.Endpoint
	{
		CreateNoteEndpoint = MakeCreateNoteEndpoint(svc)
		CreateNoteEndpoint = ratelimit.NewErroringLimiter(rate.NewLimiter(rate.Every(time.Second), RATE))(CreateNoteEndpoint)
		CreateNoteEndpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{}))(CreateNoteEndpoint)
		CreateNoteEndpoint = opentracing.TraceServer(trace, "CreateNote")(CreateNoteEndpoint)
		CreateNoteEndpoint = LoggingMiddleware(log.With(logger, "method", "CreateNote"))(CreateNoteEndpoint)
		CreateNoteEndpoint = InstrumentingMiddleware(duration.With("method", "CreateNote"))(CreateNoteEndpoint)
	}

	var DeleteNoteEndpoint endpoint.Endpoint
	{
		DeleteNoteEndpoint = MakeDeleteNoteEndpoint(svc)
		DeleteNoteEndpoint = ratelimit.NewErroringLimiter(rate.NewLimiter(rate.Every(time.Second), RATE))(DeleteNoteEndpoint)
		DeleteNoteEndpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{}))(DeleteNoteEndpoint)
		DeleteNoteEndpoint = opentracing.TraceServer(trace, "DeleteNote")(DeleteNoteEndpoint)
		DeleteNoteEndpoint = LoggingMiddleware(log.With(logger, "method", "DeleteNote"))(DeleteNoteEndpoint)
		DeleteNoteEndpoint = InstrumentingMiddleware(duration.With("method", "DeleteNote"))(DeleteNoteEndpoint)
	}

	var UpdateNoteEndpoint endpoint.Endpoint
	{
		UpdateNoteEndpoint = MakeUpdateNoteEndpoint(svc)
		UpdateNoteEndpoint = ratelimit.NewErroringLimiter(rate.NewLimiter(rate.Every(time.Second), RATE))(UpdateNoteEndpoint)
		UpdateNoteEndpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{}))(UpdateNoteEndpoint)
		UpdateNoteEndpoint = opentracing.TraceServer(trace, "UpdateNote")(UpdateNoteEndpoint)
		UpdateNoteEndpoint = LoggingMiddleware(log.With(logger, "method", "UpdateNote"))(UpdateNoteEndpoint)
		UpdateNoteEndpoint = InstrumentingMiddleware(duration.With("method", "UpdateNote"))(UpdateNoteEndpoint)
	}

	var GetNotesEndpoint endpoint.Endpoint
	{
		GetNotesEndpoint = MakeGetNotesEndpoint(svc)
		GetNotesEndpoint = ratelimit.NewErroringLimiter(rate.NewLimiter(rate.Every(time.Second), RATE))(GetNotesEndpoint)
		GetNotesEndpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{}))(GetNotesEndpoint)
		GetNotesEndpoint = opentracing.TraceServer(trace, "GetNotes")(GetNotesEndpoint)
		GetNotesEndpoint = LoggingMiddleware(log.With(logger, "method", "GetNotes"))(GetNotesEndpoint)
		GetNotesEndpoint = InstrumentingMiddleware(duration.With("method", "GetNotes"))(GetNotesEndpoint)
	}
	return Set{
		CreateNoteEndpoint: CreateNoteEndpoint,
		UpdateNoteEndpoint: UpdateNoteEndpoint,
		DeleteNoteEndpoint: DeleteNoteEndpoint,
		GetNotesEndpoint:   GetNotesEndpoint,
	}
}

// MakeCreateNoteEndpoint constructs a Sum endpoint wrapping the service.
func MakeCreateNoteEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		note := request.(*pb.Note)
		n := types.PbNoteToMongo(note)
		noteResponse, err := s.CreateNote(&ctx, &n)

		proto := noteResponse.ToProto()
		return &proto, err
	}
}

// MakeSumEndpoint constructs a Sum endpoint wrapping the service.
func MakeDeleteNoteEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		note := request.(*pb.Note)
		n := types.PbNoteToMongo(note)
		noteResponse, err := s.DeleteNote(&ctx, &n)

		proto := noteResponse.ToProto()
		return &proto, err
	}
}

// MakeSumEndpoint constructs a Sum endpoint wrapping the service.
func MakeUpdateNoteEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		note := request.(*pb.Note)
		n := types.PbNoteToMongo(note)
		noteResponse, err := s.UpdateNote(&ctx, &n)
		if err != nil {
			return nil, err
		}

		proto := noteResponse.ToProto()
		return &proto, err
	}
}

// MakeSumEndpoint constructs a Sum endpoint wrapping the service.
func MakeGetNotesEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		query := request.(*pb.Query)
		q := types.PbQueryToMongo(query)

		noteList, err := s.GetNotes(&ctx, &q)

		var pbNotes []*pb.Note
		for _, note := range noteList {
			n := note.ToProto()
			pbNotes = append(pbNotes, &n)
		}

		return pbNotes, err
	}
}
