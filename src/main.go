package main

import (
	"fmt"
	"image/color"

	"github.com/adrg/sysfont"
)

func enterText() string {
    var text string
    fmt.Printf("Enter text: ")
    fmt.Scan(&text)
    return text
}

func pickFontSize() float64 {
    var size float64
    var valid bool
    for !valid {
        fmt.Printf("Enter font size: ")
        fmt.Scan(&size)
        if size <= 0 || size > 100 {
            fmt.Println("Invalid size")
        } else {
            valid = true
        }
    }
    return size
}

func pickFont() *sysfont.Font {
    // get all fonts
    f := sysfont.NewFinder(nil)
    fonts := f.List()
    // print them
    fmt.Println("Fonts")
    for i, font := range fonts {
        fmt.Printf("%d: %s\n", i+1, font.Name)
    }
    // user selection
    fmt.Printf("Enter font number: ")
    var n int
    fmt.Scan(&n)
    return fonts[n-1]
}

func main() {
    text := enterText()
    fontSize := pickFontSize()
    outputFilePath := "/mnt/c/Users/junyi/Downloads/vertical_text.png"

    fnt := pickFont()
    font := loadFont(fnt.Filename)

    img := tategakiImage(font, fontSize, color.Black, text)
    saveToPng(img, outputFilePath)
}

