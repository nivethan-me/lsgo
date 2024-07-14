package main

import (
	"fmt"
	"github.com/ryanuber/columnize"
	"log"
	"os"
	"strings"
)

var reset = "\033[0m"
var red = "\033[31m"
var green = "\033[32m"
var yellow = "\033[33m"
var blue = "\033[34m"
var magenta = "\033[35m"
var cyan = "\033[36m"
var white = "\033[97m"
var orange = red + yellow

func main() {
	files, err := os.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Println(files)

	filesPrinted := 0

	output := []string{}

	for _, file := range files {
		// Ignore dot files
		if strings.HasPrefix(file.Name(), ".") {
			continue
		}

		fileNameHasSuffix := func(suffix string) bool {
			return strings.HasSuffix(file.Name(), suffix)
		}

		var tempNameString string

		if file.IsDir() {
			tempNameString = fmt.Sprint(green + "󰉋 " + reset + file.Name() + "/")
		} else if fileNameHasSuffix(".go") {
			tempNameString = fmt.Sprint(blue + " " + reset + file.Name())
		} else if fileNameHasSuffix(".md") {
			tempNameString = fmt.Sprint(cyan + " " + reset + file.Name())
		} else if fileNameHasSuffix(".txt") {
			tempNameString = fmt.Sprint(cyan + " " + reset + file.Name())
		} else if fileNameHasSuffix(".lua") {
			tempNameString = fmt.Sprint(blue + " " + reset + file.Name())
		} else if fileNameHasSuffix(".sh") {
			tempNameString = fmt.Sprint(white + " " + reset + file.Name())
		} else if fileNameHasSuffix(".json") {
			tempNameString = fmt.Sprint(white + "󰘦 " + reset + file.Name())
		} else if fileNameHasSuffix(".java") {
			tempNameString = fmt.Sprint(orange + "󰬷 " + reset + file.Name())
		} else if fileNameHasSuffix(".mp3") {
			tempNameString = fmt.Sprint(white + " " + reset + file.Name())
		} else if fileNameHasSuffix(".py") {
			tempNameString = fmt.Sprint(yellow + " " + reset + file.Name())
		} else if fileNameHasSuffix(".css") {
			tempNameString = fmt.Sprint(blue + " " + reset + file.Name())
		} else if fileNameHasSuffix(".js") {
			tempNameString = fmt.Sprint(yellow + " " + reset + file.Name())
		} else if fileNameHasSuffix(".mp4") {
			tempNameString = fmt.Sprint(red + " " + reset + file.Name())
		} else if fileNameHasSuffix(".html") {
			tempNameString = fmt.Sprint(orange + " " + reset + file.Name())
		} else if fileNameHasSuffix(".pdf") {
			tempNameString = fmt.Sprint(white + " " + reset + file.Name())
		} else if fileNameHasSuffix(".docx") {
			tempNameString = fmt.Sprint(white + " " + reset + file.Name())
		} else if fileNameHasSuffix(".jsx") {
			tempNameString = fmt.Sprint(blue + "" + reset + file.Name())
		} else if fileNameHasSuffix(".tsx") {
			tempNameString = fmt.Sprint(blue + " " + reset + file.Name())
		} else {
			tempNameString = fmt.Sprint(blue + " " + reset + file.Name())
		}
		output = append(output, tempNameString)

		filesPrinted++
	}
	// fmt.Println(output)
	test := groupWithSeparator(output, "|")
	// fmt.Println(test)

	config := columnize.DefaultConfig()
	config.Glue = "\t\t"

	result2 := columnize.Format(test, config)
	fmt.Println(result2)
}

func groupWithSeparator(slice []string, separator string) []string {
	result := []string{}

	for i, item := range slice {
		if i%2 == 1 {
			result = append(result, fmt.Sprintf("%s %s %s", item, separator, slice[i-1]))
		}
	}
	if len(slice)%2 == 1 {
		result = append(result, slice[len(slice)-1])
	}
	return result
}
