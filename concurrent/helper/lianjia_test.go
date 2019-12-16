package helper

import (
	"encoding/json"
	"io/ioutil"
	"regexp"
	"testing"
)

// cityListRe城市列表匹配正则
var (
	cityListRe = regexp.MustCompile(`<a\s+href="(https://\w+.lianjia.com/)">([^<]+)</a>`)
	nextRe     = regexp.MustCompile(`<div\s+class="page-box\s+house-lst-page-box"\s+comp-module='page'\s+page-url="/ershoufang/pg\{page\}"page-data='\{"totalPage":(\d+),"curPage":(\d+)\}'></div>`)
)

func TestGetFullDomainURL(t *testing.T) {
	expected := "https://aq.lianjia.com/ershoufang/"
	url := GetFullDomainURL("https://aq.lianjia.com/", "ershoufang")
	if url != expected {
		t.Errorf("result should %s url, but has %s\n", expected, url)
	}
}

func TestLinkDeduplication(t *testing.T) {
	c, err := ioutil.ReadFile("../aim/lianjia/parser/citylist_test_data.txt")
	if err != nil {
		panic(err)
	}
	const resultSize = 105
	expectedStrJSON := `[{"https://aq.lianjia.com/":"安庆"},{"https://baoji.lianjia.com/":"宝鸡"},{"https://bd.lianjia.com/":"保定"},{"https://bh.lianjia.com/":"北海"},{"https://bj.lianjia.com/":"北京"},{"https://cc.lianjia.com/":"长春"},{"https://cd.lianjia.com/":"成都"},{"https://changde.lianjia.com/":"常德"},{"https://changzhou.lianjia.com/":"常州"},{"https://cq.lianjia.com/":"重庆"},{"https://cs.lianjia.com/":"长沙"},{"https://dali.lianjia.com/":"大理"},{"https://dazhou.lianjia.com/":"达州"},{"https://dd.lianjia.com/":"丹东"},{"https://dg.lianjia.com/":"东莞"},{"https://dl.lianjia.com/":"大连"},{"https://ez.lianjia.com/":"鄂州"},{"https://fcg.lianjia.com/":"防城港"},{"https://fs.lianjia.com/":"佛山"},{"https://fz.lianjia.com/":"福州"},{"https://ganzhou.lianjia.com/":"赣州"},{"https://gl.lianjia.com/":"桂林"},{"https://gy.lianjia.com/":"贵阳"},{"https://gz.lianjia.com/":"广州"},{"https://ha.lianjia.com/":"淮安"},{"https://haimen.lianjia.com/":"海门"},{"https://hanzhong.lianjia.com/":"汉中"},{"https://hf.lianjia.com/":"合肥"},{"https://hhht.lianjia.com/":"呼和浩特"},{"https://hk.lianjia.com/":"海口"},{"https://hrb.lianjia.com/":"哈尔滨"},{"https://huangshi.lianjia.com/":"黄石"},{"https://hui.lianjia.com/":"惠州"},{"https://huzhou.lianjia.com/":"湖州"},{"https://hz.lianjia.com/":"杭州"},{"https://jh.lianjia.com/":"金华"},{"https://jian.lianjia.com/":"吉安"},{"https://jiangmen.lianjia.com/":"江门"},{"https://jiujiang.lianjia.com/":"九江"},{"https://jl.lianjia.com/":"吉林"},{"https://jn.lianjia.com/":"济南"},{"https://jx.lianjia.com/":"嘉兴"},{"https://jy.lianjia.com/":"江阴"},{"https://jz.lianjia.com/":"晋中"},{"https://kf.lianjia.com/":"开封"},{"https://km.lianjia.com/":"昆明"},{"https://ks.lianjia.com/":"昆山"},{"https://lf.lianjia.com/":"廊坊"},{"https://liangshan.lianjia.com/":"凉山"},{"https://linyi.lianjia.com/":"临沂"},{"https://liuzhou.lianjia.com/":"柳州"},{"https://luoyang.lianjia.com/":"洛阳"},{"https://lz.lianjia.com/":"兰州"},{"https://mas.lianjia.com/":"马鞍山"},{"https://mianyang.lianjia.com/":"绵阳"},{"https://nanchong.lianjia.com/":"南充"},{"https://nb.lianjia.com/":"宁波"},{"https://nc.lianjia.com/":"南昌"},{"https://nj.lianjia.com/":"南京"},{"https://nn.lianjia.com/":"南宁"},{"https://nt.lianjia.com/":"南通"},{"https://qd.lianjia.com/":"青岛"},{"https://quanzhou.lianjia.com/":"泉州"},{"https://qy.lianjia.com/":"清远"},{"https://san.lianjia.com/":"三亚"},{"https://sh.lianjia.com/":"上海"},{"https://sjz.lianjia.com/":"石家庄"},{"https://sr.lianjia.com/":"上饶"},{"https://su.lianjia.com/":"苏州"},{"https://sx.lianjia.com/":"绍兴"},{"https://sy.lianjia.com/":"沈阳"},{"https://sz.lianjia.com/":"深圳"},{"https://ta.lianjia.com/":"泰安"},{"https://taizhou.lianjia.com/":"台州"},{"https://tj.lianjia.com/":"天津"},{"https://ts.lianjia.com/":"唐山"},{"https://ty.lianjia.com/":"太原"},{"https://weihai.lianjia.com/":"威海"},{"https://wf.lianjia.com/":"潍坊"},{"https://wh.lianjia.com/":"武汉"},{"https://wuhu.lianjia.com/":"芜湖"},{"https://wx.lianjia.com/":"无锡"},{"https://wz.lianjia.com/":"温州"},{"https://xa.lianjia.com/":"西安"},{"https://xc.lianjia.com/":"许昌"},{"https://xianyang.lianjia.com/":"咸阳"},{"https://xinxiang.lianjia.com/":"新乡"},{"https://xm.lianjia.com/":"厦门"},{"https://xy.lianjia.com/":"襄阳"},{"https://xz.lianjia.com/":"徐州"},{"https://yc.lianjia.com/":"盐城"},{"https://yichang.lianjia.com/":"宜昌"},{"https://yinchuan.lianjia.com/":"银川"},{"https://yt.lianjia.com/":"烟台"},{"https://yw.lianjia.com/":"义乌"},{"https://yy.lianjia.com/":"岳阳"},{"https://zb.lianjia.com/":"淄博"},{"https://zh.lianjia.com/":"珠海"},{"https://zhangzhou.lianjia.com/":"漳州"},{"https://zhanjiang.lianjia.com/":"湛江"},{"https://zhuzhou.lianjia.com/":"株洲"},{"https://zj.lianjia.com/":"镇江"},{"https://zjk.lianjia.com/":"张家口"},{"https://zs.lianjia.com/":"中山"},{"https://zz.lianjia.com/":"郑州"}]`
	matchs := cityListRe.FindAllSubmatch(c, -1)
	parseResult := LinkDeduplication(matchs)
	if len(parseResult) != resultSize {
		t.Errorf("result should have %d items; but had %d\n", resultSize, len(parseResult))
	}
	if strJSON, _ := json.Marshal(parseResult); string(strJSON) != expectedStrJSON {
		t.Errorf("result should have %s; but had %s\n", expectedStrJSON, string(strJSON))
	}
}

func TestMustPrice(t *testing.T) {
	var expected1 float64 = 83.3
	var expected2 float64 = 86

	if price := MustPrice("83.3"); price != expected1 {
		t.Errorf("result should %v, but has %v", expected1, price)
	}

	if price := MustPrice("86"); price != expected2 {
		t.Errorf("result should %v, but has %v", expected2, price)
	}
}

func TestGenerateMultiPages(t *testing.T) {
	c, err := ioutil.ReadFile("../aim/lianjia/parser/city_test_data.txt")
	if err != nil {
		panic(err)
	}
	const resultSize = 99
	expected := `["https://aq.lianjia.com/ershoufang/pg2","https://aq.lianjia.com/ershoufang/pg3","https://aq.lianjia.com/ershoufang/pg4","https://aq.lianjia.com/ershoufang/pg5","https://aq.lianjia.com/ershoufang/pg6","https://aq.lianjia.com/ershoufang/pg7","https://aq.lianjia.com/ershoufang/pg8","https://aq.lianjia.com/ershoufang/pg9","https://aq.lianjia.com/ershoufang/pg10","https://aq.lianjia.com/ershoufang/pg11","https://aq.lianjia.com/ershoufang/pg12","https://aq.lianjia.com/ershoufang/pg13","https://aq.lianjia.com/ershoufang/pg14","https://aq.lianjia.com/ershoufang/pg15","https://aq.lianjia.com/ershoufang/pg16","https://aq.lianjia.com/ershoufang/pg17","https://aq.lianjia.com/ershoufang/pg18","https://aq.lianjia.com/ershoufang/pg19","https://aq.lianjia.com/ershoufang/pg20","https://aq.lianjia.com/ershoufang/pg21","https://aq.lianjia.com/ershoufang/pg22","https://aq.lianjia.com/ershoufang/pg23","https://aq.lianjia.com/ershoufang/pg24","https://aq.lianjia.com/ershoufang/pg25","https://aq.lianjia.com/ershoufang/pg26","https://aq.lianjia.com/ershoufang/pg27","https://aq.lianjia.com/ershoufang/pg28","https://aq.lianjia.com/ershoufang/pg29","https://aq.lianjia.com/ershoufang/pg30","https://aq.lianjia.com/ershoufang/pg31","https://aq.lianjia.com/ershoufang/pg32","https://aq.lianjia.com/ershoufang/pg33","https://aq.lianjia.com/ershoufang/pg34","https://aq.lianjia.com/ershoufang/pg35","https://aq.lianjia.com/ershoufang/pg36","https://aq.lianjia.com/ershoufang/pg37","https://aq.lianjia.com/ershoufang/pg38","https://aq.lianjia.com/ershoufang/pg39","https://aq.lianjia.com/ershoufang/pg40","https://aq.lianjia.com/ershoufang/pg41","https://aq.lianjia.com/ershoufang/pg42","https://aq.lianjia.com/ershoufang/pg43","https://aq.lianjia.com/ershoufang/pg44","https://aq.lianjia.com/ershoufang/pg45","https://aq.lianjia.com/ershoufang/pg46","https://aq.lianjia.com/ershoufang/pg47","https://aq.lianjia.com/ershoufang/pg48","https://aq.lianjia.com/ershoufang/pg49","https://aq.lianjia.com/ershoufang/pg50","https://aq.lianjia.com/ershoufang/pg51","https://aq.lianjia.com/ershoufang/pg52","https://aq.lianjia.com/ershoufang/pg53","https://aq.lianjia.com/ershoufang/pg54","https://aq.lianjia.com/ershoufang/pg55","https://aq.lianjia.com/ershoufang/pg56","https://aq.lianjia.com/ershoufang/pg57","https://aq.lianjia.com/ershoufang/pg58","https://aq.lianjia.com/ershoufang/pg59","https://aq.lianjia.com/ershoufang/pg60","https://aq.lianjia.com/ershoufang/pg61","https://aq.lianjia.com/ershoufang/pg62","https://aq.lianjia.com/ershoufang/pg63","https://aq.lianjia.com/ershoufang/pg64","https://aq.lianjia.com/ershoufang/pg65","https://aq.lianjia.com/ershoufang/pg66","https://aq.lianjia.com/ershoufang/pg67","https://aq.lianjia.com/ershoufang/pg68","https://aq.lianjia.com/ershoufang/pg69","https://aq.lianjia.com/ershoufang/pg70","https://aq.lianjia.com/ershoufang/pg71","https://aq.lianjia.com/ershoufang/pg72","https://aq.lianjia.com/ershoufang/pg73","https://aq.lianjia.com/ershoufang/pg74","https://aq.lianjia.com/ershoufang/pg75","https://aq.lianjia.com/ershoufang/pg76","https://aq.lianjia.com/ershoufang/pg77","https://aq.lianjia.com/ershoufang/pg78","https://aq.lianjia.com/ershoufang/pg79","https://aq.lianjia.com/ershoufang/pg80","https://aq.lianjia.com/ershoufang/pg81","https://aq.lianjia.com/ershoufang/pg82","https://aq.lianjia.com/ershoufang/pg83","https://aq.lianjia.com/ershoufang/pg84","https://aq.lianjia.com/ershoufang/pg85","https://aq.lianjia.com/ershoufang/pg86","https://aq.lianjia.com/ershoufang/pg87","https://aq.lianjia.com/ershoufang/pg88","https://aq.lianjia.com/ershoufang/pg89","https://aq.lianjia.com/ershoufang/pg90","https://aq.lianjia.com/ershoufang/pg91","https://aq.lianjia.com/ershoufang/pg92","https://aq.lianjia.com/ershoufang/pg93","https://aq.lianjia.com/ershoufang/pg94","https://aq.lianjia.com/ershoufang/pg95","https://aq.lianjia.com/ershoufang/pg96","https://aq.lianjia.com/ershoufang/pg97","https://aq.lianjia.com/ershoufang/pg98","https://aq.lianjia.com/ershoufang/pg99","https://aq.lianjia.com/ershoufang/pg100"]`
	submatch := nextRe.FindSubmatch(c)
	multiPages := GenerateMultiPages(submatch, "https://aq.lianjia.com/ershoufang/")
	if len(multiPages) != resultSize {
		t.Errorf("result should have %d items; but had %d\n", resultSize, len(multiPages))
	}
	if bytes, _ := json.Marshal(multiPages); string(bytes) != expected {
		t.Errorf("result should have %s; but had %s\n", expected, string(bytes))
	}
}
