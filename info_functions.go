package main

import "fmt"

func main_menu_loop_info() {
	fmt.Println("Choose function to continue")
	fmt.Println(`
1. Add packages to system
2. Check env value at *nix system
3. Bank simple CRUD	app
4. Exit
5. Test write to file
	`)

}
func bank_menu_info() {
	fmt.Println("Choose operation")
	fmt.Println(`
1. Check bank amount
2. Add money
3. Withdraw money
4. Exit
	`)

}
