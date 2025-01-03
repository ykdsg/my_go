package main

import (
	"clientgo/common"
	"context"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

/*
调用client-go
*/
func main() {

	// 1.加载配置文件
	config, err := clientcmd.BuildConfigFromFlags("", "/Users/ykdsg/.kube/config")
	if err != nil {
		panic(err.Error())
	}

	// 2.配置API 路径
	config.APIPath = "api"

	// 3.配置分组版本
	config.GroupVersion = &corev1.SchemeGroupVersion // 无名资源组 group:"" ,version:"v1"

	// 4. 配置数据编解码工具
	config.NegotiatedSerializer = scheme.Codecs

	// 5.实例化RESTClient 对象
	restClient, err := rest.RESTClientFor(config)
	if err != nil {
		panic(err.Error())
	}

	// 6. 定义接收返回值
	result := &corev1.PodList{}

	// 7. APIServer 交互
	err = restClient.Get().
		Namespace("default").                                          // 命名空间
		Resource("pods").                                              // 资源
		VersionedParams(&metav1.ListOptions{}, scheme.ParameterCodec). // 参数及参数序列化工具
		Do(context.TODO()).                                            // 触发请求
		Into(result)                                                   // 写入返回结果

	if err != nil {
		panic(err.Error())
	}

	for _, pod := range result.Items {
		common.ProcessPod(pod)
	}

}
