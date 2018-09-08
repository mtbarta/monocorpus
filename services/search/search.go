package main

import (
	"time"

	_ "github.com/lib/pq"
	micro "github.com/micro/go-micro"
	registry "github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/server"
	"github.com/spf13/viper"

	"github.com/mtbarta/monocorpus/pkg/discovery"
	"github.com/mtbarta/monocorpus/pkg/logging"
	"github.com/mtbarta/monocorpus/pkg/search"
)

var logger = logging.NewProductionLogger()

func getDBAddresses(registry registry.Registry, dbName string) ([]string, error) {
	services, err := registry.GetService(dbName)

	if err != nil {
		return nil, err
	}

	var addresses []string
	for _, service := range services {
		for _, node := range service.Nodes {
			address := node.Address
			addresses = append(addresses, address)
		}
	}

	return addresses, nil
}

func main() {
	viper.AutomaticEnv()
	searchIndex := viper.GetString("searchIndex")
	searchType := viper.GetString("searchType")

	service := micro.NewService(
		micro.Name("search"),
		micro.Version("latest"),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*15),
	)
	service.Init()

	elasticLocation, err := discovery.GetMicroService(service.Options().Registry, "elasticsearch")
	if err != nil {
		logger.Fatalf("failed to connect to elasticsearch")
	}

	handler, err := search.NewNoteManager(elasticLocation.Address, elasticLocation.Port, searchIndex, searchType)
	if err != nil {
		logger.Fatal(err)
	}

	logger.Infof("connected to elasticsearch",
		"host", elasticLocation.Address,
		"port", elasticLocation.Port,
		"index", searchIndex,
		"type", searchType)

	micro.RegisterSubscriber("note.search.put", service.Server(), handler.Put, server.SubscriberQueue("queue.search.put"))
	micro.RegisterSubscriber("note.search.update", service.Server(), handler.Update, server.SubscriberQueue("queue.search.update"))
	micro.RegisterSubscriber("note.search.delete", service.Server(), handler.Delete, server.SubscriberQueue("queue.search.delete"))

	if err := service.Run(); err != nil {
		logger.Fatal(err)
	}
}
