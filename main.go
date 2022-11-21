package main

import (
	"fmt"
)

func main(){
	var userInput string
	var customerCart [] string
	desktopNum := 10
	laptopNum := 10
	cableNum := 100
	ssdNum := 20
	hdNum := 50

	for {
		fmt.Print("Hello, please select the item you wish to retrieve\nor if you wish to exit the system")

		fmt.Print("1 - Desktop\n2 - Laptop\n3 - SSD\n4- HD\n5 - Network cable\n6 - Exit\n")
		
		fmt.Scan(&userInput)
		
		switch userInput {
		case "1":
			customerCart = append(customerCart, "Desktop")
			desktopNum -= 1
			fmt.Printf("There are %v desktops remaining\n", desktopNum)
			fmt.Printf("The items in your cart are\n")
			for index,item := range customerCart {
				fmt.Printf("Index %v - %v\n",index + 1, item)
			}
		case "2":
			customerCart = append(customerCart, "Laptop")
			laptopNum -= 1
			fmt.Printf("There are %v laptops remaining\n", laptopNum)
			fmt.Printf("The items in your cart are\n")
			for index,item := range customerCart {
				fmt.Printf("Index %v - %v\n",index + 1, item)
			}
		case "3":
			customerCart = append(customerCart, "SSD")
			ssdNum -= 1
			fmt.Printf("There are %v SSD's remaining\n",ssdNum)
			fmt.Printf("The items in your cart are\n")
			
			for index,item := range customerCart {
				fmt.Printf("Index %v - %v\n",index + 1, item)
			}

		case "4":
			customerCart = append(customerCart, "HD")
			hdNum -= 1
			fmt.Printf("There are %v HD's remaining\n",hdNum)
			fmt.Printf("The items in your cart are\n")
			for index,item := range customerCart {
				fmt.Printf("Index %v - %v\n",index + 1, item)
			}
		case "5":
			customerCart = append(customerCart, "Network cable")
			cableNum -= 1
			fmt.Printf("There are %v network cables available\n", cableNum)
			fmt.Printf("The items in your cart are\n")
		case "6":
			fmt.Print("Exiting system\n")					
		default:
			fmt.Print("Entered option is not valid\n") 
		}

		//Breaks the infinite loop
		if(userInput == "6"){
			break
		}	

		//Checks if the values are valid so that they can be reduced
	}
}