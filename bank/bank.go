package bank

import (
	"fmt"
	"os"

	"udemy.course.dev/v1/filework"
	"udemy.course.dev/v1/info"
)

// Bank CRUD functions

func Bank_crud_entrypoint() {

	for {
		var bank_amount int
		var bank_choice int
		filename := "bank.txt"
		info.Bank_menu_info()
		fmt.Print("Your choice : ")
		fmt.Scan(&bank_choice)

		if bank_choice == 1 {
			bank_amount = filework.Read_from_file(filename)
			Check_bank_amount(bank_amount)

		} else if bank_choice == 2 {

			Add_bank_amount(filename)

		} else if bank_choice == 3 {
			Withdraw_bank_amount(filename)

		} else if bank_choice == 4 {
			os.Exit(0)
		} else {
			os.Exit(0)
		}

	}

}

func Check_bank_amount(amount int) {
	fmt.Println("Your bank amount is : ", amount, "$")

}

func Add_bank_amount(filename string) int {
	amount := filework.Read_from_file(filename)
	var add_sum int
	fmt.Print("Enter cash amount :")
	fmt.Scan(&add_sum)
	if add_sum > 0 {
		amount += add_sum
		fmt.Println(amount)
		filework.Clear_file(filename)
		filework.Write_to_file(filename, amount)
		fmt.Println("Success account replenishment")
	} else {
		fmt.Println("Incorrect money output")
	}

	return amount
}

func Withdraw_bank_amount(filename string) int {
	amount := filework.Read_from_file(filename)
	var withdraw_sum int
	fmt.Print("Enter cash amount :")
	fmt.Scan(&withdraw_sum)
	if withdraw_sum <= amount {
		amount -= withdraw_sum
		fmt.Println(amount)
		filework.Clear_file(filename)
		filework.Write_to_file(filename, amount)
		fmt.Println("Success withdrawal from account")

	} else {
		fmt.Println("Incorrect money output")
	}

	return amount
}
