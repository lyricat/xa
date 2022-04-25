/*
 * The code is automatically generated by the Goland.
 * Copyright © Aquarian-Age. All Rights Reserved.
 * Licensed under MIT
 */

package gz

import (
	"github.com/Aquarian-Age/xa/pkg/pub"
	"sort"
	"strings"
)

// HuangHei 黄黑 传(月干支 日干支 或 日干支 时辰干支)
func HuangHei(gz, subgz string) string {
	arr := []string{"青龙", "金匮", "司命", "明堂", "天德", "玉堂", "天刑", "白虎", "天牢", "朱雀", "玄武", "勾陈"}
	//k:月建 v:对应arr 如正月建寅 子日青龙 亥日勾陈
	//k:日干支 v:时辰数组 如寅日 子时青龙 亥时勾陈
	xmap := map[string][]string{
		"寅": {"子", "辰", "戌", "丑", "巳", "未", "寅", "午", "申", "卯", "酉", "亥"},
		"卯": {"寅", "午", "子", "卯", "未", "酉", "辰", "申", "戌", "巳", "亥", "丑"},
		"辰": {"辰", "申", "寅", "巳", "酉", "亥", "午", "戌", "子", "未", "丑", "卯"},
		"巳": {"午", "戌", "辰", "未", "亥", "丑", "申", "子", "寅", "酉", "卯", "巳"},
		"午": {"申", "子", "午", "酉", "丑", "卯", "戌", "寅", "辰", "亥", "巳", "未"},
		"未": {"戌", "寅", "申", "亥", "卯", "巳", "子", "辰", "午", "丑", "未", "酉"},
		"申": {"子", "辰", "戌", "丑", "巳", "未", "寅", "午", "申", "卯", "酉", "亥"},
		"酉": {"寅", "午", "子", "卯", "未", "酉", "辰", "申", "戌", "巳", "亥", "丑"},
		"戌": {"辰", "申", "寅", "巳", "酉", "亥", "午", "戌", "子", "未", "丑", "卯"},
		"亥": {"午", "戌", "辰", "未", "亥", "丑", "申", "子", "寅", "酉", "卯", "巳"},
		"子": {"申", "子", "午", "酉", "丑", "卯", "戌", "寅", "辰", "亥", "巳", "未"},
		"丑": {"戌", "寅", "申", "亥", "卯", "巳", "子", "辰", "午", "丑", "未", "酉"},
	}
	zs := pub.GetZhiS(gz)
	subzs := pub.GetZhiS(subgz)
	var hhs string
	for z, subarr := range xmap {
		if strings.EqualFold(zs, z) {
			for i := 0; i < len(subarr); i++ {
				if strings.EqualFold(subzs, subarr[i]) {
					hhs = arr[i]
					break
				}
			}
			break
		}
	}
	return hhs
}

//日黄黑
func GetRiHuangHei(mgz, dgz string) string {
	zhimap := huangHeiMap()
	var hhs string
	hhs = huanHei(mgz, dgz, zhimap)
	return hhs
}

//寅申青龙起子 卯酉起寅 辰戌起辰 巳亥起午 子午起申 丑未起戌 顺行十二辰
//生成月份黄黑道地支map数组
func huangHeiMap() map[int][]string {
	zhi := []string{"子", "丑", "寅", "卯", "辰", "巳", "午", "未", "申", "酉", "戌", "亥"}
	var index []string
	var zhimap map[int][]string
	zhimap = make(map[int][]string)
	for i := 0; i < len(zhi); i++ {
		if i%2 == 0 { //只要阳支
			index = zhi[i:]
			swap := zhi[:i]
			index = append(index, swap...)
			zhimap[i] = index
		}
	}
	return zhimap
}

/*
	青龙, 明堂, 天刑, 朱雀, 金匮, 天德, 白虎, 玉堂, 天牢, 玄武, 司命, 勾陈
	寅申青龙起子 卯酉起寅 辰戌起辰 巳亥起午 子午起申 丑未起戌 顺行十二辰
	月起日则 建寅之月子日为青龙 丑日为明堂
	日起时则 子日申时起青龙 酉时为明堂 依次顺数
*/
func huanHei(gz1, gz2 string, zhimap map[int][]string) (hhs string) {
	var keys []int
	for k := range zhimap {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	//如果是月日论 这里mgz为月干支 dgz为日干支
	//如果是日时论 这里mgz为日干支 dgz为时辰干支
	mgz := gz1
	dgz := gz2
	//正月建寅
	zhi := []string{"err", "寅", "卯", "辰", "巳", "午", "未", "申", "酉", "戌", "亥", "子", "丑"}
	//神煞
	hh := []string{"青龙", "明堂", "天刑", "朱雀", "金匮", "天德", "白虎", "玉堂", "天牢", "玄武", "司命", "勾陈"}
	//地支对应黄黑道
	var hhz map[string]string
	hhz = make(map[string]string)
	var nx int
	for nx = 0; nx < len(zhi); nx++ {
		if strings.ContainsAny(mgz, zhi[nx]) {
			break
		}
	}
	mi := zhi[nx]
	//日地支配神煞
	for _, k := range keys {
		switch mi {
		case zhi[1], zhi[7]: //寅申月起子
			if k == 0 {
			I1:
				for i := 0; i < len(hh); i++ {
					hhz[zhimap[0][i]] = hh[i]
					for zi, name := range hhz {
						if strings.ContainsAny(dgz, zi) {
							hhs = name
							break I1
						}
					}
				}
				break
			}
		case zhi[2], zhi[8]: //卯酉起寅
			if k == 2 {
			I2:
				for i := 0; i < len(hh); i++ {
					hhz[zhimap[2][i]] = hh[i]
					for zi, name := range hhz {
						if strings.ContainsAny(dgz, zi) {
							hhs = name
							break I2
						}
					}
				}
				break
			}
		case zhi[3], zhi[9]:
			if k == 4 {
			I3:
				for i := 0; i < len(hh); i++ {
					hhz[zhimap[4][i]] = hh[i]
					for zi, name := range hhz {
						if strings.ContainsAny(dgz, zi) {
							hhs = name
							break I3
						}
					}
				}
				break
			}
		case zhi[4], zhi[10]:
			if k == 6 {
			I4:
				for i := 0; i < len(hh); i++ {
					hhz[zhimap[6][i]] = hh[i]
					for zi, name := range hhz {
						if strings.ContainsAny(dgz, zi) {
							hhs = name
							break I4
						}
					}
				}
				break
			}
		case zhi[5], zhi[11]:
			if k == 8 {
			I5:
				for i := 0; i < len(hh); i++ {
					hhz[zhimap[8][i]] = hh[i]
					for zi, name := range hhz {
						if strings.ContainsAny(dgz, zi) {
							hhs = name
							break I5
						}
					}
				}
				break
			}
		case zhi[6], zhi[12]:
			if k == 10 {
			I6:
				for i := 0; i < len(hh); i++ {
					hhz[zhimap[10][i]] = hh[i]
					for zi, name := range hhz {
						if strings.ContainsAny(dgz, zi) {
							hhs = name
							break I6
						}
					}
				}
				break
			}
		}
	}
	return
}
