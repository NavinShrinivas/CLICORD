/**********************************************************************************
  ____ _     ___ ____ ___  ____  ____
 / ___| |   |_ _/ ___/ _ \|  _ \|  _ \
| |   | |    | | |  | | | | |_) | | | |
| |___| |___ | | |__| |_| |  _ <| |_| |
 \____|_____|___\____\___/|_| \_\____/

Copyright (c) 2021  CLICORD

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
**********************************************************************************/

//This file contains functions that all messages have to pass through
//If the message contains any valid emojis it converts it to repective utf-8

package main

import (
	"encoding/json"
	"fmt"
	"os"
)

//refered to these to write this part of this program :
/*
 *https://github.com/enescakir/emojimap
 *https://raw.githubusercontent.com/github/gemoji/master/db/emoji.json
 *https://github.com/github/gemoji
 *
 */

type emojis struct {
	Emoji   string   `json : "emoji"`
	Aliases []string `json : "aliases"`
	Tags    []string `json : "tags"`
}

func EmojiMapGenerator() (map[string]string, error) {
	var emojimap []emojis
	json_str, read_err := os.ReadFile("./gemoji_database.json")
	if read_err != nil {
		fmt.Println("Oh oh, something went wrong. May be you are runnig the binary from else where?")
		fmt.Println("From devs : this could have been avoided, but we wanted to keep the binary size very low")
		return nil, read_err
	}
	json_err := json.Unmarshal([]byte(json_str), &emojimap)
	if json_err != nil {
		fmt.Println("Something went wrong parsing to JSON, you sure you haven'tmodified our JSON file?")
		return nil, json_err
	}
	r := make(map[string]string)
	//To understand the iterators below :
	// https://stackoverflow.com/questions/18130859/how-can-i-iterate-over-a-string-by-runes-in-go
	for _, items := range emojimap {
		for _, aliases := range items.Aliases {
			r[aliases] = items.Emoji
		}
		for _, tags := range items.Tags {
			r[tags] = items.Emoji
		}
	}
	return r, nil
}
