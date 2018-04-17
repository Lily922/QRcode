package main

import (
	"github.com/skip2/go-qrcode"
)

func main() {
	qrcode.WriteFile("https://github.com/LilyFaFa", qrcode.Medium, 256, "./blog_qrcode.png")
}
