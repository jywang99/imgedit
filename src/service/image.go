package service

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"
	"unicode/utf8"

	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/colornames"
	"golang.org/x/image/math/fixed"
)

var Colors = []color.Color{
    colornames.Black,
    colornames.White,
}
var needOffset = []rune{
    '。', '、',
}

func contains(s []rune, e rune) bool {
    for _, a := range s {
        if a == e {
            return true
        }
    }
    return false
}

func TategakiImage(font *truetype.Font, fontsize float64, fontcolor color.Color, text string) *image.RGBA {
    // Calculate the total width and height of the text
    bounds := font.Bounds(fixed.Int26_6(fontsize * 64))
    width := (bounds.Max.X - bounds.Min.X).Ceil()
    height := (bounds.Max.Y - bounds.Min.Y).Ceil()
    totalHeight := height * utf8.RuneCountInString(text)

    // Resize the image to accommodate the text
    img := image.NewRGBA(image.Rect(0, 0, width, totalHeight))
    draw.Draw(img, img.Bounds(), &image.Uniform{color.Transparent}, image.Point{}, draw.Src)

    // Create a context
    ctx := freetype.NewContext()
    ctx.SetFont(font)
    ctx.SetFontSize(fontsize)
    ctx.SetSrc(image.NewUniform(fontcolor))
    ctx.SetDst(img)
    ctx.SetClip(img.Bounds())

    // Draw the text
    x := 0
    y := height
    for _, c := range text {
        ctx.SetSrc(image.NewUniform(fontcolor))
        var pt fixed.Point26_6
        if !contains(needOffset, c) {
            pt = freetype.Pt(x, y)
        } else {
            pt = freetype.Pt(x+width/2, y-height/2)
        }
        ctx.DrawString(string(c), pt)
        y += height
    }

    return img
}

func SaveToPng(img *image.RGBA, path string) {
    // Save the image to a file
    outputFile, err := os.Create(path)
    if err != nil {
        panic(err)
    }
    defer outputFile.Close()
    if err := png.Encode(outputFile, img); err != nil {
        panic(err)
    }
}

