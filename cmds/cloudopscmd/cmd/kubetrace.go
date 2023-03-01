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
	"net"
	"reflect"
	"time"

	"github.com/spf13/cobra"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// kubetraceCmd represents the kubetrace command
var kubetraceCmd = &cobra.Command{
	Use:   "kubetrace",
	Short: kubetraceTitle,
	Long:  kubetraceTitle,
	RunE:  kubetrace,
}

func init() {

	// 添加新的flags
	kubetraceCmd.Flags().BoolP("error", "e", false, errorFunction)
	kubetraceCmd.Flags().BoolP("service", "", false, "The Trace Service Info")
	kubetraceCmd.Flags().BoolP("endpoint", "", false, "The Trace Endpoint Info")
	kubetraceCmd.Flags().Int64("timeout", 2, "Setting Service Trace Timeout")

	rootCmd.AddCommand(kubetraceCmd)
}

func kubetrace(cmd *cobra.Command, _ []string) error {
	// 获取命令行参数
	namespace, _ := cmd.Flags().GetString("namespace")
	timeout, _ := cmd.Flags().GetInt64("timeout")
	json, _ := cmd.Flags().GetBool("json")
	service, _ := cmd.Flags().GetBool("service")
	endpoint, _ := cmd.Flags().GetBool("endpoint")
	errInfo, _ := cmd.Flags().GetBool("error")
	debug, _ := cmd.Flags().GetBool("debug")
	desTurn, _ := cmd.Flags().GetBool("des")
	rmitem, _ := cmd.Flags().GetStringSlice("rmitem")

	// 获取授权方式
	clientSet := controller.HandleToCluster()

	resBody := serverTrace(clientSet, namespace, timeout, errInfo, service, endpoint, debug)

	// 判断返回数据
	if len(resBody) == 0 {
		if errorMesg == "" {
			defer fmt.Printf("No resource displayed.\n")
		} else {
			defer fmt.Printf("No resource displayed. please check environment for kubernetes. \n[ERROR]: %s\n", errorMesg)
		}
	} else {
		defer func(desTurn bool, Title string) {
			if desTurn {
				fmt.Println(Title + "\n" + WarningInfo("kubetraceContextInfo0", desTurn) + WarningInfo("Unittext", desTurn))
			}
		}(desTurn, kubetraceTitle)
	}
	// 生成表头结构体
	vReflect := reflect.TypeOf(traceBody{})
	// 加载表单信息
	table := GenTable(vReflect, resBody, rmitem, debug)

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

func serverTrace(clientSet *kubernetes.Clientset, namespaces string, timeout int64, errInfo bool, service bool, endpoint bool, debug bool) (bodyRes []map[string]string) {
	// https://vimsky.com/examples/detail/golang-ex-net---DialTimeout-function.html
	// https://vimsky.com/examples/detail/golang-ex-net---DialIP-function.html
	util.ErrPanicDebug(debug)

	// 加载数据模块
	ServiceTraceBody := serviceGet(clientSet, namespaces)
	EndpointTraceBody := endpointGet(clientSet, namespaces)

	// 汇总数据集
	var TraceBody []traceBody
	if service {
		TraceBody = append(TraceBody, ServiceTraceBody...)
	} else if endpoint {
		TraceBody = append(TraceBody, EndpointTraceBody...)
	} else {
		TraceBody = append(TraceBody, ServiceTraceBody...)
		TraceBody = append(TraceBody, EndpointTraceBody...)
	}

	// testData
	// #########################
	// exampleData := traceBody{
	// 	Name:      "nameTest",
	// 	Namespace: "namespaceTest",
	// 	TraceName: "service",
	// 	PodName:   "podTest",
	// 	PortList:  []string{"127.0.0.1:49153", "127.0.0.1:49154", "127.0.0.1:49155", "127.0.0.1:49156"},
	// }
	// TraceBody = append(TraceBody, exampleData)
	// #########################

	for _, v := range TraceBody {

		Name := v.Name
		Namespace := v.Namespace
		TraceName := v.TraceName

		for i := 0; i < len(v.PortList); i++ {
			bodyResTrace := make(map[string]string, 0)
			var PortStr string
			var NodeStr string
			var PodStr string
			if v.TraceName == "endpoint" {
				PortStr = util.StringReplice(v.PortList[i], "A")
				NodeStr = util.StringReplice(v.PortList[i], "B")
				PodStr = util.StringReplice(v.PortList[i], "C")

			} else {
				PortStr = v.PortList[i]
				NodeStr = "null"
				PodStr = "null"
			}

			// 核心探测程序
			// start := time.Now() //耗时： 开始

			// conn, err := net.DialTimeout("tcp", PortStr, util.TimeFuncDuration(timeout, "Second"))

			// if err != nil {
			// 	if debug {
			// 		log.Println(err) // bug ：socket too many file  // pid=`pidof cloudopscmd` ;cat /proc/$pid/net/socket
			// 	}
			// } else {
			// 	conn.Close()
			// }
			// end := time.Since(start) //耗时： 结束

			// var TraceStatus string
			// if err != nil {
			// 	TraceStatus = "fail[" + fmt.Sprintf("%.2f", since.Seconds()) + "s]:" + fmt.Sprint(err)
			// } else {
			// 	if errInfo {
			// 		continue
			// 	} else {
			// 		TraceStatus = "pass[" + fmt.Sprintf("%.2f", since.Seconds()) + "s]"
			// 	}

			// fix + function: start
			_, TraceStatus, _ := func(PortStr string) (conn net.Conn, TraceStatus string, err error) {
				start := time.Now() //耗时： 开始
				conn, err = net.DialTimeout("tcp", PortStr, util.TimeFuncDuration(timeout, "Second"))
				since := time.Since(start) //耗时： 结束

				if err != nil {
					TraceStatus = "fail[" + fmt.Sprintf("%.2f", since.Seconds()) + "s]:" + fmt.Sprint(err)
					if debug {
						log.Println(err) // bug ：socket too many file  // pid=`pidof cloudopscmd` ;cat /proc/$pid/net/socket
					}
				} else {
					if errC := conn.Close(); errC != nil {
						log.Println(errC)
					}
					if !errInfo {
						TraceStatus = "pass[" + fmt.Sprintf("%.2f", since.Seconds()) + "s]"
					}
				}
				return
			}(PortStr)
			// fix + function: end

			// error 开关
			if errInfo && TraceStatus == "" {
				continue
			} else {
				bodyResTrace["Name"] = Name
				bodyResTrace["Namespace"] = Namespace
				bodyResTrace["TraceName"] = TraceName
				bodyResTrace["NodeName"] = NodeStr
				bodyResTrace["PortList"] = PortStr
				bodyResTrace["Status"] = TraceStatus
				bodyResTrace["PodName"] = PodStr
				bodyRes = append(bodyRes, bodyResTrace)
			}
		}
	}
	return
}

type traceBody struct {
	TraceName string
	Name      string
	Namespace string
	PodName   string
	PortList  []string
	NodeName  string
	Status    string
}

// 初始化
func newtraceBody(Name string, Namespace string, PortList []string, TraceName string) *traceBody {
	return &traceBody{Name: Name, Namespace: Namespace, PortList: PortList, TraceName: TraceName}
}

func serviceGet(clientset *kubernetes.Clientset, namespace string) (ServiceList []traceBody) {

	ServiceList = make([]traceBody, 0)

	service, err := clientset.CoreV1().Services(namespace).List(context.TODO(), metav1.ListOptions{})

	if err != nil {
		log.Println(err)
	}

	for _, v := range service.Items {
		PortList := make([]string, 0)
		ClusterIP := v.Spec.ClusterIP
		if ClusterIP == "None" {
			continue
		}
		for _, v0 := range v.Spec.Ports {
			if v0.Protocol == "UDP" {
				continue
			}
			PortList = append(PortList, ClusterIP+":"+fmt.Sprint(v0.Port))
		}
		ServiceList = append(ServiceList, *newtraceBody(v.Name, v.Namespace, PortList, "service"))
	}
	return
}

func endpointGet(clientset *kubernetes.Clientset, namespace string) (EndpointList []traceBody) {
	EndpointList = make([]traceBody, 0)
	endpoint, err := clientset.CoreV1().Endpoints(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Println(err)
	}

	for _, v := range endpoint.Items {

		PortList := make([]string, 0)
		if len(v.Subsets) < 1 {
			continue
		}

		for _, v0 := range v.Subsets[0].Addresses {
			if nilPodName := v0.TargetRef; nilPodName == nil {
				continue
			}
			for _, v1 := range v.Subsets[0].Ports {
				if v1.Protocol == "UDP" {
					continue
				}
				if nilNodeName := v0.NodeName; nilNodeName == nil {
					continue
				}
				PortList = append(PortList, v0.IP+":"+fmt.Sprint(v1.Port)+"/"+fmt.Sprint(*v0.NodeName)+"#"+fmt.Sprint(v0.TargetRef.Name))
			}

		}
		EndpointList = append(EndpointList, *newtraceBody(v.Name, v.Namespace, PortList, "endpoint"))
	}
	return

}
