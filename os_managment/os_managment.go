package os_managment

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// Package installer function

func Install_packages() {
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

func Check_user_input() string {

	var env_key string
	fmt.Print("Enter enviroment variable key for console :")
	fmt.Scan(&env_key)
	return env_key

}

func Check_os_env(key string) string {
	fmt.Println("Trigger get env function")
	value := os.Getenv(key)

	return value
}
