package pub

import (
	"log"
	"regexp"
	"strings"
)

//取天干 传干支 返回干
func GetGanS(gz string) string {
	gan := [11]string{"甲", "乙", "丙", "丁", "戊", "己", "庚", "辛", "壬", "癸"}
	var g string
	for i := 0; i < len(gan); i++ {
		if strings.ContainsAny(gz, gan[i]) {
			g = gan[i]
			break
		}
	}
	return g
}

//取地支 传干支返回支
func GetZhiS(gz string) string {
	zhi := []string{"子", "丑", "寅", "卯", "辰", "巳", "午", "未", "申", "酉", "戌", "亥"}
	var z string
	for i := 0; i < len(zhi); i++ {
		if strings.ContainsAny(gz, zhi[i]) {
			z = zhi[i]
			break
		}
	}
	return z
}

//顺序排地支 传入对应的地支 原地支数组 返回排序后的地支数组
func SortArr(zhi string, zhiArr []string) []string {
	for i := 0; i < len(zhiArr); i++ {
		if strings.EqualFold(zhi, zhiArr[i]) {
			head := zhiArr[:i]
			end := zhiArr[i:]
			zhiArr = append(end, head...)
			break
		}
	}
	return zhiArr
}

//逆序排地支
func ReSortArr(zhi string, zhiArr []string) []string {
	for i := 0; i < len(zhiArr); i++ {
		if strings.EqualFold(zhi, zhiArr[i]) {
			head := zhiArr[:i]
			end := zhiArr[i:]
			zhiArr = append(end, head...)
			break
		}
	}
	//
	head := zhiArr[:1]
	end := zhiArr[1:]
	var rArr []string
	for i := len(end) - 1; i >= 0; i-- {
		rArr = append(rArr, end[i])
	}
	rArr = append(head, rArr...)
	return rArr
}

//顺序不变　把arr数组中source名字改为source + add 如果没有符合条件的返回原值
func ReName(arr []string, source, add string) []string {
	//fmt.Printf("arr:%s source:%s add:%s\n", arr, source, add)
	var newArr []string
	for i := 0; i < len(arr); i++ {
		if strings.EqualFold(arr[i], source) {
			source = source + add
			head := arr[:i+1]
			head = head[:len(head)-1]
			head = append(head, source)
			end := arr[i+1:]
			//fmt.Printf("\nhead:%s\nend:%s\n", head, end)
			newArr = append(head, end...)
		} else {
			newArr = arr
		}
	}
	return newArr
}

//正则替换 把s的dels元素替换成r
func ReNameS(s, dels string, r string) string {
	rs, err := regexp.Compile(dels)
	if err != nil {
		log.Fatal(err)
	}
	return rs.ReplaceAllString(s, r)
}
