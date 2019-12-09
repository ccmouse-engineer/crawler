package model

// Ershoufang二手房
type Ershoufang struct {
	Title                  string                           // 房源标题
	Price                  float64                          // 总价
	UnitPrice              float64                          // 每平单价
	CityName               string                           // 城市名称
	CommunityName          string                           // 小区名称
	AreaName               string                           // 所在区域
	BaseInfo               ErshoufangBaseInfo               // 二手房基本信息
	Images                 []map[string]string              // 二手房房源照片
	Characteristics        []map[string]string              // 二手房房源特色
	DoorModelBetweenPoints ErshoufangDoorModelBetweenPoints // 二手房户型分间

}

// Ershoufang二手房基本信息
type ErshoufangBaseInfo struct {
	Base        map[string]string // 基本属性
	Transaction map[string]string // 交易属性
}

// Ershoufang二手房户型分间
type ErshoufangDoorModelBetweenPoints struct {
	DoorModelFigure string                // 户型图
	RoomData        []map[string][]string // 户型房间数据
}
