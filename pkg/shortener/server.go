package shortener

import (
	"context"
	"log"

	"github.com/yurioganesyan/LinkShortener/pkg/api"
	"github.com/yurioganesyan/LinkShortener/pkg/database"
)

type GRPCServer struct{}

func (s *GRPCServer) Get(ctx context.Context, req *api.GetRequest) (*api.GetResponse, error) {
	err := database.GetConnection()
	if err != nil {
		log.Fatalln(err)
	}

	var url string

	linkID := database.GetShortLinkID(req.Shorturl)

	url = database.GetLink(linkID)

	return &api.GetResponse{Url: url}, nil
}

func (s *GRPCServer) Create(ctx context.Context, req *api.CreateRequest) (*api.CreateResponse, error) {
	err := database.GetConnection()
	if err != nil {
		log.Fatalln(err)
	}

	var shortUrl string

	linkID := database.GetLinkID(req.Url)

	if linkID == 0 {
		maxID := database.GetMaxID()
		linkID, err = database.CreateUrl(req.Url, maxID)
		if err != nil {
			log.Fatalln(err)
		}
		shortUrl = CutUrl()
		shortUrlID := database.GetMaxShortID() + 1
		database.SaveShortURL(shortUrlID, shortUrl, linkID)
	} else {
		shortUrl = database.GetShortURL(linkID)
	}

	return &api.CreateResponse{Shorturl: shortUrl}, nil
}
