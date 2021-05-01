package main

import (
	"fmt"
	"os"
)

func int_to_string(arr []string) {
	var strarr []string
	var str string
	m := map[byte]string{'0': "Zero", '1': "One", '2': "Two", '3': "Three", '4': "Four", '5': "Five", '6': "Six", '7': "Seven", '8': "Eight", '9': "Nine"}
	for i := range arr {
		number := arr[i]
		for j := 0; j < len(number); j++ {
			str += m[number[j]]
		}
		strarr = append(strarr, str)
		str = ""
	}
	for i := range strarr {
		print(strarr[i])
		if i != len(strarr)-1 {
			print(",")
		}
	}
	fmt.Print("\n")
}

func main() {
	// scanner := bufio.NewScanner(os.Stdin)
	// scanner.Scan()
	// input := scanner.Text()
	// arr := strings.Split(input, " ")
	arr := os.Args[1:]
	// fmt.Println(arr)
	int_to_string(arr)

}
