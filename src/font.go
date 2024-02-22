package main

import (
	"os"

	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
)

func loadFont(fontfile string) *truetype.Font {
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

