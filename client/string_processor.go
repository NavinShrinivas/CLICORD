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

package main 

import (
    "strings"
)

func Fancyfier (msg string) string {
    //As of writing this we only have emojis for fancy stuff in msg 
    var ret_str string;
    for index1 ,chars := range msg{
        if chars == ':'{
            for index2,chars2 := range msg[index1+1:]{
                if chars2 == ':'{
                    possible_alias := msg[index1+1:index1+index2+1]

                    if strings.ContainsAny(possible_alias , " "){
                        break;
                    }else if _,ok := glob_emoji_data[possible_alias] ; !ok{
                        break
                    }
                    ret_str+=glob_emoji_data[possible_alias]
                }
            }
        }
        ret_str += string(chars)
    }
    return ret_str
}
