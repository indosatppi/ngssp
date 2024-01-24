package main

import (
	"fmt"
	"os"

	"github.com/indosatppi/ngssp"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		panic("Please provide a filename")
	}

	fpath := args[1]
	f, err := os.Open(fpath)
	if err != nil {
		panic(err)
	}

	def, err := ngssp.NewNgsspPackageDefDecoder(f).Decode()
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", def)
}
