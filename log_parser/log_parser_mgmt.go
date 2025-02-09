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

	format_logs := strings.Split(string(auth_log_std), "\n")
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

			fmt.Println("Enter number of sliced strings")
			fmt.Scan(&head_count)

			format_logs_head_slice := format_logs[0:head_count]

			for i := range format_logs_head_slice {
				fmt.Println(format_logs_head_slice[i])
			}

		case 3:
			var tail_count int
			fmt_logs_range := len(format_logs)
			fmt.Println("Enter number of sliced strings")
			fmt.Scan(&tail_count)

			format_logs_tail_slice := format_logs[fmt_logs_range-tail_count-1 : fmt_logs_range]
			for i := range format_logs_tail_slice {
				fmt.Println(format_logs_tail_slice[i])
			}

		}

	}

}
