package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"skillfactory/34/calc"
	"skillfactory/34/pars"
)

func main() {
	var err error
	var data []byte
	var outputFile *os.File

	fmt.Println("Name of input file (file_name.txt):")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	inputFileName := scanner.Text()

	data, err = ioutil.ReadFile("files/" + inputFileName)
	variables := pars.Split(data)
	if err != nil {
		fmt.Println("No such file")
		log.Fatal(err)
	}
	fmt.Println("Name of output file (file_name.txt):")
	scanner.Scan()
	outputFileName := scanner.Text()
	outputFile, err = os.Create("files/" + outputFileName)
	if err != nil {
		log.Fatal(err)
	}

	wr := bufio.NewWriter(outputFile)
	for _, item := range variables {
		if len(item) < 4 {
			continue
		}
		result := calc.Calculate(item[1], item[3], item[2])
		if result == "" {
			continue
		}
		_, err = wr.Write([]byte(item[0] + result + "\n"))
		if err != nil {
			log.Fatal(err)
		}
	}

	err = wr.Flush()
	if err != nil {
		log.Fatal(err)
	}
}
