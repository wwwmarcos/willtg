package commands

import (
	"os"
	"path"
	"strings"

	t "github.com/eptaccio/willtg/types"
	"github.com/fogleman/gg"
)

// WriteImage command
func WriteImage(text, ID string, imageConfig t.ImageConfig) (*string, error) {
	im, err := gg.LoadImage(imageConfig.ImagePath)
	if err != nil {
		return nil, err
	}

	textOnImage := strings.Replace(text, imageConfig.Command, "", 1)

	dc := gg.NewContext(imageConfig.Context.Width, imageConfig.Context.Height)
	dc.LoadFontFace("./files/font.ttf", imageConfig.FontSize)
	dc.DrawImage(im, 0, 0)
	dc.SetHexColor(imageConfig.Color)

	writeBoxWidth := float64(imageConfig.Context.Width / 2)
	writeboxHeight := float64(imageConfig.Context.Height / 2)
	lineSpacing := 1.5
	lineWidth := float64(imageConfig.Context.Width - 20)

	dc.DrawStringWrapped(
		textOnImage,
		writeBoxWidth,
		writeboxHeight,
		0.5,
		imageConfig.AY,
		lineWidth,
		lineSpacing,
		gg.AlignCenter)

	fileName := ID + ".png"
	filePath := path.Join(os.TempDir(), fileName)
	dc.SavePNG(filePath)

	return &filePath, nil
}
