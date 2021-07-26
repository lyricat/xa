/*
 * The code is automatically generated by the Goland.
 * Copyright © Aquarian-Age. All Rights Reserved.
 * Licensed under MIT
 */
package x

import (
	"strings"
	"time"
)

var (
	lnmap = map[string]string{
		"天哭": "流年天哭星入出命 主孝服官非口舌、疾病、灾厄 有喜免哭 无喜要防 忌正、四、七月",
		"福德": "流年福德星入命 主吉多凶少 官非口舌尽消除 有喜无喜俱美 忌二、五、八月小阻",
		"黑煞": "流年黑煞星入命 主官非口舌 刑克 丧服 有喜可押身 无喜当防 忌正、四、五月",
		"文昌": "流年文昌星入命 主有吉兆士者 名标金榜 庶人招财 女人有喜 忌二、五、八月",
		"火孛": "流年火孛星入命 主官非口舌 破财 刑克孝服 损六畜 小人侵害 有喜可防 忌二、八、九、十二月",
		"暗曜": "流年暗曜星入命 主刑克破财损六畜 防盗贼 有喜可免 无喜当保可免 忌二、五、六、十二月",
		"天扫": "流年天扫星入命 主有孝服 官非口舌疾病 有不测之祸 有喜可免 无喜当防 忌六、八、十二月",
		"太阴": "流年太阴星入命 主有喜事临门 家中百事俱吉 女人有喜定是男儿 无喜亦吉",
		"天喜": "流年天喜星入命 主有喜庆 家道吉祥四季得财 但主有血刃之灾 女人有喜是男儿 忌四、六、九、十二月",
		"无常": "流年无常星入命 主有不测之灾及刑克孝服 有喜可免 无喜须防 忌二、六、七、九、十二月",
		"龙德": "流年龙德得入命 主喜庆临门家中吉庆百事享通 纵有血刃之灾亦无防",
		"锦秀": "流年锦绣星入命 主有喜 事事遂意 女人有喜定是男儿 忌六、九、十二月",
	}
	lnmap1 = map[string]string{
		"太岁": "太岁当头照，诸神不敢当，自己有刑克，须防克亲人",
		"太阳": "太阳星吉照，求财得珠宝，阴阳须和合，福禄自然来",
		"丧门": "丧门入命来，内外孝服来，破财并刑克，灾难又重临",
		"太阴": "太阴入命来，见喜又见财，贵人相接引，万事必和谐",
		"官符": "官符入命来，祸患必无常，破财并外孝，祈保免灾殃",
		"岁破": "岁破入命来，孝服哭哀哀，父母有刑克，损妻又损财",
		"死符": "死符入命来，病疾主有灾，若无官符事，脓血积生害",
		"龙德": "龙德贵人生，发财旺人丁，小灾不为害，家宅最光明",
		"白虎": "白虎星辰赛，春夏防灾厄，秋冬平平过，祈保无邪惑",
		"福星": "福星入命游，一岁进田牛，添丁又进契，凡事吉安康",
		"吊客": "吊客入命凶，内外孝服重，破财并损口，祈保免灾殃",
		"病符": "病符入命来，疾病必沾身，时运多颠倒，要好必求神",
	}
)

// 流年表 传入出生年
func LiuNianBiao(by int) (string, string) {
	x := time.Now().Year() - by + 1                                                         //虚岁＝今年－出生年+1
	arr := []string{"天哭", "福德", "黑煞", "文昌", "火孛", "暗曜", "天扫", "太阴", "天喜", "无常", "龙德", "锦绣"} //流年表
	arr1 := []string{"太岁", "太阳", "丧门", "太阴", "官符", "岁破", "死符", "龙德", "白虎", "福星", "吊客", "病符"}
	mod := x % 12
	var lns, lns1 string
	for i := 0; i < len(arr); i++ {
		if i == mod-1 {
			lns = arr[i]
			lns1 = arr1[i]
			break
		}
	}
	for k, v := range lnmap {
		if strings.EqualFold(k, lns) {
			lns = v
			break
		}
	}
	for k1, v1 := range lnmap1 {
		if strings.EqualFold(k1, lns1) {
			lns1 = v1
			break
		}
	}
	return lns, lns1
}
