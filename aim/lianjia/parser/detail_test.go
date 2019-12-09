package parser

import (
	"crawler/model"
	"encoding/json"
	"io/ioutil"
	"testing"
)

func TestParseDetail(t *testing.T) {
	c, err := ioutil.ReadFile("detail_test_data.html")
	if err != nil {
		t.Fatalf("fetcher.Fetch error: %s\n", err)
	}
	parseResult := ParseDetail(c, "房屋户型方正，布局合理采光好，布局合理，视野好", "安庆")
	if len(parseResult.Items) != 1 {
		t.Errorf("Items should contain 1 element; but was %v\n", parseResult.Items)
	}

	expected := `{"Title":"房屋户型方正，布局合理采光好，布局合理，视野好","Price":81,"UnitPrice":8708,"CityName":"安庆","CommunityName":"天下名筑","AreaName":"迎江区-老峰镇","BaseInfo":{"Base":{"产权年限":"产权年限70年","套内面积":"套内面积暂无数据","建筑类型":"建筑类型暂无数据","建筑结构":"建筑结构未知结构","建筑面积":"建筑面积93.02㎡","户型结构":"户型结构平层","房屋户型":"房屋户型3室2厅1厨1卫","房屋朝向":"房屋朝向南 北","所在楼层":"所在楼层中楼层 (共34层)","梯户比例":"梯户比例两梯四户","装修情况":"装修情况毛坯","配备电梯":"配备电梯有"},"Transaction":{"上次交易":"2017-10-02","交易权属":"商品房","产权所属":"共有","房屋年限":"满两年","房屋用途":"普通住宅","房本备件":"已上传房本照片","抵押信息":"无抵押","挂牌时间":"2019-11-23"}},"Images":[{"https://image1.ljcdn.com/340800-inspection/apc_V1FfW36ed_1.jpg.710x400.jpg":"客厅"},{"https://image1.ljcdn.com/340800-inspection/pc0_lhYYz4QIc_1.jpg.710x400.jpg":"餐厅"},{"https://image1.ljcdn.com/x-se/hdic-frame/standard_73ff5e79-37b0-4896-a920-2e53b3497225.png.533x400.jpg":"户型图"},{"https://image1.ljcdn.com/340800-inspection/apc_9auPEWvSy_1.jpg.710x400.jpg":"卧室A"},{"https://image1.ljcdn.com/340800-inspection/apc_4jjKD8cEJ_1.jpg.710x400.jpg":"卧室B"},{"https://image1.ljcdn.com/340800-inspection/apc_ehYkc7jSS_1.jpg.710x400.jpg":"卧室C"},{"https://image1.ljcdn.com/340800-inspection/pc0_3VxQ4vdSu_1.jpg.710x400.jpg":"厨房"},{"https://image1.ljcdn.com/340800-inspection/pc0_OGizLBXNm_1.jpg.710x400.jpg":"卫生间"}],"Characteristics":[{"房源标签":"VR房源|随时看房"},{"税费解析":"此房屋不满2年，需要缴纳5.3%的增值税和1%的个人所得税，契税首套90平米以下1%，首套90平米以上1.5%，二套90平米以下1%，90平米以上2%。"},{"户型介绍":"此房屋为三室两厅户型，中间户型，客厅朝南，主卧朝南，采光好，布局合理"},{"核心卖点":"房屋户型方正，布局合理采光好，日照时间长，布局合理，视野好"}],"DoorModelBetweenPoints":{"DoorModelFigure":"https://image1.ljcdn.com/x-se/hdic-frame/standard_73ff5e79-37b0-4896-a920-2e53b3497225.png.960x640.jpg","RoomData":[{"客厅":["17.01平米","南","普通窗"]},{"餐厅":["6.03平米","无","无窗"]},{"卧室A":["5.6平米","北","普通窗"]},{"卧室B":["12.34平米","无","未知窗户类型"]},{"卧室C":["9.9平米","无","未知窗户类型"]},{"厨房":["6.93平米","东","普通窗"]},{"卫生间":["4.31平米","北","普通窗"]},{"阳台A":["2.68平米","北","普通窗"]},{"阳台B":["4.42平米","南","普通窗"]}]}}`

	ershoufang, ok := parseResult.Items[0].(*model.Ershoufang)
	if !ok {
		t.Errorf("incorrect item value type: %+v\n", parseResult.Items[0])
	}

	bytes, err := json.Marshal(ershoufang)
	if err != nil {
		panic(err)
	}

	if expected != string(bytes) {
		t.Errorf("expected %+v, \n but was %+v\n", expected, string(bytes))
	}
}
