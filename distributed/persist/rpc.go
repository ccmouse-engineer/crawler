package persist

import (
	"crawler/concurrent/engine"
	"crawler/concurrent/persist"

	"github.com/olivere/elastic"
)

// ItemSaverService定义RPC保存爬虫记录服务
type ItemSaverService struct {
	Client *elastic.Client
	Index  string
}

// Save定义RPC保存爬虫记录服务方法
func (s *ItemSaverService) Save(item engine.Item, result *string) error {
	err := persist.Save(item, s.Client, s.Index)
	if err == nil {
		*result = "ok"
	}
	return err
}
