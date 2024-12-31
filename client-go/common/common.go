package common

import (
	"fmt"
	corev1 "k8s.io/api/core/v1"
	"strings"
)

func ProcessPod(pod corev1.Pod) {
	if strings.Contains(pod.Name, "spring") {
		fmt.Printf("namespace:%v,name:%v\n", pod.Namespace, pod.Name)
	}
}
