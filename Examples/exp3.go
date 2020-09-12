package main

import "fmt"

func main3() {
	var ilkSayi, ikinciSayi int
	fmt.Print("İlk Sayıyı Giriniz: ")
	fmt.Scanln(&ilkSayi)
	fmt.Print("İkinci Sayıyı Giriniz: ")
	fmt.Scanln(&ikinciSayi)

	t := ilkSayi + ikinciSayi
	if((ilkSayi > 100) && (ikinciSayi > 100)) {
		fmt.Print("Büyük değer girdiniz, 100'den büyük olamaz")
	} else {
		fmt.Print("Toplam: ", t)
	}	
}

// Kullanıcıdan alınan 100'den küçük 2 sayının toplamını verir.