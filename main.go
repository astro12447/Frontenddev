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

func GetDirectoryFromTerminal(src string, dst string) (string, string) {
	var source *string
	var destination *string
	source = flag.String(src, "None", "")
	destination = flag.String(dst, "None", "")
	flag.Parse()
	return *source, *destination
}
func ReadDataFromfile(filename string) string {
	//read file lines one by one line in memory
	f, err := os.Open(filename)
	if err != nil {
		fmt.Printf("anable to Open File: %v", err)
	}
	defer f.Close()
	buf := make([]byte, 1024)
	var FileContent string
	for {
		n, err := f.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Print(err)
		}
		if n > 0 {

			FileContent = string(buf[:n])
		}
	}
	return FileContent
}
func Createfile(filename string) *os.File {
	file, err := os.Create(filename)
	if err != nil {
		log.Fatalf("%v", file)
	}
	fmt.Printf(file.Name(), " Was been created succesfully...")
	fmt.Println()
	return file
}
func RequestFromClient(src string, destination string) {
	datafile := ReadDataFromfile(src)
	lines := strings.Split(datafile, "\n")
	for _, item := range lines {
		client := http.Client{}
		resp, err := client.Get(string(item))
		if err != nil {
			continue
		}
		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Print(err)
		}
		f := Createfile(destination)
		defer f.Close()
		f.WriteString(string(body))
		defer f.Close()
		fmt.Println(f.Name(), " Was Writed Succesfully...")
	}
}

func main() {
	start := time.Now()
	srcflag := "src"
	dstflag := "dst"
	s1, s2 := GetDirectoryFromTerminal(srcflag, dstflag)
	if s1 == "None" || s2 == "None" || s1 == "" || s2 == "" {
		fmt.Println("->Introduce correct Command line:(--src=./file.txt  --dst=./google.txt)")
	} else {
		RequestFromClient(s1, s2)
	}
	end := time.Now()
	elapse := end.Sub(start)
	fmt.Println("Duration time elapse:", elapse)
}
