package main

import (
  "github.com/MindTwister/llano"
  "flag"
)

func main() {
  address := flag.String("http","127.0.0.1:2020","Interface to listen on")
  default200 := flag.String("body","OK","Default response for /200")
  flag.Parse()
  llano.Standalone(*address,*default200)
}
