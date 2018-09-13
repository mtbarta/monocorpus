package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/mtbarta/monocorpus/pkg/auth"
	"github.com/mtbarta/monocorpus/pkg/discovery"

	graphql "github.com/graph-gophers/graphql-go"
	"github.com/micro/go-micro"

	app "github.com/mtbarta/monocorpus/pkg/gateway"
	"github.com/mtbarta/monocorpus/pkg/healthcheck"
	"github.com/mtbarta/monocorpus/pkg/logging"
	notesPB "github.com/mtbarta/monocorpus/pkg/notes"
	"github.com/mtbarta/monocorpus/pkg/search"

	"github.com/gorilla/mux"
	"github.com/spf13/viper"
)

func main() {
	viper.AutomaticEnv()

	searchIndex := viper.GetString("searchIndex")
	searchType := viper.GetString("searchType")
	clientID := viper.GetString("clientID")
	port := viper.GetString("HttpPort")
	issuerURL := viper.GetString("issuerURL")
	jwksURL := viper.GetString("jwksURL")

	logger := logging.NewProductionLogger()

	r := mux.NewRouter()

	notesService := micro.NewService(
		micro.Name("notes"),
		micro.Version("latest"),
	)
	notesService.Init()
	notesClient := notesPB.NewNotesService("notes", notesService.Client())

	elasticLocation, err := discovery.GetMicroService(notesService.Options().Registry, "elasticsearch")
	if err != nil {
		fmt.Print(err)
		logger.Fatalf("failed to connect to elasticsearch")
	}

	searchClient, err := search.NewNoteSearcher(elasticLocation.Address, elasticLocation.Port, searchIndex, searchType)
	if err != nil {
		logger.Fatalf("failed to create search client")
	}

	putSink := micro.NewPublisher("notes.search.put", notesService.Client())
	updateSink := micro.NewPublisher("notes.search.update", notesService.Client())
	deleteSink := micro.NewPublisher("notes.search.delete", notesService.Client())

	graphqlResolver := app.NewResolver(notesClient, searchClient, putSink, updateSink, deleteSink)
	schema := graphql.MustParseSchema(app.Schema, graphqlResolver)

	var notesHandler http.Handler
	{
		notesHandler = app.MakeNotesHandler(schema)
		notesHandler = accessControl(notesHandler)
		notesHandler = auth.TeamAuthMiddleware(notesHandler)
		notesHandler = auth.NewJWTMiddleware(clientID, "user", jwksURL, issuerURL, notesHandler)
	}

	service := micro.NewService(
		micro.Name("gateway"),
	)
	service.Init()

	logger.Infof("creating service", "service", "gateway")
	// basically hack the service discovery mechanism to register the gateway.
	service.Server().Register()
	defer service.Server().Deregister()

	r.Handle("/notes", notesHandler)
	r.HandleFunc("/health", healthcheck.HealthCheckHandlerFunc)

	// Interrupt handler.
	errc := make(chan error)
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errc <- fmt.Errorf("%s", <-c)
	}()

	// HTTP transport.
	go func() {
		// logger.Fatalf("transport", "HTTP", "addr", port)
		errc <- http.ListenAndServe(":"+port, r)
	}()

	// logger.Log("error", http.ListenAndServe(":8080", nil))

	// Run!
	logger.Fatalf("exit", <-errc)
}

//login

func accessControl(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "localhost")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Authorization, Credentials")

		if r.Method == "OPTIONS" {
			return
		}

		h.ServeHTTP(w, r)
	})
	// origins := handlers.AllowedOrigins([]string{"*"})
	// return handlers.CORS(origins)(h)
}
