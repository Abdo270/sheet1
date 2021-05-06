
package main

import (
	"fmt"
)

func isPerfect(num uint64) {
	sum := calcDiv(num)
	if sum == num {
		fmt.Println(num)
	}
}

func calcDiv(num uint64) (sum uint64) {
	var i uint64
	for i = 1; i <= num / 2; i++ {
		if num % i == 0 {
			sum += i
		}
	}
	return
}


func main() {
	var start, end uint64
	start = 1
	end = 999999999
	for ; start <= end; start++ {
		isPerfect(start) 
		// go isPerfect(start) // concurrent execution
	}
}
