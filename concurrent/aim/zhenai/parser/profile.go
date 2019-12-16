package parser

import (
	"bytes"
	"crawler/concurrent/engine"
	"crawler/concurrent/model"
	"regexp"
	"strconv"

	"github.com/PuerkitoBio/goquery"
	"github.com/sirupsen/logrus"
)

var (
	ageRe             = regexp.MustCompile(`(\d+)岁`)
	heightRe          = regexp.MustCompile(`(\d+)[cmCM]{2}`)
	weightRe          = regexp.MustCompile(`(\d+)[kgKG]{2}`)
	workingLocationRe = regexp.MustCompile(`(工作地[^<]+)`)
	incomeRe          = regexp.MustCompile(`(月收入[^<]+)`)
	marriageRe        = regexp.MustCompile(`(离异|离婚|未婚|丧偶)`)
	educationRe       = regexp.MustCompile(`(高中及以下|中专|大专|大学本科|硕士|博士)`)
	occupationRe      = regexp.MustCompile(`(广告客户经理|广告客户专员|广告设计经理|广告设计专员|广告策划|市场营销经理|市场营销专员|市场策划|市场调研与分析|市场拓展|公关经理|公关专员|媒介经理|媒介专员|品牌经理|品牌专员|生物工程|药品生产|临床研究|医疗器械|医药代表|化工工程师|投资|保险|金融|银行|证券|律师|律师助理|法务经理|法务专员|知识产权专员|销售|客户服务|计算机/互联网|通信/电子|生产/制造|物流/仓储|商贸/采购|人事/行政|高级管理|广告/市场|传媒/艺术|生物/制药|医疗/护理|金融/银行/保险|建筑/房地产|咨询/顾问|法律|财会/审计|教育/科研|服务业|交通运输|政府机构|军人/警察|农林牧渔|自由职业|在校学生|待业|销售总监|销售经理|销售主管|销售专员|渠道/分销管理|渠道/分销专员|经销商|客户经理|客户代表|商务经理|商务专员|采购经理|采购专员|外贸经理|外贸专员|业务跟单|报关员|建筑师|工程师|规划师|景观设计|房地产策划|房地产交易|物业管理|教授|讲师/助教|中学教师|小学教师|幼师|教务管理人员|职业技术教师|培训师|科研管理人员|科研人员|公务员|客服经理|客服主管|客服专员|客服协调|客服技术支持|IT技术总监|IT技术经理|IT工程师|系统管理员|测试专员|运营管理|网页设计|网站编辑|网站产品经理|医疗管理|医生|心理医生|药剂师|护士|兽医|专业顾问|咨询经理|咨询师|培训师|通信技术|电子技术|物流经理|物流主管|物流专员|仓库经理|仓库管理员|货运代理|集装箱业务|海关事物管理|报单员|快递员|主编|编辑|作家|撰稿人|文案策划|出版发行|导演|记者|主持人|演员|模特|经纪人|摄影师|影视后期制作|设计师|画家|音乐家|舞蹈|财务总监|财务经理|财务主管|会计|注册会计师|审计师|税务经理|税务专员|成本经理|餐饮管理|厨师|餐厅服务员|酒店管理|大堂经理|酒店服务员|导游|美容师|健身教练|商场经理|零售店店长|店员|保安经理|保安人员|家政服务|飞行员|空乘人员|地勤人员|列车司机|乘务员|船长|船员|司机|工厂经理|工程师|项目主管|营运经理|营运主管|车间主任|物料管理|生产领班|操作工人|安全管理|人事总监|人事经理|人事主管|人事专员|招聘经理|招聘专员|培训经理|培训专员|秘书|文员|后勤|总经理|副总经理|合伙人|总监|经理|总裁助理)`)
	hometownRe        = regexp.MustCompile(`(籍贯[^<]+)`)
	nationRe          = regexp.MustCompile(`([^>]+族)`)
	constellationRe   = regexp.MustCompile(`([^>]+\d+.\d+-\d+.\d+[^<]+)`)
	houseRe           = regexp.MustCompile(`(和家人同住|已购房|租房|打算婚后购房|住在单位宿舍)`)
	carRe             = regexp.MustCompile(`(已买车|未买车)`)
	drinkingRe        = regexp.MustCompile(`(不喝酒|稍微喝一点酒|酒喝得很多|社交场合会喝酒)`)
	smokingRe         = regexp.MustCompile(`(不吸烟|稍微抽一点烟|烟抽得很多|社交场合会抽烟)`)
)

// ParseProfile解析HTTP响应内容用户详情页
func ParseProfile(contents []byte, name, gender string) engine.ParseResult {
	parseResult := engine.ParseResult{}
	contents, err := ParseProfileForInfo(contents)
	if err != nil {
		panic(err)
	}
	profile := model.Profile{}
	profile.Name = name
	profile.Gender = gender
	profile.Age = ParseProfileInfoFieldToInt(contents, ageRe)
	profile.Height = ParseProfileInfoFieldToInt(contents, heightRe)
	profile.Weight = ParseProfileInfoFieldToInt(contents, weightRe)
	profile.WorkingLocation = ParseProfileInfoFieldToStr(contents, workingLocationRe)
	profile.Income = ParseProfileInfoFieldToStr(contents, incomeRe)
	profile.Marriage = ParseProfileInfoFieldToStr(contents, marriageRe)
	profile.Education = ParseProfileInfoFieldToStr(contents, educationRe)
	profile.Occupation = ParseProfileInfoFieldToStr(contents, occupationRe)
	profile.Hometown = ParseProfileInfoFieldToStr(contents, hometownRe)
	profile.Nation = ParseProfileInfoFieldToStr(contents, nationRe)
	profile.Constellation = ParseProfileInfoFieldToStr(contents, constellationRe)
	profile.House = ParseProfileInfoFieldToStr(contents, houseRe)
	profile.Car = ParseProfileInfoFieldToStr(contents, carRe)
	profile.Smoking = ParseProfileInfoFieldToStr(contents, smokingRe)
	profile.Drinking = ParseProfileInfoFieldToStr(contents, drinkingRe)
	parseResult.Items = append(parseResult.Items, profile)
	return parseResult
}

// ParseProfileInfoFieldToInt 解析用户个人资料信息字段到字符串
func ParseProfileInfoFieldToStr(contents []byte, re *regexp.Regexp) string {
	find := re.Find(contents)
	if string(find) != "" {
		return string(find)
	}
	return ""
}

// ParseProfileInfoFieldToInt 解析用户个人资料信息字段到整型
func ParseProfileInfoFieldToInt(contents []byte, re *regexp.Regexp) int {
	ageByte := re.FindSubmatch(contents)
	if len(ageByte) != 0 {
		age, err := strconv.Atoi(string(ageByte[1]))
		if err != nil {
			logrus.Errorf("type string to convert integer fault")
		}
		return age
	}
	return 0
}

// ParseProfileForInfo获取用户个人资料信息部分HTML
func ParseProfileForInfo(contents []byte) (info []byte, err error) {
	reader := bytes.NewReader(contents)
	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		return nil, err
	}
	var findStr string
	doc.Find(".CONTAINER .f-fl div .m-userInfoDetail").Each(func(index int, selection *goquery.Selection) {
		findStr, err = selection.Children().Next().Next().Next().Html()
	})
	info = []byte(findStr)
	return info, nil
}
