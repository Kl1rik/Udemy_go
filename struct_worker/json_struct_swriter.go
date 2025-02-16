package struct_worker

import (
	"bufio"
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

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter username")
	fmt.Scan(&username)

	fmt.Println("Enter note")
	text_message, _ = reader.ReadString('\n')

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

func Map_managment() {
	bucket_keys := map[string]string{"PROD_KEY": "s0m@ h9sh", "TEST_KEY": "6e$t"}
	fmt.Println(bucket_keys["PROD_KEY"])
	fmt.Println(bucket_keys)

	bucket_keys["STAGE_KEY"] = "st@g1n6"
	fmt.Println(bucket_keys)

	delete(bucket_keys, "TEST_KEY")
	fmt.Println(bucket_keys)
}
