package rds

import (
	"fmt"

	"rootship.co.jp/kinmuhyo/config"

	_ "github.com/go-sql-driver/mysql" // for mysql
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

// Client DBクライアント
type Client struct {
	DB *gorm.DB
}

// NewClient RDSへの接続を確立
func NewClient(conf *config.Config) (*Client, error) {
	RDS := conf.Client.RDS
	me := new(Client)

	// コネクション生成
	db, err := gorm.Open(RDS.Dialect, RDS.Dsn)
	if err != nil {
		return nil, errors.Wrapf(err, fmt.Sprintf("Unable to connect databese. Dialect: %v DSN: %v", "Dialect", "DSN"))
	}

	// プール設定
	db.DB().SetMaxIdleConns(RDS.IdleConn)
	db.DB().SetMaxOpenConns(RDS.MaxConn)

	// ログモード設定
	db.LogMode(RDS.Debug)

	me.DB = db
	return me, nil
}

// Close RDSのコネクションクローズ
func (c *Client) Close() {
	if c.DB != nil {
		c.DB.Close()
	}
}
