/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"CloudDevKubernetes/controller"
	"CloudDevKubernetes/model"
	"CloudDevKubernetes/util"
	"fmt"
	"log"
	"reflect"

	"github.com/spf13/cobra"
	"k8s.io/client-go/kubernetes"
	metrics "k8s.io/metrics/pkg/client/clientset/versioned"
)

// kubediagsCmd represents the kubediags command
var kubediagsCmd = &cobra.Command{
	Use:   "kubediags",
	Short: kubediagsTitle,
	Long:  kubediagsTitle,
	RunE:  kubediag,
}

func init() {

	// 添加新的flags
	kubediagsCmd.Flags().BoolP("deployments", "d", false, "The Deployments Diagsinfo .")
	kubediagsCmd.Flags().BoolP("statefulsets", "f", false, "The Statefulsets Diagsinfo .")
	kubediagsCmd.Flags().BoolP("systemns", "", false, "Open Action The System_namespaces Diagsinfo .")

	rootCmd.AddCommand(kubediagsCmd)
}

type DiagsRes struct {
	ID           string
	Case         string
	Point        string
	CurrentState string
	Severity     string
	Describe     string
}

func kubediag(cmd *cobra.Command, _ []string) error {

	// 获取命令行参数
	ns, _ := cmd.Flags().GetString("namespace")
	rmitem, _ := cmd.Flags().GetStringSlice("rmitem")
	flagdeployment, _ := cmd.Flags().GetBool("deployments")
	flagstatefulsets, _ := cmd.Flags().GetBool("statefulsets")
	systemns, _ := cmd.Flags().GetBool("systemns")
	json, _ := cmd.Flags().GetBool("json")
	debug, _ := cmd.Flags().GetBool("debug")
	desTurn, _ := cmd.Flags().GetBool("des")
	utilAction, _ := cmd.Flags().GetBool("util")

	// 获取授权方式
	clientSet, mClientset := controller.HandleToK8sAndMetric(debug)

	// 判断查询对象类型
	var flagtext string
	if flagdeployment {
		flagtext = "deployments"
	} else if flagstatefulsets {
		flagtext = "statefulsets"
	} else {
		flagtext = ""
	}

	kubediagBody := resourceDiag(clientSet, mClientset, ns, flagtext, systemns, debug, utilAction, json)

	// 生成表单的结构体
	vReflect := reflect.TypeOf(DiagsRes{})

	if len(kubediagBody) == 0 {
		if errorMesg == "" {
			defer fmt.Printf("No resource displayed.\n")
		} else {
			defer fmt.Printf("No resource displayed. please check environment for kubernetes. \n[ERROR]: %s\n", errorMesg)
		}
	} else {
		defer func(desTurn bool, Title string) {
			if desTurn {
				fmt.Println(Title + "\n" + WarningInfo("DiagsSeverity", desTurn) + "\n" + WarningInfo("Unittext", desTurn))
			}
		}(desTurn, kubediagsTitle)

	}

	// 加载表单信息
	table := GenTable(vReflect, kubediagBody, rmitem, debug)

	// json format
	if json {
		jsonStr, _ := table.Json(2)
		fmt.Println(jsonStr)
		// return nil
	} else {
		table.PrintTable()
	}
	if debug && errorMesg != "" {
		log.Println(errorMesg)
	}
	return nil
}

//定义指标
func resourceDiag(clientSet *kubernetes.Clientset, mClientset *metrics.Clientset, ns string, flagtext string, systemns bool, debug bool, utilAction bool, json bool) (bodyRes []map[string]string) {
	util.ErrPanicDebug(debug)
	configDir := "" // 默认为当前目录

	// 加载配置文件
	error := model.InitDefaultConfig().GetConf(configDir).Error
	if debug && error != nil {
		log.Println(error)
	}

	PODUnAvailableWarning := model.InitDefaultConfig().GetConf(configDir).Kubediags.PODUnAvailableWarning
	PODUnAvailableError := model.InitDefaultConfig().GetConf(configDir).Kubediags.PODUnAvailableError
	CPULimitInfo := model.InitDefaultConfig().GetConf(configDir).Kubediags.CPULimitInfo
	CPULimitWarning := model.InitDefaultConfig().GetConf(configDir).Kubediags.CPULimitWarning
	MEMLimitInfo := model.InitDefaultConfig().GetConf(configDir).Kubediags.MEMLimitInfo
	MEMLimitWarning := model.InitDefaultConfig().GetConf(configDir).Kubediags.MEMLimitWarning
	MEMOversoldNodeInfo := model.InitDefaultConfig().GetConf(configDir).Kubediags.MEMOversoldNodeInfo
	MEMOversoldNodeWarning := model.InitDefaultConfig().GetConf(configDir).Kubediags.MEMOversoldNodeWarning

	RequestCPUInfo := model.InitDefaultConfig().GetConf(configDir).Kubediags.RequestCPUInfo
	RequestCPUWarning := model.InitDefaultConfig().GetConf(configDir).Kubediags.RequestCPUWarning
	RequestMemInfo := model.InitDefaultConfig().GetConf(configDir).Kubediags.RequestMemInfo
	RequestMemWarning := model.InitDefaultConfig().GetConf(configDir).Kubediags.RequestMemWarning
	PODCapacityNodeWarning := model.InitDefaultConfig().GetConf(configDir).Kubediags.PODCapacityNodeWarning
	PODErrorNodeWarning := model.InitDefaultConfig().GetConf(configDir).Kubediags.PODErrorNodeWarning

	SystemNamespace := model.InitDefaultConfig().GetConf(configDir).Kubediags.SystemNamespace

	var bodyResDeployDeployment = ResourceBodyDeploy(clientSet, ns, flagtext, debug)

	for _, v := range bodyResDeployDeployment {

		// 不检查系统命名空间
		if !systemns && util.SliceContains(SystemNamespace, v["NameSpace"]) {
			continue
		}

		if util.StringRepliceRes00(v["Replices"]) > PODUnAvailableWarning {
			bodyRes = append(bodyRes, DiagContext(bodyRes, "PODUnAvailableWarning", PODUnAvailableWarning, v["Replices"], v["TYPE"]+"/"+v["NameSpace"]+"/"+v["ResourceName"]+"/"+v["ContainerName"]))
		}
		if util.StringRepliceRes00(v["Replices"]) > PODUnAvailableError {
			bodyRes = append(bodyRes, DiagContext(bodyRes, "PODUnAvailableError", PODUnAvailableError, v["Replices"], v["TYPE"]+"/"+v["NameSpace"]+"/"+v["ResourceName"]+"/"+v["ContainerName"]))
		}
		if util.StringRepliceRes00(v["LimitCPU"]) > CPULimitInfo {
			bodyRes = append(bodyRes, DiagContext(bodyRes, "CPULimitInfo", CPULimitInfo, v["LimitCPU"], v["TYPE"]+"/"+v["NameSpace"]+"/"+v["ResourceName"]+"/"+v["ContainerName"]))
		}
		if util.StringRepliceRes00(v["LimitCPU"]) > CPULimitWarning {
			bodyRes = append(bodyRes, DiagContext(bodyRes, "CPULimitWarning", CPULimitWarning, v["LimitCPU"], v["TYPE"]+"/"+v["NameSpace"]+"/"+v["ResourceName"]+"/"+v["ContainerName"]))
		}
		if util.StringRepliceRes00(v["LimitMem"]) > MEMLimitInfo {
			bodyRes = append(bodyRes, DiagContext(bodyRes, "MEMLimitInfo", MEMLimitInfo, v["LimitMem"], v["TYPE"]+"/"+v["NameSpace"]+"/"+v["ResourceName"]+"/"+v["ContainerName"]))
		}
		if util.StringRepliceRes00(v["LimitMem"]) > MEMLimitWarning {
			bodyRes = append(bodyRes, DiagContext(bodyRes, "MEMLimitWarning", MEMLimitWarning, v["LimitMem"], v["TYPE"]+"/"+v["NameSpace"]+"/"+v["ResourceName"]+"/"+v["ContainerName"]))
		}
	}

	// node 资源不需要查询ns,该查询项放在最后

	var bodyResNode []map[string]string
	if ns != "" { // 判断，如果指定命名空间，则不查询 node维度指标信息
		bodyResNode = nil
	} else {
		bodyResNode, _ = KubeResHandle(clientSet, mClientset, "", debug, utilAction, json) // node资源，不需要携带命名空间参数
	}

	for _, v := range bodyResNode {

		if v["NodeTaints"] != "success" {
			bodyRes = append(bodyRes, DiagContext(bodyRes, "nodeNotready", 0, v["NodeTaints"], "node/"+v["NodeName"]))
		}
		if util.StringRepliceRes00(v["RequestCPU"]) > RequestCPUInfo {
			bodyRes = append(bodyRes, DiagContext(bodyRes, "RequestCPUInfo", RequestCPUInfo, v["RequestCPU"], "node/"+v["NodeName"]))
		}
		if util.StringRepliceRes00(v["RequestCPU"]) > RequestCPUWarning {
			bodyRes = append(bodyRes, DiagContext(bodyRes, "RequestCPUWarning", RequestCPUWarning, v["RequestCPU"], "node/"+v["NodeName"]))
		}
		if util.StringRepliceRes00(v["RequestMem"]) > RequestMemInfo {
			bodyRes = append(bodyRes, DiagContext(bodyRes, "RequestMemInfo", RequestMemInfo, v["RequestMem"], "node/"+v["NodeName"]))
		}
		if util.StringRepliceRes00(v["RequestMem"]) > RequestMemWarning {
			bodyRes = append(bodyRes, DiagContext(bodyRes, "RequestMemWarning", RequestMemWarning, v["RequestMem"], "node/"+v["NodeName"]))
		}
		if util.StringRepliceRes00(v["LimitMem"]) > MEMOversoldNodeInfo {
			bodyRes = append(bodyRes, DiagContext(bodyRes, "MEMOversoldNodeInfo", MEMOversoldNodeInfo, v["LimitMem"], "node/"+v["NodeName"]))
		}
		if util.StringRepliceRes00(v["LimitMem"]) > MEMOversoldNodeWarning {
			bodyRes = append(bodyRes, DiagContext(bodyRes, "MEMOversoldNodeWarning", MEMOversoldNodeWarning, v["LimitMem"], "node/"+v["NodeName"]))
		}
		if util.StringRepliceRes01(v["PodCal"]) > PODCapacityNodeWarning {
			bodyRes = append(bodyRes, DiagContext(bodyRes, "PODCapacityNodeWarning", PODCapacityNodeWarning, v["PodCal"], "node/"+v["NodeName"]))
		}
		if util.StringRepliceRes00(v["PodCal"]) > PODErrorNodeWarning {
			bodyRes = append(bodyRes, DiagContext(bodyRes, "PODErrorNodeWarning", PODErrorNodeWarning, v["PodCal"], "node/"+v["NodeName"]))
		}
	}

	return bodyRes
}

//定义告警标准
func DiagContext(bodyRes []map[string]string, bodyResType string, alertNum float64, Status string, Point string) (diagcontext map[string]string) {
	diagcontext = make(map[string]string, 0)
	// Severity： 1、严重，2、警告，3、预警、4、提示
	switch {
	// node diags
	case bodyResType == "nodeNotready":
		diagcontext["ID"] = fmt.Sprint(len(bodyRes))
		diagcontext["Case"] = bodyResType + "_" + fmt.Sprint(alertNum)
		diagcontext["Severity"] = "严重"
		diagcontext["CurrentState"] = Status
		diagcontext["Describe"] = "节点状态非Ready"
		diagcontext["Point"] = Point
	case bodyResType == "RequestCPUInfo":
		diagcontext["ID"] = fmt.Sprint(len(bodyRes))
		diagcontext["Case"] = bodyResType + "_" + fmt.Sprint(alertNum)
		diagcontext["Severity"] = "提示"
		diagcontext["CurrentState"] = Status
		diagcontext["Describe"] = "节点CPU分配率高于:" + fmt.Sprint(alertNum)
		diagcontext["Point"] = Point
	case bodyResType == "RequestMemInfo":
		diagcontext["ID"] = fmt.Sprint(len(bodyRes))
		diagcontext["Case"] = bodyResType + "_" + fmt.Sprint(alertNum)
		diagcontext["Severity"] = "提示"
		diagcontext["CurrentState"] = Status
		diagcontext["Describe"] = "节点内存分配率高于:" + fmt.Sprint(alertNum)
		diagcontext["Point"] = Point
	case bodyResType == "RequestCPUWarning":
		diagcontext["ID"] = fmt.Sprint(len(bodyRes))
		diagcontext["Case"] = bodyResType + "_" + fmt.Sprint(alertNum)
		diagcontext["Severity"] = "预警"
		diagcontext["CurrentState"] = Status
		diagcontext["Describe"] = "节点CPU分配率高于:" + fmt.Sprint(alertNum)
		diagcontext["Point"] = Point
	case bodyResType == "RequestMemWarning":
		diagcontext["ID"] = fmt.Sprint(len(bodyRes))
		diagcontext["Case"] = bodyResType + "_" + fmt.Sprint(alertNum)
		diagcontext["Severity"] = "预警"
		diagcontext["CurrentState"] = Status
		diagcontext["Describe"] = "节点内存分配率高于:" + fmt.Sprint(alertNum)
		diagcontext["Point"] = Point
	case bodyResType == "MEMOversoldNodeInfo":
		diagcontext["ID"] = fmt.Sprint(len(bodyRes))
		diagcontext["Case"] = bodyResType + "_" + fmt.Sprint(alertNum)
		diagcontext["Severity"] = "提示"
		diagcontext["CurrentState"] = Status
		diagcontext["Describe"] = "节点内存超售高于:" + fmt.Sprint(alertNum)
		diagcontext["Point"] = Point
	case bodyResType == "MEMOversoldNodeWarning":
		diagcontext["ID"] = fmt.Sprint(len(bodyRes))
		diagcontext["Case"] = bodyResType + "_" + fmt.Sprint(alertNum)
		diagcontext["Severity"] = "预警"
		diagcontext["CurrentState"] = Status
		diagcontext["Describe"] = "节点内存超售高于:" + fmt.Sprint(alertNum)
		diagcontext["Point"] = Point
	case bodyResType == "PODCapacityNodeWarning":
		diagcontext["ID"] = fmt.Sprint(len(bodyRes))
		diagcontext["Case"] = bodyResType + "_" + fmt.Sprint(alertNum)
		diagcontext["Severity"] = "预警"
		diagcontext["CurrentState"] = Status
		diagcontext["Describe"] = "节点POD数高于:" + fmt.Sprint(alertNum)
		diagcontext["Point"] = Point
	case bodyResType == "PODErrorNodeWarning":
		diagcontext["ID"] = fmt.Sprint(len(bodyRes))
		diagcontext["Case"] = bodyResType + "_" + fmt.Sprint(alertNum)
		diagcontext["Severity"] = "提示"
		diagcontext["CurrentState"] = Status
		diagcontext["Describe"] = "节点POD数高于:" + fmt.Sprint(alertNum)
		diagcontext["Point"] = Point
	// depoly diags
	case bodyResType == "PODUnAvailableWarning":
		diagcontext["ID"] = fmt.Sprint(len(bodyRes))
		diagcontext["Case"] = bodyResType + "_" + fmt.Sprint(alertNum)
		diagcontext["Severity"] = "预警"
		diagcontext["CurrentState"] = Status
		diagcontext["Describe"] = "应用副本异常高于:" + fmt.Sprint(alertNum)
		diagcontext["Point"] = Point
	case bodyResType == "PODUnAvailableError":
		diagcontext["ID"] = fmt.Sprint(len(bodyRes))
		diagcontext["Case"] = bodyResType + "_" + fmt.Sprint(alertNum)
		diagcontext["Severity"] = "警告"
		diagcontext["CurrentState"] = Status
		diagcontext["Describe"] = "应用副本异常高于:" + fmt.Sprint(alertNum)
		diagcontext["Point"] = Point
	case bodyResType == "MEMLimitInfo":
		diagcontext["ID"] = fmt.Sprint(len(bodyRes))
		diagcontext["Case"] = bodyResType + "_" + fmt.Sprint(alertNum)
		diagcontext["Severity"] = "提示"
		diagcontext["CurrentState"] = Status
		diagcontext["Describe"] = "容器内存超售高于:" + fmt.Sprint(alertNum)
		diagcontext["Point"] = Point
	case bodyResType == "MEMLimitWarning":
		diagcontext["ID"] = fmt.Sprint(len(bodyRes))
		diagcontext["Case"] = bodyResType + "_" + fmt.Sprint(alertNum)
		diagcontext["Severity"] = "预警"
		diagcontext["CurrentState"] = Status
		diagcontext["Describe"] = "容器内存超售高于:" + fmt.Sprint(alertNum)
		diagcontext["Point"] = Point
	case bodyResType == "CPULimitInfo":
		diagcontext["ID"] = fmt.Sprint(len(bodyRes))
		diagcontext["Case"] = bodyResType + "_" + fmt.Sprint(alertNum)
		diagcontext["Severity"] = "提示"
		diagcontext["CurrentState"] = Status
		diagcontext["Describe"] = "容器CPU超售高于:" + fmt.Sprint(alertNum)
		diagcontext["Point"] = Point
	case bodyResType == "CPULimitWarning":
		diagcontext["ID"] = fmt.Sprint(len(bodyRes))
		diagcontext["Case"] = bodyResType + "_" + fmt.Sprint(alertNum)
		diagcontext["Severity"] = "预警"
		diagcontext["CurrentState"] = Status
		diagcontext["Describe"] = "容器CPU超售高于:" + fmt.Sprint(alertNum)
		diagcontext["Point"] = Point
	}
	return
}
