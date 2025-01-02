package main

import (
	"fmt"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/discovery"
	"k8s.io/client-go/tools/clientcmd"
)

/*
前3 种客户端，都是针对与资源对象管理的
DiscoveryClient 则是针对于资源的。用于查看当前 Kubernetes 集群支持那些资源组、资源版本、资源信息。
*/
func main() {
	// 1.加载配置文件
	config, err := clientcmd.BuildConfigFromFlags("", "/Users/ykdsg/.kube/config")
	if err != nil {
		panic(err.Error())
	}

	discoveryClient, err := discovery.NewDiscoveryClientForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	_, apiResourceLists, err := discoveryClient.ServerGroupsAndResources()
	if err != nil {
		panic(err.Error())
	}

	for _, apiResourceList := range apiResourceLists {
		gv, err := schema.ParseGroupVersion(apiResourceList.GroupVersion)
		if err != nil {
			panic(err.Error())
		}
		for _, apiResource := range apiResourceList.APIResources {
			fmt.Printf("name: %v, group: %v, version: %v\n", apiResource.Name, gv.Group, gv.Version)
		}
	}
}
