package main

import (
    _  "fmt"
    "image"
    _ "image/draw"
    "image/png"
    "os"
)

func main(){
  rect := image.Rect(0,0,100,100)
  img := image.NewRGBA(rect)
  
  file, _ := os.Create("test.png")
  defer file.Close()
  
  png.Encode(file, img)
}
