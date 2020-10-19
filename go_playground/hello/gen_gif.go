package main

import (
 "image"
 "image/color"
 "image/gif"
 "io"
 "math"
 "math/rand"
 "os"
)

const (
 BLACK_INDEX = 0
 WHITE_INDEX = 1
 )
var palette = []color.Color{color.Black, color.White}

func main() {
 //fmt.Println(WHITE_INDEX + 1)
 //f, _ := os.Open("haha.gif")
 //lissajous(os.Stdout)
 lissajous(os.Stdout)

}

func lissajous(out io.Writer) {
 const (
  cycles  = 5     // number of complete x oscillator revolutions
  res     = 0.001 // angular resolution
  size    = 100   // image canvas covers [-size..+size]
  nframes = 64    // number of animation frames
  delay   = 8     // delay between frames in 10ms units
 )
 freq := rand.Float64() * 3
 anim := gif.GIF{LoopCount:nframes}
 phase := 0.0
 for i := 0; i < nframes; i++ {
  rect := image.Rect(0, 0, 2*size+1, 2*size+1)
  img := image.NewPaletted(rect, palette)
  for t := 0.0; t < math.Pi*2 * cycles; t += res {
    x := math.Sin(t)
    y := math.Sin(t * freq + phase)
    img.SetColorIndex(size + int(x * size + 0.5), size + int(y * size + 0.5), BLACK_INDEX)
  }
  phase += 0.1
  anim.Delay = append(anim.Delay, delay)
  anim.Image = append(anim.Image, img)
 }
 _ = gif.EncodeAll(out, &anim)
}