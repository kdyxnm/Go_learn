package main

// package name is the same as the last element of the import path
import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("My favorate number is", rand.Intn(10))
}
