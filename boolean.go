//Author : Nemuel Wainaina
//This is a simple program that prints out the Truth Tables for &&, || and ! operators

package main

import "fmt"

func main(){
    fmt.Println("\t\tTruth Tables :")
    fmt.Println("--------------------")
    fmt.Println("AND Operator (&&) :")
    fmt.Println("    true && true : ", true&&true)
    fmt.Println("    true && false : ", true&&false)
    fmt.Println("    false && true : ", false&&true)
    fmt.Println("    false && false : ", false&&false)
    
    fmt.Println("\nOR Operator (||) :")
    fmt.Println("    true || true : ", true||true)
    fmt.Println("    true || false : ", true||false)
    fmt.Println("    false || true : ", false||true)
    fmt.Println("    false || false : ", false||false)
    
    fmt.Println("\nNOT Operator (!) :")
    fmt.Println("    !true : ", !true)
    fmt.Println("    !false : ", !false)
}