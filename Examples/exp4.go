package main

import "fmt"

func main4() {
	t := 0
	var k int
	fmt.Print("Sayı Giriniz:")
	fmt.Scanln(&k)
	for i :=0; i < k; i++ {
		t += i
	}
	fmt.Println("Toplam:", t)
}

// Kullanıcıdan alınan sayıya kadar olan sayıların toplamı.