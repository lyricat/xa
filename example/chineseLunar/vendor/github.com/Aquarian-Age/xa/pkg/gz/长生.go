package gz

var (
	jiaChangShengMap  = map[string]string{"丑": "冠带", "亥": "长生", "午": "死", "卯": "帝旺", "子": "沐浴", "寅": "临官", "巳": "病", "戌": "养", "未": "墓", "申": "绝", "辰": "衰", "酉": "胎"}
	yiChangShengMap   = map[string]string{"丑": "衰", "亥": "死", "午": "长生", "卯": "临官", "子": "病", "寅": "帝旺", "巳": "沐浴", "戌": "墓", "未": "养", "申": "胎", "辰": "冠带", "酉": "绝"}
	bingChangShengMap = map[string]string{"丑": "养", "亥": "绝", "午": "帝旺", "卯": "沐浴", "子": "胎", "寅": "长生", "巳": "临官", "戌": "墓", "未": "衰", "申": "病", "辰": "冠带", "酉": "死"}
	dingChangShengMap = map[string]string{"丑": "墓", "亥": "胎", "午": "临官", "卯": "病", "子": "绝", "寅": "死", "巳": "帝旺", "戌": "养", "未": "冠带", "申": "沐浴", "辰": "衰", "酉": "长生"}
	wuChangShengMap   = map[string]string{"丑": "养", "亥": "绝", "午": "帝旺", "卯": "沐浴", "子": "胎", "寅": "长生", "巳": "临官", "戌": "墓", "未": "衰", "申": "病", "辰": "冠带", "酉": "死"}
	jiChangShengMap   = map[string]string{"丑": "墓", "亥": "胎", "午": "临官", "卯": "病", "子": "绝", "寅": "死", "巳": "帝旺", "戌": "养", "未": "冠带", "申": "沐浴", "辰": "衰", "酉": "长生"}
	gengChangShengMap = map[string]string{"丑": "墓", "亥": "病", "午": "沐浴", "卯": "胎", "子": "死", "寅": "绝", "巳": "长生", "戌": "衰", "未": "冠带", "申": "临官", "辰": "养", "酉": "帝旺"}
	xinChangShengMap  = map[string]string{"丑": "养", "亥": "沐浴", "午": "病", "卯": "绝", "子": "长生", "寅": "胎", "巳": "死", "戌": "冠带", "未": "衰", "申": "帝旺", "辰": "墓", "酉": "临官"}
	renChangShengMap  = map[string]string{"丑": "衰", "亥": "临官", "午": "胎", "卯": "死", "子": "帝旺", "寅": "病", "巳": "绝", "戌": "冠带", "未": "养", "申": "长生", "辰": "墓", "酉": "沐浴"}
	guiChangShengMap  = map[string]string{"丑": "冠带", "亥": "帝旺", "午": "绝", "卯": "长生", "子": "临官", "寅": "沐浴", "巳": "胎", "戌": "衰", "未": "墓", "申": "死", "辰": "养", "酉": "病"}
)

// GanZhiChangSheng 干支长生 返回干与支在十二长生的关系 比如甲子 甲在子位沐浴 则甲临沐浴
func GanZhiChangSheng(xgz string) string {
	gs := xgz[:3]
	zs := xgz[3:6]
	cs := ChangSheng(gs, zs)
	return gs + "临" + cs
}

// ChangSheng 十干长生 传入干 支 返回干在支的十二长生名称
func ChangSheng(gan, zhi string) string {
	var s string
	switch gan {
	case "甲":
		s = jiaChangShengMap[zhi]
	case "乙":
		s = yiChangShengMap[zhi]
	case "丙":
		s = bingChangShengMap[zhi]
	case "丁":
		s = dingChangShengMap[zhi]
	case "戊":
		s = wuChangShengMap[zhi]
	case "己":
		s = jiChangShengMap[zhi]
	case "庚":
		s = gengChangShengMap[zhi]
	case "辛":
		s = xinChangShengMap[zhi]
	case "壬":
		s = renChangShengMap[zhi]
	case "癸":
		s = guiChangShengMap[zhi]
	}
	return s
}
