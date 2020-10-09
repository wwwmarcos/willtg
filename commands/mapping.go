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
}
