package client

import (
	"crawler/concurrent/engine"
	"crawler/distributed/config"
	"crawler/distributed/suport"
	"encoding/json"
	"log"
)

// 存储爬取项
func ItemSaver(host string) (chan engine.Item, error) {
	// NewClient连接到RPC服务端服务
	cli, err := suport.NewClient(host)
	if err != nil {
		return nil, err
	}

	// 数据存储通道
	saver := make(chan engine.Item)

	// 异步接收数据并保存
	go func(saver chan engine.Item) {
		var itemNum int = 1
		for {
			// 接收数据
			var item = <-saver
			bytes, _ := json.Marshal(item)
			log.Printf("Item saver: got item #%d: %+v\n", itemNum, string(bytes))
			itemNum++

			// 调用ItemSaverServiceRPC存储数据到Elasticsearch
			result := ""
			err := cli.Call(config.ItemSaverRPC, item, &result)
			if err != nil || result != "ok" {
				bytes, _ := json.Marshal(item)
				log.Printf("Item saver: error saving item %v: %v\n", err, string(bytes))
			}
		}
	}(saver)
	return saver, nil
}
