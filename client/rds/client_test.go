package rds

import (
	"testing"

	"rootship.co.jp/kinmuhyo/config"
)

func TestNewClient(t *testing.T) {
	conf := config.Config{
		Client: &config.Client{
			RDS: &config.RDS{
				Dialect:  "mysql",
				Dsn:      "root:Ah4vn3tetsuya@/mydb?charset=utf8&parseTime=True&loc=Local",
				IdleConn: 10,
				MaxConn:  10,
				Debug:    true,
			},
		},
	}

	client, err := NewClient(&conf)
	defer client.Close()

	if err != nil {
		t.Errorf("Test failed. Detail: %v", err)
	}

	if client == nil {
		t.Error("Test failed. Client object must not be nil.")
	}
}
