package main

import "fmt"

func main2() {
	var sayi int 
	fmt.Print("Sayi giriniz:")
	fmt.Scanln(&sayi)
	
	for x := 2; x < sayi; x++ {
		if (sayi % x  == 0) {
			fmt.Println("Asal Degildir")
			break
		} else {
			fmt.Println("Asaldir..")
			break
		}
	}
}

// Kullanıcıdan alınan sayının asallık durumuna bakılır.