package persist

import (
	"context"
	"crawler/concurrent/engine"
	"encoding/json"
	"log"

	"github.com/olivere/elastic"
)

// 字段属性配置
const mappings = `{
  "mappings": {
	"dynamic": false,
    "properties": {
      "id": {
        "type": "text",
        "fields": {
          "keyword": {
            "type": "keyword",
            "ignore_above": 256
          }
        }
      },
      "payload": {
        "properties": {
          "area_name": {
            "type": "text",
            "fields": {
              "keyword": {
                "type": "keyword",
                "ignore_above": 256
              }
            }
          },
          "base_info": {
            "properties": {
              "base": {
                "properties": {
                  "产权年限": {
                    "type": "text",
                    "fields": {
                      "keyword": {
                        "type": "keyword",
                        "ignore_above": 256
                      }
                    }
                  },
                  "套内面积": {
                    "type": "text",
                    "fields": {
                      "keyword": {
                        "type": "keyword",
                        "ignore_above": 256
                      }
                    }
                  },
                  "建筑类型": {
                    "type": "text",
                    "fields": {
                      "keyword": {
                        "type": "keyword",
                        "ignore_above": 256
                      }
                    }
                  },
                  "建筑结构": {
                    "type": "text",
                    "fields": {
                      "keyword": {
                        "type": "keyword",
                        "ignore_above": 256
                      }
                    }
                  },
                  "建筑面积": {
                    "type": "text",
                    "fields": {
                      "keyword": {
                        "type": "keyword",
                        "ignore_above": 256
                      }
                    }
                  },
                  "户型结构": {
                    "type": "text",
                    "fields": {
                      "keyword": {
                        "type": "keyword",
                        "ignore_above": 256
                      }
                    }
                  },
                  "房屋户型": {
                    "type": "text",
                    "fields": {
                      "keyword": {
                        "type": "keyword",
                        "ignore_above": 256
                      }
                    }
                  },
                  "房屋朝向": {
                    "type": "text",
                    "fields": {
                      "keyword": {
                        "type": "keyword",
                        "ignore_above": 256
                      }
                    }
                  },
                  "所在楼层": {
                    "type": "text",
                    "fields": {
                      "keyword": {
                        "type": "keyword",
                        "ignore_above": 256
                      }
                    }
                  },
                  "梯户比例": {
                    "type": "text",
                    "fields": {
                      "keyword": {
                        "type": "keyword",
                        "ignore_above": 256
                      }
                    }
                  },
                  "装修情况": {
                    "type": "text",
                    "fields": {
                      "keyword": {
                        "type": "keyword",
                        "ignore_above": 256
                      }
                    }
                  },
                  "配备电梯": {
                    "type": "text",
                    "fields": {
                      "keyword": {
                        "type": "keyword",
                        "ignore_above": 256
                      }
                    }
                  }
                }
              },
              "transaction": {
                "properties": {
                  "上次交易": {
                    "type": "text",
                    "fields": {
                      "keyword": {
                        "type": "keyword",
                        "ignore_above": 256
                      }
                    }
                  },
                  "交易权属": {
                    "type": "text",
                    "fields": {
                      "keyword": {
                        "type": "keyword",
                        "ignore_above": 256
                      }
                    }
                  },
                  "产权所属": {
                    "type": "text",
                    "fields": {
                      "keyword": {
                        "type": "keyword",
                        "ignore_above": 256
                      }
                    }
                  },
                  "房屋年限": {
                    "type": "text",
                    "fields": {
                      "keyword": {
                        "type": "keyword",
                        "ignore_above": 256
                      }
                    }
                  },
                  "房屋用途": {
                    "type": "text",
                    "fields": {
                      "keyword": {
                        "type": "keyword",
                        "ignore_above": 256
                      }
                    }
                  },
                  "房本备件": {
                    "type": "text",
                    "fields": {
                      "keyword": {
                        "type": "keyword",
                        "ignore_above": 256
                      }
                    }
                  },
                  "抵押信息": {
                    "type": "text",
                    "fields": {
                      "keyword": {
                        "type": "keyword",
                        "ignore_above": 256
                      }
                    }
                  },
                  "挂牌时间": {
                    "type": "text",
                    "fields": {
                      "keyword": {
                        "type": "keyword",
                        "ignore_above": 256
                      }
                    }
                  }
                }
              }
            }
          },
          "characteristics": {
            "properties": {
              "户型介绍": {
                "type": "text",
                "fields": {
                  "keyword": {
                    "type": "keyword",
                    "ignore_above": 256
                  }
                }
              },
              "房源标签": {
                "type": "text",
                "fields": {
                  "keyword": {
                    "type": "keyword",
                    "ignore_above": 256
                  }
                }
              },
              "核心卖点": {
                "type": "text",
                "fields": {
                  "keyword": {
                    "type": "keyword",
                    "ignore_above": 256
                  }
                }
              },
              "税费解析": {
                "type": "text",
                "fields": {
                  "keyword": {
                    "type": "keyword",
                    "ignore_above": 256
                  }
                }
              }
            }
          },
          "city_name": {
            "type": "text",
            "fields": {
              "keyword": {
                "type": "keyword",
                "ignore_above": 256
              }
            }
          },
          "community_name": {
            "type": "text",
            "fields": {
              "keyword": {
                "type": "keyword",
                "ignore_above": 256
              }
            }
          },
          "price": {
            "type": "long"
          },
          "title": {
            "type": "text",
            "fields": {
              "keyword": {
                "type": "keyword",
                "ignore_above": 256
              }
            }
          },
          "unit_price": {
            "type": "long"
          }
        }
      },
      "url": {
        "type": "text",
        "fields": {
          "keyword": {
            "type": "keyword",
            "ignore_above": 256
          }
        }
      }
    }
  }
}`

// 存储爬取项
func ItemSaver(index string) (chan engine.Item, error) {
	// 创建elastic客户端
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		return nil, err
	}

	// 数据存储通道
	saver := make(chan engine.Item)

	// 异步接收数据并保存
	go func(saver chan engine.Item, client *elastic.Client, index string) {
		var itemNum int = 1
		for {
			// 接收数据
			var item = <-saver
			bytes, _ := json.Marshal(item)
			log.Printf("Item saver: got item #%d: %+v\n", itemNum, string(bytes))
			itemNum++

			// 存储数据到 Elasticsearch
			err := Save(item, client, index)
			if err != nil {
				bytes, _ := json.Marshal(item)
				log.Printf("Item saver: error saving item %v: %v\n", err, string(bytes))
			}
		}
	}(saver, client, index)
	return saver, nil
}

// 存储数据到 Elasticsearch
func Save(item engine.Item, client *elastic.Client, index string) error {
	// 获取非空上下文
	ctx := context.Background()

	// 转为JSON
	bytes, err := json.Marshal(item)
	if err != nil {
		return err
	}

	// 检测索引库是否存在，不存在则创建
	exists, err := client.IndexExists(index).Do(ctx)
	if err != nil {
		return err
	}
	if !exists {
		_, err := client.CreateIndex(index).BodyString(mappings).Do(ctx)
		if err != nil {
			return err
		}
	}

	// 写入记录到elastic
	indexService := client.Index().Index(index).BodyString(string(bytes))
	if item.Id != "" {
		indexService.Id(item.Id)
	}
	_, err = indexService.Do(ctx)
	if err != nil {
		return err
	}
	return nil
}
