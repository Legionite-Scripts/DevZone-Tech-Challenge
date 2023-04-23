package main

import (
	"fmt" //fmt package to print out statements
)

func main() {
	var purchasedItems = []int{} //empty array to show purchased items

	fmt.Println("Hello User!\nWelcome to Legionite Shopping App\nPlease input your name : ")
	var nameInput string
	fmt.Scanln(&nameInput)

	fmt.Print("How much would you like to fund your wallet with? : ")
	var cashInput int
	fmt.Scanln(&cashInput)

	for {
		fmt.Println("\nWelcome " + nameInput + " Please choose the item you wish to purchase : ")
		fmt.Println("1. Shoes ~ $300")
		fmt.Println("2. Mascara ~ $200")
		fmt.Println("3. Bread ~ $100")
		fmt.Println("4. Dell laptop ~ $800")
		fmt.Println("5. Pringles ~ $10")
		fmt.Println("6. Display Collective Price selected items items")
		fmt.Println("7. Pay for all items")
		fmt.Println("8. Exit App")
		fmt.Println("\nEnter Choice : ")
		var itemInput int
		fmt.Scanln(&itemInput)

		switch itemInput {
		case 1:
			fmt.Println("You have added shoes to the cart, Your bill has been recorded")
			purchasedItems = append(purchasedItems, 300)

		case 2:
			fmt.Println("You have added mascara to the cart, Your bill has been recorded")
			purchasedItems = append(purchasedItems, 200)

		case 3:
			fmt.Println("You have added Bread to the cart, Your bill has been recorded")
			purchasedItems = append(purchasedItems, 100)

		case 4:
			fmt.Println("You have added a Dell Laptop to the cart, Your bill has been recorded")
			purchasedItems = append(purchasedItems, 800)

		case 5:
			fmt.Println("You have added a can of Pringles to the cart, Your bill has been recorded")
			purchasedItems = append(purchasedItems, 10)

		case 6:
			sum := 0                             //initial value of sum
			for _, num := range purchasedItems { //for loop to print out sum of items 
				sum += num
			}
			fmt.Println(sum)
		case 7:

			sum := 0
			for _, num := range purchasedItems {
				sum += num
			}
			if cashInput >= sum {
				fmt.Print("You have successfully paid : ")
				fmt.Print(sum)
				fmt.Println("\nYour wallet balance is : ")
				fmt.Print(cashInput - sum)
			} else {
				fmt.Println("Dear user\n Your current Wallet balance is currently unable to afford this item")
			}
		case 8:
			fmt.Println("GoodBye!")
			return
			break
		default:
			fmt.Println("Incorrect Input\nPlease try again.")

			fmt.Println() // print statement to run the for loop continuously
		}
	}
}