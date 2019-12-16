package model

type Car struct {
	Name             string            // 车型名称
	ImageURL         string            // 车型图片
	VendorGuidePrice string            // 厂商指导价
	Base             map[string]string // 基本信息
	Body             map[string]string // 车身信息
}
