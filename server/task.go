package server

import (
	"context"
	"fmt"
	"github.com/RichardKnop/machinery/v2"
	redisbackend "github.com/RichardKnop/machinery/v2/backends/redis"
	"github.com/RichardKnop/machinery/v2/backends/result"
	redisbroker "github.com/RichardKnop/machinery/v2/brokers/redis"
	"github.com/RichardKnop/machinery/v2/config"
	redislock "github.com/RichardKnop/machinery/v2/locks/redis"
	"github.com/RichardKnop/machinery/v2/tasks"
	"github.com/centrifugal/centrifuge-go"
	"github.com/launchboxio/launchbox/api"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	v12 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"os"
	"path/filepath"
	"strconv"
)

type Task struct {
	Name string
}

type CreateNamespaceTaskOptions struct {
	Application *api.Application
}

func createTask(name string, args []tasks.Arg, refObject string, refId uint) (*result.AsyncResult, error) {
	task := &tasks.Signature{
		Name: name,
		Args: args,
	}
	res, err := taskServer.SendTask(task)
	if err != nil {
		return nil, err
	}
	record := &api.Task{
		TaskId:          res.Signature.UUID,
		ReferenceObject: refObject,
		ReferenceId:     refId,
		TaskName:        res.Signature.Name,
	}
	database.Create(record)
	return res, nil
}

func createNamespaceTask(applicationId uint) (*result.AsyncResult, error) {
	return createTask("namespace.create", []tasks.Arg{{
		Type:  "uint",
		Value: applicationId,
	}}, "Application", applicationId)
}

func deleteNamespaceTask(applicationId uint) (*result.AsyncResult, error) {
	return createTask("namespace.delete", []tasks.Arg{{
		Type:  "uint",
		Value: applicationId,
	}}, "Application", applicationId)
}

func createServiceTask(applicationId uint, projectId uint) (*result.AsyncResult, error) {
	return createTask("service.create", []tasks.Arg{{
		Type:  "uint",
		Value: applicationId,
	}, {
		Type:  "uint",
		Value: projectId,
	}}, "Project", projectId)
}

func Tasks() map[string]interface{} {
	apiClient, err := api.New()
	if err != nil {
		panic(err)
	}
	// TODO: This should also support in-cluster configuration
	kubeconfig := filepath.Join(
		os.Getenv("HOME"), ".kube", "config",
	)
	conf, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		panic(err)
	}
	clientset, err := kubernetes.NewForConfig(conf)

	centrifugeClient := centrifuge.NewJsonClient("ws://localhost:8000/connection/websocket", centrifuge.DefaultConfig())

	th := TaskHandler{
		kubeClient: clientset,
		apiClient:  apiClient,
		centrifuge: centrifugeClient,
	}
	return map[string]interface{}{
		"namespace.create": th.syncNamespace,
		"namespace.sync":   th.syncNamespace,
		"namespace.delete": th.deleteNamespace,
	}
}

type TaskHandler struct {
	kubeClient *kubernetes.Clientset
	apiClient  *api.Client
	centrifuge *centrifuge.Client
}

func (th *TaskHandler) syncNamespace(applicationId uint) error {

	app, err := th.apiClient.Apps().Find(applicationId, &api.ApplicationFindOptions{})
	if err != nil {
		return err
	}

	stringApplicationId := strconv.Itoa(int(app.ID))

	ns := &v1.Namespace{
		ObjectMeta: v12.ObjectMeta{
			Name: app.Namespace,
			Labels: map[string]string{
				"launchbox.io/application.id":     stringApplicationId,
				"openservicemesh.io/monitored-by": "osm",
				"openservicemesh.io/metrics":      "enabled",
			},
			Annotations: map[string]string{
				"openservicemesh.io/sidecar-injection": "enabled",
			},
		},
	}
	_, err = th.kubeClient.CoreV1().Namespaces().Get(context.TODO(), app.Namespace, v12.GetOptions{})
	if err != nil {
		if errors.IsNotFound(err) {
			fmt.Println("Creating namespace")
			_, err = th.kubeClient.CoreV1().Namespaces().Create(context.TODO(), ns, v12.CreateOptions{})
			if err != nil {
				return err
			}
			th.publishApplicationEvent(stringApplicationId, []byte{})
			return err
		}
	}
	fmt.Println("Updating namespace")
	_, err = th.kubeClient.CoreV1().Namespaces().Update(context.TODO(), ns, v12.UpdateOptions{})
	if err != nil {
		return err
	}

	return err
}

func (th *TaskHandler) deleteNamespace(applicationId uint) error {
	app, err := th.apiClient.Apps().Find(applicationId, &api.ApplicationFindOptions{Deleted: true})
	if err != nil {
		return err
	}
	_, err = th.kubeClient.CoreV1().Namespaces().Get(context.TODO(), app.Namespace, v12.GetOptions{})
	if err != nil {
		if errors.IsNotFound(err) {
			return nil
		}
		return err
	}
	return th.kubeClient.CoreV1().Namespaces().Delete(context.TODO(), app.Namespace, v12.DeleteOptions{})
}

func (th *TaskHandler) publishApplicationEvent(appChannel string, data []byte) {
	_, err := th.centrifuge.Publish(appChannel, data)
	if err != nil {
		fmt.Println(err)
	}
}

type TaskServerConfig struct {
	RedisUrl string
}

func NewTaskServer(opts *TaskServerConfig) (*machinery.Server, error) {

	cnf := &config.Config{
		DefaultQueue:    "machinery_tasks",
		ResultsExpireIn: 3600,
	}

	// TODO: Create actual instances of the following interfaces
	broker := redisbroker.New(cnf, opts.RedisUrl, "", "", 1)
	backend := redisbackend.New(cnf, opts.RedisUrl, "", "", 1)
	lock := redislock.New(cnf, []string{opts.RedisUrl}, 1, 1)
	server := machinery.NewServer(cnf, broker, backend, lock)

	return server, server.RegisterTasks(Tasks())
}
