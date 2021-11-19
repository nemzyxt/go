//Author : Nemuel Wainaina
//This is a Program to print the numbers 1-10 using a for loop

package main

import "fmt"

func main(){
    fmt.Println("The For Loop")
    
    //method one
    fmt.Println("Using Approach 1 : ")
    i:=1//initialize variable i
    for i<=10{
        fmt.Println(i)
        i++
    }
    
    //method two
    fmt.Println("Using Approach 2 : ")
    for i=1;i<=10;i++ {
        fmt.Println(i)
    }
    
}