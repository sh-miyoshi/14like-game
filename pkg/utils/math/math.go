package math

import (
	"math"

	"github.com/sh-miyoshi/14like-game/pkg/utils/point"
)

const (
	// パッケージの読み込みを簡単化するためここにも定義しておく
	Pi = math.Pi
)

// Rotate関数は、(bx, by)を中心に(x, y)をangle度回転させた座標を返します。
func Rotate(base, pos point.Point, angle float64) point.Point {
	x := float64(pos.X) - float64(base.X)
	y := float64(pos.Y) - float64(base.Y)
	rx := math.Cos(angle)*x - math.Sin(angle)*y
	ry := math.Sin(angle)*x + math.Cos(angle)*y
	return point.Point{
		X: base.X + int(rx),
		Y: base.Y + int(ry),
	}
}

func MountainIndex(i, max int) int {
	if i >= max/2 {
		return max - i - 1
	} else {
		return i
	}
}
