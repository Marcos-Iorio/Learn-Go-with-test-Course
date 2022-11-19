package integers

import "fmt"

// The function `Add` takes two integers, `a` and `b`, and returns the sum of the two integers
func Add(a, b int) int {
	sum := a + b

	return sum
}

func main() {
	fmt.Println(Add(2, 2))
}
