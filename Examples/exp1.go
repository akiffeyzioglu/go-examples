package main

import "fmt"

func main1() {
	var sayi int
	fmt.Print("Sayı giriniz:")
	fmt.Scanln(&sayi)
	for i := 1; i <= sayi; i++ {
		if (sayi % i == 0) {
			fmt.Println(i)
		}
	}
}

// Girilen sayının tam bölenlerini verir. 