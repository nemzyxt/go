//Author : Nemuel Wainaina
//This program takes age as input from the user and tells them whether or not they are eligible to vote

package main
import "fmt"

func main(){
	var age int
	fmt.Println("Enter your age : ")
	fmt.Scanf("%d", &age)
	if(age>=18){
		fmt.Println("You are eligible to vote\n")
	} else{
		fmt.Println("You are too young to vote")
	}
}