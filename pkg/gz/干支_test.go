/*
 * The code is automatically generated by the Goland.
 * Copyright © Aquarian-Age. All Rights Reserved.
 * Licensed under MIT
 */

package gz

import (
	"fmt"
	"strings"
	"testing"
	"time"
)

func TestNewGanZhi(t *testing.T) {
	y, m, d, h := 2021, 2, 3, 21 //庚子 己丑 壬午 辛亥
	//y, m, d, h = 2021, 2, 3, 22  //辛丑 庚寅 壬午 辛亥
	//y, m, d, h = 2033, 2, 3, 19  //壬子 癸丑 乙酉 丙戌
	//y, m, d, h = 2033, 2, 3, 20  //癸丑 甲寅 乙酉 丙戌
	//y, m, d, h = 2022, 2, 3, 0 //辛丑 辛丑 丁亥 庚子
	//y, m, d, h = 2022, 2, 4, 0 //辛丑 辛丑 戊子 壬子
	//y, m, d, h = 2022, 2, 4, 4//壬寅 壬寅 戊子 甲寅
	y, m, d, h = 2022, 3, 4, 0 //壬寅 壬寅 丙辰 戊子
	y, m, d, h = 2022, 3, 5, 0 //壬寅 癸卯 丁巳 庚子
	gzo := NewGanZhi(y, m, d, h)
	fmt.Println(gzo.YGZ, gzo.MGZ, gzo.DGZ, gzo.HGZ)

	//cust := time.Date(y, time.Month(m), d, h, 0, 0, 0, time.Local)
	//lcb, lct := fixLiChun(y, cust)
	//yg, yz := yearGZ(y, lcb)
	//ygz := yg + yz
	//arrT, _ := getJieArr(y)
	//jieqib, index := findJie(cust, arrT)
	//mgz := monthGZ(cust, lcb, lct, jieqib, index)
	//dgz, gn := dayGanZhi(y, m, d)
	//hgz := GetHourGZ(gn, h)
	//fmt.Println(ygz, mgz, dgz, hgz)

}

//立春布尔值
func TestFixLiChun(t *testing.T) {
	year := 2021
	year = 2033
	cust := time.Date(year, time.Month(2), 3, 22, 0, 0, 0, time.Local) //精确到时
	cust = time.Date(year, time.Month(2), 3, 21, 0, 0, 0, time.Local)
	cust = time.Date(year, time.Month(2), 3, 19, 0, 0, 0, time.Local)
	cust = time.Date(year, time.Month(2), 3, 20, 0, 0, 0, time.Local)
	lcb, _ := fixLiChun(year, cust)
	fmt.Println(lcb)
}

//年干 年支
func TestYearGZ(t *testing.T) {
	year := 2021
	lcb := true //干:辛 支:丑
	lcb = false //干:庚 支:子
	year = 2033
	lcb = true  //干:癸 支:丑
	lcb = false //干:壬 支:子
	g, z := yearGZ(year, lcb)
	fmt.Printf("干:%s 支:%s\n", g, z)
}

//年干支
func TestGetYGZ(t *testing.T) {
	y, m, d := 2021, 2, 3 //庚子
	y, m, d = 2021, 2, 3  //辛丑
	y, m, d = 2033, 2, 3  //壬子
	y, m, d = 2033, 2, 3  //癸丑
	gz := GetYGZ(y, m, d)
	fmt.Println(gz)
}

//2年的节气 中气
func TestGetJieArr(t *testing.T) {
	year := 2021
	getJieArr(year)
	/* 12节时间戳: [
	2021-01-05 11:23:16.924079954 +0800 CST
	2021-02-03 22:58:37.499354481 +0800 CST
	2021-03-05 16:53:32.036919593 +0800 CST
	2021-04-04 21:34:57.242134809 +0800 CST
	2021-05-05 14:47:00.883324742 +0800 CST
	2021-06-05 18:51:56.499531269 +0800 CST
	2021-07-07 05:05:18.752901256 +0800 CST
	2021-08-07 14:53:47.871581912 +0800 CST
	2021-09-07 17:52:46.014263033 +0800 CST
	2021-10-08 09:38:52.155123353 +0800 CST
	2021-11-07 12:58:36.09108746 +0800 CST
	2021-12-07 05:56:54.075279235 +0800 CST
	2022-01-05 17:13:53.057212829 +0800 CST
	2022-02-04 04:50:27.707844972 +0800 CST
	2022-03-05 22:43:34.501436054 +0800 CST
	2022-04-05 03:20:03.713084757 +0800 CST
	2022-05-05 20:25:46.525521576 +0800 CST
	2022-06-06 00:25:37.778542935 +0800 CST
	2022-07-07 10:37:49.296710193 +0800 CST
	2022-08-07 20:28:58.208247721 +0800 CST
	2022-09-07 23:32:07.361194789 +0800 CST
	2022-10-08 15:22:16.706931889 +0800 CST
	2022-11-07 18:45:18.361879885 +0800 CST
	2022-12-07 11:46:04.664840698 +0800 CST
	] len=24
	 12中气时间戳:[
	 2020-12-21 18:02:11.643059849 +0800 CST
	 2021-01-20 04:39:42.671713829 +0800 CST
	 2021-02-18 18:43:46.703435182 +0800 CST
	 2021-03-20 17:37:19.524659514 +0800 CST
	 2021-04-20 04:33:14.733846187 +0800 CST
	 2021-05-21 03:36:57.461840808 +0800 CST
	 2021-06-21 11:32:00.398005843 +0800 CST
	 2021-07-22 22:26:15.868888199 +0800 CST
	 2021-08-23 05:34:48.650563359 +0800 CST
	 2021-09-23 03:20:55.759392678 +0800 CST
	 2021-10-23 12:51:00.527555644 +0800 CST
	 2021-11-22 10:33:34.356350898 +0800 CST
	 2021-12-21 23:59:08.787242174 +0800 CST
	 2022-01-20 10:38:55.942174494 +0800 CST
	 2022-02-19 00:42:44.283978939 +0800 CST
	 2022-03-20 23:33:14.363366067 +0800 CST
	 2022-04-20 10:24:05.944789052 +0800 CST
	 2022-05-21 09:22:24.211819767 +0800 CST
	 2022-06-21 17:13:40.219503343 +0800 CST
	 2022-07-23 04:06:48.266571164 +0800 CST
	 2022-08-23 11:15:58.711329102 +0800 CST
	 2022-09-23 09:03:30.910141468 +0800 CST
	 2022-10-23 18:35:30.812180042 +0800 CST
	 2022-11-22 16:20:18.264108896 +0800 CST
	 2022-12-22 05:48:01.263524293 +0800 CST
	 ] len=25 */
}

//月干支
func TestMonthGZ(t *testing.T) {
	arrT := []struct {
		y, m, d, h int
	}{
		{2021, 1, 4, 22}, {2021, 1, 5, 22}, {2021, 2, 2, 22}, {2021, 2, 3, 22}, {2021, 3, 4, 22}, {2021, 3, 5, 22}, {2021, 4, 3, 22}, {2021, 4, 4, 22}, {2021, 5, 4, 22}, {2021, 5, 5, 22}, {2021, 6, 4, 22}, {2021, 6, 5, 22}, {2021, 7, 6, 22}, {2021, 7, 7, 22}, {2021, 8, 6, 22}, {2021, 8, 7, 22}, {2021, 9, 6, 22}, {2021, 9, 7, 22}, {2021, 10, 7, 22}, {2021, 10, 8, 22}, {2021, 11, 6, 22}, {2021, 11, 7, 22}, {2021, 12, 6, 22}, {2021, 12, 7, 22}, {2022, 1, 4, 22}, {2022, 1, 5, 22}, {2022, 2, 3, 22}, {2022, 2, 4, 22}, {2020, 5, 5, 22}, {2020, 6, 4, 22}, {2020, 6, 5, 22}, {2033, 1, 4, 22}, {2033, 1, 5, 22}, {2033, 12, 6, 22}, {2033, 12, 7, 22}, {2034, 1, 4, 22}, {2034, 1, 5, 22},
	}
	wantMgzArr := []string{"戊子", "己丑", "己丑", "庚寅", "庚寅", "辛卯", "辛卯", "壬辰", "壬辰", "癸巳", "癸巳", "甲午", "甲午", "乙未", "乙未", "丙申", "丙申", "丁酉", "丁酉", "戊戌", "戊戌", "己亥", "己亥", "庚子", "庚子", "辛丑", "辛丑", "壬寅", "辛巳", "辛巳", "壬午", "壬子", "癸丑", "癸亥", "甲子", "甲子", "乙丑"}
	for i := 0; i < len(arrT); i++ {
		tx := arrT[i]
		gz := GetMonthGZ(tx.y, tx.m, tx.d, tx.h)
		if !strings.EqualFold(gz, wantMgzArr[i]) {
			t.Errorf("func MonthGZ(%d %d %d %d)=%s want:%s\n", tx.y, tx.m, tx.d, tx.h, gz, wantMgzArr[i])
		}
	}
}

/*	y, m, d, h := 2021, 12, 7, 22
	y, m, d, h = 2021, 1, 4, 22 //年干:庚 index:10 月干支:戊子
	y, m, d, h = 2021, 1, 5, 22 //年干:庚 index:11 月干支:己丑
	y, m, d, h = 2021, 2, 2, 22 //年干:庚 index:11 月干支:己丑
	y, m, d, h = 2021, 2, 3, 22 //年干:辛 index:0 月干支:庚寅
	y, m, d, h = 2021, 3, 4, 22 //年干:辛 index:0 月干支:庚寅
	y, m, d, h = 2021, 3, 5, 22 //年干:辛 index:1 月干支:辛卯
	y, m, d, h = 2021, 4, 3, 22 //年干:辛 index:1 月干支:辛卯
	y, m, d, h = 2021, 4, 4, 22 //年干:辛 index:2 月干支:壬辰
	y, m, d, h = 2021, 5, 4, 22 //年干:辛 index:2 月干支:壬辰
	y, m, d, h = 2021, 5, 5, 22 //年干:辛 index:3 月干支:癸巳
	y, m, d, h = 2021, 6, 4, 22 //年干:辛 index:3 月干支:癸巳
	y, m, d, h = 2021, 6, 5, 22 //年干:辛 index:4 月干支:甲午
	y, m, d, h = 2021, 7, 6, 22 //年干:辛 index:4 月干支:甲午
	y, m, d, h = 2021, 7, 7, 22 //年干:辛 index:5 月干支:乙未
	y, m, d, h = 2021, 8, 6, 22 //年干:辛 index:5 月干支:乙未
	y, m, d, h = 2021, 8, 7, 22 //年干:辛 index:6 月干支:丙申
	y, m, d, h = 2021, 9, 6, 22 //年干:辛 index:6 月干支:丙申
	y, m, d, h = 2021, 9, 7, 22 //年干:辛 index:7 月干支:丁酉
	y, m, d, h = 2021, 10, 7, 22 //年干:辛 index:7 月干支:丁酉
	y, m, d, h = 2021, 10, 8, 22 //年干:辛 index:8 月干支:戊戌
	y, m, d, h = 2021, 11, 6, 22 //年干:辛 index:8 月干支:戊戌
	y, m, d, h = 2021, 11, 7, 22 //年干:辛 index:9 月干支:己亥
	y, m, d, h = 2021, 12, 6, 22 //年干:辛 index:9 月干支:己亥
	y, m, d, h = 2021, 12, 7, 22 //年干:辛 index:10 月干支:庚子
	y, m, d, h = 2022, 1, 4, 22 //年干:辛 index:10 月干支:庚子
	y, m, d, h = 2022, 1, 5, 22 //年干:辛 index:11 月干支:辛丑
	y, m, d, h = 2022, 2, 3, 22 //年干:辛 index:11 月干支:辛丑
	y, m, d, h = 2022, 2, 4, 22 //年干:壬 index:0 月干支:壬寅
	////
	y, m, d, h = 2020, 5, 5, 22 //年干:庚 index:3 月干支:辛巳
	y, m, d, h = 2020, 6, 4, 22 //年干:庚 index:3 月干支:辛巳
	y, m, d, h = 2020, 6, 5, 22 //年干:庚 index:4 月干支:壬午
	//
	y, m, d, h = 2033, 1, 4, 22 //年干:壬 index:10 月干支:壬子
	y, m, d, h = 2033, 1, 5, 22 //年干:壬 index:11 月干支:癸丑
	y, m, d, h = 2033, 12, 6, 22 //年干:癸 index:9 月干支:癸亥
	y, m, d, h = 2033, 12, 7, 22 //年干:癸 index:10 月干支:甲子
	y, m, d, h = 2034, 1, 4, 22 //年干:癸 index:10 月干支:甲子
	y, m, d, h = 2034, 1, 5, 22 //年干:癸 index:11 月干支:乙丑
	MonthGZ(y, m, d, h)*/
//日干支
func TestDayGZ(t *testing.T) {
	arr2020 := []struct {
		y int
		m int
		d int
	}{
		{2020, 11, 26}, {2020, 12, 7}, {2020, 12, 22}, {2020, 12, 26}, {2020, 1, 6}, {2020, 1, 20}, {2020, 1, 25}, {2020, 2, 4}, {2020, 2, 19}, {2020, 2, 23}, {2020, 3, 5}, {2020, 3, 20}, {2020, 3, 24}, {2020, 4, 4}, {2020, 4, 19}, {2020, 4, 23}, {2020, 5, 5}, {2020, 5, 20}, {2020, 5, 23}, {2020, 6, 5}, {2020, 6, 21}, {2020, 6, 21}, {2020, 7, 6}, {2020, 7, 21}, {2020, 7, 22}, {2020, 8, 7}, {2020, 8, 19}, {2020, 8, 22}, {2020, 9, 7}, {2020, 9, 17}, {2020, 9, 22}, {2020, 10, 8}, {2020, 10, 17}, {2020, 10, 23}, {2020, 11, 7}, {2020, 11, 15}, {2020, 11, 22}, {2020, 12, 7}, {2020, 12, 15},
	}
	wantArr := []string{"癸酉", "甲申", "己亥", "癸卯", "戊申", "壬戌", "丁卯", "丁丑", "壬辰", "丙申", "丁未", "壬戌", "丙寅", "丁丑", "壬辰", "丙申", "戊申", "癸亥", "丙寅", "己卯", "乙未", "乙未", "庚戌", "乙丑", "丙寅", "壬午", "甲午", "丁酉", "癸丑", "癸亥", "戊辰", "甲申", "癸巳", "己亥", "甲寅", "壬戌", "己巳", "甲申", "壬辰"}
	//
	md2033 := []struct {
		y, m, d int
	}{
		{2033, 12, 3}, {2033, 12, 6}, {2033, 12, 7}, {2033, 12, 21}, {2033, 1, 1}, {2033, 1, 5}, {2033, 1, 20}, {2033, 1, 31}, {2033, 2, 3}, {2033, 2, 18}, {2033, 3, 1}, {2033, 3, 5}, {2033, 3, 20}, {2033, 3, 31}, {2033, 4, 4}, {2033, 4, 20}, {2033, 4, 29}, {2033, 5, 5}, {2033, 5, 21}, {2033, 5, 28}, {2033, 6, 5}, {2033, 6, 21}, {2033, 6, 27}, {2033, 7, 7}, {2033, 7, 22}, {2033, 7, 26}, {2033, 8, 7}, {2033, 8, 23}, {2033, 8, 25}, {2033, 9, 7}, {2033, 9, 23}, {2033, 9, 23}, {2033, 10, 8}, {2033, 10, 23}, {2033, 10, 23}, {2033, 11, 7}, {2033, 11, 22}, {2033, 11, 22}, {2033, 12, 7},
	}
	want2033 := []string{"戊子", "辛卯", "壬辰", "丙午", "壬子", "丙辰", "辛未", "壬午", "乙酉", "庚子", "辛亥", "乙卯", "庚午", "辛巳", "乙酉", "辛丑", "庚戌", "丙辰", "壬申", "己卯", "丁亥", "癸卯", "己酉", "己未", "甲戌", "戊寅", "庚寅", "丙午", "戊申", "辛酉", "丁丑", "丁丑", "壬辰", "丁未", "丁未", "壬戌", "丁丑", "丁丑", "壬辰"}
	for i := 0; i < len(arr2020); i++ {
		tx := arr2020[i]
		dgz, _ := dayGanZhi(tx.y, tx.m, tx.d)
		if !strings.EqualFold(dgz, wantArr[i]) {
			t.Errorf("func dayGanZhi(%d %d %d)=%s want:%s", tx.y, tx.m, tx.d, dgz, wantArr[i])
		}
	}
	//
	for i := 0; i < len(md2033); i++ {
		tx := md2033[i]
		dgz, _ := dayGanZhi(tx.y, tx.m, tx.d)
		if !strings.EqualFold(dgz, want2033[i]) {
			t.Errorf("func dayGanZhi(%d %d %d)=%s want:%s", tx.y, tx.m, tx.d, dgz, want2033[i])
		}
	}
}
func TestHourGZ(t *testing.T) {
	gn := "甲" //1 //[甲子 乙丑 丙寅 丁卯 戊辰 己巳 庚午 辛未 壬申 癸酉 甲戌 乙亥]
	gn = "丙"  //2  //[戊子 己丑 庚寅 辛卯 壬辰 癸巳 甲午 乙未 丙申 丁酉 戊戌 己亥]
	gn = "戊"  // 3  //[壬子 癸丑 甲寅 乙卯 丙辰 丁巳 戊午 己未 庚申 辛酉 壬戌 癸亥]
	gn = "庚"  //4  //[丙子 丁丑 戊寅 己卯 庚辰 辛巳 壬午 癸未 甲申 乙酉 丙戌 丁亥]
	gn = "壬"  //5  //[庚子 辛丑 壬寅 癸卯 甲辰 乙巳 丙午 丁未 戊申 己酉 庚戌 辛亥]
	hour := 0
	//arr := hgzArr(gn)
	//fmt.Println(arr)
	hgz := GetHourGZ(gn, hour)
	fmt.Println(hgz)
}

//月将
func TestGanZhi_GetYueJiangName(t *testing.T) {
	y, m, d, h := 2021, 7, 22, 22
	obj := NewGanZhi(y, m, d, h)
	yjzhi, name, zt, _ := obj.YueJiang()
	fmt.Println(yjzhi, name, zt) //未 小吉 2021-06-21 11:32:00.398005843 +0800 CST
}
