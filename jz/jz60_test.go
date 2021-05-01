package jz

import (
	"reflect"
	"testing"
)

func TestGetJzArr(t *testing.T) {
	arr:=[]string{
		"甲子","乙丑","丙寅","丁卯","戊辰","己巳","庚午","辛未","壬申","癸酉",
		"甲戌","乙亥","丙子","丁丑","戊寅","己卯","庚辰","辛巳","壬午","癸未",
		"甲申","乙酉","丙戌","丁亥","戊子","己丑","庚寅","辛卯","壬辰","癸巳",
		"甲午","乙未","丙申","丁酉","戊戌","己亥","庚子","辛丑","壬寅","癸卯",
		"甲辰","乙巳","丙午","丁未","戊申","己酉","庚戌","辛亥","壬子","癸丑",
		"甲寅","乙卯","丙辰","丁巳","戊午","己未","庚申","辛酉","壬戌","癸亥"}
	arrx:= GetJzArr()
	if !reflect.DeepEqual(arr,arrx){
		t.Errorf("func (GetJzArr()=%s want:%s",arrx,arr)
	}
}
