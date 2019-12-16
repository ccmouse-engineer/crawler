package main

import (
	"crawler/distributed/config"
	"crawler/distributed/persist"
	"crawler/distributed/suport"
	"fmt"
	"log"

	"github.com/olivere/elastic"
)

func main() {
	log.Fatal(serverRPC(fmt.Sprintf(":%d", config.ItemSaverPort), config.IterSaverElasticsearchIndex))
}

// serverRPC启动RPC服务器端服务
func serverRPC(host, index string) error {
	// 获取Elasticsearch客户端
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		return err
	}

	// SeverRPC开启RPC服务器端服务
	err = suport.SeverRPC(host, &persist.ItemSaverService{
		Client: client,
		Index:  index,
	})
	if err != nil {
		return err
	}
	return nil
}
