package main

import (
	"clientgo/common"
	"context"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

/*
clientSet 基于RESTClient 的封装，可对k8s 内置的资源对象进行操作
*/
func main() {
	// 1.加载配置文件
	config, err := clientcmd.BuildConfigFromFlags("", "/Users/ykdsg/.kube/config")
	if err != nil {
		panic(err.Error())
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	pods, err := clientset.CoreV1().Pods("default").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}

	for _, pod := range pods.Items {
		common.ProcessPod(pod)
	}
}
