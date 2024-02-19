package config

import (
	"testing"
)

func TestGetConfig(t *testing.T) {

	config := GetConfig()
	t.Logf("config: %+v", config)
}
