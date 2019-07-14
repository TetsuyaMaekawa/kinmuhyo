package rds

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

// Client DBクライアント
type Client struct {
	DB *gorm.DB
}

// NewClient RDSへの接続を確立
func NewClient() (*Client, error) {
	me := new(Client)

	// コネクション生成
	db, err := gorm.Open("Dialect", "DSN")
	if err != nil {
		return nil, errors.Wrapf(err, fmt.Sprintf("Unable to connect databese. Dialect: %v DSN: %v", "Dialect", "DSN"))
	}

	// プール設定
	db.DB().SetMaxIdleConns(1)
	db.DB().SetMaxOpenConns(1)

	// ログモード設定
	db.LogMode(true)

	me.DB = db
	return me, nil
}

// Close RDSのコネクションクローズ
func (c *Client) Close() {
	if c.DB != nil {
		c.DB.Close()
	}
}
