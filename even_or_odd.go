//Author : Nemuel Wainaina
//This is a simple Go program that loops from 1 to 10 and prints whether the number is even or odd for each of them

package main

import "fmt"

func main(){
	for i:=1;i<=10;i++{
		if i%2==0{
			fmt.Println(i,"even")
		}
		else{
			fmt.Println(i,"odd")
		}
	}
}