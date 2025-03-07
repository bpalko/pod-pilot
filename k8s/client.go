package k8s

import (
	"context"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func getClient() (*kubernetes.Clientset, error) {
	config, err := clientcmd.BuildConfigFromFlags("", clientcmd.RecommendedHomeFile)
	if err != nil {
		return nil, err
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}
	return clientset, nil
}

func CreatePod(pod *v1.Pod) (*v1.Pod, error) {
	clientset, err := getClient()
	if err != nil {
		return nil, err
	}
	return clientset.CoreV1().Pods(pod.Namespace).Create(context.Background(), pod, metav1.CreateOptions{})
}

func GetPod(name, namespace string) (*v1.Pod, error) {
	clientset, err := getClient()
	if err != nil {
		return nil, err
	}
	return clientset.CoreV1().Pods(namespace).Get(context.Background(), name, metav1.GetOptions{})
}

func DeletePod(name, namespace string) error {
	clientset, err := getClient()
	if err != nil {
		return err
	}
	return clientset.CoreV1().Pods(namespace).Delete(context.Background(), name, metav1.DeleteOptions{})
}
