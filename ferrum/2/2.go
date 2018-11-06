package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	Http200  = "200"
	HttpGet  = `"GET`
	GifUpper = ".GIF"
	GifLower = ".gif"
)

// input  GET /shuttle/countdown/video/livevideo.gif HTTP/1.0
func getResource(input string) string {
	layers := strings.Split(input, "/")
	res := layers[len(layers)-1]
	return res
}

func main() {
	// read the string filename
	var filename string
	fmt.Scanf("%s", &filename)
	// read file
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// output filename as gifs_filename
	outputFile, err := os.Create("gifs_" + filename)
	if err != nil {
		log.Fatal(err)
	}
	defer outputFile.Close()

	w := bufio.NewWriter(outputFile)
	dup := map[string]bool{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		strs := strings.Split(line, " ")
		// split line with " "
		httpMethod := strs[5]
		code := strs[8]
		httpReq := strs[6]
		fmt.Println(httpReq)
		fmt.Println(code)
		// check file type, HTTP GET and response == 200
		if code != Http200 {
			continue
		}
		if httpMethod != HttpGet {
			continue
		}
		resourceName := getResource(httpReq)
		fmt.Println("res:", resourceName)
		if strings.HasSuffix(resourceName, GifUpper) || strings.HasSuffix(resourceName, GifLower) {
			if _, ok := dup[resourceName]; !ok {
				fmt.Println(resourceName)
				w.WriteString(resourceName + "\n")
				dup[resourceName] = true
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	w.Flush()
}
