package main

import (
    "fmt"
    "os"
)

func main() {
	//defer is the last function to run
    defer fmt.Println("!")

    os.Exit(3)
}