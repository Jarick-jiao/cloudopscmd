/**
* @description :
* @author : Jarick
* @Date : 2022-12-11
* @Url : http://CloudWebOps
 */

package service

import (
	"CloudDevKubernetes/util"
	"context"
	"log"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"

	corev1 "k8s.io/api/core/v1"

	resourcehelper "k8s.io/kubectl/pkg/util/resource"
)

// pod list  数据结构实现方式
// podRespone Handle 提供第三方API接口，用于给其他模块调用
type PodResponse struct {
	PodResponse []PodResponseItem
}

type PodResponseItem struct {
	PodName       string
	Namespace     string
	NodeName      string
	RequestCPU    int64
	LimitCPU      int64
	RequestMem    int64
	LimitMem      int64
	SpecContainer ContainerResponseItem
}
type ContainerResponseItem struct {
	ContainerName string
	RequestCPU    int64
	LimitCPU      int64
	RequestMem    int64
	LimitMem      int64
}

func NewPodResponse() *PodResponse {
	return &PodResponse{}
}

func (p *PodResponse) GetPodResponse(clientSet *kubernetes.Clientset, namespace string, LabelSelector string, Containers bool, debug bool) (PodListRes *corev1.PodList) {
	// ns string, LabelSelector string,
	util.ErrPanicDebug(debug)

	NewPodResponse()

	PodListRes, err := clientSet.CoreV1().Pods(namespace).List(context.Background(), v1.ListOptions{
		LabelSelector: LabelSelector,
	})
	if err != nil {
		log.Println(err)
		if debug {
			log.Panic(err)
		}
	}

	podItem := new(PodResponseItem)

	for _, podv := range PodListRes.Items {

		if !Containers {
			podCore := podv                                             // Using a reference for the variable on range scope `pod` (scopelint)
			req, limit := resourcehelper.PodRequestsAndLimits(&podCore) // 官方工具获取资源信息
			cpuReq, cpuLimit, memoryReq, memoryLimit := req[corev1.ResourceCPU], limit[corev1.ResourceCPU], req[corev1.ResourceMemory], limit[corev1.ResourceMemory]

			podItem.LimitCPU = cpuLimit.MilliValue()
			podItem.LimitMem = memoryLimit.MilliValue()
			podItem.RequestCPU = cpuReq.MilliValue()
			podItem.RequestMem = memoryReq.MilliValue()
			podItem.PodName = podv.Name
			podItem.Namespace = podv.Namespace
			podItem.NodeName = podv.Spec.NodeName

			p.PodResponse = append(p.PodResponse, *podItem)
		} else {
			for i := 0; i < len(podv.Spec.Containers); i++ {
				podItem.PodName = podv.Name
				podItem.Namespace = podv.Namespace
				podItem.NodeName = podv.Spec.NodeName
				podItem.SpecContainer.ContainerName = podv.Spec.Containers[i].Name
				podItem.SpecContainer.LimitCPU = podv.Spec.Containers[i].Resources.Limits.Cpu().MilliValue()
				podItem.SpecContainer.LimitMem = podv.Spec.Containers[i].Resources.Limits.Memory().MilliValue()
				podItem.SpecContainer.RequestCPU = podv.Spec.Containers[i].Resources.Requests.Cpu().MilliValue()
				podItem.SpecContainer.RequestMem = podv.Spec.Containers[i].Resources.Requests.Memory().MilliValue()
				p.PodResponse = append(p.PodResponse, *podItem)
			}
		}
	}
	return
}
