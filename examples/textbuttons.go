package main

import (
	"image/color"
	"log"
	"os"
	"os/signal"

	sdeck "github.com/dh1tw/streamdeck"
	"github.com/gobuffalo/packr"
	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
)

var monoFont *truetype.Font

func main() {

	fontBox := packr.NewBox("fonts")

	var err error

	// Load the font
	monoFont, err = freetype.ParseFont(fontBox.Bytes("mplus-1m-regular.ttf"))
	if err != nil {
		log.Panic(err)
	}

	lineLabel := sdeck.TextLine{
		Font:      monoFont,
		FontColor: color.RGBA{255, 255, 0, 0}, // Yellow
		FontSize:  22,
		PosX:      10,
		PosY:      5,
		Text:      "STATE",
	}

	linePressed := sdeck.TextLine{
		Font:      monoFont,
		FontColor: color.RGBA{255, 255, 255, 0}, // White
		FontSize:  14,
		PosX:      12,
		PosY:      30,
		Text:      "PRESSED",
	}

	lineReleased := sdeck.TextLine{
		Font:      monoFont,
		FontColor: color.RGBA{255, 0, 0, 0}, // Red
		FontSize:  14,
		PosX:      9,
		PosY:      30,
		Text:      "RELEASED",
	}

	pressedText := sdeck.TextButton{
		BgColor: color.RGBA{0, 0, 0, 0},
		Lines:   []sdeck.TextLine{lineLabel, linePressed},
	}

	releasedText := sdeck.TextButton{
		BgColor: color.RGBA{0, 0, 0, 0},
		Lines:   []sdeck.TextLine{lineLabel, lineReleased},
	}

	sd, err := sdeck.NewStreamDeck()
	if err != nil {
		log.Panic(err)
	}
	defer sdeck.ClearAllBtns()

	for i := 0; i < 15; i++ {
		sdeck.WriteText(i, releasedText)
	}

	btnEvtCb := func(btnIndex int, state sdeck.BtnState) {
		if state == sdeck.BtnPressed {
			sd.WriteText(btnIndex, pressedText)
		} else {
			sdeck.WriteText(btnIndex, releasedText)
		}
	}

	sdeck.SetBtnEventCb(btnEvtCb)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	<-c
}
