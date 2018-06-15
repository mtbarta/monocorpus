package types

import (
	"time"

	pb "github.com/mtbarta/monocorpus/pkg/notes"
)

type Query struct {
	IDs      []string  `json:"IDs,omitempty"`
	Title    string    `json:"title,omitempty"`
	Team     string    `json:"team,omitempty"`
	Authors  []string  `json:"authors,omitempty"`
	Todate   time.Time `json:"todate,omitempty"`
	Fromdate time.Time `json:"fromdate,omitempty"`
	Tags     []string  `json:"tags,omitempty"`
}

func PbQueryToMongo(query *pb.Query) Query {
	response := Query{
		IDs:     query.IDs,
		Title:   query.Title,
		Team:    query.Team,
		Authors: query.Authors,
		Tags:    query.Tags,
	}

	if query.Todate != nil {
		response.Todate = time.Unix(query.Todate.GetSeconds(), 0)
	} else {
		response.Todate = time.Time{}
	}

	if query.Fromdate != nil {
		response.Fromdate = time.Unix(query.Fromdate.GetSeconds(), 0)
	} else {
		response.Todate = time.Time{}
	}
	return response
}
