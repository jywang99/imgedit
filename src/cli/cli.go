package cli

import (
	"fmt"
	"image/color"
	"imgedit/src/service"

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
    // print them
    fmt.Println("Fonts")
    for i, font := range service.Fonts {
        fmt.Printf("%d: %s\n", i+1, font.Name)
    }
    // user selection
    fmt.Printf("Enter font number: ")
    var n int
    fmt.Scan(&n)
    return service.Fonts[n-1]
}

func RunCli() {
    text := enterText()
    fontSize := pickFontSize()
    outputFilePath := "/mnt/c/Users/junyi/Downloads/vertical_text.png"

    fnt := pickFont()
    font := service.LoadFont(fnt.Filename)

    img := service.TategakiImage(font, fontSize, color.Black, text)
    service.SaveToPng(img, outputFilePath)
}
