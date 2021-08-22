package service

import "0x_mt109/application/model"

type IScreenService interface {
	RemoveBlackPixel(screen model.Screen) model.Screen
}