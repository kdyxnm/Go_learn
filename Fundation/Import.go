package main

// We can also write multiple import statements like this:
// import "fmt"
// import "math"
// But it is not recommended.
import (
	"fmt"
	"math"
)

func main() {
	// If we want to use fmt like %g, wes should use Printf instead of Println
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))
}
