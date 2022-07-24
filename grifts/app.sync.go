package grifts

import (
	"flag"
	"fmt"
	"launchbox/models"
	"path/filepath"

	. "github.com/markbates/grift/grift"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

var _ = Namespace("app", func() {

	Desc("sync", "Task Description")
	Add("sync", func(c *Context) error {
		var kubeconfig *string
		if home := homedir.HomeDir(); home != "" {
			kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
		} else {
			kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
		}
		flag.Parse()

		// use the current context in kubeconfig
		config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
		if err != nil {
			panic(err.Error())
		}

		// create the clientset
		_, err = kubernetes.NewForConfig(config)
		if err != nil {
			panic(err.Error())
		}

		applications := []models.Application{}
		err = models.DB.All(&applications)
		if err != nil {
			panic(err.Error())
		}

		for _, app := range applications {
			fmt.Println(app.Namespace)
		}

		return nil
	})
})
