/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"CloudDevKubernetes/model"
	"CloudDevKubernetes/util"
	"fmt"
	"log"
	"os"
	"reflect"
	"time"

	"github.com/liushuochen/gotable"
	"github.com/liushuochen/gotable/table"
	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
)

// common type info

// unit
const (
	MemUnitG int64 = 1024 * 1024 * 1024 * 1000 //(G)
	CPUUnitC int64 = 1000                      //(C)
	MemUnitM int64 = 1024 * 1024 * 1000        //(M)
	CPUUnitM int64 = 1                         //(c)
	DisUnitM int64 = 1020 * 1024 * 1024        //(c)
)

// cmdcontext message
const (

	// cmds Title
	rootTitle       = "Display The Kubernetes Resources."
	rootLongTitle   = "CloudDevKubernetes.com: Display The Kubernetes Messages of Kinds Resources. \nMore Information At: https://kubernetes.io/docs/reference/kubectl/overview/."
	kuberesTitle    = "Display Node Resources And Metric For Usage."
	kubetopTitle    = "Display POD/Container CPU/MEM Resources And Usage Info."
	kubensTitle     = "Display Deployment And Statefulset Resources By The Namespaces."
	kubedeployTitle = "Display Deployment And Statefulset Resources Info."
	kubediagsTitle  = "Kubediags For Diagnostic Kubernetes Issues And Risks."
	kubetraceTitle  = "kubetrace For Diags Networks of Kubernetes Clusters Business Application."

	versionTitle = "The CloudDevKubernetes.com Information."

	// Global Function
	debugFunction     = "Debug mode"
	jsonFunction      = "Output In JSON Format, Default Table Format"
	desFunction       = "Output Parameter Description Info"
	docFunction       = "Generated Using The Cloudopscmd Document"
	docDirFunction    = "Point The Cloudopscmd Document Dir Path"
	genconfigFunction = "Generated Sing The Cloudopscmd Config File"
	configDirFunction = "The Genconfig Dir Path, If Using The Config File Needs To Be Moved To The Current Path(./config)"
	rmitemFunction    = "Remove Data Item. Only One Type Is Allowed To Be Specified"
	errorFunction     = "Only Display The Error Info."

	// Gloabl Args
	namespaceArgs = "If Present, The Namespace Scope For This CLI Request"

	// WarningInfo
	UnitContext = "[Note]: Unit: cpu(m), mem(Mi), ImageSize(G), CapMem(Gi), CapCPU(c), Replices[replices(err)], PodCal[pod(err)]\n[PrintTime]:"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "cloudopscmd",
	Short: rootTitle,
	Long:  "\n" + rootLongTitle,
	Run: func(cmd *cobra.Command, args []string) {
		docTurn, _ := cmd.Flags().GetBool("doc")
		docDirP, _ := cmd.Flags().GetString("docDir")
		configTurn, _ := cmd.Flags().GetBool("genconfig")
		configDirP, _ := cmd.Flags().GetString("configDir")

		// 判断是否携带参数，否则打印help
		if len(args) == 0 && !docTurn && !configTurn || len(args) == 0 && docTurn && configTurn {
			err := cmd.Help()
			if err != nil {
				log.Panicln(err)
			}
			return
		} else if len(args) == 0 && docTurn {
			// 生成文档代码
			// rootCmd.DisableAutoGenTag = true

			var docDir string = docDirP
			if docDirP == "" {
				docDir = "./cloudops/docs"
				log.Println("the cloudopscmd document for default dir [./cloudops/docs]")
			}
			util.ExistDir(docDir)
			err := doc.GenMarkdownTree(cmd, docDir)
			if err == nil {
				log.Println("the cloudopscmd document is gen to ", docDir)
			} else {
				log.Println(err)
			}

		} else if len(args) == 0 && configTurn {
			// 生成配置默认配置文件
			var configDir string = configDirP
			if configDirP == "" {
				configDir = "./cloudops/config"
				log.Println("the cloudopscmd config for default dir [./cloudops/config]")
			}
			util.ExistDir(configDir)
			err := model.Genconfig(configDir)
			if err != nil {
				log.Println(err)
			} else {
				log.Println("the cloudopscmd config is gen :", configDir+"/config.yaml")
			}
		}

	},
}

func Execute() {
	rootCmd.CompletionOptions.DisableDefaultCmd = true
	rootCmd.CompletionOptions.DisableNoDescFlag = true
	rootCmd.CompletionOptions.DisableDescriptions = true

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Global Flags:

	// KubernetesConfigFlags := genericclioptions.NewConfigFlags(false)  //集成 kubectl 命令行参数，目前意义不大
	// KubernetesConfigFlags.AddFlags(rootCmd.PersistentFlags())

	rootCmd.PersistentFlags().BoolP("debug", "", false, debugFunction)
	rootCmd.PersistentFlags().BoolP("json", "", false, jsonFunction)
	rootCmd.PersistentFlags().BoolP("des", "", false, desFunction)

	// rootCmd.PersistentFlags().String("namespace", "", "usage string")
	rootCmd.PersistentFlags().StringP("namespace", "n", "", namespaceArgs)

	rootCmd.Flags().BoolP("doc", "", false, docFunction)
	rootCmd.Flags().String("docDir", "", docDirFunction)

	rootCmd.Flags().BoolP("genconfig", "", false, genconfigFunction)
	rootCmd.Flags().String("configDir", "", configDirFunction)

	rootCmd.PersistentFlags().StringSlice("rmitem", []string{}, rmitemFunction)
}

//error meeages
var errorMesg string

func GenTable(v reflect.Type, mapList []map[string]string, rmitem []string, debug bool) *table.Table {

	var title []string

	util.DebugPrint(debug, "")

	// 反射，形成表头
	// v = reflect.TypeOf(StructtType{}.NodeClsResources)
	for i := 0; i < v.NumField(); i++ {
		// Remove data item. Only one type is allowed to be specified

		// 删除列名

		if len(rmitem) > 0 {
			if util.SliceContains(rmitem, v.Field(i).Name) { // 删除表头需要剔除的表头
				continue
			}
		}

		title = append(title, v.Field(i).Name)
	}

	// Remove data item. Only one type is allowed to be specified
	// 删除items_vlaues 列数据
	for _, v := range mapList {
		if len(rmitem) > 0 { //rmitem turn
			for i := 0; i < len(rmitem); i++ {
				delete(v, rmitem[i])
			}
		} else { // // 判断输入错误的items输入
			// log.Println("rmitem is error:", rmitem)
			errorMesg = fmt.Sprint("rmitem is error:", rmitem)
		}
	}

	t, err := gotable.Create(title...)

	//所有列都向左靠
	for _, v := range title {
		t.Align(v, gotable.Left)
	}

	if err != nil {
		log.Printf("create table error: %s", err.Error())
		return nil
	}

	t.AddRows(mapList)

	return t
}

func WarningInfo(cmdName string, des bool) (text string) {
	NowDate := time.Now().Format("2006-01-02 15:04:05")
	if des {
		switch {
		case cmdName == "Unittext":
			text = fmt.Sprintf("%v%v", UnitContext, NowDate)
		case cmdName == "kubernetesDefalutKind":
			text = "[Note]: Default resource messages is about The deployment , help more info."
		case cmdName == "DiagsSeverity":
			text = "[Note]: Severity：1、严重，2、警告，3、预警、4、提示."
		// case cmdName == "2":
		// 	text = "[Note]: \n"
		case cmdName == "nowdate":
			text = "[" + time.Now().Format("2006-01-02 15:04:05") + "]\n"
		case cmdName == "kubetraceContextInfo0":
			text = "kubetrace INFO:\n[1] Connection timed out\n表示TCP路由不正常，原因有很多，可能是服务器无法ping通，可能是服务器（防火墙等）丢弃了该请求报文包，也可能是服务器应答太慢，又或者存在间歇性的问题（这种情况很难从日志文件中排查问题）。\n[2] Connection refused\n表示从本地客户端到目标IP地址的路由是正常的，但是该目标端口没有进程在监听，然后服务端拒绝掉了连接。一个成功的tcp链接将会看到Syn，Syn-Ack，Ack，这也就是我们预期的TCP三次握手。当使用tcpdump或wireshark抓包工具来探测发送过来的请求报文包时，Connection refused将会看到Syn,Rst。"
			return text
		}
	}
	return
}
