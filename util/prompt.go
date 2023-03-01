/**
* @description :
* @author : Jarick
* @Date : 2022-06-27
* @Url : http://CloudWebOps
 */
package util

import (
	"bufio"
	"fmt"
	"os"
)

// 回车表示确认
func Prompt() {
	fmt.Printf("-> Press Return key to continue.")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		break
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	fmt.Println()
}

// 按键表示确认
func PromptConfirm() {
	var PromptConfirmOk string
	fmt.Println("请确认是否执行(yes/no)：")
	_, err := fmt.Scan(&PromptConfirmOk)
	if err != nil {
		os.Exit(0)
	}
	if PromptConfirmOk != "yes" {
		os.Exit(0)
	}
}
