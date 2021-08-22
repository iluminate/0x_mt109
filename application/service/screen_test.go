package service

import (
	"0x_mt109/application/model"
	"testing"
)

var service = NewScreenService()

func TestRemoveBlackPixelCase1(t *testing.T)  {
	screen := model.Screen{
		Pixels: [][]int{
			{ 1, 0, 1, 0, 0 },
			{ 0, 1, 0, 1, 1 },
			{ 0, 1, 0, 1, 0 },
			{ 0, 1, 0, 1, 0 },
			{ 0, 0, 0, 0, 1 },
		},
	}
	expect := model.Screen{
		Pixels: [][]int{
			{ 1, 0, 1, 0, 0 },
			{ 0, 0, 0, 1, 1 },
			{ 0, 0, 0, 1, 0 },
			{ 0, 0, 0, 1, 0 },
			{ 0, 0, 0, 0, 1 },
		},
	}
	result := service.RemoveBlackPixel(screen)
	if expect.String() != result.String() {
		t.Errorf("AssertionError: \nExpected :%s\nActual   :%s", expect.String(), result.String())
	}
}

func TestRemoveBlackPixelCase2(t *testing.T)  {
	screen := model.Screen{
		Pixels: [][]int{
			{ 1, 0, 0, 0, 0, 0 },
			{ 0, 1, 0, 1, 1, 1 },
			{ 0, 0, 1, 0, 1, 0 },
			{ 1, 1, 0, 0, 1, 0 },
			{ 1, 0, 1, 1, 0, 0 },
			{ 1, 0, 0, 0, 0, 1 },
		},
	}
	expect := model.Screen{
		Pixels: [][]int{
			{ 1, 0, 0, 0, 0, 0 },
			{ 0, 1, 0, 1, 1, 1 },
			{ 0, 0, 1, 0, 1, 0 },
			{ 1, 1, 0, 0, 1, 0 },
			{ 1, 0, 0, 0, 0, 0 },
			{ 1, 0, 0, 0, 0, 1 },

		},
	}
	result := service.RemoveBlackPixel(screen)
	if expect.String() != result.String() {
		t.Errorf("AssertionError: \nExpected :%s\nActual   :%s", expect.String(), result.String())
	}
}

func TestRemoveBlackPixelCase3(t *testing.T)  {
	screen := model.Screen{
		Pixels: [][]int{
			{ 1, 0, 0, 0, 0, 0 },
			{ 0, 1, 0, 1, 1, 1 },
			{ 0, 1, 0, 1, 1, 0 },
			{ 0, 1, 1, 0, 1, 0 },
			{ 0, 1, 0, 1, 0, 0 },
			{ 0, 0, 0, 0, 0, 1 },
		},
	}
	expect := model.Screen{
		Pixels: [][]int{
			{ 1, 0, 0, 0, 0, 0 },
			{ 0, 0, 0, 1, 1, 1 },
			{ 0, 0, 0, 1, 1, 0 },
			{ 0, 0, 0, 0, 1, 0 },
			{ 0, 0, 0, 1, 0, 0 },
			{ 0, 0, 0, 0, 0, 1 },
		},
	}
	result := service.RemoveBlackPixel(screen)
	if expect.String() != result.String() {
		t.Errorf("AssertionError: \nExpected :%s\nActual   :%s", expect.String(), result.String())
	}
}

func TestRemoveBlackPixelCase4(t *testing.T)  {
	screen := model.Screen{
		Pixels: [][]int{
			{ 1, 0, 0, 0, 0, 0 },
			{ 0, 1, 0, 1, 1, 1 },
			{ 0, 1, 0, 1, 1, 0 },
			{ 0, 1, 0, 0, 1, 0 },
			{ 0, 1, 1, 1, 1, 0 },
			{ 0, 0, 0, 0, 0, 1 },
		},
	}
	expect := model.Screen{
		Pixels: [][]int{
			{ 1, 0, 0, 0, 0, 0 },
			{ 0, 1, 0, 1, 1, 1 },
			{ 0, 1, 0, 1, 1, 0 },
			{ 0, 1, 0, 0, 1, 0 },
			{ 0, 1, 1, 1, 1, 0 },
			{ 0, 0, 0, 0, 0, 1 },
		},
	}
	result := service.RemoveBlackPixel(screen)
	if expect.String() != result.String() {
		t.Errorf("AssertionError: \nExpected :%s\nActual   :%s", expect.String(), result.String())
	}
}