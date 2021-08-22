package service

import (
	"0x_mt109/application/model"
)

type ScreenService struct{}

func NewScreenService() *ScreenService {
	return &ScreenService{}
}

func (s *ScreenService) RemoveBlackPixel(screen model.Screen) model.Screen {
	//fmt.Println(screen.String())
	for i := 1; i < screen.GetLimitRows(); i++ {
		for j := 1; j < screen.GetLimitCols(); j++ {
			pixel := screen.Pixels[i][j]
			if pixel == 0 {
				continue
			}
			screen.ResetReview()
			if screen.Point(i, j).HaveLink() && !touchEnd(screen.Pixels, i, j, false) {
				screen.Pixels[i][j] = 2
				continue
			} else {
				screen.Pixels[i][j] = -1
				continue
			}
		}
		//fmt.Printf("%d:\n%s\n", i, screen.String())
	}

	for i := 1; i < screen.GetLimitRows(); i++ {
		for j := 1; j < screen.GetLimitCols(); j++ {
			if screen.Pixels[i][j] == 2 {
				screen.Pixels[i][j] = 0
			}
			if screen.Pixels[i][j] == -1 {
				screen.Pixels[i][j] = 1
			}
		}
	}

	return screen
}

func touchEnd(matriz [][]int, x int, y int, touch bool) bool {
	//fmt.Printf("(%d,%d): touch(%v)\n", x, y, touch)
	if y > 0 && y < len(matriz[0])-1 && x > 0 && x < len(matriz)-1 {
		//Flag Revisado (3)
		matriz[x][y] = 3
		if matriz[x][y-1] == 1 {
			touch = touchEnd(matriz, x, y-1, true)
		} else if matriz[x][y+1] == 1 {
			touch = touchEnd(matriz, x, y+1, true)
		} else if matriz[x-1][y] == 1 {
			touch = touchEnd(matriz, x-1, y, true)
		} else if matriz[x+1][y] == 1 {
			touch = touchEnd(matriz, x+1, y, true)
		} else {
			touch = false
		}
	}
	return touch
}
