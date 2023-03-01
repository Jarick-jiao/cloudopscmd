/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/

package cmd

import (
	"CloudDevKubernetes/controller"
	service "CloudDevKubernetes/service/kubernetesHandle"
	"CloudDevKubernetes/util"
	"fmt"
	"reflect"

	"github.com/spf13/cobra"
	"k8s.io/client-go/kubernetes"
	metrics "k8s.io/metrics/pkg/client/clientset/versioned"
)

// kubetopCmd represents the kubetop command
var kubetopCmd = &cobra.Command{
	Use:   "kubetop",
	Short: kubetopTitle,
	Long:  kubetopTitle,
	RunE:  kubeTop,
}

func init() {

	// 创建子命令
	kubetopCmd.Flags().BoolP("error", "e", false, errorFunction)
	kubetopCmd.Flags().BoolP("containers", "", false, "Show All Containers Resources.")
	kubetopCmd.Flags().BoolP("extUsage", "", false, "Only Display The Correct Containers Resources Configuration")
	// kubetopCmd.Flags().String("namespaceLabels", "", "service trace timeout")
	// kubetopCmd.Flags().String("nodeLabels", "", "service trace timeout")
	kubetopCmd.Flags().StringP("labelSelector", "l", "", "Select Pod From LabelSelector")

	rootCmd.AddCommand(kubetopCmd)
}

func kubeTop(cmd *cobra.Command, _ []string) error {

	// 获取执行参数
	namespace, _ := cmd.Flags().GetString("namespace")
	labelSelector, _ := cmd.Flags().GetString("labelSelector")
	// namespaceLabels, _ := cmd.Flags().GetString("namespaceLabels")
	// nodeLabels, _ := cmd.Flags().GetString("nodeLabels")
	debug, _ := cmd.Flags().GetBool("debug")
	extUsage, _ := cmd.Flags().GetBool("extUsage")
	containers, _ := cmd.Flags().GetBool("containers")
	desTurn, _ := cmd.Flags().GetBool("des")
	rmitem, _ := cmd.Flags().GetStringSlice("rmitem")
	json, _ := cmd.Flags().GetBool("json")

	util.ErrPanicDebug(debug)

	// 获取授权方式
	clientSet, mClientset := controller.HandleToK8sAndMetric(debug)

	// 获取查询指标数据
	bodyRes := KubeTopHandle(clientSet, mClientset, namespace, labelSelector, containers, extUsage, debug)

	// 判断返回数据
	if len(bodyRes) == 0 {
		if errorMesg == "" {
			defer fmt.Printf("No resource displayed.\n")
		} else {
			defer fmt.Printf("No resource displayed. please check environment for kubernetes. \n[ERROR]: %s\n", errorMesg)
		}
	} else {
		defer func(desTurn bool, Title string) {
			if desTurn {
				Title02 := "[Note]: metric(%)=(UsageCPU - RequestCPU）/（LimitCPU - RequestCPU), (Limit == 0 && Request == 0 || Request > Limit) = 100"
				fmt.Println(Title + "\n" + Title02 + "\n" + WarningInfo("Unittext", desTurn))
			}
		}(desTurn, kubetopTitle)
	}
	// 生成表头结构体
	vReflect := reflect.TypeOf(DisPlayKubeTopItem{})

	// 定制表头
	if !containers {
		rmitem = append(rmitem, "ContainerName")
	}
	// 加载表单信息
	table := GenTable(vReflect, bodyRes, rmitem, debug)

	// 命令行输出选项
	if json { //json
		jsonStr, _ := table.Json(2)
		fmt.Println(jsonStr)
	} else {
		table.PrintTable()
	}
	// if debug && errorMesg != "" {
	// 	log.Println(errorMesg)
	// }
	util.DebugPrint(debug, errorMesg)

	return nil
}

// display 数据结构实现方式
type DisPlayKubeTop struct {
	DisPlayKubeTop []DisPlayKubeTopItem
}

type DisPlayKubeTopItem struct {
	ID            int
	Type          string
	ResourceName  string
	Namespace     string
	NodeName      string
	ContainerName string
	RequestCPU    int64
	LimitCPU      int64
	RequestMem    int64
	LimitMem      int64
	UsageCPU      int64
	UsageMem      int64
}

// init DisPlayKubeTop
func NewDisPlayKubeTop() *DisPlayKubeTop {
	return &DisPlayKubeTop{}
}

func KubeTopHandle(clientSet *kubernetes.Clientset, mClientset *metrics.Clientset, namespace string, LabelSelector string, Containers bool, extUsage bool, debug bool) (bodyRes []map[string]string) {

	util.ErrPanicDebug(debug)

	// init DisPlayKubeTop
	NewDisPlayKubeTop()

	var KubeTopBody DisPlayKubeTop
	var top DisPlayKubeTopItem

	// bodyRes []map[string]string
	// debug := false
	// Containers := false
	// LabelSelector = "app.kubernetes.io/instance=kube-prometheus-0"

	// 获取pod list 数据
	var podItem service.PodResponse
	podItem.GetPodResponse(clientSet, namespace, LabelSelector, Containers, debug)
	podItemdb := podItem.PodResponse

	// 获取metric pod 数据
	var metricPodItem service.MetricData
	metricPodItem.PodMetric(mClientset, namespace, LabelSelector, Containers, debug)
	metricPodItemdb := metricPodItem.PodMetricData

	for i := 0; i < len(podItemdb); i++ {
		if !Containers {
			for _, v := range metricPodItemdb {
				top.ID = i
				top.Type = "POD"
				top.ResourceName = podItemdb[i].PodName
				top.Namespace = podItemdb[i].Namespace
				top.NodeName = podItemdb[i].NodeName
				top.LimitCPU = podItemdb[i].LimitCPU
				top.LimitMem = podItemdb[i].LimitMem
				top.RequestCPU = podItemdb[i].RequestCPU
				top.RequestMem = podItemdb[i].RequestMem
				top.ContainerName = ""
				if v.Namespace == podItemdb[i].Namespace && v.Name == podItemdb[i].PodName {
					top.UsageCPU = v.CPUSum
					top.UsageMem = v.MEMSum
				}
			}
			KubeTopBody.DisPlayKubeTop = append(KubeTopBody.DisPlayKubeTop, top)

		} else {
			top.ID = i
			top.Type = "container"
			top.ResourceName = podItemdb[i].PodName
			top.Namespace = podItemdb[i].Namespace
			top.NodeName = podItemdb[i].NodeName
			top.LimitCPU = podItemdb[i].SpecContainer.LimitCPU
			top.LimitMem = podItemdb[i].SpecContainer.LimitMem
			top.RequestCPU = podItemdb[i].SpecContainer.RequestCPU
			top.RequestMem = podItemdb[i].SpecContainer.RequestMem
			for _, v := range metricPodItemdb {
				if v.Namespace == podItemdb[i].Namespace && v.Name == podItemdb[i].PodName && v.ContainerName == podItemdb[i].SpecContainer.ContainerName {
					top.UsageCPU = v.Meta.CPU
					top.UsageMem = v.Meta.MEM
					top.ContainerName = v.ContainerName
				}
			}
			KubeTopBody.DisPlayKubeTop = append(KubeTopBody.DisPlayKubeTop, top)

		}
	}

	// 直接赋值的方式，struct 转换map，支持定制，性能和速度相对快一些

	for _, v := range KubeTopBody.DisPlayKubeTop {
		// kubeTopRes 初始化数据样式
		kubeTopRes := make(map[string]string, 0)

		kubeTopRes["ID"] = fmt.Sprint(v.ID)
		kubeTopRes["ResourceName"] = v.ResourceName
		kubeTopRes["Namespace"] = v.Namespace
		kubeTopRes["Type"] = v.Type
		kubeTopRes["ContainerName"] = v.ContainerName
		kubeTopRes["NodeName"] = v.NodeName
		kubeTopRes["RequestCPU"] = fmt.Sprint(util.StrconvFloat(v.RequestCPU, CPUUnitM))
		kubeTopRes["LimitCPU"] = fmt.Sprint(util.StrconvFloat(v.LimitCPU, CPUUnitM))
		kubeTopRes["RequestMem"] = fmt.Sprint(util.StrconvFloat(v.RequestMem, MemUnitM))
		kubeTopRes["LimitMem"] = fmt.Sprint(util.StrconvFloat(v.LimitMem, MemUnitM))

		extUsageCPU := UsageStrconvFloat(v.UsageCPU, v.RequestCPU, v.LimitCPU)
		extUsageMem := UsageStrconvFloat(v.UsageMem, v.RequestMem, v.LimitMem)

		// 剔除异常resource配置资源,参考函数：UsageStrconvFloat
		if extUsage {
			if extUsageCPU == 100 || extUsageMem == 100 {
				continue
			}
		}
		kubeTopRes["UsageCPU"] = fmt.Sprint(util.StrconvFloat(v.UsageCPU, CPUUnitM)) + "(" + fmt.Sprint(extUsageCPU) + ")"
		kubeTopRes["UsageMem"] = fmt.Sprint(util.StrconvFloat(v.UsageMem, MemUnitM)) + "(" + fmt.Sprint(extUsageMem) + ")"

		bodyRes = append(bodyRes, kubeTopRes)
	}
	return
}

// metric 使用率百分比：（UsageCPU - RequestCPU）/（LimitCPU - RequestCPU）
func UsageStrconvFloat(Usage, Request, Limit int64) (value float64) {
	if Limit == 0 && Request == 0 || Request > Limit { // 不设置resources的资源，占比100
		value = 100
	} else if Request > 0 && Request == Limit { // 对应高优先级，limit=request，计算usage/requesg
		value = util.StrconvFloat(Usage, Request)
	} else if Limit > 0 && Request == 0 { // 对应不设置request，只设置limit，则计算limit/requesg，算异常值
		value = util.StrconvFloat(Usage, Limit)
	} else {
		value = util.StrconvFloat(Usage-Request, Limit-Request)
	}
	return
}
