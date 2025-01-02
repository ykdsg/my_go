package main

import (
	"context"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	// 1.加载配置文件
	config, err := clientcmd.BuildConfigFromFlags("", "/Users/ykdsg/.kube/config")
	if err != nil {
		panic(err.Error())
	}
	// 2. 实例化客户端
	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	watch, err := clientSet.AppsV1().Deployments("default").Watch(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("watch start...")
	for {
		select {
		case event := <-watch.ResultChan():
			fmt.Println(event.Type, event.Object) // Type 表示事件变化类型,Object 表示变化后的数据
		}
	}
}
