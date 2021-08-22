package model

import (
	"0x_mt109/application/helper"
	"strconv"
	"strings"
)

type Screen struct {
	x int
	y int
	Pixels [][]int
}

func (s Screen) GetLimitRows() int {
	return len(s.Pixels) - 1
}

func (s Screen) GetLimitCols() int {
	return len(s.Pixels[0]) - 1
}

func (s Screen) Point(x int, y int) Screen {
	s.x = x
	s.y = y
	return s
}

func (s Screen) HaveLink() bool {
	if s.Pixels[s.x][s.y-1] == -1 ||
		s.Pixels[s.x][s.y+1] == -1 ||
		s.Pixels[s.x-1][s.y] == -1 ||
		s.Pixels[s.x+1][s.y] == -1 {
		return false
	}
	return s.Pixels[s.x][s.y-1] > helper.PixelOff ||
		s.Pixels[s.x][s.y+1] > helper.PixelOff ||
		s.Pixels[s.x-1][s.y] > helper.PixelOff ||
		s.Pixels[s.x+1][s.y] > helper.PixelOff
}

func (s Screen) ResetPixelsRevised() {
	for i := 1; i < s.GetLimitRows(); i++ {
		for j := 1; j < s.GetLimitCols(); j++ {
			if s.Pixels[i][j] == helper.PixelRevised {
				s.Pixels[i][j] = helper.PixelOn
			}
		}
	}
}

func (s Screen) String() string {
	str := ""
	for i := 0; i < len(s.Pixels); i++ {
		str = str + "[ "
		row := make([]string, 0)
		for j := 0; j < len(s.Pixels[0]); j++ {
			row = append(row, strconv.Itoa(s.Pixels[i][j]))
		}
		str = str + strings.Join(row, ", ")
		str = str + " ]\n"
	}
	return str
}