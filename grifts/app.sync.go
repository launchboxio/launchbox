package grifts

import (
	"context"
	"flag"
	"fmt"
	"github.com/launchboxio/operator/api/v1alpha1"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
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
		clientset, err := kubernetes.NewForConfig(config)
		if err != nil {
			panic(err.Error())
		}

		applications := []models.Application{}
		err = models.DB.All(&applications)
		if err != nil {
			panic(err.Error())
		}

		v1alpha1.AddToScheme(scheme.Scheme)
		crdConfig := *config
		crdConfig.ContentConfig.GroupVersion = &v1alpha1.GroupVersion
		crdConfig.APIPath = "/apis"
		crdConfig.NegotiatedSerializer = serializer.NewCodecFactory(scheme.Scheme)
		crdConfig.UserAgent = rest.DefaultKubernetesUserAgent()
		exampleRestClient, err := rest.UnversionedRESTClientFor(&crdConfig)
		if err != nil {
			panic(err.Error())
		}

		// Sync applications
		for _, app := range applications {
			_, err = clientset.CoreV1().Namespaces().Get(context.TODO(), app.Namespace, metav1.GetOptions{})
			if err != nil && errors.IsNotFound(err) {
				// Namespace not found, go ahead and create
				_, err := clientset.CoreV1().Namespaces().Create(context.TODO(), &v1.Namespace{
					ObjectMeta: metav1.ObjectMeta{
						Name: app.Namespace,
						Labels: map[string]string{
							"openservicemesh.io/monitored-by": "osm",
						},
						Annotations: map[string]string{
							"openservicemesh.io/sidecar-injection": "enabled",
						},
					},
				}, metav1.CreateOptions{})
				if err != nil {
					return err
				}
			} else if err != nil {
				return err
			}

			// Create the "application" resource in the namespace
			projects := []models.Project{}
			err = models.DB.All(&projects)
			if err != nil {
				panic(err.Error())
			}

			for _, proj := range projects {
				project := &v1alpha1.Project{
					ObjectMeta: metav1.ObjectMeta{
						Name:      proj.Slug,
						Namespace: app.Namespace,
					},
					TypeMeta: metav1.TypeMeta{},
					Spec: v1alpha1.ProjectSpec{
						Secrets:   nil,
						Resources: v1.ResourceRequirements{},
					},
				}

				found := &v1alpha1.Project{}
				err := exampleRestClient.
					Get().
					Resource("projects").
					Name(proj.Slug).
					Namespace(app.Namespace).Do(context.TODO()).Into(found)
				fmt.Println(err)
				if err != nil && errors.IsNotFound(err) {
					fmt.Println("Creating new project")
					err := exampleRestClient.
						Post().
						Namespace(app.Namespace).
						Resource("projects").
						Body(project).
						Do(context.TODO()).
						Into(project)
					fmt.Println(err)
					fmt.Println(project)
				} else if err != nil {
					fmt.Println("Failed finding project")
					panic(err.Error())
				} else {
					fmt.Println("Updating project")

					err = exampleRestClient.
						Put().
						Namespace(app.Namespace).
						Resource("projects").
						Name(proj.Slug).
						Body(project).
						Do(context.TODO()).
						Into(project)
				}

				revisions := []models.Revision{}
				err = models.DB.All(&revisions)
				if err != nil {
					panic(err.Error())
				}

				for _, rev := range revisions {
					revision := &v1alpha1.Revision{
						ObjectMeta: metav1.ObjectMeta{
							Name:      rev.ID.String(),
							Namespace: app.Namespace,
						},
						TypeMeta: metav1.TypeMeta{},
						Spec: v1alpha1.RevisionSpec{
							Project:   proj.ID.String(),
							Resources: v1.ResourceRequirements{},
						},
					}

					found := &v1alpha1.Revision{}
					err = exampleRestClient.
						Get().
						Resource("revisions").
						Name(rev.ID.String()).
						Namespace(app.Namespace).Do(context.TODO()).Into(found)
					if err != nil && errors.IsNotFound(err) {
						fmt.Println("Creating new revision")
						err := exampleRestClient.
							Post().
							Namespace(app.Namespace).
							Resource("revisions").
							Body(revision).
							Do(context.TODO()).
							Into(project)
						fmt.Println(err)
						fmt.Println(revision)
					}
				}
			}
		}

		return nil
	})
})
