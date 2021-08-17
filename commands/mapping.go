package commands

import (
	t "github.com/eptaccio/willtg/types"
)

// CommandMapping all commands
var CommandMapping = []t.ImageConfig{
	{
		Context: t.ImageContext{
			Height: 900,
			Width:  1200,
		},
		FontSize:  50,
		ImagePath: "./files/bc.jpg",
		Color:     "#FFFFFF",
		Command:   "/will",
		AY:        1,
	},
	{
		Context: t.ImageContext{
			Height: 720,
			Width:  1280,
		},
		FontSize:  50,
		ImagePath: "./files/michael-scott.jpg",
		Color:     "#FFFFFF",
		Command:   "/maico",
		AY:        1,
	},
	{
		Context: t.ImageContext{
			Height: 335,
			Width:  560,
		},
		FontSize:  35,
		ImagePath: "./files/dwight.jpg",
		Color:     "#FFFFFF",
		Command:   "/du",
		AY:        1,
	},
	{
		Context: t.ImageContext{
			Height: 1080,
			Width:  1920,
		},
		FontSize:  100,
		ImagePath: "./files/sasuske.jpg",
		Color:     "#FFFFFF",
		Command:   "/sasu",
		AY:        1,
	},
	{
		Context: t.ImageContext{
			Height: 381,
			Width:  680,
		},
		FontSize:  50,
		ImagePath: "./files/bolos.jpg",
		Color:     "#FFFFFF",
		Command:   "/bolo",
		AY:        1,
	},
	{
		Context: t.ImageContext{
			Height: 500,
			Width:  500,
		},
		FontSize:  20,
		ImagePath: "./files/dvd.jpg",
		Color:     "#000000",
		Command:   "/dvd",
		AY:        -5,
	},
}
