package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
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
		line := 1
		if !number {
			dat, err := ioutil.ReadFile(file)
			check(err)
			fmt.Print(string(dat))
		} else {
			fp, err := os.Open(file)
			check(err)
			scanner := bufio.NewScanner(fp)
			for scanner.Scan() {
				fmt.Println(line, " ", scanner.Text())
				line = line + 1
			}
		}
	}
}
