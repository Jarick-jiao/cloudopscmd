/**
* @description :
* @author : Jarick
* @Date : 2022-06-26
* @Url : http://CloudWebOps
 */
package main

import (
	"CloudDevKubernetes/cmds/cloudopscmd/cmd"
	"log"
	"os"
	"runtime/trace"

	"github.com/pkg/profile"
)

func Cloudopscmd() {
	cmd.Execute()
}

func main() {

	// Function: tarce alpha

	// debug tool trace && pprof //export GOTRACE="pprof"
	if envGOTRACE := os.Getenv("GOTRACE"); len(envGOTRACE) > 0 {
		log.Printf("Mode GOTRACE = %s\n", envGOTRACE)

		if envGOTRACE == "trace" {
			file, _ := os.Create("./trace")
			err := trace.Start(file)
			if err != nil {
				log.Println(err)
			}
			log.Println("trace: trace enabled, ./trace")
			defer trace.Stop() //go tool trace ./trace  // 通过该方法直接访问web
		} else if envGOTRACE == "pprof" {

			defer profile.Start().Stop() // 开启pprof，会在tmp生成文件，go tool pprof /tmp/profile2104646588/cpu.pprof； svg // https://www.cnblogs.com/landv/p/11274877.html
		} else {
			log.Printf("Mode GOTRACE = %s  Invalid argument .\n", envGOTRACE)
		}
	}

	// Function: main
	Cloudopscmd()
}
