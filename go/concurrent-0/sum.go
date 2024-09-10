package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

type sumStruct struct {
	sum  int64
	path string
}

// read a file from a filepath and return a slice of bytes
func readFile(filePath string) ([]byte, error) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Printf("Error reading file %s: %v", filePath, err)
		return nil, err
	}
	return data, nil
}

// sum all bytes of a file
func sum(path string, sumCh chan sumStruct) {
	data, err := readFile(path)
	if err != nil {
		fmt.Printf(err.Error())
		return
	}

	var _sum int64
	_sum = 0
	for _, b := range data {
		_sum += int64(b)
	}

	sumCh <- sumStruct{_sum, path}
}

// print the totalSum for all files and the files with equal sum
func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <file1> <file2> ...")
		return
	}

	var totalSum int64

	sumCh := make(chan sumStruct)

	sums := make(map[int64][]string)
	for _, path := range os.Args[1:] {
		go sum(path, sumCh)
	}

	for i := 0; i < len(os.Args[1:]) ; i++ {
		sumstruct := <-sumCh

		totalSum += sumstruct.sum

		sums[sumstruct.sum] = append(sums[sumstruct.sum], sumstruct.path)
	}

	fmt.Println(totalSum)

	for sum, files := range sums {
		if len(files) > 1 {
			fmt.Printf("Sum %d: %v\n", sum, files)
		}
	}

}
