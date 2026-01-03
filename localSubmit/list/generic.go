package main

import (
	"os"
)

const filePath = "/Users/soas/Work/templateWork/proj/go/goLearn/localSubmit/inOut"

func IOES() {

    in, err := os.Open(filePath+"/input.txt")
	if err != nil {
		panic(err)
	}
	out, err := os.Create(filePath+"/output.txt")
	if err != nil {
		panic(err)
	}

	os.Stdin = in
	os.Stdout = out
}

