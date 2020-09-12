package main

import "fmt"

func main() {
	for x := 100; x >= 0; x= x - 5 {
		fmt.Println(x)
	} 
}

// 100'den 0'a kadar 5'er 5'er ekrana yazar. 