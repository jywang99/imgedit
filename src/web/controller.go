package web

import (
	"bytes"
	"encoding/base64"
	"image/color"
	"image/png"
	"imgedit/src/service"
	"net/http"
	"strconv"

	"github.com/adrg/sysfont"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
    renderTemplate(w, "index", nil)
}

type formOpts struct {
    Fonts *[]*sysfont.Font
    Colors *map[string]color.Color
}

type inputForm struct {
    text string
    fontIndex int
    fontSize float64
    colorName string
}

func (f *inputForm) validate() bool {
    l := len(f.text)
    if l < 1 || l > 100 {
        return false
    }
    if f.fontIndex < 0 || f.fontIndex >= len(service.Fonts) {
        return false
    }
    if f.fontSize <= 0 || f.fontSize > 100 {
        return false
    }
    if _, ok := service.Colors[f.colorName]; !ok {
        return false
    }
    return true
}

func newInputForm(text string, fontStr string, fontSizeStr string, colorName string) *inputForm {
    fontIndex, _ := strconv.Atoi(fontStr)
    fontSize, _ := strconv.ParseFloat(fontSizeStr, 64)
    return &inputForm{
        text: text,
        fontIndex: fontIndex,
        fontSize: fontSize,
        colorName: colorName,
    }
}

func generateHandler(w http.ResponseWriter, r *http.Request) {
    // GET
    if r.Method == http.MethodGet {
        renderTemplate(w, "input", &formOpts{
            Fonts: &service.Fonts,
            Colors: &service.Colors,
        })
        return
    }

    // other http methods
    if r.Method != http.MethodPost {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    // POST
    inputs := newInputForm(r.FormValue("text"), r.FormValue("font"), r.FormValue("fontSize"), r.FormValue("color"))
    if !inputs.validate() {
        http.Error(w, "Invalid input", http.StatusBadRequest)
        return
    }
    font := service.LoadFont(service.Fonts[inputs.fontIndex].Filename)
    color := service.Colors[inputs.colorName]
    // generate, encode and render
    img := service.TategakiImage(font, inputs.fontSize, color, inputs.text)
    var imgBytes bytes.Buffer
    png.Encode(&imgBytes, img)
    imgBase64Str := base64.StdEncoding.EncodeToString(imgBytes.Bytes())
    renderTemplate(w, "output", imgBase64Str)
}
