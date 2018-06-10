package gateway

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-kit/kit/circuitbreaker"
	"github.com/go-kit/kit/log"
	graphql "github.com/graph-gophers/graphql-go"
	"github.com/sony/gobreaker"

	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
)

//create the notes handler to serve at /notes
func MakeNotesHandler(schema *graphql.Schema, logger *log.Logger) *httptransport.Server {
	endpoint := makeGraphQLEndpoint(schema, logger)

	endpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{}))(endpoint)

	return httptransport.NewServer(
		endpoint,
		decodeGraphQLRequest,
		encodeGraphQLResponse,
		// httptransport.ServerBefore(httptransport.PopulateRequestContext),
	)
}

func makeGraphQLEndpoint(schema *graphql.Schema, _ *log.Logger) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {

		// _, ok := ctx.Value(httptransport.ContextKeyRequestAuthorization).(string)
		// if !ok {
		// 	return nil, AuthError{}
		// }
		r := request.(GraphQLRequest)

		resp := schema.Exec(ctx, r.Query, r.OperationName, r.Variables)

		return resp, nil
	}
}

type GraphQLRequest struct {
	Query         string                 `json:"query"`
	OperationName string                 `json:"operationName"`
	Variables     map[string]interface{} `json:"variables"`
}

func encodeGraphQLResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	responseJSON, err := json.Marshal(response)
	if err != nil {
		// http.Error(w, err.Error(), http.StatusInternalServerError)
		return err
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(responseJSON)
	return nil
}

func decodeGraphQLRequest(ctx context.Context, req *http.Request) (interface{}, error) {
	var params GraphQLRequest
	// fmt.Println(json.Marshal(req))

	if err := json.NewDecoder(req.Body).Decode(&params); err != nil {
		// http.Error(w, err.Error(), http.StatusBadRequest)
		return nil, err
	}

	return params, nil
}
