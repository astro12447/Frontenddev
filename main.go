package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

func GetIdfile(ptrflag string, value string, description string) (*string, *string) {
	src := flag.String(ptrflag, value, description)
	dst := flag.String(ptrflag, value, description)
	flag.Parse()
	return src, dst
}
func ReadAllfile(name string) []byte {
	f, err := os.ReadFile(name)
	if err != nil {
		log.Fatalf("Error reading file...!%v", err)
	}
	return f
}
func CreateFile(Newfile string) {
	f, err := os.Create(Newfile)
	if err != nil {
		log.Fatalf("%v", f)
	}
	defer f.Close()
	fmt.Print("file was created...!")
}

// func IsUrl(pathName string) bool {

// }

func GetRequest(src *string, dst string) {
	bytes := ReadAllfile(*src)
	lines := strings.Split(string(bytes), "\n")

	for _, item := range lines {
		response, err := http.Get(item)
		if err != nil {
			panic(err)
		}
		defer response.Body.Close()
		body, err := io.ReadAll(response.Body)
		google, err := os.Create(*&dst)
		if err != nil {
			panic(err)
		}
		defer google.Close()
		r, err := google.WriteString(string(body))
		if err != nil {
			panic(err)
		}
		fmt.Print(r)
	}

}

func main() {

	f1 := flag.String("src", "", "")
	f2 := flag.String("dst", "", "")
	start := time.Now()
	flag.Parse()
	GetRequest(f1, *f2)
	end := time.Now()
	elapse := end.Sub(start)
	fmt.Println("Duration time elapse:", elapse)
}
