package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	var help bool
	var inputPath, outputPath string

	flag.BoolVar(&help, "h", false, "tampilkan cara menggunakan")
	flag.BoolVar(&help, "help", false, "tampilkan cara menggunakan")
	flag.StringVar(&inputPath, "i", "", "lokasi file JSON sebagai input")
	flag.StringVar(&inputPath, "input", "", "lokasi file JSON sebagai input")
	flag.StringVar(&outputPath, "o", "", "Lokasi file JSON sebagai output")
	flag.StringVar(&outputPath, "output", "", "Lokasi file JSON sebagai output")

	flag.Parse()
	if help || inputPath == "" || outputPath == "" {
		printUsage()
		//fmt.Println("mock data -i input.json -o output.json")
		os.Exit(0)
	}
	// if inputPath == "" {
	// 	fmt.Println("input wajib diisi")
	// 	os.Exit(0)
	// }
	// if outputPath == "" {
	// 	fmt.Println("output wajib diisi")
	// 	os.Exit(0)
	// }

}

func printUsage() {
	fmt.Println("Usage: mockdata [-i | --input] <input file> [-o | --output] <output file>")
	fmt.Println("-i --input: File input berupa JSON sebagai template")
	fmt.Println("-o --output: File input berupa JSON sebagai hasil")
}
