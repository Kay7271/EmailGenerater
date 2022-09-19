package conf

import "converter/methods"

//HY表示-
//Dot表示.

type Method func(pinyin []string) string

var (
	Methods = []string{
		"张三丰 --> zhangsanfeng",
		"张三丰 --> zhangsf",
		"张三丰 --> zsf",
		"张三丰 --> sanfengzhang",
		"张三丰 --> sfzhang",
		"张三丰 --> sanfeng.zhang",
		"张三丰 --> zhang.sanfeng",
		"张三丰 --> sf.zhang",
		"张三丰 --> zhang.sf",
		"张三丰 --> s.fzhang",
		"张三丰 --> san-feng.zhang",
		"张三丰 --> sanfeng_zhang",
		"张三丰 --> sf_zhang",
		"张三丰 --> zhang_sanfeng",
		"张三丰 --> zhang_sf",
		"张三丰 --> san_feng_zhang",
		"张三丰 --> zhang_san_feng",
		"张三丰 --> zhang-sanfeng",
		"张三丰 --> sanfeng-zhang",
	}
	MethodMap = map[string]Method{
		"张三丰 --> zhangsanfeng":   methods.ZhangSanFeng,
		"张三丰 --> zhangsf":        methods.ZhangSF,
		"张三丰 --> zsf":            methods.ZSF,
		"张三丰 --> sanfengzhang":   methods.SanFengZhang,
		"张三丰 --> sfzhang":        methods.SFZhang,
		"张三丰 --> sanfeng.zhang":  methods.SanFengDotZhang,
		"张三丰 --> zhang.sanfeng":  methods.ZhangDotSanFeng,
		"张三丰 --> sf.zhang":       methods.SFDotZhang,
		"张三丰 --> zhang.sf":       methods.ZhangDotSF,
		"张三丰 --> s.fzhang":       methods.SDotFZhang,
		"张三丰 --> san-feng.zhang": methods.SanHYFengDotZhang,
		"张三丰 --> sanfeng_zhang":  methods.SanFeng_Zhang,
		"张三丰 --> sf_zhang":       methods.SF_Zhang,
		"张三丰 --> zhang_sanfeng":  methods.Zhang_SanFeng,
		"张三丰 --> zhang_sf":       methods.Zhang_SF,
		"张三丰 --> san_feng_zhang": methods.San_Feng_Zhang,
		"张三丰 --> zhang_san_feng": methods.Zhang_San_Feng,
		"张三丰 --> zhang-sanfeng":  methods.ZhangHYSanFeng,
		"张三丰 --> sanfeng-zhang":  methods.SanFengHYZhang,
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
