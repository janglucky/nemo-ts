package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main()  {
	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	fmt.Println(dir)
 	fmt.Println("hello world")
}