/**
* @description :
* @author : Jarick
* @Date : 2022-08-19
* @Url : http://CloudWebOps
 */
package util

import (
	"fmt"
	"log"
	"os"
)

func ExistDir(path string) {

	// check
	if _, err := os.Stat(path); err == nil {
		fmt.Printf("")
	} else {
		log.Println("The path is not exists:", path)
		err := os.MkdirAll(path, 0711)

		if err != nil {
			log.Println("Error creating directory:", path)
			log.Println(err)
			return
		}
	}

	// check again
	if _, err := os.Stat(path); err == nil {
		fmt.Printf("")
	} else {
		log.Println("The path is error: ", err)
	}
}

// slice 是否包含元素
func SliceContains(elems []string, v string) bool {
	for _, s := range elems {
		if v == s {
			return true
		}
	}
	return false
}
