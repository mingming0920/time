package locale

import (
	"fmt"
	"testing"
)

func TestZHTranslator(t *testing.T) {
	trans := Translator{
		WeekMap:  ZH_CN.WeekMap,
		MonthMap: ZH_CN.MonthMap,
	}
	if trans.WT(1) != "星期一" {
		t.Error("test failed, ZH_CN translator")
	}
	fmt.Println(trans)
	fmt.Println(trans.WT(1))
}

func TestENTranslator(t *testing.T) {
	trans := Translator{
		WeekMap:  EN.WeekMap,
		MonthMap: EN.MonthMap,
	}
	if trans.WT(1) != "Monday" {
		t.Error("test failed, EN translator")
	}
	fmt.Println(trans)
	fmt.Println(trans.WT(1))
}
