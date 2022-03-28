package sqlstore_test

import (
	"log"
	"os"
	"testing"

	"github.com/BurntSushi/toml"
	"github.com/wrs-news/bfb-user-microservice/internal/config"
)

var (
	TestConfig = config.NewConfig()
)

func TestMain(m *testing.M) {
	if _, err := toml.DecodeFile("../../../config/config.local.toml", TestConfig); err != nil {
		log.Println(err)
		os.Exit(1)
	}

	os.Exit(m.Run())
}
