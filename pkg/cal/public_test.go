package cal

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
	"time"
)

// func TestJQArr_YueJiang(t *testing.T) {
// 	y := 2020
// 	jq := NewYueJiangJQ(YueJiangJQT(y))
// 	yj := jq.YueJiang(time.Now().Local())
// 	fmt.Println(yj)
// 	jqtoday := jq.TodayJQ(y, time.Now().Local())
// 	fmt.Println(jqtoday)
// }

func TestNewGanZhiInfo(t *testing.T) {
	md2033 := []struct {
		y, m, d int
	}{
		{2033, 12, 3},
		{2033, 12, 6},
		{2033, 12, 7},
		{2033, 12, 21},
		{2033, 1, 1},
		{2033, 1, 5},
		{2033, 1, 20},
		{2033, 1, 31},
		{2033, 2, 3},
		{2033, 2, 18},
		{2033, 3, 1},
		{2033, 3, 5},
		{2033, 3, 20},
		{2033, 3, 31},
		{2033, 4, 4},
		{2033, 4, 20},
		{2033, 4, 29},
		{2033, 5, 5},
		{2033, 5, 21},
		{2033, 5, 28},
		{2033, 6, 5},
		{2033, 6, 21},
		{2033, 6, 27},
		{2033, 7, 7},
		{2033, 7, 22},
		{2033, 7, 26},
		{2033, 8, 7},
		{2033, 8, 23},
		{2033, 8, 25},
		{2033, 9, 7},
		{2033, 9, 23},
		{2033, 9, 23},
		{2033, 10, 8},
		{2033, 10, 23},
		{2033, 10, 23},
		{2033, 11, 7},
		{2033, 11, 22},
		{2033, 11, 22},
		{2033, 12, 7},
	}
	want2033 := []struct {
		yg, mg, dg string
	}{
		{"癸丑", "癸亥", "戊子"},
		{"癸丑", "癸亥", "辛卯"},
		{"癸丑", "甲子", "壬辰"},
		{"癸丑", "甲子", "丙午"},
		{"壬子", "壬子", "壬子}"},
		{"壬子", "癸丑", "丙辰"},
		{"壬子", "癸丑", "辛未"},
		{"壬子", "癸丑", "壬午"},
		{"癸丑", "甲寅", "乙酉"},
		{"癸丑", "甲寅", "庚子"},
		{"癸丑", "甲寅", "辛亥"},
		{"癸丑", "乙卯", "乙卯"},
		{"癸丑", "乙卯", "庚午"},
		{"癸丑", "乙卯", "辛巳"},
		{"癸丑", "丙辰", "乙酉"},
		{"癸丑", "丙辰", "辛丑"},
		{"癸丑", "丙辰", "庚戌"},
		{"癸丑", "丁巳", "丙辰"},
		{"癸丑", "丁巳", "壬申"},
		{"癸丑", "丁巳", "己卯"},
		{"癸丑", "戊午", "丁亥"},
		{"癸丑", "戊午", "癸卯"},
		{"癸丑", "戊午", "己酉"},
		{"癸丑", "己未", "己未"},
		{"癸丑", "己未", "甲戌"},
		{"癸丑", "己未", "戊寅"},
		{"癸丑", "庚申", "庚寅"},
		{"癸丑", "庚申", "丙午"},
		{"癸丑", "庚申", "戊申"},
		{"癸丑", "辛酉", "辛酉"},
		{"癸丑", "辛酉", "丁丑"},
		{"癸丑", "辛酉", "丁丑"},
		{"癸丑", "壬戌", "壬辰"},
		{"癸丑", "壬戌", "丁未"},
		{"癸丑", "壬戌", "丁未"},
		{"癸丑", "癸亥", "壬戌"},
		{"癸丑", "癸亥", "丁丑"},
		{"癸丑", "癸亥", "丁丑"},
		{"癸丑", "甲子", "壬辰"},
	}

	for i := 0; i < len(md2033); i++ {
		gz := NewCal(md2033[i].y, md2033[i].m, md2033[i].d, 12)
		yg := gz.YearGZ
		mg := gz.MonthGZ
		dg := gz.DayGZ
		if !strings.EqualFold(yg, want2033[i].yg) &&
			!strings.EqualFold(mg, want2033[i].mg) &&
			!strings.EqualFold(dg, want2033[i].dg) {
			t.Errorf("NewCal(%d %d %d, 12)=%s %s %s want:%s %s %s\n",
				md2033[i].y, md2033[i].m, md2033[i].d, yg, mg, dg, want2033[i].yg, want2033[i].mg, want2033[i].dg)
		}

	}
}

/*
2033年12月3日--癸丑-癸亥-戊子
2033年12月6日--癸丑-癸亥-辛卯
2033年12月7日--癸丑-甲子壬辰
2033年12月21日--癸丑-甲子-丙午
2033年1月1日--壬子-壬子-壬子  特殊日
2033年1月5日--壬子-癸丑-丙辰
2033年1月20日--壬子-癸丑-辛未
2033年1月31日--壬子-癸丑-壬午
2033年2月3日--癸丑-甲寅-乙酉
2033年2月18日--癸丑-甲寅-庚子
2033年3月1日--癸丑-甲寅-辛亥
2033年3月5日--癸丑-乙卯-乙卯
2033年3月20日--癸丑-乙卯-庚午
2033年3月31日--癸丑-乙卯-辛巳
2033年4月4日--癸丑-丙辰-乙酉
2033年4月20日--癸丑-丙辰-辛丑
2033年4月29日--癸丑-丙辰-庚戌
2033年5月5日--癸丑-丁巳-丙辰
2033年5月21日--癸丑-丁巳-壬申
2033年5月28日--癸丑-丁巳-己卯
2033年6月5日--癸丑-戊午-丁亥
2033年6月21日--癸丑-戊午-癸卯
2033年6月27日--癸丑-戊午-己酉
2033年7月7日--癸丑-己未-己未
2033年7月22日--癸丑-己未-甲戌
2033年7月26日--癸丑-己未-戊寅
2033年8月7日--癸丑-庚申-庚寅
2033年8月23日--癸丑-庚申-丙午
2033年8月25日--癸丑-庚申-戊申
2033年9月7日--癸丑-辛酉-辛酉
2033年9月23日--癸丑-辛酉-丁丑
2033年9月23日--癸丑-辛酉-丁丑
2033年10月8日--癸丑-壬戌-壬辰
2033年10月23日--癸丑-壬戌-丁未
2033年10月23日--癸丑-壬戌-丁未
2033年11月7日--癸丑-癸亥-壬戌
2033年11月22日--癸丑-癸亥-丁丑
2033年11月22日--癸丑-癸亥-丁丑
2033年12月7日--癸丑-甲子-壬辰
*/
/*
2021年12月15日--辛丑-庚子-丁酉
2021年12月21日--辛丑-庚子-癸卯

2021年1月5日--庚子-己丑-癸丑
2021年1月13日--庚子-己丑-辛酉
2021年1月20日--庚子-己丑-戊辰

2021年2月3日--辛丑-庚寅-壬午
2021年2月12日--辛丑-庚寅-辛卯
2021年2月18日--辛丑-庚寅-丁酉

2021年3月5日--辛丑-辛卯-壬子
2021年3月13日--辛丑-辛卯-庚申
2021年3月20日--辛丑-辛卯-丁卯

2021年4月4日--辛丑-壬辰-壬午
2021年4月12日--辛丑-壬辰-庚寅
2021年4月20日--辛丑-壬辰-戊戌

2021年5月5日--辛丑-癸巳-癸丑
2021年5月12日--辛丑-癸巳-庚申
2021年5月21日--辛丑-癸巳-己巳

2021年6月5日--辛丑-甲午-甲申
2021年6月10日--辛丑-甲午-己丑
2021年6月21日--辛丑-甲午-庚子

2021年7月7日--辛丑-乙未-丙辰
2021年7月10日--辛丑-乙未-己未
2021年7月22日--辛丑-乙未-辛未

2021年8月7日--辛丑-丙申-丁亥
2021年8月8日--辛丑-丙申-戊子
2021年8月23日--辛丑-丙申-癸卯

2021年9月7日--辛丑-丁酉-戊午
2021年9月7日--辛丑-丁酉-戊午
2021年9月23日--辛丑-丁酉-甲戌
2021年10月6日--辛丑-丁酉-丁亥

2021年10月8日--辛丑-戊戌-己丑
2021年10月23日--辛丑-戊戌-甲辰
2021年11月5日--辛丑-戊戌-丁巳

2021年11月7日--辛丑-己亥-己未
2021年11月22日--辛丑-己亥-甲戌

2021年12月4日--辛丑-己亥-丙戌
2021年12月7日--辛丑-庚子-己丑

*/
/*
2020年11月26日--庚子-丁亥-癸酉

2020年12月7日--庚子-戊子-甲申
2020年12月22日--庚子-戊子-己亥
2020年12月26日--庚子-戊子-癸卯

2020年1月6日--己亥-丁丑-戊申
2020年1月20日--己亥-丁丑-壬戌
2020年1月25日--己亥-丁丑-丁卯

2020年2月4日--庚子-戊寅-丁丑
2020年2月19日--庚子-戊寅-壬辰
2020年2月23日--庚子-戊寅-丙申

2020年3月5日--庚子-己卯-丁未
2020年3月20日--庚子-己卯-壬戌
2020年3月24日--庚子-己卯-丙寅

2020年4月4日--庚子-庚辰-丁丑
2020年4月19日--庚子-庚辰-壬辰
2020年4月23日--庚子-庚辰-丙申

2020年5月5日--庚子-辛巳-戊申
2020年5月20日--庚子-辛巳-癸亥
2020年5月23日--庚子-辛巳-丙寅

2020年6月5日--庚子-壬午-己卯
2020年6月21日--庚子-壬午-乙未
2020年6月21日--庚子-壬午-乙未

2020年7月6日--庚子-癸未-庚戌
2020年7月21日--庚子-癸未-乙丑
2020年7月22日--庚子-癸未-丙寅

2020年8月7日--庚子-甲申-壬午
2020年8月19日--庚子-甲申-甲午
2020年8月22日--庚子-甲申-丁酉

2020年9月7日--庚子-乙酉-癸丑
2020年9月17日--庚子-乙酉-癸亥
2020年9月22日--庚子-乙酉-戊辰

2020年10月8日--庚子-丙戌-甲申
2020年10月17日--庚子-丙戌-癸巳
2020年10月23日--庚子-丙戌-己亥

2020年11月7日--庚子-丁亥-甲寅
2020年11月15日--庚子-丁亥-壬戌
2020年11月22日--庚子-丁亥-己巳

2020年12月7日--庚子-戊子-甲申
2020年12月15日--庚子-戊子-壬辰

*/
func TestJqT(t *testing.T) {

	y, m, d := 2033, 1, 31 //2033, 12, 22 //2034, 1, 1  // 2020, 6, 23 //2020, 4, 22  //2033, 12, 22
	//清明前一天//6, 4 //芒种润四月 //11, 7 //立东前一天//1, 4 //小寒节前一天
	h := 0
	leapmb := LeapmB(y)
	fmt.Println("两冬至间含闰月?:", leapmb)
	allZqarrt := AllZhongQi(y) //中气
	allShuo := AllShuo(y)
	fmt.Printf("lenzq:%d allshuo:%d\n", len(allZqarrt), len(allShuo))

	//arrT := JqT(y)
	//fmt.Println(len(arrT)) //49
	//fmt.Println(arrT[0])
	jqarrOBj := NewJQArr(y)
	///fmt.Println(jqarrOBj)
	//lichunt := jqarrOBj.Time[3]
	//dongzhit := jqarrOBj.Time[24]
	//fmt.Printf("立春时间:%v 冬至时间:%v\n", lichunt, dongzhit)

	cust := time.Date(y, time.Month(m), d, h, 0, 0, 0, time.Local)
	b, _ := jqarrOBj.LinChuT(cust) //time.Now().Local()
	//fmt.Println(b)
	//节气数组 0小寒 1立春 2惊蛰 3清明 4立夏 5芒种 6小暑 7立秋 8白露 9 寒露 10立东 11大雪
	jie12arrT := jqarrOBj.Jie12ArrT()
	///fmt.Println(jie12arrT)
	zqArrT := jqarrOBj.ZhongQiArrT() //len:13
	fmt.Printf("中气:%v\n len=%d\n", zqArrT, len(zqArrT))

	yearg, yearz := YearGZ(y, b)
	fmt.Printf("年干支:%s%s\n", yearg, yearz)

	//
	index, dz0t, dz0jd := m1(Data(y))
	fmt.Printf("JQHS数组索引值:%d 第一个冬至前朔的时间戳:%v JD:%f\n", index, dz0t, dz0jd)
	//fmt.Println(len(Data(y))) //len(Data(y)-index=60 60/4=15

	shuoWangFArr := NewShuoWangF(index, Data(y))
	moonSWFarr := MoonShuoWangF(shuoWangFArr)
	//fmt.Println("朔望jd len=", len(moonSWFarr))//15

	//ydxArr := YueDaXiao(moonSWFarr)
	//fmt.Printf("%d--月大小:%d\n", len(ydxArr), ydxArr) //14-月大小: [30 30 29 30 30 30 29 30 29 29 30 29 30 29]

	moonSwArrT := MoonShuoWangT(moonSWFarr)
	//fmt.Println("朔望总数:", len(moonSwArrT)) //15

	//===========================================================
	leapmN, leapmT := findLeapmN(leapmb, zqArrT, moonSwArrT)
	fmt.Printf("闰月数字%d 朔的阳历时间:%v\n", leapmN, leapmT)

	mtarr := findLunarMN(allShuo)
	lunaryN, lunarmN := LunarMN(cust, mtarr, leapmb, leapmT)
	fmt.Printf("农历-->%d年-%d月\n", lunaryN, lunarmN)
	///===========================================================

	moonswj12arrT := MoonShuoTJ12ArrT(jie12arrT, moonSwArrT)
	//m, d := 4, 4 //清明前一天//6, 4 //芒种润四月 //11, 7 //立东前一天//1, 4 //小寒节前一天
	diffjietB, indexMGZ, indexSWN := DiffJieT(y, m, d, moonswj12arrT)
	fmt.Printf("%t 月干支索引数字:%d 朔望索引:%d\n", diffjietB, indexMGZ, indexSWN)

	//////////////////
	//swtObj, swtsObj := NewShuoWangT(diffjietB, indexSWN, moonSwArrT)
	//fmt.Printf("阳历年月日对应农历本月朔望时间戳:%v\n阳历年月日对应农历本月朔望时间戳String格式:%v\n", swtObj, swtsObj)
	///////////////////

	mgz, err := MonthGZ(yearg, indexMGZ, diffjietB)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("月干支:", mgz)

	dgz, daygN := DayGZ(y, m, d)
	fmt.Println("日干支:", dgz, daygN)

	HGZ := HourGZ(daygN, h)
	fmt.Println("时干支", HGZ)

	//2458873.404964 农历正月初一 转阳历时间戳 2020-01-25 05:43:08 +0000 UTC
	//2458903.148031 农历二月初一 转阳历时家戳 2020-02-23 23:33:09 +0000 UTC
	//fmt.Printf("%v\n", JdToLocalTime(2458873.404964))
	//m1 := JdToLocalTime(2458873.404964)
	//m2 := JdToLocalTime(2458903.148031)
	////m3 := JdToLocalTime(2458932.895393)
	////精确到天计算农历月大小
	//m1 = time.Date(m1.Year(), m1.Month(), m1.Day(), 0, 0, 0, 0, time.Local)
	//m2 = time.Date(m2.Year(), m2.Month(), m2.Day(), 0, 0, 0, 0, time.Local)
	////m3 = time.Date(m3.Year(), m3.Month(), m3.Day(), 0, 0, 0, 0, time.Local)
	//mh := m2.Sub(m1).Hours() / 24
	//fmt.Printf("%f\n", mh)
	//h12:=covnH24Toh12(22)

}

//
func TestNewLunarMonth(t *testing.T) {
	ymd := []struct {
		y, m, d int
	}{
		{2033, 1, 4},
		{2033, 3, 4},
		{2033, 6, 4},
		{2033, 7, 6},
		{2033, 8, 6},
		{2033, 10, 7},
		{2033, 11, 6},
		{2033, 12, 6},
		{2033, 12, 22},
		{2033, 12, 22},
		{2033, 12, 23},
		{2033, 12, 23},
		{2033, 12, 24},
		{2033, 12, 24},
		{2033, 12, 25},
		{2033, 12, 25},
		{2033, 12, 26},
		{2033, 12, 26},
		{2033, 12, 27},
		{2033, 12, 27},
		{2033, 12, 28},
		{2033, 12, 28},
		{2033, 12, 29},
		{2033, 12, 29},
	}
	want := []struct {
		ly, lm, ld int
	}{
		{2032, 12, 4},
		{2033, 2, 4},
		{2033, 5, 8},
		{2033, 6, 10},
		{2033, 7, 12},
		{2033, 9, 15},
		{2033, 10, 15},
		{2033, 11, 15},
		{2033, 11, 1},
		{2033, 11, 1},
		{2033, 11, 2},
		{2033, 11, 2},
		{2033, 11, 3},
		{2033, 11, 3},
		{2033, 11, 4},
		{2033, 11, 4},
		{2033, 11, 5},
		{2033, 11, 5},
		{2033, 11, 6},
		{2033, 11, 6},
		{2033, 11, 7},
		{2033, 11, 7},
		{2033, 11, 8},
		{2033, 11, 8},
	}

	for i := 0; i < len(ymd); i++ {
		lunar := NewLunar(ymd[i].y, ymd[i].m, ymd[i].d)
		y := lunar.LY
		m := lunar.LM
		d := conv(lunar.LRmc)
		lunarobj := struct {
			ly, lm, ld int
		}{
			ly: y,
			lm: m,
			ld: d,
		}
		if !reflect.DeepEqual(lunarobj, want[i]) {
			t.Errorf("NewLunar(%d %d %d)=%v want:%d,%d %d\n",
				ymd[i].y, ymd[i].m, ymd[i].d, lunarobj, want[i].ly, want[i].lm, want[i].ld)
		}

	}
	//y, m, d := 2020, 1, 25
	//lunar := NewLunar(y, m, d)
	//want := "初一"
	//if !strings.EqualFold(want, lunar.LRmc) {
	//	t.Errorf("NewLunar(%d %d %d)=%s want:%s\n", y, m, d, lunar.LRmc, want)
	//}

}
func conv(rmc string) int {
	Rmc := []string{
		"初一", "初二", "初三", "初四", "初五", "初六", "初七", "初八", "初九", "初十",
		"十一", "十二", "十三", "十四", "十五", "十六", "十七", "十八", "十九", "二十",
		"廿一", "廿二", "廿三", "廿四", "廿五", "廿六", "廿七", "廿八", "廿九", "三十"}
	var index int
	for i := 0; i < len(Rmc); i++ {
		if strings.EqualFold(Rmc[i], rmc) {
			index = i + 1
			break
		}
	}
	return index
}

//
func TestNewLunarM(t *testing.T) {
	//ymd := []struct {
	//	y, m, d int
	//}{
	//	{2020, 1, 23},
	//	{2020, 1, 25},
	//	{2020, 2, 22},
	//	{2020, 2, 23},
	//	{2020, 3, 23},
	//	{2020, 3, 24},
	//	{2020, 4, 22},
	//	{2020, 4, 23},
	//	{2020, 5, 22},
	//	{2020, 5, 23},
	//	{2020, 6, 20},
	//	{2020, 6, 21},
	//	{2020, 9, 14},
	//	{2020, 9, 17},
	//	{2020, 10, 16},
	//	{2020, 10, 17},
	//	{2020, 11, 14},
	//	{2020, 11, 15},
	//	{2020, 12, 14},
	//	{2020, 12, 15},
	//}
	//want := []int{12, 1, 1, 2, 2, 3, 3, 4, 4, 4, 4, 5, 7, 8, 8, 9, 9, 10, 10, 11}
	//NewLunar(2020 1 23)=11 want:12
	//
	//ymd := []struct {
	//	y, m, d int
	//}{
	//	{2033, 12, 3},
	//	{2033, 1, 1},
	//	{2033, 1, 31},
	//	{2033, 3, 1},
	//	{2033, 3, 31},
	//	{2033, 4, 29},
	//	{2033, 5, 28},
	//	{2033, 6, 27},
	//	{2033, 7, 26},
	//	{2033, 8, 25},
	//	{2033, 9, 23},
	//	{2033, 10, 23},
	//	{2033, 11, 22},
	//	{2033, 12, 22},
	//}
	//want := []int{11, 12, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 11}
	///NewLunar(2033 12 22)=12 want:11
	//ymd := []struct {
	//	y, m, d int
	//}{
	//	{2021, 12, 15},
	//	{2021, 1, 13},
	//	{2021, 2, 12},
	//	{2021, 3, 13},
	//	{2021, 4, 12},
	//	{2021, 5, 12},
	//	{2021, 6, 10},
	//	{2021, 7, 10},
	//	{2021, 8, 8},
	//	{2021, 9, 7},
	//	{2021, 10, 6},
	//	{2021, 11, 5},
	//	{2021, 12, 4},
	//}
	//want := []int{11, 12, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	//for i := 0; i < len(ymd); i++ {
	//	_, res := NewLunar(ymd[i].y, ymd[i].m, ymd[i].d)
	//	for j := i; j < len(want); j++ {
	//		if res != want[j] {
	//			t.Errorf("NewLunar(%d %d %d)=%d want:%d\n", ymd[i].y, ymd[i].m, ymd[i].d, res, want[j])
	//		}
	//		break
	//	}
	//}

	y, m, d := 2020, 3, 23 //2033, 12, 22
	lunar := NewLunar(y, m, d)
	fmt.Printf("农历:%d年-%d月%s日\n", lunar.LY, lunar.LM, lunar.LRmc)

}

/*中气
2019-12-22 12:20:35 +0000 UTC 冬至
2020-01-20 22:55:49 +0000 UTC 大寒
2020-02-19 12:58:09 +0000 UTC 雨水
2020-03-20 11:50:46 +0000 UTC 春分
2020-04-19 22:46:38 +0000 UTC 谷雨
2020-05-20 21:50:26 +0000 UTC 小满
2020-06-21 05:44:50 +0000 UTC 夏至
2020-07-22 16:38:01 +0000 UTC 大暑
2020-08-22 23:46:05 +0000 UTC 处暑
2020-09-22 21:31:48 +0000 UTC 秋分
2020-10-23 07:00:41 +0000 UTC 霜降
2020-11-22 04:40:54 +0000 UTC 小雪
2020-12-21 18:03:29 +0000 UTC 冬至
*/

/* MoonShuoTJ12ArrT()
j=0 2019-11-26 23:06:44 +0000 UTC 2020-12-07 00:10:39 +0000 UTC
j=1 2019-12-26 13:14:16 +0000 UTC 2020-01-06 05:31:15 +0000 UTC
j=2 2020-01-25 05:43:08 +0000 UTC 2020-02-04 17:04:29 +0000 UTC
j=3 2020-02-23 23:33:09 +0000 UTC 2020-03-05 10:58:02 +0000 UTC
j=4 2020-03-24 17:29:21 +0000 UTC 2020-04-04 15:39:19 +0000 UTC
j=5 2020-04-23 10:27:00 +0000 UTC 2020-05-05 08:52:32 +0000 UTC
j=6 2020-05-23 01:40:00 +0000 UTC 2020-06-05 12:59:35 +0000 UTC
j=7 2020-06-21 14:42:36 +0000 UTC 2020-07-06 23:15:36 +0000 UTC
j=8 2020-07-21 01:34:06 +0000 UTC 2020-08-07 09:07:20 +0000 UTC
j=9 2020-08-19 10:42:48 +0000 UTC 2020-09-07 12:09:12 +0000 UTC
j=10 2020-09-17 19:01:21 +0000 UTC 2020-10-08 03:56:25 +0000 UTC
j=11 2020-10-17 03:32:12 +0000 UTC 2020-11-07 07:15:04 +0000 UTC
*/
/*y=2020
0==>&{2019-11-26 23:06:44 +0000 UTC 2019-12-04 14:59:22 +0000 UTC 2019-12-12 13:13:25 +0000 UTC 2019-12-19 12:58:14 +0000 UTC}
1==>&{2019-12-26 13:14:16 +0000 UTC 2020-01-03 12:46:33 +0000 UTC 2020-01-11 03:22:27 +0000 UTC 2020-01-17 20:59:33 +0000 UTC}
2==>&{2020-01-25 05:43:08 +0000 UTC 2020-02-02 09:42:49 +0000 UTC 2020-02-09 15:34:26 +0000 UTC 2020-02-16 06:18:21 +0000 UTC}
3==>&{2020-02-23 23:33:09 +0000 UTC 2020-03-03 03:58:32 +0000 UTC 2020-03-10 01:48:53 +0000 UTC 2020-03-16 17:35:20 +0000 UTC}
4==>&{2020-03-24 17:29:21 +0000 UTC 2020-04-01 18:22:24 +0000 UTC 2020-04-08 10:36:14 +0000 UTC 2020-04-15 06:57:17 +0000 UTC}
5==>&{2020-04-23 10:27:00 +0000 UTC 2020-05-01 04:39:30 +0000 UTC 2020-05-07 18:46:22 +0000 UTC 2020-05-14 22:03:51 +0000 UTC}
6==>&{2020-05-23 01:40:00 +0000 UTC 2020-05-30 11:31:03 +0000 UTC 2020-06-06 03:13:32 +0000 UTC 2020-06-13 14:24:50 +0000 UTC}
7==>&{2020-06-21 14:42:36 +0000 UTC 2020-06-28 16:16:50 +0000 UTC 2020-07-05 12:45:33 +0000 UTC 2020-07-13 07:30:09 +0000 UTC}
8==>&{2020-07-21 01:34:06 +0000 UTC 2020-07-27 20:33:42 +0000 UTC 2020-08-03 23:59:55 +0000 UTC 2020-08-12 00:45:56 +0000 UTC}
9==>&{2020-08-19 10:42:48 +0000 UTC 2020-08-26 01:58:47 +0000 UTC 2020-09-02 13:23:13 +0000 UTC 2020-09-10 17:26:52 +0000 UTC}
10==>&{2020-09-17 19:01:21 +0000 UTC 2020-09-24 09:56:01 +0000 UTC 2020-10-02 05:06:25 +0000 UTC 2020-10-10 08:40:41 +0000 UTC}
11==>&{2020-10-17 03:32:12 +0000 UTC 2020-10-23 21:24:05 +0000 UTC 2020-10-31 22:50:18 +0000 UTC 2020-11-08 21:47:15 +0000 UTC}
12==>&{2020-11-15 13:08:20 +0000 UTC 2020-11-22 12:46:10 +0000 UTC 2020-11-30 17:30:50 +0000 UTC 2020-12-08 08:37:45 +0000 UTC}
13==>&{2020-12-15 00:17:44 +0000 UTC 2020-12-22 07:42:22 +0000 UTC 2020-12-30 11:29:21 +0000 UTC 2021-01-06 17:38:22 +0000 UTC}
14==>&{2021-01-13 13:01:19 +0000 UTC 2021-01-21 05:02:45 +0000 UTC 2021-01-29 03:17:23 +0000 UTC 2021-02-05 01:38:15 +0000 UTC}
*/

/*第一列值就是当月农历初一(朔) 第二列上弦 第三列望 第四列下弦
i:3=>{2458814.129681 2458821.791232 2458829.717652 2458836.707108}  11月
i:7=>{2458843.718252 2458851.699002 2458859.307263 2458866.041364}	12月
i:11=>{2458873.404964 2458881.571406 2458888.815579 2458895.429417}	正月
i:15=>{2458903.148031 2458911.332319 2458918.242289 2458924.899546}	二月
i:19=>{2458932.895393 2458940.932227 2458947.608500 2458954.456456}
i:23=>{2458962.602088 2458970.360770 2458976.948873 2458984.086016}
i:27=>{2458992.236120 2458999.646574 2459006.301072 2459013.767247}
i:31=>{2459021.779591 2459028.845029 2459035.698309 2459043.479276}
i:35=>{2459051.232020 2459058.023413 2459065.166611 2459073.198571}
i:39=>{2459080.613066 2459087.249158 2459094.724464 2459102.893669}
i:43=>{2459109.959280 2459116.580572 2459124.379457 2459132.528254}
i:47=>{2459139.314030 2459146.058395 2459154.118269 2459162.074482}
i:51=>{2459168.714126 2459175.698730 2459183.896415 2459191.526223}

i:55=>{2459198.178984 2459205.487759 2459213.645392 2459220.901651}
i:59=>{2459227.709258 2459235.376921 2459243.303742 2459250.234899}
*/
/* shuoWangFArr
i:0=>{2458814.129681 0.000000 0.000000 0.000000}
i:1=>{2458814.129681 2458821.791232 0.000000 0.000000}
i:2=>{2458814.129681 2458821.791232 2458829.717652 0.000000}
i:3=>{2458814.129681 2458821.791232 2458829.717652 2458836.707108}
i:4=>{2458843.718252 2458821.791232 2458829.717652 2458836.707108}
i:5=>{2458843.718252 2458851.699002 2458829.717652 2458836.707108}
i:6=>{2458843.718252 2458851.699002 2458859.307263 2458836.707108}
i:7=>{2458843.718252 2458851.699002 2458859.307263 2458866.041364}
i:8=>{2458873.404964 2458851.699002 2458859.307263 2458866.041364}
i:9=>{2458873.404964 2458881.571406 2458859.307263 2458866.041364}
i:10=>{2458873.404964 2458881.571406 2458888.815579 2458866.041364}
i:11=>{2458873.404964 2458881.571406 2458888.815579 2458895.429417}
i:12=>{2458903.148031 2458881.571406 2458888.815579 2458895.429417}
i:13=>{2458903.148031 2458911.332319 2458888.815579 2458895.429417}
i:14=>{2458903.148031 2458911.332319 2458918.242289 2458895.429417}
i:15=>{2458903.148031 2458911.332319 2458918.242289 2458924.899546}
i:16=>{2458932.895393 2458911.332319 2458918.242289 2458924.899546}
i:17=>{2458932.895393 2458940.932227 2458918.242289 2458924.899546}
i:18=>{2458932.895393 2458940.932227 2458947.608500 2458924.899546}
i:19=>{2458932.895393 2458940.932227 2458947.608500 2458954.456456}
i:20=>{2458962.602088 2458940.932227 2458947.608500 2458954.456456}
i:21=>{2458962.602088 2458970.360770 2458947.608500 2458954.456456}
i:22=>{2458962.602088 2458970.360770 2458976.948873 2458954.456456}
i:23=>{2458962.602088 2458970.360770 2458976.948873 2458984.086016}
i:24=>{2458992.236120 2458970.360770 2458976.948873 2458984.086016}
i:25=>{2458992.236120 2458999.646574 2458976.948873 2458984.086016}
i:26=>{2458992.236120 2458999.646574 2459006.301072 2458984.086016}
i:27=>{2458992.236120 2458999.646574 2459006.301072 2459013.767247}
i:28=>{2459021.779591 2458999.646574 2459006.301072 2459013.767247}
i:29=>{2459021.779591 2459028.845029 2459006.301072 2459013.767247}
i:30=>{2459021.779591 2459028.845029 2459035.698309 2459013.767247}
i:31=>{2459021.779591 2459028.845029 2459035.698309 2459043.479276}
i:32=>{2459051.232020 2459028.845029 2459035.698309 2459043.479276}
i:33=>{2459051.232020 2459058.023413 2459035.698309 2459043.479276}
i:34=>{2459051.232020 2459058.023413 2459065.166611 2459043.479276}
i:35=>{2459051.232020 2459058.023413 2459065.166611 2459073.198571}
i:36=>{2459080.613066 2459058.023413 2459065.166611 2459073.198571}
i:37=>{2459080.613066 2459087.249158 2459065.166611 2459073.198571}
i:38=>{2459080.613066 2459087.249158 2459094.724464 2459073.198571}
i:39=>{2459080.613066 2459087.249158 2459094.724464 2459102.893669}
i:40=>{2459109.959280 2459087.249158 2459094.724464 2459102.893669}
i:41=>{2459109.959280 2459116.580572 2459094.724464 2459102.893669}
i:42=>{2459109.959280 2459116.580572 2459124.379457 2459102.893669}
i:43=>{2459109.959280 2459116.580572 2459124.379457 2459132.528254}
i:44=>{2459139.314030 2459116.580572 2459124.379457 2459132.528254}
i:45=>{2459139.314030 2459146.058395 2459124.379457 2459132.528254}
i:46=>{2459139.314030 2459146.058395 2459154.118269 2459132.528254}
i:47=>{2459139.314030 2459146.058395 2459154.118269 2459162.074482}
i:48=>{2459168.714126 2459146.058395 2459154.118269 2459162.074482}
i:49=>{2459168.714126 2459175.698730 2459154.118269 2459162.074482}
i:50=>{2459168.714126 2459175.698730 2459183.896415 2459162.074482}
i:51=>{2459168.714126 2459175.698730 2459183.896415 2459191.526223}
i:52=>{2459198.178984 2459175.698730 2459183.896415 2459191.526223}
i:53=>{2459198.178984 2459205.487759 2459183.896415 2459191.526223}
i:54=>{2459198.178984 2459205.487759 2459213.645392 2459191.526223}
i:55=>{2459198.178984 2459205.487759 2459213.645392 2459220.901651}
i:56=>{2459227.709258 2459205.487759 2459213.645392 2459220.901651}
i:57=>{2459227.709258 2459235.376921 2459213.645392 2459220.901651}
i:58=>{2459227.709258 2459235.376921 2459243.303742 2459220.901651}
i:59=>{2459227.709258 2459235.376921 2459243.303742 2459250.234899}
2019-11-26 23:06:44 +0000 UTC*/