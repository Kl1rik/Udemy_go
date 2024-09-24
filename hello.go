package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"udemy.course.dev/v1/filework"
	"udemy.course.dev/v1/info"
	"udemy.course.dev/v1/struct_worker"
)

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

	for {
		var bank_amount int
		var bank_choice int
		filename := "bank.txt"
		info.Bank_menu_info()
		fmt.Print("Your choice : ")
		fmt.Scan(&bank_choice)

		if bank_choice == 1 {
			bank_amount = filework.Read_from_file(filename)
			check_bank_amount(bank_amount)

		} else if bank_choice == 2 {

			add_bank_amount(filename)

		} else if bank_choice == 3 {
			withdraw_bank_amount(filename)

		} else if bank_choice == 4 {
			os.Exit(0)
		} else {
			os.Exit(0)
		}

	}

}

func check_bank_amount(amount int) {
	fmt.Println("Your bank amount is : ", amount, "$")

}

func add_bank_amount(filename string) int {
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

func withdraw_bank_amount(filename string) int {
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

// Entrypoint

func main() {

	for {
		var choice int
		info.Main_menu_loop_info()
		fmt.Print("Your choice : ")
		fmt.Scan(&choice)
		if choice > 0 {
			switch choice {
			case 1:
				install_packages()

			case 2:
				input_key := check_user_input()
				input_value := check_os_env(input_key)
				fmt.Println("Your value :", input_value)

			case 3:
				bank_crud_entrypoint()
			case 4:
				os.Exit(0)
			case 5:
				err := os.WriteFile("bank.txt", []byte("1000"), 0644)

				if err != nil {
					fmt.Println("Unable to write file")
					log.Fatal(err)
				}
			case 6:
				pointer_test()

			case 7:
				struct_worker.Struct_input()
			}

		} else {
			fmt.Println("Incorrect value")
		}

	}

}
