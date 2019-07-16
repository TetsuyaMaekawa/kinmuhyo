package main

import (
	"log"
	"net/http"

	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"

	"rootship.co.jp/kinmuhyo/handler/kinmuhyo"
	"rootship.co.jp/kinmuhyo/system"
)

const configPath = "../../kinmuhyo.toml"

func setup() {
	context, err := system.NewContext(configPath)
	if err != nil {
		log.Fatal(err)
		return
	}

	handler := kinmuhyo.Handler{Rds: context.Rds, Line: context.Line}

	// // Postのルーティング
	goji.Post("/callback", func(context web.C, writer http.ResponseWriter, request *http.Request) {
		handler.Handler(request)
	})

}
