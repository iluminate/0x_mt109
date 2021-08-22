package service

import (
	"0x_mt109/application/helper"
	"0x_mt109/application/model"
)

type ScreenService struct{}

func NewScreenService() *ScreenService {
	return &ScreenService{}
}

func (s *ScreenService) RemoveBlackPixel(screen model.Screen) model.Screen {
	for i := 1; i < screen.GetLimitRows(); i++ {
		for j := 1; j < screen.GetLimitCols(); j++ {
			pixel := screen.Pixels[i][j]
			if pixel == helper.PixelOff {
				continue
			}
			screen.ResetPixelsRevised()
			if screen.Point(i, j).HaveLink() && !touchEnd(screen.Pixels, i, j, false) {
				screen.Pixels[i][j] = helper.PixelFading
			} else {
				screen.Pixels[i][j] = helper.PixelTurning
			}
		}
	}
	for i := 1; i < screen.GetLimitRows(); i++ {
		for j := 1; j < screen.GetLimitCols(); j++ {
			if screen.Pixels[i][j] == helper.PixelFading {
				screen.Pixels[i][j] = helper.PixelOff
			}
			if screen.Pixels[i][j] == helper.PixelTurning {
				screen.Pixels[i][j] = helper.PixelOn
			}
		}
	}
	return screen
}

func touchEnd(matrizInput [][]int, x int, y int, touch bool) bool {
	if y > 0 && y < len(matrizInput[0])-1 && x > 0 && x < len(matrizInput)-1 {
		matrizInput[x][y] = helper.PixelRevised
		if matrizInput[x][y-1] == helper.PixelOn {
			touch = touchEnd(matrizInput, x, y-1, true)
		} else if matrizInput[x][y+1] == helper.PixelOn {
			touch = touchEnd(matrizInput, x, y+1, true)
		} else if matrizInput[x-1][y] == helper.PixelOn {
			touch = touchEnd(matrizInput, x-1, y, true)
		} else if matrizInput[x+1][y] == helper.PixelOn {
			touch = touchEnd(matrizInput, x+1, y, true)
		} else {
			touch = false
		}
	}
	return touch
}
