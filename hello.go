package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

var bank_amount int = 1000

// Package installer function

func install_packages() {
	var package_input string

	ubuntu_flag := true

	os_id := exec.Command("lsb_release", "-a")
	os_output, _ := os_id.Output()
	str_os_output := string(os_output[:])

	fmt.Println(str_os_output)

	contain := strings.Contains(str_os_output, "Ubuntu")

	if contain == ubuntu_flag {

		fmt.Println("Find Ubuntu distro")
		fmt.Print("Enter package name (divide by space)")
		fmt.Scan(&package_input)
		fmt.Println("Package list :", package_input)

	} else {
		fmt.Println("Find unknown distro")

	}

}

// Env value check functions

func check_user_input() string {

	var env_key string
	fmt.Print("Enter enviroment variable key for console :")
	fmt.Scan(&env_key)
	return env_key

}

func check_os_env(key string) string {
	fmt.Println("Trigger get env function")
	value := os.Getenv(key)

	return value
}

// Bank CRUD functions

func bank_crud_entrypoint() {

	var bank_choice int

	fmt.Println("Choose operation")
	fmt.Println(`
	1. Check bank amount
	2. Add money
	3. Withdraw money
	4. Exit
	`)
	fmt.Print("Your choice : ")
	fmt.Scan(&bank_choice)

	if bank_choice == 1 {
		check_bank_amount(bank_amount)
	} else if bank_choice == 2 {
		add_bank_amount(bank_amount)
	} else if bank_choice == 3 {
		withdraw_bank_amount(bank_amount)
	} else if bank_choice == 4 {
		os.Exit(0)
	}

}

func check_bank_amount(amount int) {
	fmt.Println("Your bank amount is : ", amount, "$")

}

func add_bank_amount(amount int) {
	var add_sum int
	fmt.Print("Enter cash amount :")
	fmt.Scan(&add_sum)
	amount += add_sum
	fmt.Println("Success account replenishment")
}

func withdraw_bank_amount(amount int) {
	var withdraw_sum int
	fmt.Print("Enter cash amount :")
	fmt.Scan(&withdraw_sum)
	amount -= withdraw_sum
	fmt.Println("Success withdrawal from account")
}

// Entrypoint

func main() {

	var choice int

	fmt.Println("Choose function to continue")
	fmt.Println(`
	1. Add packages to system
	2. Check env value at *nix system
	3. Bank simple CRUD	app
	4. Exit
	
	`)
	fmt.Print("Your choice : ")
	fmt.Scan(&choice)

	if choice == 1 {

		install_packages()

	} else if choice == 2 {

		input_key := check_user_input()
		input_value := check_os_env(input_key)
		fmt.Println("Your value :", input_value)

	} else if choice == 3 {

		bank_crud_entrypoint()

	} else if choice == 4 {

		os.Exit(0)
	}

}
