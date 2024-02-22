package service

import (
	"os"

	"github.com/adrg/sysfont"
	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
)

var Fonts = sysfont.NewFinder(nil).List()

func LoadFont(fontfile string) *truetype.Font {
    // Load font
    fontBytes, err := os.ReadFile(fontfile)
    if err != nil {
        panic(err)
    }
    font, err := freetype.ParseFont(fontBytes)
    if err != nil {
        panic(err)
    }
    return font
}

