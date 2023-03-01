/**
* @description :
* @author : Jarick
* @Date : 2022-12-11
* @Url : http://CloudWebOps
 */

package service

import (
	"context"
	"log"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	metrics "k8s.io/metrics/pkg/client/clientset/versioned"
)

// metric 数据结构实现方式
type MetricData struct {
	NodeMetricData []NodeMetricBody
	PodMetricData  []PodMetricBody
	errMsg         []error
}

type NodeMetricBody struct {
	NodeName string
	Meta     MetricMeta
}
type PodMetricBody struct {
	Name          string
	Namespace     string
	CPUSum        int64
	MEMSum        int64
	ContainerName string
	Meta          MetricMeta
}
type MetricMeta struct {
	CPU int64
	MEM int64
}

// init NewMetricData
func NewMetricData() *MetricData {
	return &MetricData{} // 初始化结构体
}

// 获取node metric数据
func (nm *MetricData) NodeMetric(mClientset *metrics.Clientset, nodeLabels string, debug bool) {

	// init NewMetric
	NewMetricData()

	nmList, err := mClientset.MetricsV1beta1().NodeMetricses().List(context.TODO(), metav1.ListOptions{
		LabelSelector: nodeLabels,
	})
	if err != nil {
		if debug {
			log.Printf("Error getting Node Metrics: %v\n", err)
			log.Println("For this to work, metrics-server needs to be running in your cluster")
		}
		// os.Exit(7)
		nm.errMsg = append(nm.errMsg, err)
	}

	for i := 0; i < len(nmList.Items); i++ {
		var y NodeMetricBody
		y.NodeName = nmList.Items[i].Name
		y.Meta.CPU = nmList.Items[i].Usage.Cpu().MilliValue()
		y.Meta.MEM = nmList.Items[i].Usage.Memory().MilliValue()
		nm.NodeMetricData = append(nm.NodeMetricData, y)
	}

}

// 获取pod metric数据
func (nm *MetricData) PodMetric(mClientset *metrics.Clientset, namespace string, LabelSelector string, Containers bool, debug bool) {

	// init NewMetric
	NewMetricData()

	pmList, err := mClientset.MetricsV1beta1().PodMetricses(namespace).List(context.TODO(), metav1.ListOptions{
		LabelSelector: LabelSelector, // app.kubernetes.io/instance=kube-prometheus-0
	})

	if err != nil {
		if debug {
			log.Printf("Error getting Pod Metrics: %v\n", err)
			log.Println("For this to work, metrics-server needs to be running in your cluster")
			// os.Exit(6)
		}
		nm.errMsg = append(nm.errMsg, err)
	}

	y := new(PodMetricBody)
	for i := 0; i < len(pmList.Items); i++ {

		if Containers {
			for j := 0; j < len(pmList.Items[i].Containers); j++ {
				y.Name = pmList.Items[i].Name
				y.Namespace = pmList.Items[i].Namespace
				y.ContainerName = pmList.Items[i].Containers[j].Name
				y.Meta.CPU = pmList.Items[i].Containers[j].Usage.Cpu().MilliValue()
				y.Meta.MEM = pmList.Items[i].Containers[j].Usage.Memory().MilliValue()
				nm.PodMetricData = append(nm.PodMetricData, *y)
			}
		} else {
			// init CPUSum,MEMSum
			CPUSum := new(int64)
			MEMSum := new(int64)
			for j := 0; j < len(pmList.Items[i].Containers); j++ {
				func() {
					*CPUSum = *CPUSum + pmList.Items[i].Containers[j].Usage.Cpu().MilliValue()
					*MEMSum = *MEMSum + pmList.Items[i].Containers[j].Usage.Memory().MilliValue()
				}()
			}
			y.Name = pmList.Items[i].Name
			y.Namespace = pmList.Items[i].Namespace
			y.CPUSum = *CPUSum
			y.MEMSum = *MEMSum
			nm.PodMetricData = append(nm.PodMetricData, *y)
		}
	}
}

// metric 提供第三方API接口，用于给其他模块调用
func GetMetricDBHandle(mClientset *metrics.Clientset, metricType string, nodeLabels string, namespace string, LabelSelector string, Containers bool, debug bool) *MetricData {
	// init NewMetric
	NewMetricData()

	MetricDB := new(MetricData)

	if metricType == "nodeMetric" {
		MetricDB.NodeMetric(mClientset, nodeLabels, debug)
	} else if metricType == "podMetric" {
		MetricDB.PodMetric(mClientset, namespace, LabelSelector, Containers, debug)
	}
	return MetricDB
}
