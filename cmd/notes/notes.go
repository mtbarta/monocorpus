package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	mgo "gopkg.in/mgo.v2"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/metrics"
	"github.com/go-kit/kit/metrics/prometheus"
	// "github.com/go-kit/kit/transport/grpc"
	"github.com/oklog/oklog/pkg/group"
	stdopentracing "github.com/opentracing/opentracing-go"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
	"github.com/spf13/viper"
	"google.golang.org/grpc"

	pb "github.com/mtbarta/monocorpus/pkg/notes"
	"github.com/mtbarta/monocorpus/pkg/notes/endpoints"
	"github.com/mtbarta/monocorpus/pkg/notes/service"
	"github.com/mtbarta/monocorpus/pkg/notes/transport"
)

func main() {
	viper.AutomaticEnv()
	// viper.SetConfigName("notes")
	// viper.AddConfigPath("/home/matt/go/src/github.com/mtbarta/monocorpus/notes/configs")

	// err := viper.ReadInConfig() // Find and read the config file
	// if err != nil { // Handle errors reading the config file
	// 	panic(fmt.Errorf("Fatal error config file: %s \n", err))
	// }

	grpcAddr := viper.GetString("port")
	httpAddr := viper.GetString("httpPort")
	dbHost := viper.GetString("dbHost")
	dbPort := viper.GetString("dbPort")
	dbName := viper.GetString("dbName")
	collection := viper.GetString("collection")

	// Create a single logger, which we'll use and give to other components.
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stdout)
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}

	//https://github.com/go-kit/kit/blob/master/examples/addsvc/cmd/addsvc/addsvc.go
	var tracer stdopentracing.Tracer
	logger.Log("tracer", "none")
	tracer = stdopentracing.GlobalTracer() // no-op

	connStr := fmt.Sprintf("%s:%s", dbHost, dbPort)
	session, err := mgo.Dial(connStr)
	if err != nil {
		logger.Log("db", err)
		os.Exit(1)
	} else {
		logger.Log("db", connStr)
	}

	index := mgo.Index{
		Key: []string{"$text:author", "$text:team", "$text:tags.text"},
	}

	mongoColl := session.DB(dbName).C(collection)
	err = mongoColl.EnsureIndex(index)
	if err != nil {
		logger.Log("mongoIndex", err)
	}

	var duration metrics.Histogram
	{
		// Endpoint-level metrics.
		duration = prometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
			Namespace: "example",
			Subsystem: "addsvc",
			Name:      "request_duration_seconds",
			Help:      "Request duration in seconds.",
		}, []string{"method", "success"})
	}

	r := mux.NewRouter()
	// r.Handle("/metrics", promhttp.Handler())

	r.HandleFunc("/health", service.HealthCheckHandlerFunc)

	var (
		service    = service.New(logger, mongoColl)
		endpoints  = endpoints.New(service, logger, duration, tracer)
		grpcServer = transport.MakeGRPCServer(endpoints, logger)
	)

	var g group.Group
	{
		// The gRPC listener mounts the Go kit gRPC server we created.
		grpcListener, err := net.Listen("tcp", ":"+grpcAddr)
		if err != nil {
			logger.Log("transport", "gRPC", "during", "Listen", "err", err)
			os.Exit(1)
		}
		g.Add(func() error {
			logger.Log("transport", "gRPC", "addr", grpcAddr)
			baseServer := grpc.NewServer()

			pb.RegisterNotesServer(baseServer, grpcServer)

			return baseServer.Serve(grpcListener)
		}, func(error) {
			grpcListener.Close()
		})
	}
	{
		// g.Add(func() error {
		// 	logger.Log("transport", "HTTP", "addr", httpAddr)
		// 	errc <- http.ListenAndServe(":"+httpAddr, r)
		// })
		// The HTTP listener mounts the Go kit HTTP handler we created.
		httpListener, err := net.Listen("tcp", ":"+httpAddr)
		if err != nil {
			logger.Log("transport", "HTTP", "during", "Listen", "err", err)
			os.Exit(1)
		}
		g.Add(func() error {
			logger.Log("transport", "HTTP", "addr", httpAddr)
			return http.Serve(httpListener, r)
		}, func(error) {
			httpListener.Close()
		})
	}
	{
		// This function just sits and waits for ctrl-C.
		cancelInterrupt := make(chan struct{})
		g.Add(func() error {
			c := make(chan os.Signal, 1)
			signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
			select {
			case sig := <-c:
				return fmt.Errorf("received signal %s", sig)
			case <-cancelInterrupt:
				return nil
			}
		}, func(error) {
			close(cancelInterrupt)
		})
	}
	logger.Log("exit", g.Run())
}
