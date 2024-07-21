package main

import (
	"fmt"
	"github.com/otiai10/opengraph/v2"
)

func main() {
	og, err := opengraph.Fetch("https://www.youtube.com/watch?v=n0o5utReXXw")
	fmt.Println("title:", og.Title)
	fmt.Println("type:", og.Type)
	fmt.Println("type:", og.Image[0].URL)
	fmt.Println("error:", err)
}