package main

import "fmt"

func main() {
	for x := 100; x >= 0; x -= 10 {
		fmt.Println(x)
	} 
}

// 100'den 0'a kadar 10'ar 10'ar ekrana yazar. 