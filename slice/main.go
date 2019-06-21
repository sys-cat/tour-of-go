package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	p := make([][]uint8, dy)
	for i := range p {
		p[i] = make([]uint8, dx)
	}

	for y, row := range p { //縦のラインの描画
		for x := range row { // 横のラインの描画
			row[x] = uint8((x + y)/2) // x * y, x + y など式の変化で画像パターンが変わる
			// X軸を移動させながら値を置くことで描画を行っている。インクジェットプリンタの出力方法に近い。
		}
	}

	return p
}

func main() {
	pic.Show(Pic)
}