package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

func FileAsByte(path string) []byte {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println("Possibly wrong path!! : ", err)
		time.Sleep(3 * time.Second)
		return nil
	}
	fmt.Print(bytes)
	return bytes
}

func FileUtils(path_init string) (string, string) {
	var last_slash int
	for i, chars := range path_init {
		if chars == '/' {
			last_slash = i
		}
	}
	fullname := path_init[last_slash+1:]
	var last_dot int
	for i, chars := range fullname {
		if chars == '.' {
			last_dot = i
		}
	}
	res_name := fullname[:last_dot]
	res_type := fullname[last_dot:]
	return res_type, res_name
}

func FileDownloader(file_id int) {
	copy_nodes := msg_list.head
	for copy_nodes != nil {
		var msgstruct MsgStruct
		var json_str = copy_nodes.Key
		marshal_err := json.Unmarshal([]byte(json_str), &msgstruct)
		if marshal_err != nil {
			fmt.Println("Something went wrong,Donwload er.")
			time.Sleep(3 * time.Second)
			return
		}
		if msgstruct.Is_file {
			if copy_nodes.File_id == file_id {
				home_dir, _ := os.UserHomeDir()
				f, err := os.OpenFile(home_dir+"/"+msgstruct.File_name+msgstruct.File_type, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0777)
				if err != nil {
					fmt.Println("Something went wrong using Donwload, Possibly permission errors :", err)
					time.Sleep(3 * time.Second)
					return
				}
				_, write_err := f.Write(msgstruct.File_content)
				if write_err != nil {
					fmt.Println("Something went wrong using Donwload, Possibly permission errors :", write_err)
					time.Sleep(3 * time.Second)
					return
				} else {
					fmt.Println("File written to home dir!!")
					time.Sleep(3 * time.Second)
					return
				}
			}

		}
		copy_nodes = copy_nodes.next
	}
	fmt.Println("You sneaky fellow, There is no file with that id :)) ")
	time.Sleep(3 * time.Second)
}
