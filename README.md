# Golang Exercises  


## Table of contents
* [General info](#general-info)
* [Setup](#setup)
* [Inspiration](#inspiration)
* [Contact](#contact)

## General info
This repository includes the exercises I did while learning go. I will add as I learn.


## Setup
- First Setup your Go Enviroment [here](https://golang.org/dl/)
- Open Examples Folder
- Build Code 
```sh 
go build exp.go
```
- Run It 
```sh 
./exp
```

## Code Examples

## [For Loop Example]((https://github.com/akiffeyzioglu/go-examples/blob/master/Examples/exp.go))
```go
package main

import "fmt"

func main() {
	for x := 100; x >= 0; x -= 10 {
		fmt.Println(x)
	} 
}
// 100'den 0'a kadar 10'ar 10'ar ekrana yazar. 
```
## [For Loop Example-2](https://github.com/akiffeyzioglu/go-examples/blob/master/Examples/exp1.go)
```go
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
```

## Contact
Created by [@akiffeyzioğlu](https://github.com/akiffeyzioglu) - [@muratmirgun](https://github.com/muratmirgun) - feel free to contact us!