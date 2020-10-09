package commands

import (
	t "github.com/eptaccio/willtg/types"
)

// CommandMapping all commands
var CommandMapping = []t.ImageConfig{
	t.ImageConfig{
		Context: t.ImageContext{
			Height: 900,
			Width:  1200,
		},
		FontSize:  50,
		ImagePath: "./files/bc.jpg",
		Color:     "#FFFFFF",
		Command:   "/will",
	},
	t.ImageConfig{
		Context: t.ImageContext{
			Height: 720,
			Width:  1280,
		},
		FontSize:  50,
		ImagePath: "./files/michael-scott.jpg",
		Color:     "#FFFFFF",
		Command:   "/maico",
	},
	t.ImageConfig{
		Context: t.ImageContext{
			Height: 335,
			Width:  560,
		},
		FontSize:  35,
		ImagePath: "./files/dwight.jpg",
		Color:     "#FFFFFF",
		Command:   "/du",
	},
	t.ImageConfig{
		Context: t.ImageContext{
			Height: 1080,
			Width:  1920,
		},
		FontSize:  50,
		ImagePath: "./files/sasuske.jpg",
		Color:     "#FFFFFF",
		Command:   "/du",
	},
}
