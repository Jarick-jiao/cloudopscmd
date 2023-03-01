/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"CloudDevKubernetes/controller"
	service "CloudDevKubernetes/service/kubernetesHandle"
	"CloudDevKubernetes/util"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"reflect"

	"github.com/spf13/cobra"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	resourcehelper "k8s.io/kubectl/pkg/util/resource"
	metrics "k8s.io/metrics/pkg/client/clientset/versioned"
)

// kuberesCmd represents the kuberes command
var kuberesCmd = &cobra.Command{
	Use:   "kuberes",
	Short: kuberesTitle,
	Long:  kuberesTitle,
	RunE:  resourcesNode,
}

func init() {

	// 创建子命令
	kuberesCmd.Flags().BoolP("error", "e", false, errorFunction)
	kuberesCmd.Flags().BoolP("usage", "u", false, "Output The resources Usage Info")

	rootCmd.AddCommand(kuberesCmd)
}

var ResAPIHandle *DisplayKubeRes

func resourcesNode(cmd *cobra.Command, _ []string) error {

	// 获取命令行参数
	ns, _ := cmd.Flags().GetString("namespace")
	errInfo, _ := cmd.Flags().GetBool("error")
	usageAction, _ := cmd.Flags().GetBool("usage")
	json, _ := cmd.Flags().GetBool("json")
	debug, _ := cmd.Flags().GetBool("debug")
	desTurn, _ := cmd.Flags().GetBool("des")
	rmitem, _ := cmd.Flags().GetStringSlice("rmitem")

	// 获取授权方式
	clientSet, mClientset := controller.HandleToK8sAndMetric(debug)

	// var Cache ResCache
	// Cache.ResCacheReq(clientSet, ns, debug)
	// 获取查询指标数据
	ResNode, ResAPINode := KubeResHandle(clientSet, mClientset, ns, debug, usageAction, json)

	// 申明展示数据样式
	// var ResData = make([]map[string]string, 0)
	var ResData []map[string]string

	// show debug error // 只显示错误信息
	if errInfo {
		ResNodeerr := make([]map[string]string, 0)
		for _, v := range ResNode {
			ss := new(util.StringTypeInt)
			err := ss.StringRepliceRes(v["PodCal"], "float64", "()")
			if err != nil {
				log.Println(err)
			}
			if ss.Float64 == 0 {
				continue
			}
			ResNodeerr = append(ResNodeerr, v)
		}
		ResData = ResNodeerr
	} else {
		ResData = ResNode
	}

	// 判断返回数据
	if len(ResData) == 0 {
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
		}(desTurn, kuberesTitle)
	}

	// 生成表头结构体
	vReflect := reflect.TypeOf(DisplayKubeResItem{})

	// 加载表单信息
	if !usageAction {
		rmitem = append(rmitem, "UsageCPU", "UsageMem", "ErrPod", "ErrCount")
	} else {
		rmitem = append(rmitem, "NodeTaints", "ImageSize", "ErrPod", "ErrCount")
	}
	table := GenTable(vReflect, ResData, rmitem, debug)

	// 命令行输出选项
	if json { //json
		// jsonStr, _ := table.Json(2)
		// fmt.Print(string(jsonStr))
		fmt.Print(string(ResAPINode))
	} else {
		table.PrintTable()
	}
	// if debug && errorMesg != "" {
	// 	log.Println(errorMesg)
	// }
	util.DebugPrint(debug, errorMesg)
	return nil
}

type DisplayKubeRes struct {
	DisplayKubeResItem []DisplayKubeResItem
}

type DisplayKubeResItem struct {
	NodeName   string
	PodCal     int64
	ErrCount   int64
	CapMem     int64
	CapCPU     int64
	RequestCPU int64
	LimitCPU   int64
	UsageCPU   int64
	RequestMem int64
	LimitMem   int64
	UsageMem   int64
	NodeTaints string
	ImageSize  int64
}

// init resources
func NewDisplayKubeResHandle() *DisplayKubeRes {
	if ResAPIHandle == nil {
		ResAPIHandle = new(DisplayKubeRes)
	}
	return &DisplayKubeRes{
		DisplayKubeResItem: make([]DisplayKubeResItem, 0),
	}
}

// 请求API获取node，pod 所有信息

// cache registry
type ResCache struct {
	nodeinfo *corev1.NodeList
	PODList  *corev1.PodList
	errMsg   []error
}

// http get res for pod and node resource
func (c *ResCache) ResCacheReq(clientSet *kubernetes.Clientset, ns string, debug bool) {
	util.ErrPanicDebug(debug)

	var err error

	if c.nodeinfo == nil {
		c.nodeinfo, err = clientSet.CoreV1().Nodes().List(context.Background(), v1.ListOptions{})
		if err != nil {
			c.errMsg = append(c.errMsg, err)
			errorMesg = fmt.Sprintln(err)
		}

	}
	if c.PODList == nil {
		c.PODList, err = clientSet.CoreV1().Pods(ns).List(context.Background(), v1.ListOptions{})
		if err != nil {
			c.errMsg = append(c.errMsg, err)
			errorMesg = fmt.Sprintln(err)
		}
	}
}

func KubeResHandle(clientSet *kubernetes.Clientset, mClientset *metrics.Clientset, ns string, debug bool, usageAction bool, jsonAction bool) (bodyRes []map[string]string, ResAPINode []byte) {
	util.ErrPanicDebug(debug)

	// 初始化
	NewDisplayKubeResHandle()

	var nodeRes []map[string]string

	// 加载缓存数据
	var Cache ResCache
	// nodeinfo, err := clientSet.CoreV1().Nodes().List(context.Background(), v1.ListOptions{})
	if Cache.nodeinfo == nil {
		Cache.ResCacheReq(clientSet, ns, debug)
	}

	// 获取node，pod 资源信息
	nodeinfo := Cache.nodeinfo
	PODList := Cache.PODList

	// 获取util资源
	var MetricDB *service.MetricData
	if usageAction {
		MetricDB = service.GetMetricDBHandle(mClientset, "nodeMetric", "", ns, "", false, debug)
	}

	for _, v := range nodeinfo.Items {

		// init res struct
		// CIL 标准输出格式
		nodesingleRes := make(map[string]string, 0)
		nodeResources := new(DisplayKubeResItem) // 针对每一个节点，初始化一次

		// 计算所有POD资源总量
		nodeResources.PodCals(PODList, v.Name)

		// 1、节点容量资源量大小
		// vaule_Allocatable_cpu := v.Status.Allocatable.Cpu().MilliValue()
		// vaule_Allocatableem := v.Status.Allocatable.Memory().MilliValue()

		// 2、镜像卷大小
		var vauleImageSize int64 = 0
		for _, v := range v.Status.Images {
			vauleImageSize = vauleImageSize + v.SizeBytes
		}

		// 3、节点可用标识
		// var vauleNodeTaints string
		// if len(v.Spec.Taints) > 1 {
		// 	vauleNodeTaints = "fail"
		// } else if len(v.Spec.Taints) == 1 && v.Spec.Taints[0].Key == "node-role.kubernetes.io/master" { //判断master污点
		// 	vauleNodeTaints = "success"
		// } else { //不存在污点的
		// 	vauleNodeTaints = "success"
		// }

		vauleNodeTaints := func(sliceTaints []corev1.Taint) (StatusTaints string) {
			if len(sliceTaints) > 1 {
				StatusTaints = "fail"
			} else if len(sliceTaints) == 1 && sliceTaints[0].Key == "node.kubernetes.io/unreachable" {
				StatusTaints = "fail"
			} else if len(sliceTaints) == 1 && sliceTaints[0].Key == "node-role.kubernetes.io/master" { //判断master污点
				StatusTaints = "success"
			} else { //不存在污点的
				StatusTaints = "success"
			}
			return
		}(v.Spec.Taints)

		// 4.0 获取节点metric值
		var mnNodeMetricBody service.NodeMetricBody
		if usageAction {
			mnNodeMetricBody = func(nodeName string) (v service.NodeMetricBody) {
				for _, v = range MetricDB.NodeMetricData {
					if v.NodeName == nodeName {
						return v
					}
				}
				return
			}(v.Name)
		}
		// 4.1、node节点指标赋值汇总
		// nodeResources.AllocatableCpu = vaule_Allocatable_cpu
		// nodeResources.AllocatableMem = vaule_Allocatableem
		nodeResources.NodeName = v.Name
		nodeResources.CapCPU = v.Status.Capacity.Cpu().MilliValue()
		nodeResources.CapMem = v.Status.Capacity.Memory().MilliValue()
		nodeResources.NodeTaints = vauleNodeTaints
		nodeResources.ImageSize = vauleImageSize
		nodeResources.UsageCPU = mnNodeMetricBody.Meta.CPU
		nodeResources.UsageMem = mnNodeMetricBody.Meta.MEM

		// 5、CLI输出格式
		nodesingleRes["CapCPU"] = fmt.Sprint(util.StrconvFloat(nodeResources.CapCPU, CPUUnitC))
		nodesingleRes["CapMem"] = fmt.Sprint(util.StrconvFloat(nodeResources.CapMem, MemUnitG))
		nodesingleRes["LimitCPU"] = fmt.Sprint(util.StrconvFloat(nodeResources.LimitCPU, CPUUnitM)) + "(" + fmt.Sprint(util.StrconvFloat(nodeResources.LimitCPU, nodeResources.RequestCPU)) + ")"
		nodesingleRes["LimitMem"] = fmt.Sprint(util.StrconvFloat(nodeResources.LimitMem, MemUnitM)) + "(" + fmt.Sprint(util.StrconvFloat(nodeResources.LimitMem, nodeResources.RequestMem)) + ")"
		nodesingleRes["RequestCPU"] = fmt.Sprint(util.StrconvFloat(nodeResources.RequestCPU, CPUUnitM)) + "(" + fmt.Sprint(util.StrconvFloat(nodeResources.RequestCPU, nodeResources.CapCPU)) + ")"
		nodesingleRes["RequestMem"] = fmt.Sprint(util.StrconvFloat(nodeResources.RequestMem, MemUnitM)) + "(" + fmt.Sprint(util.StrconvFloat(nodeResources.RequestMem, nodeResources.CapMem)) + ")"
		nodesingleRes["PodCal"] = fmt.Sprint(nodeResources.PodCal) + "(" + fmt.Sprint(nodeResources.ErrCount) + ")"
		nodesingleRes["NodeName"] = nodeResources.NodeName
		nodesingleRes["NodeTaints"] = fmt.Sprint(nodeResources.NodeTaints)
		nodesingleRes["ImageSize"] = fmt.Sprint(util.StrconvFloat(nodeResources.ImageSize, DisUnitM))
		nodesingleRes["UsageCPU"] = fmt.Sprint(util.StrconvFloat(nodeResources.UsageCPU, CPUUnitM)) + "(" + fmt.Sprint(util.StrconvFloat(nodeResources.UsageCPU, nodeResources.CapCPU)) + ")"
		nodesingleRes["UsageMem"] = fmt.Sprint(util.StrconvFloat(nodeResources.UsageMem, MemUnitM)) + "(" + fmt.Sprint(util.StrconvFloat(nodeResources.UsageMem, nodeResources.CapMem)) + ")"
		nodeRes = append(nodeRes, nodesingleRes)

		// 6、API标准输出格式
		ResAPIHandle.DisplayKubeResItem = append(ResAPIHandle.DisplayKubeResItem, *nodeResources)
	}

	// API输出-json

	if jsonAction {
		ResAPINode, _ = json.Marshal(&ResAPIHandle.DisplayKubeResItem)
		// if err != nil {
		// 	errorMesg = fmt.Sprint(err)
		// }
	} else {
		ResAPINode = nil
	}

	// CLI输出
	bodyRes = nodeRes
	return
}

// node资源求和
func (r *DisplayKubeResItem) calNodeResources(LimitCPUValue, LimitMemValue, RequestCPUValue, RequestMemValue, ErrCountValue, PodCalValue int64) {
	r.LimitCPU = r.LimitCPU + LimitCPUValue
	r.LimitMem = r.LimitMem + LimitMemValue
	r.RequestCPU = r.RequestCPU + RequestCPUValue
	r.RequestMem = r.RequestMem + RequestMemValue
	r.ErrCount = r.ErrCount + ErrCountValue
	r.PodCal = r.PodCal + PodCalValue
}

// pod维度资源归类和总计
func (r *DisplayKubeResItem) PodCals(PODList *corev1.PodList, NodeName string) {
	// func (r *NodeClsResources) PodCals(clientSet *kubernetes.Clientset, ns string, NodeName string, debug bool) {

	// PODList = Cache.PODList
	for _, pod := range PODList.Items {

		podCore := pod // Using a reference for the variable on range scope `pod` (scopelint)

		if pod.Spec.NodeName != NodeName { //bug： fix nodename ==！ hostip  // 遍历节点
			continue
		}

		// 判断POD状态计数器
		var podNumValue int64 = 0
		var ErrCountValue int64 = 0
		var cpuReq, cpuLimit, memoryReq, memoryLimit resource.Quantity

		// 检查POD有异常的Container，则作为ErrCount计数
		PodContainerFalseStatuses := true
		for i := 0; i < len(pod.Status.ContainerStatuses); i++ {
			if pod.Status.ContainerStatuses[i].Ready {
				PodContainerFalseStatuses = false
				continue
			}
		}

		// pod状态归类
		if pod.Status.Phase == "Running" { // fix: pod.Status.Phase == "Succeeded"

			podNumValue = +1

			req, limit := resourcehelper.PodRequestsAndLimits(&podCore) // 官方工具获取资源信息
			cpuReq, cpuLimit, memoryReq, memoryLimit = req[corev1.ResourceCPU], limit[corev1.ResourceCPU], req[corev1.ResourceMemory], limit[corev1.ResourceMemory]

		} else if pod.Status.Phase == "Pending" { // fix pending 状态需要算占用资源。和describe node相比

			ErrCountValue = +1
			req, limit := resourcehelper.PodRequestsAndLimits(&podCore) // 官方工具获取资源信息
			cpuReq, cpuLimit, memoryReq, memoryLimit = req[corev1.ResourceCPU], limit[corev1.ResourceCPU], req[corev1.ResourceMemory], limit[corev1.ResourceMemory]

		} else if pod.Status.Phase == "Succeeded" { //fix: PodContainerFalseStatuses and pod.Status.Phase == "Succeeded"
			continue
		} else if pod.Status.Phase == "Failed" || pod.Status.Phase == "Unknown" || PodContainerFalseStatuses { // fix 该状态不需要核算资源量

			ErrCountValue = +1
			// fmt.Println("xxxxxxxxxxxxxxxxxx:", pod.Name, pod.Status.Phase)
		} else {
			continue
		}

		// 汇总函数
		r.calNodeResources(cpuLimit.MilliValue(), memoryLimit.MilliValue(), cpuReq.MilliValue(), memoryReq.MilliValue(), ErrCountValue, podNumValue)
	}
}
