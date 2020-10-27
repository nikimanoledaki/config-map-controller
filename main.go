package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/klog"
)

const annotation = "x-k8s.io/curl-me-that"

func main() {
	loadRules := clientcmd.NewDefaultClientConfigLoadingRules()

	cfg, err := loadRules.Load()
	if err != nil {
		klog.Errorf("failed to build config from flags: %v", err)
		return
	}

	clientConfig, err := clientcmd.NewDefaultClientConfig(
		*cfg,
		&clientcmd.ConfigOverrides{},
	).ClientConfig()
	if err != nil {
		klog.Errorf("failed to build config from flags: %v", err)
		return
	}

	client := kubernetes.NewForConfigOrDie(clientConfig)

	cm, err := client.CoreV1().ConfigMaps("default").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		klog.Errorf("failed to get configmaps: %v", err)
	}
	fmt.Printf("configmaps %v\n", cm.Items)
	fmt.Printf("there are %d configmaps in this cluster\n", len(cm.Items))

	watcher, err := client.CoreV1().ConfigMaps("default").Watch(context.TODO(), metav1.ListOptions{})
	if err != nil {
		klog.Errorf("failed to get configmaps: %v", err)
	}

	for ev := range watcher.ResultChan() {
		cm, ok := ev.Object.(*corev1.ConfigMap)
		if !ok {
			klog.Errorf("not a corev1.ConfigMap, instead got %T\n", ev.Object)

		}
		switch ev.Type {
		case watch.Added, watch.Modified:
			fmt.Printf("event %s cm: %#v\n", ev.Type, cm)
			nameAndURL, ok := cm.Annotations[annotation]
			if !ok {
				continue
			}

			items := strings.SplitN(nameAndURL, "=", 2)
			if len(items) != 2 {
				fmt.Printf("annotation '%s' should have value of the form 'joke=curl-a-joke.herokuapp.com' but got '%v'\n", annotation, nameAndURL)
				continue
			}

			name := items[0]
			url := "https://" + items[1]

			if _, ok := cm.Data[name]; ok {
				fmt.Printf("the field '%v' already exists, not updating\n", name)
				continue
			}

			resp, err := http.Get(url)
			if err != nil {
				klog.Errorf("GET '%v' failed: %v", "http://"+url, err)
				continue
			}

			// x-k8s.io/curl-me-that: joke=curl-a-joke.herokuapp.com
			bytes, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				klog.Errorf("GET '%v' failed: %v", url, err)
				continue
			}

			cm.Data[name] = string(bytes)

			_, err = client.CoreV1().ConfigMaps("default").Update(context.TODO(), cm, metav1.UpdateOptions{})
			if err != nil {
				klog.Errorf("GET '%v' failed: %v", url, err)
				continue
			}

		default:
			continue
		}
	}

	watcher.Stop()
}
