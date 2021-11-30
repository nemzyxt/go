//Author : Nemuel Wainaina
//The program takes a number as input from the user, and prints the corresponding day of the week based on the user's input

package main
import "fmt"

func main(){
	var dayOfWeek int
	fmt.Println("Enter a number between 1 and 7 : ")
	fmt.Scanf("%d", &dayOfWeek)
	switch dayOfWeek{
	case 1:
		fmt.Println("Today is Monday")
	case 2:
		fmt.Println("Today is Tuesday")
	case 3:
		fmt.Println("Today is Wednesday")
	case 4:
		fmt.Println("Today is Thursday")
	case 5:
		fmt.Println("Today is Friday")
	case 6:
		fmt.Println("Today is Saturday")
	case 7:
		fmt.Println("Today is Sunday")
	default:
		fmt.Println("Invalid Input")
	}
}