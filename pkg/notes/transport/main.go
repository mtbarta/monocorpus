package transport

import (
	"context"
	"errors"

	pb "github.com/mtbarta/monocorpus/pkg/notes"
	"github.com/mtbarta/monocorpus/pkg/notes/endpoints"

	"github.com/go-kit/kit/auth/jwt"
	"github.com/go-kit/kit/log"
	grpctransport "github.com/go-kit/kit/transport/grpc"
)

type grpcServer struct {
	createNote grpctransport.Handler
	deleteNote grpctransport.Handler
	updateNote grpctransport.Handler
	getNotes   grpctransport.Handler
}

//MakeGRPCServer to serve notes.
func MakeGRPCServer(endpoints endpoints.Set, logger log.Logger) pb.NotesServer {
	options := []grpctransport.ServerOption{grpctransport.ServerErrorLogger(logger)}

	return &grpcServer{
		createNote: grpctransport.NewServer(
			endpoints.CreateNoteEndpoint,
			DecodeGRPCNoteRequest,
			EncodeGRPCNoteResponse,
			append(options, grpctransport.ServerBefore(jwt.GRPCToContext()))...,
		),
		deleteNote: grpctransport.NewServer(
			endpoints.DeleteNoteEndpoint,
			DecodeGRPCNoteRequest,
			EncodeGRPCNoteResponse,
			append(options, grpctransport.ServerBefore(jwt.GRPCToContext()))...,
		),
		updateNote: grpctransport.NewServer(
			endpoints.UpdateNoteEndpoint,
			DecodeGRPCNoteRequest,
			EncodeGRPCNoteResponse,
			append(options, grpctransport.ServerBefore(jwt.GRPCToContext()))...,
		),
		getNotes: grpctransport.NewServer(
			endpoints.GetNotesEndpoint,
			DecodeGRPCQueryRequest,
			EncodeGRPCGetNotesResponse,
			append(options, grpctransport.ServerBefore(jwt.GRPCToContext()))...,
		),
	}
}

func (s *grpcServer) CreateNote(ctx context.Context, req *pb.Note) (*pb.Note, error) {
	_, note, err := s.createNote.ServeGRPC(ctx, req)
	return note.(*pb.Note), err
}

func (s *grpcServer) DeleteNote(ctx context.Context, req *pb.Note) (*pb.Note, error) {
	_, note, err := s.deleteNote.ServeGRPC(ctx, req)

	return note.(*pb.Note), err
}

func (s *grpcServer) UpdateNote(ctx context.Context, req *pb.Note) (*pb.Note, error) {
	_, note, err := s.updateNote.ServeGRPC(ctx, req)
	if note != nil {
		return note.(*pb.Note), nil
	}
	return nil, err
}

func (s *grpcServer) GetNotes(ctx context.Context, req *pb.Query) (*pb.NoteList, error) {
	_, response, err := s.getNotes.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}

	// noteList := pb.NoteList{Notes: response.([]*pb.Note)}

	return response.(*pb.NoteList), err
}

// encodeGRPCSumResponse is a transport/grpc.EncodeResponseFunc that converts a
// user-domain sum response to a gRPC sum reply. Primarily useful in a server.
func EncodeGRPCGetNotesResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.([]*pb.Note)

	noteList := pb.NoteList{Notes: resp}

	return &noteList, nil

	// var notes pb.NoteList
	// for _, note := range resp.Notes {
	// 	notes.Notes = append(notes.Notes, &pb.Note{
	// 		ID:           note.ID,
	// 		Author:       note.Author,
	// 		DateCreated:  note.DateCreated,
	// 		DateModified: note.DateModified,
	// 		Title:        note.Title,
	// 		Body:         note.Body,
	// 		Type:         note.Type,
	// 		Team:         note.Team,
	// 		Link:         note.Link,
	// 	})
	// notes.Notes = append(notes.Notes, proto.Marshal(note))
	// }

	// return notes, nil
	// return resp, nil
}

func EncodeGRPCEmptyResponse(_ context.Context, response interface{}) (interface{}, error) {
	return pb.Empty{}, nil
}

func EncodeGRPCNoteResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.Note)

	return resp, nil // proto.Unmarshal(resp, &pb.Note)
}

// decodeGRPCSumResponse is a transport/grpc.DecodeResponseFunc that converts a
// gRPC sum reply to a user-domain sum response. Primarily useful in a client.
// func decodeGRPCSumResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
// 	reply := grpcReply.(*pb.SumReply)
// 	return endpoints.LoginResponse{V: int(reply.V), Err: str2err(reply.Err)}, nil
// }

// decodeGRPCSumRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC sum request to a user-domain sum request. Primarily useful in a server.
func DecodeGRPCNoteRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.Note)
	return req, nil
}

func DecodeGRPCQueryRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.Query)
	return req, nil
}

// func DecodeGRPCCreateUserRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
// 	req := grpcReq.(*pb.CreateUserRequest)
// 	return api.CreateUserRequest{Email: req.Email, Password: req.Password}, nil
// }

// These annoying helper functions are required to translate Go error types to
// and from strings, which is the type we use in our IDLs to represent errors.
// There is special casing to treat empty strings as nil errors.

func str2err(s string) error {
	if s == "" {
		return nil
	}
	return errors.New(s)
}

func err2str(err error) string {
	if err == nil {
		return ""
	}
	return err.Error()
}
