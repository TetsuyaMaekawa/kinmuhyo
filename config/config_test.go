package config

import "testing"

var configPath = "../kinmuhyo.toml"

func TestNewConfigSuccess(t *testing.T) {
	config, err := NewConfig(configPath)
	if err != nil {
		t.Errorf("Test failed. Dettail :%v", err)
	}

	if config == nil {
		t.Errorf("Test failed. Return config object return nil.")
	}

	if config.AppName != "kinmuhyo" {
		t.Error("Test failed. app_name not found.")
	}

	if config.LogLevel != "debug" {
		t.Error("Test failed. LogLevel not found.")
	}

	RDS := config.Client.RDS
	if RDS == nil {
		t.Error("Test failed. RDS object is nil.")
	}

	if RDS.Dialect == "" {
		t.Error("Test failed. RDS.Dialect is empty.")
	}

	if RDS.Dsn == "" {
		t.Error("Test failed. RDS.Dsn is empty.")
	}

	if RDS.IdleConn != 10 {
		t.Error("Test failed. RDS.IdleConn not found.")
	}

	if RDS.MaxConn != 10 {
		t.Error("Test failed. RDS.MaxConn not found.")
	}

	if !RDS.Debug {
		t.Error("Test failed. RDS.Debug not found.")
	}

}

func TestNewConfigWrogFilePath(t *testing.T) {
	config, err := NewConfig("wrong/file/path/config.toml")
	if err == nil {
		t.Errorf("Test failed. Error object is nil.")
	}
	if config != nil {
		t.Errorf("Test failed. Return config object must be nil.")
	}
}
