/*
 * The code is automatically generated by the Goland.
 * Copyright © Aquarian-Age. All Rights Reserved.
 * Licensed under MIT.
 */

package gz

import (
	"fmt"
	"github.com/starainrt/astro/calendar"
	"testing"
	"time"
)

func TestAliasZhi(t *testing.T) {
	tx := time.Now().Local()
	for i := 0; i <= 31; i++ {
		xt := time.Date(tx.Year(), tx.Month()+1, i, 0, 0, 0, 0, time.Local)
		jd := calendar.Date2JDE(xt)
		n := AliasZhi(jd)
		fmt.Println(xt.String()[:10], n)
	}
}

/*
2022-02-28 子
2022-03-01 丑
2022-03-02 寅
2022-03-03 卯
2022-03-04 辰
2022-03-05 巳
2022-03-06 午
2022-03-07 未
2022-03-08 申
2022-03-09 酉
2022-03-10 戌
2022-03-11 亥
2022-03-12 子
2022-03-13 丑
2022-03-14 寅
2022-03-15 卯
2022-03-16 辰
2022-03-17 巳
2022-03-18 午
2022-03-19 未
2022-03-20 申
2022-03-21 酉
2022-03-22 戌
2022-03-23 亥
2022-03-24 子
2022-03-25 丑
2022-03-26 寅
2022-03-27 卯
2022-03-28 辰
2022-03-29 巳
2022-03-30 午
2022-03-31 未
*/
