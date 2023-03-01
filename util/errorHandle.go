/**
* @description :
* @author : Jarick
* @Date : 2022-08-15
* @Url : http://CloudWebOps
 */
package util

import "log"

//  Error Handling for panic recover
// func ErrPanic() {

// 	var PanicDebug = false //debug panic model,only test

// 	if !PanicDebug {
// 		if err := recover(); err != nil {
// 			log.Printf("[Panic Error INFO]: %s\n", err)
// 		}
// 	}
// }

func ErrPanicDebug(PanicDebug bool) {

	// var PanicDebug = false //debug panic model,only test

	if !PanicDebug {
		if err := recover(); err != nil {
			log.Printf("[Panic Error INFO]: %s\n", err)
		}
	}
	// log.Println("Debug mode is being used")
	// log.Println("")

}

//  Error Handling for panic recover
func ErrPanic() {

	var PanicDebug = false //debug panic model,only debug_test change true

	if !PanicDebug {
		if err := recover(); err != nil {
			log.Printf("[Panic Error INFO]: %s\n", err)
		}
	}
}

func ErrorHandle(err interface{}, text string) {
	if err != nil {
		log.Printf("[ERROR INFO: %s]: %s\n", text, err)
		panic(err)
	}
}

func DebugPrint(debug bool, message interface{}) {
	if debug {
		log.Println("Debug mode is open.")
		if message != nil {
			log.Println(message)
		}
	}
}
