/**
* @description :
* @author : Jarick
* @Date : 2022-11-19
* @Url : http://CloudWebOps
 */

package util

import (
	"fmt"
	"log"
	"regexp"
)

//
//  @headerTitle :  RegexpMatchTrue
//  @description :
//  @function : 正则方案匹配，Pattern： 正则表达式， keyString：匹配文本。Debug： 是否开启debug输出。
func RegexpMatchTrue(Pattern string, keyString string, Debug bool) (ok bool, err error) {

	// Web input regexp pattern
	// https://www.lddgo.net/string/golangregex
	var PatternName string
	var PatternLen int
	var Msg string
	switch Pattern {
	case "username":
		PatternName = `^[a-zA-Z0-9]+$`
		PatternLen = 10
	case "password":
		PatternName = `^[a-z0-9A-Z!@#$%-]+$`
		PatternLen = 20
	case "telephone":
		PatternName = `^(1[3|4|5|8][0-9]\d{4,8})$`
		PatternLen = 11
	case "email":
		PatternName = `^[0-9a-z][_.0-9a-z-]{0,31}@([0-9a-z][0-9a-z-]{0,30}[0-9a-z]\.){1,4}[a-z]{2,4}$`
		PatternLen = 20
	case "url":
		PatternName = ""
		PatternLen = 20
	case "number":
		PatternName = ""
		PatternLen = 10
	case "ipv4":
		PatternName = ""
		PatternLen = 10
	case "ipv6":
		PatternName = ""
		PatternLen = 10
	case "IDCard": // //验证18位身份证，18位前17位为数字，最后一位是校验位，可能为数字或字符X。
		PatternName = `^(\d{17})([0-9]|X)$`
		PatternLen = 10
	case "token": // 、
		PatternName = ``
		PatternLen = 64
	case "status": // 、
		PatternName = ``
		PatternLen = 5
	}

	if len(keyString) == 0 { // 空值判断
		ok = false
		err = fmt.Errorf("%v is null  [The parameter is null !]", Pattern)
	} else if len(keyString) > PatternLen { // 超出字符长度判断
		ok = false
		err = fmt.Errorf("%v is TooBig [The parameter length cannot be greater than %v !]", Pattern, PatternLen)
	} else { // 字符正则匹配判断
		if m, _ := regexp.MatchString(PatternName, keyString); !m {
			err = fmt.Errorf("%v=%v+ [The input parameter is invalid !]", Pattern, keyString)
			ok = false
		} else {
			Msg = Pattern + "=" + keyString
			ok = true
			err = nil
		}
	}

	// debug 模式输出
	if Debug {
		if ok {
			log.Printf("Messages: %v\n", Msg)
		} else {
			log.Printf("Error: %v\n", err)
		}
	}

	return ok, err
}
