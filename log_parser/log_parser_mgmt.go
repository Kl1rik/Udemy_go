package log_parser

import (
	"fmt"
	"os"
	"strings"

	"udemy.course.dev/v1/info"
)

var err_arr []string
var sddm_arr []string
var systemd_arr []string
var filter_choice int

func Auth_log_parser() {

	fmt.Println("Enter function")

	file_path := "/var/log/auth.log"
	auth_log_std, err := os.ReadFile(file_path)

	if nil != err {
		fmt.Println(err)
	}

	fmt.Println("log file content")
	// fmt.Print(string(auth_log_std))
	//Task 1
	format_logs := strings.Split(string(auth_log_std), "\n")

	//Task 2
	//fmt.Println(format_logs[0])
	//slice_logs := format_logs[1:3]
	//fmt.Println(slice_logs)

	info.Log_menu_info()

	fmt.Scan(&filter_choice)

	if filter_choice > 0 {
		switch filter_choice {

		case 1:
			var slice_count int
			for i := range format_logs {

				if strings.Contains(format_logs[i], "failed") || strings.Contains(format_logs[i], "Failed") {
					err_arr = append(err_arr, format_logs[i])

				} else if strings.Contains(format_logs[i], "sddm-helper") {
					sddm_arr = append(sddm_arr, format_logs[i])
				} else if strings.Contains(format_logs[i], "systemd") {
					systemd_arr = append(systemd_arr, format_logs[i])
				}
			}

			info.Slice_menu_info()
			fmt.Scan(&slice_count)

			switch slice_count {
			case 1:
				for i := range err_arr {
					fmt.Println(err_arr[i])
				}
			case 2:
				for i := range sddm_arr {
					fmt.Println(sddm_arr[i])
				}
			case 3:
				for i := range systemd_arr {
					fmt.Println(systemd_arr[i])
				}

			}

		case 2:
			var head_count int

			//Task 3.1

			fmt.Println("Enter number of sliced strings")
			fmt.Scan(&head_count)

			format_logs_head_slice := format_logs[0:head_count]

			for i := range format_logs_head_slice {
				fmt.Println(format_logs_head_slice[i])
			}

		case 3:
			var tail_count int

			//Task 3.2

			fmt_logs_range := len(format_logs)
			fmt.Println("Enter number of sliced strings")
			fmt.Scan(&tail_count)

			format_logs_tail_slice := format_logs[fmt_logs_range-tail_count-1 : fmt_logs_range]

			//Task 4
			//len_logs := len(format_logs)
			//special_slice := format_logs[2]
			//special_slice := format_logs[len_logs - 1]
			//format_logs_tail_slice = append(format_logs_tail_slice,special_slice)

			//task 5-6 -//-
			for i := range format_logs_tail_slice {
				fmt.Println(format_logs_tail_slice[i])
			}
		case 4:
			check_struct_arr()
		}

	}

}

//Task 7

func check_struct_arr() {

	type course struct {
		title string
		price float64
		owner string
	}
	struct_arr := [2]course{{"Udemy GO", 2.99, "Alex"}, {"Udemy k8s", 10.99, "Tyler"}}
	fmt.Println(struct_arr[0])
}
