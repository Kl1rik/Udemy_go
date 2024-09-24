package struct_worker

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"
)

type Note struct {
	User string
	Text string
	Time time.Time
}

func Struct_input() {
	var username string
	var text_message string
	var time_stamp time.Time
	var struct_object Note
	fmt.Println("Enter username")
	fmt.Scan(&username)

	fmt.Println("Enter note")
	fmt.Scan(&text_message)

	fmt.Println("Enter timestamp(auto)")
	time_stamp = time.Now()
	fmt.Println(time_stamp)

	fmt.Println("Write to struct")
	struct_object = Note{username, text_message, time_stamp}
	fmt.Println(struct_object)

	fmt.Println("Convert struct for json")
	struct_json, err := json.MarshalIndent(struct_object, "", " ")
	_ = os.WriteFile("struct.json", struct_json, 0777)

	if nil != err {
		log.Fatal(err)
	}

	fmt.Println(struct_json)

}
