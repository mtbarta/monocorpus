package main

import (
	"context"
	"time"

	"github.com/mtbarta/monocorpus/pkg/notes"

	_ "github.com/lib/pq"
	micro "github.com/micro/go-micro"
	registry "github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/server"
	"github.com/spf13/viper"

	"github.com/mtbarta/monocorpus/pkg/discovery"
	"github.com/mtbarta/monocorpus/pkg/logging"
	"github.com/mtbarta/monocorpus/pkg/routing"
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

	micro.RegisterSubscriber(routing.NOTE_PUT_CHANNEL, 
		service.Server(), 
		func(ctx context.Context, note *notes.Note) error {
		logging.Logger.Debug("publishing note")
		err := handler.Put(ctx, note)
		if err != nil {
			logger.Errorf("failed to put note", "err", err)
		}
		return nil
	}, server.SubscriberQueue("queue.search.update")), server.SubscriberQueue("queue.search.put"))
	micro.RegisterSubscriber(routing.NOTE_UPDATE_CHANNEL,
		service.Server(),
		func(ctx context.Context, note *notes.Note) error {
			logging.Logger.Debug("update note")
			err := handler.Update(ctx, note)
			if err != nil {
				logger.Errorf("failed to update note", "err", err)
			}
			return nil
		}, server.SubscriberQueue("queue.search.update"))
	micro.RegisterSubscriber(routing.NOTE_DELETE_CHANNEL, 
		service.Server(), 
		func(ctx context.Context, note *notes.Note) error {
		logging.Logger.Debug("delete note")
		err := handler.Update(ctx, note)
		if err != nil {
			logger.Errorf("failed to delete note", "err", err)
		}
		return nil
	}, server.SubscriberQueue("queue.search.update")), server.SubscriberQueue("queue.search.delete"))

	if err := service.Run(); err != nil {
		logger.Fatal(err)
	}
}
