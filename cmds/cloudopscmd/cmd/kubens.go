/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"CloudDevKubernetes/controller"
	"CloudDevKubernetes/util"
	"context"
	"fmt"
	"log"
	"reflect"

	"github.com/spf13/cobra"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// kubensCmd represents the kubens command
var kubensCmd = &cobra.Command{
	Use:   "kubens",
	Short: kubensTitle,
	Long:  kubensTitle,
	RunE:  resourceNs,
}

func init() {

	// 添加新的flags
	kubensCmd.Flags().BoolP("deployments", "d", false, "The Deployments Resources.")
	kubensCmd.Flags().BoolP("statefulsets", "f", false, "The Statefulsets Resources.")

	rootCmd.AddCommand(kubensCmd)
}

type NsClsResources struct {
	Namespace  string
	Container  int64
	Pod        int64
	ErrCount   int64
	RequestCPU int64
	RequestMem int64
	LimitCPU   int64
	LimitMem   int64
}

func resourceNs(cmd *cobra.Command, _ []string) error {
	// startT := time.Now()

	// 获取授权方式
	clientSet := controller.HandleToCluster()

	// 获取命令行参数
	json, _ := cmd.Flags().GetBool("json")
	flagdeployment, _ := cmd.Flags().GetBool("deployments")
	flagstatefulsets, _ := cmd.Flags().GetBool("statefulsets")
	debug, _ := cmd.Flags().GetBool("debug")
	desTurn, _ := cmd.Flags().GetBool("des")
	rmitem, _ := cmd.Flags().GetStringSlice("rmitem")

	// 判断查询对象类型
	var flagtext string
	if flagdeployment {
		flagtext = "deployments"
	} else if flagstatefulsets {
		flagtext = "statefulsets"
	} else {
		flagtext = ""
	}

	// 获取查询指标数据
	var n *NsClsResources

	nsRes := n.ResourceBodyNs(clientSet, flagtext, debug)
	// 判断返回数据
	if len(nsRes) == 0 {
		if errorMesg == "" {
			defer fmt.Printf("No resource displayed.\n")
		} else {
			defer fmt.Printf("No resource displayed. please check environment for kubernetes. \n[ERROR]: %s\n", errorMesg)
		}
	} else {
		defer func(desTurn bool, Title string) {
			if desTurn {
				fmt.Println(Title + "\n" + WarningInfo("Unittext", desTurn))
			}
		}(desTurn, kubensTitle)
	}
	// 生成表头结构体
	vReflect := reflect.TypeOf(NsClsResources{})
	// 加载表单信息
	table := GenTable(vReflect, nsRes, rmitem, debug)

	// 命令行输出选项
	if json { //json
		jsonStr, _ := table.Json(2)
		fmt.Println(jsonStr)
	} else {
		table.PrintTable()
	}
	if debug && errorMesg != "" {
		log.Println(errorMesg)
	}
	return nil
}

func (n *NsClsResources) ResourceBodyNs(clientSet *kubernetes.Clientset, flagtext string, debug bool) (bodyRes []map[string]string) {
	util.ErrPanicDebug(debug)
	var nsRes []map[string]string

	// get all nsinfo
	nsInfo, err := clientSet.CoreV1().Namespaces().List(context.TODO(), v1.ListOptions{})
	if err != nil {
		fmt.Println(err)
	}
	// get all namespaces type resource
	ResourceBodyDeployNs := ResourceBodyDeploy(clientSet, "", flagtext, debug)

	for _, ns := range nsInfo.Items {
		nssingleRes := make(map[string]string, 0)
		nsResources := new(NsClsResources) // 针对每一个节点，初始化一次
		PodList := make([]string, 0)

		for _, v := range ResourceBodyDeployNs {

			if ns.Name == v["NameSpace"] {
				replicesNum := util.StringRepliceRes02(v["Replices"])

				if !util.SliceContains(PodList, v["ResourceName"]) {
					PodList = append(PodList, v["ResourceName"])
					nsResources.Pod = nsResources.Pod + replicesNum
				}
				nsResources.Container = nsResources.Container + replicesNum
				nsResources.LimitMem = nsResources.LimitMem + util.StringRepliceRes02(v["LimitMem"])*replicesNum
				nsResources.LimitCPU = nsResources.LimitCPU + util.StringRepliceRes02(v["LimitCPU"])*replicesNum
				nsResources.RequestCPU = nsResources.RequestCPU + util.StringRepliceRes03(v["RequestCPU"])*replicesNum
				nsResources.RequestMem = nsResources.RequestMem + util.StringRepliceRes03(v["RequestMem"])*replicesNum

				// 5、CLI输出格式
				nssingleRes["Container"] = fmt.Sprint(nsResources.Container)
				nssingleRes["LimitMem"] = fmt.Sprint(nsResources.LimitMem)
				nssingleRes["LimitCPU"] = fmt.Sprint(nsResources.LimitCPU)
				nssingleRes["RequestCPU"] = fmt.Sprint(nsResources.RequestCPU)
				nssingleRes["RequestMem"] = fmt.Sprint(nsResources.RequestMem)
				nssingleRes["Pod"] = fmt.Sprint(nsResources.Pod)
				nssingleRes["Namespace"] = v["NameSpace"]
			}
		}
		if len(nssingleRes) < 1 {
			continue
		}

		nsRes = append(nsRes, nssingleRes)

	}
	// srcText := "12310.1"
	// dstType := "int64"
	// fmtType := ""
	// ss := new(StringTypeInt)
	// ss.StringRepliceRes(srcText, dstType, fmtType)
	// fmt.Println("xxxxxxxxxxxxxxxxxx", ss.Int64)

	bodyRes = nsRes
	return
}
