package main

import (
	"bufio"
	"errors"
	"flag"
	"flati/internal/finfo"
	"flati/internal/util"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

func parseFlag() (in string, out string, err error) {
	input := flag.String("input", "input.txt", "input file with links")
	output := flag.String("output", "output.txt", "output file with parsed data")

	flag.Parse()

	if *input == "" {
		return "", "", errors.New("input file is required")
	}

	if *output == "" {
		return "", "", errors.New("output file is required")
	}

	return *input, *output, nil
}

func main() {
	input, output, err := parseFlag()
	if err != nil {
		log.Fatal(err)
	}

	inputFile, err := os.Open(input)
	if err != nil {
		log.Printf("File %s open error: %v", input, err)
		os.Exit(1)
	}
	defer inputFile.Close()

	fs := bufio.NewScanner(inputFile)
	fs.Split(bufio.ScanLines)

	var links []string
	for fs.Scan() {
		links = append(links, fs.Text())
	}
	fmt.Println(links)

	f, err := os.Create(output)
	if err != nil {
		log.Printf("File %s create error: %v", output, err)
		os.Exit(1)
	}
	defer f.Close()

	for _, link := range links {
		log.Printf("process link = %s\n", link)

		n, err := finfo.NewEntry(link)
		if err != nil {
			log.Printf("Parsing entry error: %v\n", err)
		}

		s := fmt.Sprintf("%s|%d|%s|Цена %s|Активен %s\n",
			n.Provider, n.ID, n.Address, n.Price, util.FormatBool(n.IsActive))
		_, err = f.WriteString(s)
		if err != nil {
			log.Printf("File %s write error: %v", output, err)
		}
		log.Printf("processed entry = %v\n", n)

		time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
	}

	err = f.Sync()
	if err != nil {
		log.Printf("File %s sync error: %v", output, err)
	}
}
