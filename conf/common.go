package conf

const (
	SanFengZhang = "张三丰 --> sanfeng.zhang"
	ZhangSanFeng = "张三丰 --> zhangsanfeng"
	ZhangSF      = "张三丰 --> zhangsf"
	ZSF          = "张三丰 --> zsf"
)

var (
	Methods = []string{
		SanFengZhang,
		ZhangSanFeng,
		ZhangSF,
		ZSF,
	}
	MethodMap = map[string]string{
		"张三丰 --> sanfeng.zhang": "SanFengZhang",
		"张三丰 --> zhangsanfeng":  "ZhangSanFeng",
		"张三丰 --> zhangsf":       "ZhangSF",
		"张三丰 --> zsf":           "ZSF",
	}
)
var (
	DuoYinZiMap = map[string]string{
		"覃": "qin",
		"种": "chong",
		"单": "shan",
		"仇": "qiu",
		"区": "ou",
		"查": "zha",
		"朴": "piao",
		"解": "xie",
		"黑": "he",
		"盖": "ge",
		"曾": "zeng",
		"缪": "miao",
		"乐": "yue",
		"员": "yun",
		"谌": "shen",
		"翟": "zhai",
	}
	Suffixs = []string{
		"woqutech.com",
		"squids.cn",
		"163.com",
		"qq.com",
		"sina.cn",
	}
)
