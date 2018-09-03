package main

import (
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/lib/pq"
	micro "github.com/micro/go-micro"
	"github.com/spf13/viper"
	mgo "gopkg.in/mgo.v2"

	notes "github.com/mtbarta/monocorpus/pkg/notes"
	svc "github.com/mtbarta/monocorpus/pkg/notes/service"
)

func instantiateMongo(dbHost, dbPort, dbName, collection string) *mgo.Collection {
	connStr := fmt.Sprintf("%s:%s", dbHost, dbPort)
	session, err := mgo.Dial(connStr)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	index := mgo.Index{
		Key: []string{"$text:author", "$text:team", "$text:tags.text"},
	}

	mongoColl := session.DB(dbName).C(collection)
	err = mongoColl.EnsureIndex(index)
	if err != nil {
		log.Fatal(err)
	}

	return mongoColl
}

func main() {
	dbHost := viper.GetString("dbHost")
	dbPort := viper.GetString("dbPort")
	dbName := viper.GetString("dbName")
	collection := viper.GetString("collection")

	mongoColl := instantiateMongo(dbHost, dbPort, dbName, collection)
	handler := svc.NoteService{Collection: mongoColl}

	service := micro.NewService(
		micro.Name("notes"),
		micro.Version("latest"),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*15),
	)

	service.Init()
	notes.RegisterNotesHandler(service.Server(), handler)

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
