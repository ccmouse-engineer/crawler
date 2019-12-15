package model

// Ershoufang二手房
type Ershoufang struct {
	Title                  string                           `json:"title"`                     // 房源标题
	Price                  float64                          `json:"price"`                     // 总价
	UnitPrice              float64                          `json:"unit_price"`                // 每平单价
	CityName               string                           `json:"city_name"`                 // 城市名称
	CommunityName          string                           `json:"community_name"`            // 小区名称
	AreaName               string                           `json:"area_name"`                 // 所在区域
	BaseInfo               ErshoufangBaseInfo               `json:"base_info"`                 // 二手房基本信息
	Images                 []map[string]string              `json:"images"`                    // 二手房房源照片
	Characteristics        []map[string]string              `json:"characteristics"`           // 二手房房源特色
	DoorModelBetweenPoints ErshoufangDoorModelBetweenPoints `json:"door_model_between_points"` // 二手房户型分间
}

// Ershoufang二手房基本信息
type ErshoufangBaseInfo struct {
	Base        map[string]string `json:"base"`        // 基本属性
	Transaction map[string]string `json:"transaction"` // 交易属性
}

// Ershoufang二手房户型分间
type ErshoufangDoorModelBetweenPoints struct {
	DoorModelFigure string                `json:"door_model_figure"` // 户型图
	RoomData        []map[string][]string `json:"room_data"`         // 户型房间数据
}
