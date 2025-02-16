package info

import "fmt"

func Main_menu_loop_info() {
	fmt.Println("Choose function to continue :")
	fmt.Println(`
1. Add packages to system
2. Check env value at *nix system
3. Bank simple CRUD	app
4. Exit
5. Test write to file
6. Test pointer value
7. Work with note struct
8. Check auth.log file 
9. Work with maps
	`)

}
func Bank_menu_info() {
	fmt.Println("Choose operation :")
	fmt.Println(`
1. Check bank amount
2. Add money
3. Withdraw money
4. Exit
	`)

}

func Pointer_menu_info() {
	fmt.Println("Choose operation for pointer :")
	fmt.Println(`
1. Check pointer value
2. Add value
3. Withdraw value
	`)

}

func Slice_menu_info() {
	fmt.Println("Choose slice:")
	fmt.Println(`
1. Failed logs
2. Sddm logs
3. Systemd logs

	`)

}

func Log_menu_info() {
	fmt.Println("Choose option(1-3):")
	fmt.Println(`
1. Slice of logs
2. Head logs
3. Tail logs

	`)

}
