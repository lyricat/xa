/*
 * The code is automatically generated by the Goland.
 * Copyright © Aquarian-Age. All Rights Reserved.
 * Licensed under MIT
 */

package gz

import (
	"strings"
	"testing"
)

func TestLiuHe(t *testing.T) {
	ygz := "甲子"
	lh := "丑"
	s := LiuHe(ygz)
	if !strings.EqualFold(s, lh) {
		t.Errorf("func LiuHe(%s)=%s want:%s", ygz, s, lh)
	}
}
