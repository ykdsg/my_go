package main

import (
	"clientgo/common"
	"context"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	// 1.加载配置文件
	config, err := clientcmd.BuildConfigFromFlags("", "/Users/ykdsg/.kube/config")
	if err != nil {
		panic(err.Error())
	}
	dynamicClient, err := dynamic.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	// 设置要请求的 GVR
	gvr := schema.GroupVersionResource{
		Group:    "",
		Version:  "v1",
		Resource: "pods",
	}
	// 发送请求，并得到返回结果
	unstructuredList, err := dynamicClient.Resource(gvr).Namespace("default").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}

	// 使用反射将 unstructuredList 的数据转成对应的结构体类型，例如这是是转成 v1.PodList 类型
	podList := &corev1.PodList{}
	err = runtime.DefaultUnstructuredConverter.FromUnstructured(unstructuredList.UnstructuredContent(), podList)
	if err != nil {
		panic(err.Error())
	}

	for _, pod := range podList.Items {
		common.ProcessPod(pod)
	}

}
