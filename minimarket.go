package main

import (
	"fmt"
	"os"
	"time"
)

type products struct {
	name  string
	price float32
}

var basket []string

func main() {
	fmt.Println("\nWELCOME TO THE MINIMARKET!")
	functionality()

}

func functionality() {
	var inp int
	fmt.Println("\n~~~~MAIN MENU~~~~")
	fmt.Println("1) Check shopping basket")
	fmt.Println("2) Select product")
	fmt.Println("3) Pay")
	fmt.Println("4) Exit the store")
	fmt.Fscan(os.Stdin, &inp)
	if inp == 1 {
		fmt.Println(basket)
		functionality()

	} else if inp == 2 {
		choice()
	} else if inp == 3 {

		if basket != nil {
			writeToFile()
		} else {
			go fmt.Println("You can`t pay without products!")
			time.Sleep(1 * time.Millisecond)
			functionality()
		}
	} else if inp == 4 {
		go os.Exit(0)
		fmt.Println("Bye!")
	} else {
		go fmt.Println("Invalid value! Enter 1-4")
		time.Sleep(1 * time.Millisecond)
		functionality()
	}
}

func choice() {
	var inp int
	fmt.Println("Select a category (1-7):")
	fmt.Println("1) Drinkables")
	fmt.Println("2) Sweets")
	fmt.Println("3) Dairy")
	fmt.Println("4) Meat")
	fmt.Println("5) Bakery Products")
	fmt.Println("6) Other Products")
	fmt.Println("7) Back to the main menu")
	fmt.Fscan(os.Stdin, &inp)
	switch inp {
	case 1:
		{
			watercheck := "Water, 0.75"
			colacheck := "Cola, 1.5"
			lemonadecheck := "Lemonade, 1"
			water := products{"Water", 0.75}
			cola := products{"Cola", 1.5}
			lemonade := products{"Lemonade", 1}
			caseFunc(water, cola, lemonade, watercheck, colacheck, lemonadecheck)
		}
	case 2:
		{
			candiescheck := "Candies, 3"
			chocolatecheck := "Chocolate, 2"
			chipscheck := "Chips, 2"
			candies := products{"Candies", 3}
			chocolate := products{"Chocolate", 2}
			chips := products{"Chips", 2}
			caseFunc(candies, chocolate, chips, candiescheck, chocolatecheck, chipscheck)
		}
	case 3:
		{
			cheesecheck := "Cheese, 5"
			kefircheck := "Kefir, 1.5"
			milkcheck := "Milk, 1.5"
			cheese := products{"Cheese", 5}
			kefir := products{"Kefir", 1.5}
			milk := products{"Milk", 1.5}
			caseFunc(cheese, kefir, milk, cheesecheck, kefircheck, milkcheck)
		}
	case 4:
		{
			pizzacheck := "Steak, 8"
			breadcheck := "Pork, 2.5"
			baguettecheck := "Chicken, 3"
			steak := products{"Steak", 8}
			pork := products{"Pork", 2.5}
			chicken := products{"Chicken", 3}
			caseFunc(steak, pork, chicken, pizzacheck, breadcheck, baguettecheck)

		}
	case 5:
		{
			pizzacheck := "Pizza, 5"
			breadcheck := "Bread, 0.5"
			baguettecheck := "Baguette, 1"
			pizza := products{"Pizza", 5}
			bread := products{"Bread", 0.5}
			baguette := products{"Baguette", 1}
			caseFunc(pizza, bread, baguette, pizzacheck, breadcheck, baguettecheck)
		}

	case 6:
		{
			sugarcheck := "Sugar, 1"
			saltcheck := "Salt, 0.5"
			peppercheck := "Pepper, 1.5"
			sugar := products{"Sugar", 1}
			salt := products{"Salt", 0.5}
			pepper := products{"Pepper", 1.5}
			caseFunc(sugar, salt, pepper, sugarcheck, saltcheck, peppercheck)

		}
	case 7:
		functionality()
	default:
		go fmt.Println("Invalid value! Enter 1-7")
		time.Sleep(1 * time.Millisecond)
		functionality()
	}
}

func caseFunc(first products, second products, third products, firstcheck string, secondcheck string, thirdcheck string) {
	var str string
	var inp int
	fmt.Println("____________________________")
	fmt.Println("Number  Name  Price(dollars)")
	fmt.Println("1", first.name, first.price)
	fmt.Println("2", second.name, second.price)
	fmt.Println("3", third.name, third.price)
	fmt.Println("Would you like to add something to the shopping basket?(write yes or no)")
	fmt.Fscan(os.Stdin, &str)
	if str == "yes" || str == "YES" {
		fmt.Println("What exactly do you want (1-3)?")
		fmt.Fscan(os.Stdin, &inp)
		if inp == 1 {
			basket = append(basket, firstcheck)
			fmt.Println("\n", firstcheck, " added to shopping basket")
		} else if inp == 2 {
			basket = append(basket, secondcheck)
			fmt.Println("\n", secondcheck, " added to shopping basket")
		} else if inp == 3 {
			basket = append(basket, thirdcheck)
			fmt.Println("\n", thirdcheck, " added to shopping basket")
		} else {
			go fmt.Println("Invalid number! Enter 1-3")
			time.Sleep(1 * time.Millisecond)
		}
		functionality()
	} else if str == "no" || str == "NO" {
		functionality()
	} else {
		go fmt.Println("Invalid text! Just <<yes>> or <<not>>")
		time.Sleep(1 * time.Millisecond)
		functionality()
	}

}

func writeToFile() {
	for i := 0; i < len(basket); i++ {
		f, err := os.OpenFile("check.txt", os.O_APPEND|os.O_WRONLY, 0600)
		if err != nil {
			panic(err)
		}
		defer f.Close()
		f.WriteString("\n")
		f.WriteString(basket[i])

	}
	fmt.Println("\nCheck printed in check.txt")
}
