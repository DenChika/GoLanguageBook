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
	"strconv"
	"time"
)

var palette3 = []color.Color{color.White, color.Black}
var cycles, size, nframes, delay = 5, 100, 64, 8
var res = 0.001

const (
	whiteIndex3 = 0
	blackIndex3 = 1
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		parse(r)
		lissajous3(w)
	})
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func parse(r *http.Request) {
	q := r.URL.Query()
	if q.Has("cycles") {
		var err error
		cycles, err = strconv.Atoi(q["cycles"][0])
		if err != nil {
			panic(err)
		}
	}
	if q.Has("size") {
		var err error
		size, err = strconv.Atoi(q["size"][0])
		if err != nil {
			panic(err)
		}
	}
	if q.Has("nframes") {
		var err error
		nframes, err = strconv.Atoi(q["nframes"][0])
		if err != nil {
			panic(err)
		}
	}
	if q.Has("delay") {
		var err error
		delay, err = strconv.Atoi(q["delay"][0])
		if err != nil {
			panic(err)
		}
	}
	if q.Has("res") {
		var err error
		res, err = strconv.ParseFloat(q["res"][0], 64)
		if err != nil {
			panic(err)
		}
	}
}

func lissajous3(out io.Writer) {
	rand.Seed(time.Now().UTC().UnixNano())
	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette3)
		for t := 0.0; t < float64(cycles)*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*float64(size)+0.5), size+int(y*float64(size)+0.5), blackIndex3)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
