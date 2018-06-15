package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-kit/kit/log"
	graphql "github.com/graph-gophers/graphql-go"
	"google.golang.org/grpc"

	app "github.com/mtbarta/monocorpus/pkg/gateway"
	"github.com/mtbarta/monocorpus/pkg/healthcheck"
	notesPB "github.com/mtbarta/monocorpus/pkg/notes"
	"github.com/mtbarta/monocorpus/pkg/notes/search"

	"github.com/gorilla/mux"
	"github.com/spf13/viper"
)

func main() {
	viper.AutomaticEnv()

	notesHost := viper.GetString("notesHost")
	notesPort := viper.GetString("notesPort")

	searchHost := viper.GetString("searchHost")
	searchPort := viper.GetString("searchPort")
	searchIndex := viper.GetString("searchIndex")
	searchType := viper.GetString("searchType")
	port := viper.GetString("HttpPort")

	pubKey := viper.GetString("pubKey")

	// Logging domain.
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stdout)
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}

	r := mux.NewRouter()

	logger.Log("notesserver", notesHost+":"+notesPort)
	notesConn, err := grpc.Dial(notesHost+":"+notesPort, grpc.WithInsecure())
	if err != nil {
		logger.Log("did not connect to notes:", err)
		os.Exit(1)
	}
	defer notesConn.Close()

	searchClient, err := search.NewNoteSearcher(searchHost, searchPort, searchIndex, searchType)
	if err != nil {
		logger.Log("error", err)
		fmt.Println("failed to connect to search server")
	}
	if searchClient != nil {
		searchClient.EnsureIndex()
	}

	logger.Log("searchHost", searchHost, "searchPort", searchPort, "searchIndex", searchIndex, "searchType", searchType)

	notesClient := notesPB.NewNotesClient(notesConn)

	graphqlResolver := app.NewResolver(notesClient, searchClient, &logger)
	schema := graphql.MustParseSchema(app.Schema, graphqlResolver)

	jwtMiddleware := app.NewJWTMiddleware([]byte(pubKey))
	var notesHandler http.Handler
	{
		notesHandler = app.MakeNotesHandler(schema, &logger)
		notesHandler = accessControl(notesHandler)
		notesHandler = app.TeamAuthMiddleware(notesHandler)
		notesHandler = jwtMiddleware.Handler(notesHandler)
	}

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
		logger.Log("transport", "HTTP", "addr", port)
		errc <- http.ListenAndServe(":"+port, r)
	}()

	// logger.Log("error", http.ListenAndServe(":8080", nil))

	// Run!
	logger.Log("exit", <-errc)
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
