package service

import (
	"strings"

	"github.com/line/line-bot-sdk-go/linebot"
	"rootship.co.jp/kinmuhyo/client/rds"
	"rootship.co.jp/kinmuhyo/static"
	"rootship.co.jp/kinmuhyo/util"
)

// DistributeAction コマンド毎に処理を振り分ける
func DistributeAction(rds *rds.Client, event *linebot.TextMessage) {
	message := event.Text
	command := strings.Split(message, " ")

	if command[0] == "/s" {
		insertShukkin(command, event)
	} else if command[0] == "/t" {
		insertTaikin(command, event)
	} else if command[0] == "/m" {
		insertMastar(command, event)
	} else {
		// 不明コマンドエラーを返却
	}
}

func insertShukkin(command []string, event *linebot.TextMessage) {
	if util.IsArgsNum(command, static.MCommandArgsNum) {
		// 引数エラーを返却
	}

}

func insertTaikin(command []string, event *linebot.TextMessage) {
	if util.IsArgsNum(command, static.MCommandArgsNum) {
		// 引数エラーを返却
	}

}

func insertMastar(command []string, event *linebot.TextMessage) {
	if util.IsArgsNum(command, static.MCommandArgsNum) {
		// 引数エラーを返却
	}

}
