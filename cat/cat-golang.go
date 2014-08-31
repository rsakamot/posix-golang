package main

import (
	"flag"
	"fmt"
	"io/ioutil"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	var number bool
	flag.BoolVar(&number, "n", false, "a bool")
	flag.Parse()

	files := flag.Args()
	fmt.Println("files: ", files)

	fmt.Println("number flag: ", number)
	for n := range files {
		file := files[n]
		fmt.Println("file: ", file)
		dat, err := ioutil.ReadFile(file)
		check(err)
		fmt.Print(string(dat))
	}
}
