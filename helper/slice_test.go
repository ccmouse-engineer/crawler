package helper

import (
	"fmt"
	"testing"
)

func TestSliceDeduplication(t *testing.T) {
	expected := []string{
		"VR房源",
		"有轻轨 公交25路 137 路 22路 了 交通便利 临火车站",
		"有胜利公园 轻轨站 联合书城 欧亚超市 临火车站小区安静好 楼层好",
		"2002年建成 南北格局 三室二厅 有大阳台 卫生间 格局大 大厅大",
		"此房满五 南北通透 两厅 开间四米五 临欧亚卖场 胜利公园",
	}

	slices := []string{
		"VR房源",
		"有轻轨 公交25路 137 路 22路 了 交通便利 临火车站",
		"有轻轨 公交25路 137 路 22路 了 交通便利 临火车站",
		"有胜利公园 轻轨站 联合书城 欧亚超市 临火车站小区安静好 楼层好",
		"有胜利公园 轻轨站 联合书城 欧亚超市 临火车站小区安静好 楼层好",
		"2002年建成 南北格局 三室二厅 有大阳台 卫生间 格局大 大厅大",
		"2002年建成 南北格局 三室二厅 有大阳台 卫生间 格局大 大厅大",
		"此房满五 南北通透 两厅 开间四米五 临欧亚卖场 胜利公园",
	}

	actual := SliceDeduplication(slices)
	if len(actual) != len(expected) {
		fmt.Printf("expected len: %d, got len: %d\n", len(actual), len(expected))
	}
}
