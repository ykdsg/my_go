package main

import (
	"fmt"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"time"
)

func main() {
	// 	使用clientSet 演示 sharedInformer
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
	// 	初始化informer
	sharedInformerFactory := informers.NewSharedInformerFactory(clientSet, time.Second*30)

	// 查询pod 数据，生成PodInformer 对象
	podInformer := sharedInformerFactory.Core().V1().Pods()

	// 生成indexer，便于数据查询
	indexer := podInformer.Lister()

	// 	启动informer
	sharedInformerFactory.Start(nil)
	// 	等待数据同步完成
	sharedInformerFactory.WaitForCacheSync(nil)

	// 	利用indexer 获取数据
	pods, err := indexer.List(labels.Everything())
	if err != nil {
		panic(err.Error())
	}
	for _, pod := range pods {

		fmt.Printf("namespace: %v,name: %v \n ", pod.GetNamespace(), pod.GetName())
	}

}
