package system

import (
	"testing"
)

func TestNewContext(t *testing.T) {
	confPath := "../kinmuhyo.toml"
	context, err := NewContext(confPath)
	if err != nil {
		t.Errorf("Test failed. Detail: %v", err)
	}

	if context.Conf == nil {
		t.Error("Test failed. context.Conf must not be nil.")
	}

	if context.Rds == nil {
		t.Error("Test failed. context.Rds must not be nil.")
	}
}
