package global

import (
	"log"
	"os"
	"time"

	"github.com/comail/colog"
	"rootship.co.jp/kinmuhyo/config"
)

// NewLogger ログ準備
func NewLogger(conf *config.Config) {
	file, err := os.OpenFile(conf.LogPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		panic(err)
	}
	// 標準仕様のログをcologに変更
	colog.Register()
	colog.SetOutput(file)
	colog.ParseFields(true)
	// ログに出力するフォーマット設定
	colog.SetFormatter(&colog.JSONFormatter{
		TimeFormat: time.RFC3339,
		Flag:       log.Lshortfile,
	})
}
