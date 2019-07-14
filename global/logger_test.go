package global

import (
	"log"
	"os"
	"testing"

	"rootship.co.jp/kinmuhyo/config"
)

// このテストを実施した場合同一ディレクトリ内に
// test.logがゴミとして残るため注意
// 手作業で削除
// os.Removeで削除しようとすると失敗
func TestNewLogger(t *testing.T) {
	conf := config.Config{
		LogPath: "test.log",
	}

	if err := NewLogger(&conf); err != nil {
		t.Errorf("Test failed. Detail: %v", err)
	}

	log.Print("test")

	if _, err := os.Stat("test.log"); err != nil {
		t.Error("Test failed. LogFile not found.")
	}
}
