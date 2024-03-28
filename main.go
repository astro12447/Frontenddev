package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {

	start := time.Now()
	source := flag.String("src", "", "sources urls")
	dest := flag.String("dst", "", "urls destination")
	flag.Parse()

	urls, err := ioutil.ReadFile(*source)
	if err != nil {
		log.Fatalf("Error reading file...%v!", *source)
	}
	lines := strings.Split(string(urls), "\n")

	for _, item := range lines {
		client := http.Client{}
		res, err := client.Get(string(item))
		if err != nil {
			continue
		} else {
			defer res.Body.Close()
			body, err := io.ReadAll(res.Body)
			if err != nil {
				log.Fatal(err)
			}
			google, err := os.Create(*dest)
			if err != nil {
				log.Fatalf("Error creating file...!%v", google)
			}
			defer google.Close()

			_, err1 := google.WriteString(string(body))

			if err1 != nil {
				log.Fatal(err1)
			}
		}
	}

	end := time.Now()
	elapse := end.Sub(start)

	fmt.Println("Duration time elapse:", elapse)
}
