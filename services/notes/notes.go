package main

import (
	"os"
	"time"

	micro "github.com/micro/go-micro"
	"github.com/spf13/viper"
	mgo "gopkg.in/mgo.v2"

	"github.com/mtbarta/monocorpus/pkg/discovery"
	"github.com/mtbarta/monocorpus/pkg/logging"
	notes "github.com/mtbarta/monocorpus/pkg/notes"
	svc "github.com/mtbarta/monocorpus/pkg/notes/service"
)

var logger = logging.NewProductionLogger()

func instantiateMongo(dialInfo *mgo.DialInfo, collection string) *mgo.Collection {
	session, err := mgo.DialWithInfo(dialInfo)
	if err != nil {
		logger.Fatal(err)
		os.Exit(1)
	}

	index := mgo.Index{
		Key: []string{"$text:author", "$text:team", "$text:tags.text"},
	}

	mongoColl := session.DB("").C(collection)
	err = mongoColl.EnsureIndex(index)
	if err != nil {
		logger.Fatal(err)
	}

	return mongoColl
}

func createMongoDialInfo(address string, dbName string) mgo.DialInfo {
	addressList := make([]string, 1)
	addressList = append(addressList, address)

	dialInfo := mgo.DialInfo{
		Addrs:    addressList,
		Database: dbName,
	}

	return dialInfo
}

func main() {
	viper.AutomaticEnv()
	dbName := viper.GetString("dbName")
	collection := viper.GetString("collection")

	service := micro.NewService(
		micro.Name("notes"),
		micro.Version("latest"),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*15),
	)
	service.Init()

	mongoLocation, err := discovery.GetMicroService(service.Options().Registry, "mongo")
	if err != nil {
		logger.Fatal(err)
	}

	dialInfo := createMongoDialInfo(mongoLocation.Address, dbName)
	mongoColl := instantiateMongo(&dialInfo, collection)

	logger.Infof("connected to mongo db",
		"collection", mongoColl.FullName,
	)

	handler := svc.NoteService{Collection: mongoColl}
	notes.RegisterNotesHandler(service.Server(), handler)

	if err := service.Run(); err != nil {
		logger.Fatal(err)
	}
}
