package line

import "rootship.co.jp/kinmuhyo/config"

// Client lineクライアント
type Client struct {
	ChannelSenret string
	AccessToken   string
}

// NewClient lineクライアント生成
func NewClient(conf *config.Config) *Client {
	line := conf.Client.LINE
	me := new(Client)
	me.AccessToken = line.AccessToken
	me.ChannelSenret = line.ChannelSecret
	return me
}
