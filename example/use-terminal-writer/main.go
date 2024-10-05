package main

import (
	"github.com/codedoga/go-qrcode//writer/terminal"
	"github.com/codedoga/go-qrcode/v2"
)

func main() {
	qrc, err := qrcode.New("with_terminal")
	if err != nil {
		panic(err)
	}

	w := terminal.New()
	if err = qrc.Save(w); err != nil {
		panic(err)
	}

	//time.Sleep(5 * time.Second)
}
