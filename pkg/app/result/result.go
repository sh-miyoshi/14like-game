package result

import (
	"github.com/sh-miyoshi/14like-game/pkg/app/config"
	"github.com/sh-miyoshi/14like-game/pkg/dxlib"
)

type Result struct {
	hitCount   int
	fontHandle int
}

func NewResult() *Result {
	return &Result{
		fontHandle: dxlib.CreateFontToHandle(dxlib.CreateFontToHandleOption{
			Size: dxlib.Int32Ptr(48),
		}),
		hitCount: 0,
	}
}

func (r *Result) SetValues(hitCount int) {
	r.hitCount = hitCount
}

func (r *Result) Draw() {
	dxlib.DrawFormatString(130, 100, dxlib.GetColor(255, 255, 255), "お疲れ様でした!")
	dxlib.DrawFormatString(130, 130, dxlib.GetColor(255, 255, 255), "あなたの成績は・・・    被弾回数: %d回", r.hitCount)

	var str string
	if r.hitCount == 0 {
		str = "おめでとう！！！"
	} else {
		str = "ざんねん・・・"
	}
	ofs := dxlib.GetDrawStringWidth(str, len(str)) * 48 / 16
	dxlib.DrawStringToHandle(config.ScreenSizeX/2-ofs/2, 230, dxlib.GetColor(255, 255, 255), r.fontHandle, str)

	str = "再プレイはできないのでXボタンかEscキーで閉じて再度やり直してね"
	ofs = dxlib.GetDrawStringWidth(str, len(str))
	dxlib.DrawFormatString(config.ScreenSizeX/2-ofs/2, 330, dxlib.GetColor(255, 255, 255), str)
}

func (r *Result) Update() {
	// Nothing to do
}