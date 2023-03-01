/**
* @description :
* @author : Jarick
* @Date : 2022-07-10
* @Url : http://CloudWebOps
 */

/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"context"
	"fmt"
	"log"
	"reflect"

	"CloudDevKubernetes/controller"
	"CloudDevKubernetes/util"

	"github.com/spf13/cobra"
	kv1 "k8s.io/api/apps/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// kubederesCmd represents the kubederes command
var kubederesCmd = &cobra.Command{
	Use:   "kubedeploy",
	Short: kubedeployTitle,
	Long:  kubedeployTitle,
	RunE:  resourceDeploy,
}

func init() {

	// 添加新的flags
	kubederesCmd.Flags().BoolP("deployments", "d", false, "The Deployments Resources")
	kubederesCmd.Flags().BoolP("statefulsets", "f", false, "The Statefulsets Resources")
	kubederesCmd.Flags().BoolP("error", "e", false, errorFunction)

	rootCmd.AddCommand(kubederesCmd)
}

type DeployRes struct {
	NameSpace     string
	TYPE          string
	ResourceName  string
	ContainerName string
	Replices      string
	// ErrPod         int64
	RequestCPU int64
	RequestMem int64
	LimitCPU   int64
	LimitMem   int64
	IMAGE      string
}

func resourceDeploy(cmd *cobra.Command, _ []string) error {

	// 获取授权方式
	clientSet := controller.HandleToCluster()

	// 获取命令行参数
	ns, _ := cmd.Flags().GetString("namespace")
	json, _ := cmd.Flags().GetBool("json")
	errInfo, _ := cmd.Flags().GetBool("error")
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
	// defer fmt.Println(WarningInfo("kubernetesDefalutKind"))

	// 加载数据
	ResDeploy := ResourceBodyDeploy(clientSet, ns, flagtext, debug)

	var ResData []map[string]string
	// var ResData = make([]map[string]string, 0)
	// show debug error // 只显示错误信息
	if errInfo {
		ResDeployerr := make([]map[string]string, 0)
		for _, v := range ResDeploy {
			ss := new(util.StringTypeInt)
			err := ss.StringRepliceRes(v["Replices"], "float64", "()")
			if err != nil {
				log.Println(err)
			}
			if ss.Float64 == 0 {
				continue
			}
			ResDeployerr = append(ResDeployerr, v)

		}
		ResData = ResDeployerr
	} else {
		ResData = ResDeploy
	}

	// 生成表单的结构体
	vReflect := reflect.TypeOf(DeployRes{})

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
		}(desTurn, kubedeployTitle)
	}

	// 加载表单信息
	table := GenTable(vReflect, ResData, rmitem, debug)

	// json format
	if json {
		jsonStr, _ := table.Json(2)
		fmt.Println(jsonStr)
		// return nil     // cmds/cloudopscmd/cmd/kubedeploy.go:130:9: `if` block ends with a `return` statement, so drop this `else` and outdent its block (golint)
	} else {
		table.PrintTable()
	}

	if debug && errorMesg != "" {
		log.Println(errorMesg)
	}
	return nil
}

func ResourceBodyDeploy(clientSet *kubernetes.Clientset, ns string, flagtext string, debug bool) (bodyRes []map[string]string) {
	util.ErrPanicDebug(debug)

	// 生成一个全局资源列表
	var rList []interface{}

	// 判断访问资源类型
	if flagtext == "deployments" {
		deployList, err := clientSet.AppsV1().Deployments(ns).List(context.Background(), v1.ListOptions{})
		if err != nil {
			errorMesg = fmt.Sprint(err)
		}
		rList = append(rList, deployList)
	} else if flagtext == "statefulsets" {
		deployList, err := clientSet.AppsV1().StatefulSets(ns).List(context.Background(), v1.ListOptions{})

		if err != nil {
			errorMesg = fmt.Sprint(err)
		}
		rList = append(rList, deployList)
	} else {
		deployList0, err0 := clientSet.AppsV1().Deployments(ns).List(context.Background(), v1.ListOptions{})
		deployList1, err1 := clientSet.AppsV1().StatefulSets(ns).List(context.Background(), v1.ListOptions{})
		if err0 != nil || err1 != nil {
			errorMesg = fmt.Sprint(err0)
			errorMesg = fmt.Sprint(err1)
		}
		rList = append(rList, deployList0)
		rList = append(rList, deployList1)
	}

	deployMapList := make([]map[string]string, 0)
	for i := 0; i < len(rList); i++ {
		switch t := rList[i].(type) {
		case *kv1.DeploymentList:
			for k := 0; k < len(t.Items); k++ {
				for j := 0; j < len(t.Items[k].Spec.Template.Spec.Containers); j++ {
					deployMap := make(map[string]string)
					deployMap["NameSpace"] = ns
					if ns == "" {
						deployMap["NameSpace"] = t.Items[k].Namespace
					}
					deployMap["TYPE"] = "deployment"
					deployMap["ResourceName"] = t.Items[k].GetName()
					deployMap["ContainerName"] = t.Items[k].Spec.Template.Spec.Containers[j].Name
					deployMap["IMAGE"] = util.StringReplice(t.Items[k].Spec.Template.Spec.Containers[j].Image, "D")
					deployMap["Replices"] = fmt.Sprint(*t.Items[k].Spec.Replicas) + "(" + fmt.Sprint(*t.Items[k].Spec.Replicas-t.Items[k].Status.ReadyReplicas) + ")"
					// deployMap["ErrPod"] = fmt.Sprint(*t.Items[k].Spec.Replicas - t.Items[k].Status.ReadyReplicas)
					vauleLimitCPU := t.Items[k].Spec.Template.Spec.Containers[j].Resources.Limits.Cpu().MilliValue()
					vauleLimitMem := t.Items[k].Spec.Template.Spec.Containers[j].Resources.Limits.Memory().MilliValue()
					vauleRequestCPU := t.Items[k].Spec.Template.Spec.Containers[j].Resources.Requests.Cpu().MilliValue()
					vauleRequestMem := t.Items[k].Spec.Template.Spec.Containers[j].Resources.Requests.Memory().MilliValue()
					deployMap["LimitCPU"] = fmt.Sprint(vauleLimitCPU) + "(" + fmt.Sprint(util.StrconvFloat(vauleLimitCPU, vauleRequestCPU)) + ")"
					deployMap["LimitMem"] = fmt.Sprint(vauleLimitMem/MemUnitM) + "(" + fmt.Sprint(util.StrconvFloat(vauleLimitMem, vauleRequestMem)) + ")"
					deployMap["RequestCPU"] = fmt.Sprint(vauleRequestCPU)
					deployMap["RequestMem"] = fmt.Sprint(vauleRequestMem / MemUnitM)
					deployMapList = append(deployMapList, deployMap)
				}
			}
		case *kv1.StatefulSetList:
			for k := 0; k < len(t.Items); k++ {
				for j := 0; j < len(t.Items[k].Spec.Template.Spec.Containers); j++ {
					deployMap := make(map[string]string)
					deployMap["NameSpace"] = ns
					if ns == "" {
						deployMap["NameSpace"] = t.Items[k].Namespace
					}
					deployMap["TYPE"] = "statefulset"
					deployMap["ResourceName"] = t.Items[k].GetName()
					deployMap["ContainerName"] = t.Items[k].Spec.Template.Spec.Containers[j].Name
					deployMap["IMAGE"] = util.StringReplice(t.Items[k].Spec.Template.Spec.Containers[j].Image, "D")
					deployMap["Replices"] = fmt.Sprint(*t.Items[k].Spec.Replicas) + "(" + fmt.Sprint(*t.Items[k].Spec.Replicas-t.Items[k].Status.ReadyReplicas) + ")"
					// deployMap["ErrPod"] = fmt.Sprint(*t.Items[k].Spec.Replicas - t.Items[k].Status.ReadyReplicas)
					vauleLimitCPU := t.Items[k].Spec.Template.Spec.Containers[j].Resources.Limits.Cpu().MilliValue()
					vauleLimitMem := t.Items[k].Spec.Template.Spec.Containers[j].Resources.Limits.Memory().MilliValue()
					vauleRequestCPU := t.Items[k].Spec.Template.Spec.Containers[j].Resources.Requests.Cpu().MilliValue()
					vauleRequestMem := t.Items[k].Spec.Template.Spec.Containers[j].Resources.Requests.Memory().MilliValue()
					deployMap["LimitCPU"] = fmt.Sprint(vauleLimitCPU) + "(" + fmt.Sprint(util.StrconvFloat(vauleLimitCPU, vauleRequestCPU)) + ")"
					deployMap["LimitMem"] = fmt.Sprint(vauleLimitMem/MemUnitM) + "(" + fmt.Sprint(util.StrconvFloat(vauleLimitMem, vauleRequestMem)) + ")"
					deployMap["RequestCPU"] = fmt.Sprint(vauleRequestCPU)
					deployMap["RequestMem"] = fmt.Sprint(vauleRequestMem / MemUnitM)
					deployMapList = append(deployMapList, deployMap)
				}
			}
		}
	}
	bodyRes = deployMapList
	return
}
