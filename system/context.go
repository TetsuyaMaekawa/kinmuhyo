package system

import (
	"rootship.co.jp/kinmuhyo/client/line"
	"rootship.co.jp/kinmuhyo/client/rds"
	"rootship.co.jp/kinmuhyo/config"
	"rootship.co.jp/kinmuhyo/global"
)

// Context コンテキスト
type Context struct {
	Conf *config.Config
	Rds  *rds.Client
	Line *line.Client
}

// NewContext コンテキストの生成
func NewContext(confPath string) (*Context, error) {

	me := new(Context)
	conf, err := config.NewConfig(confPath)
	if err != nil {
		return nil, err
	}

	if err := global.NewLogger(conf); err != nil {
		return nil, err
	}

	rds, err := rds.NewClient(conf)
	if err != nil {
		return nil, err
	}

	line := line.NewClient(conf)

	me.Conf = conf
	me.Rds = rds
	me.Line = line

	return me, nil
}
