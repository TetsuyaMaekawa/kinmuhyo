package kinmuhyo

import (
	"log"
	"net/http"

	"github.com/line/line-bot-sdk-go/linebot"

	"rootship.co.jp/kinmuhyo/client/line"
	"rootship.co.jp/kinmuhyo/client/rds"
)

// Handler Handler設定情報
type Handler struct {
	Rds  *rds.Client
	Line *line.Client
}

// Handler リクエストハンドル処理
func (h *Handler) Handler(request *http.Request) {

	events, err := h.Line.Bot.ParseRequest(request)
	if err != nil {
		log.Printf("error: Unable to ParseRequest. Detail: %v", err)
		return
	}

	for _, event := range events {
		if event.Type == linebot.EventTypeMessage {

		}
	}
}
