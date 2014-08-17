package main

import "fmt"
import "os"

func main() {
	er := os.MkdirAll(os.Args[1], 0777)
	if er != nil {
		fmt.Println(er)
	}
}