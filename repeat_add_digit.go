package main

import "fmt"

func main() {
	result := Add(38)
	fmt.Println(result)

}

func Add(num int) int {
	var sum int = 0
	var res int = 0
	for num > 0 {
		res = num % 10
		sum = sum + res
		num = num / 10

	}
	if sum > 9 {
		sum = Add(sum)
	}
	return sum
}
