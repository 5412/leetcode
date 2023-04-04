package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	intpus := strings.Split(input, " ")

	var obj = make(map[string]int)
	var max = 0

	for _, v := range intpus {
		// fmt.Println("input= ", v)
		if _, ok := obj[v]; !ok {
			max = max + 1
			obj[v] = max
		}
	}


	var total = 0
	for i := 0; i < len(intpus); i++ {
		total += obj[intpus[i]]
	}

	fmt.Println(total)
}
