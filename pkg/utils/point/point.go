package point

import "fmt"

type Point struct {
	X int
	Y int
}

func (p Point) String() string {
	return fmt.Sprintf("(%d, %d)", p.X, p.Y)
}

func (p Point) Equal(a Point) bool {
	return p.X == a.X && p.Y == a.Y
}

// 点a, b間の距離の2乗を返す
func Distance2(a, b Point) int {
	return (a.X-b.X)*(a.X-b.X) + (a.Y-b.Y)*(a.Y-b.Y)
}
