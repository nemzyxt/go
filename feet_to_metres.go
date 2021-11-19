//Author : Nemuel Wainaina
//This is a simple program that takes the length from the user, in feet, converts the length to metres, and prints it to the user on the console screen

package main

import "fmt"

func main(){
    fmt.Println("Enter the length in Feet : ")
    var feet float64
    fmt.Scanf("%f", &feet)
    var metres = feet * 0.3048
    fmt.Println("The length is ", metres, " metres")
}