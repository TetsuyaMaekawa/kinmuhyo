package global

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/comail/colog"
	"github.com/pkg/errors"
	"rootship.co.jp/kinmuhyo/config"
)

// NewLogger ログ準備
func NewLogger(conf *config.Config) error {
	file, err := os.OpenFile(conf.LogPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		return errors.Wrapf(err, fmt.Sprintf("Unable to create logfile. logPath: %v", conf.LogPath))
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

	return nil
}
