package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"log"

	"github.com/yurioganesyan/LinkShortener/pkg/api"
	"google.golang.org/grpc"
)

func main() {
	flag.Parse()
	if flag.NArg() < 1 {
		log.Fatalln("argument is missing")
	}

	command := flag.Arg(0)

	connection, error := grpc.Dial(":9080", grpc.WithInsecure())
	if error != nil {
		log.Fatalln(error)
	}
	defer connection.Close()

	conn := api.NewShortenerClient(connection)

	fmt.Println("Please, type the link:")

	var link string
	fmt.Scanln(&link)

	if command == "getlink" && command != "" {
		res, error := conn.Get(context.Background(), &api.GetRequest{Shorturl: link})
		if error != nil {
			log.Fatalln(error)
		}
		log.Println(res.GetUrl())
	} else if command == "createlink" && command != "" {
		res, error := conn.Create(context.Background(), &api.CreateRequest{Url: link})
		if error != nil {
			log.Fatalln(error)
		}
		log.Println(res.GetShorturl())
	} else {
		log.Fatalln(errors.New("wrong command, please try again"))
	}

}
