//Author : Nemuel Wainaina
//This is a simple program that takes inout from the user, which is temperature in Fahrenheit, converts it to Celsius, and prints back to the user

package main

import "fmt"

func main(){
    fmt.Println("Enter the temperature in Fahrenheit : ")
    var f_heit float64
    fmt.Scanf("%f", &f_heit)
    var celsius = (f_heit-32) * 5/9
    fmt.Println("The temperature is ", celsius, " degrees Celsius")
}