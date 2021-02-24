package commands

import (
	"os"
	"path"
	"strings"

	t "github.com/eptaccio/willtg/types"
	"github.com/fogleman/gg"

	tb "gopkg.in/tucnak/telebot.v2"
)

// WriteImage command
func WriteImage(m *tb.Message, b *tb.Bot, imageConfig t.ImageConfig) {
	im, err := gg.LoadImage(imageConfig.ImagePath)
	if err != nil {
		b.Send(m.Sender, err.Error())
		return
	}

	textOnImage := strings.Replace(m.Text, imageConfig.Command, "", 1)

	dc := gg.NewContext(imageConfig.Context.Width, imageConfig.Context.Height)
	dc.LoadFontFace("./files/font.ttf", imageConfig.FontSize)
	dc.DrawImage(im, 0, 0)
	dc.SetHexColor(imageConfig.Color)

	writeBoxWidth := float64(imageConfig.Context.Width) / 2
	writeboxHeight := float64(imageConfig.Context.Height) / 2
	lineSpacing := 1.5
	lineWidth := float64(500)

	dc.DrawStringWrapped(
		textOnImage,
		writeBoxWidth,
		writeboxHeight,
		0.5, 1, // I will discover wtf is this
		lineWidth,
		lineSpacing,
		gg.AlignCenter)

	fileName := string(m.ID) + ".png"
	filePath := path.Join(os.TempDir(), fileName)
	dc.SavePNG(filePath)

	photo := &tb.Photo{File: tb.FromDisk(filePath)}
	_, err = b.SendAlbum(m.Chat, tb.Album{photo})
	if err != nil {
		b.Send(m.Sender, err.Error())
	}
}
