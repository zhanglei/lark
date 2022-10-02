package main

/*
xpack.security.enabled: false

docker-compose -f docker-compose-dev.yaml up
*/
import (
	"lark/examples/config"
	"lark/pkg/common/xes"
)

func init() {

}

func main() {
	cfg := config.GetConfig()
	xes.NewElasticsearchClient(cfg.Elasticsearch)
}
