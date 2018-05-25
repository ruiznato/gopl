// Lissajous generates GIF animations of random Lissajous figures
package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"strings"
	"time"
)

var palette = []color.Color{
	color.Black,
	color.RGBA{0xff, 0x00, 0x00, 0xff},
	color.RGBA{0x00, 0xff, 0x00, 0xff},
	color.RGBA{0x00, 0x00, 0xff, 0xff},
}

type fields struct {
	Cycles  int
	Res     float64
	Size    int
	Nframes int
	Delay   int
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		lissajous(w, parseQuery(r.URL.Query()))
	})
	log.Fatal(http.ListenAndServe("localhost:8000", nil))

}

func parseQuery(query url.Values) fields {
	params := fields{
		Cycles:  5,
		Res:     0.001,
		Size:    100,
		Nframes: 64,
		Delay:   8,
	}

	s := reflect.ValueOf(&params).Elem()
	structType := s.Type()

	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		fieldName := strings.ToLower(structType.Field(i).Name)
		if f.Type().String() == "int" {
			if value, err := strconv.Atoi(query.Get(fieldName)); err == nil {
				f.SetInt(int64(value))
			}
		}
		if f.Type().String() == "float64" {
			if value, err := strconv.ParseFloat(query.Get(fieldName), 64); err == nil {
				f.SetFloat(value)
			}
		}
	}
	return params
}

func lissajous(out io.Writer, params fields) {
	cycles := params.Cycles
	res := params.Res
	size := params.Size
	nframes := params.Nframes
	delay := params.Delay

	// Log the parameters used
	log.Println("lissajous: rendering params:", params)

	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0

	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		lineColor := uint8(rand.Intn(len(palette)-1) + 1)
		for t := 0.0; t < float64(cycles)*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*float64(size)+0.5), size+int(y*float64(size)+0.5), lineColor)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}
