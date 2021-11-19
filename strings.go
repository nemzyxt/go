//Author : Nemuel Wainaina
//A simple program to showcase some operations on Strings

package main

import "fmt"

func main(){
    fmt.Println("Strings : \n")
    fmt.Println("String : Hello world!")
    
    //The length of the string
    fmt.Println("Length of string : ", len("Hello world!"))
    
    //accessing characters in the string, in this case the first one
    //it prints 72(H) since each character is represented as a byte which is an integer
    fmt.Println("First Character : ", "Hello world!"[0])
    
    //string concatenation
    fmt.Println("Hello " + "world!")
}