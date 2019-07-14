package line

import (
	"testing"

	"rootship.co.jp/kinmuhyo/config"
)

func TestNewClient(t *testing.T) {
	conf := config.Config{
		Client: &config.Client{
			LINE: config.LINE{
				AccessToken:   "testAccessToken",
				ChannelSecret: "testChannelSecret",
			},
		},
	}

	line := NewClient(&conf)

	if line.AccessToken != "testAccessToken" {
		t.Error("Test failed. line.AccessToken must be testAccessToken.")
	}

	if line.ChannelSenret != "testChannelSecret" {
		t.Error("Test failed. line.ChannelSenret must be testChannelSecret.")
	}
}
