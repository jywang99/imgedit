package web

import (
	"bytes"
	"encoding/base64"
	"image/png"
	"imgedit/src/service"
	"net/http"
	"strconv"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
    renderTemplate(w, "index", nil)
}

type inputForm struct {
    text string
    fontIndex int
    fontSize float64
    colorIndex int
}

func generateHandler(w http.ResponseWriter, r *http.Request) {
    // GET
    if r.Method == http.MethodGet {
        renderTemplate(w, "input", nil)
        return
    }
    // other methods
    if r.Method != http.MethodPost {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }
    // POST
    fi, _ := strconv.Atoi(r.FormValue("fontIndex"))
    fs, _ := strconv.ParseFloat(r.FormValue("fontSize"), 64)
    ci, _ := strconv.Atoi(r.FormValue("colorIndex"))
    inputs := &inputForm{
        text: r.FormValue("text"),
        fontIndex: fi,
        fontSize: fs,
        colorIndex: ci,
    }
    font := service.LoadFont(service.Fonts[inputs.fontIndex].Filename)
    color := service.Colors[inputs.colorIndex]
    // generate, encode and render
    img := service.TategakiImage(font, inputs.fontSize, color, inputs.text)
    var imgBytes bytes.Buffer
    png.Encode(&imgBytes, img)
    imgBase64Str := base64.StdEncoding.EncodeToString(imgBytes.Bytes())
    renderTemplate(w, "output", imgBase64Str)
}
