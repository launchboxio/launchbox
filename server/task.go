package server

import (
	"context"
	"github.com/RichardKnop/machinery/v2"
	redisbackend "github.com/RichardKnop/machinery/v2/backends/redis"
	"github.com/RichardKnop/machinery/v2/backends/result"
	redisbroker "github.com/RichardKnop/machinery/v2/brokers/redis"
	"github.com/RichardKnop/machinery/v2/config"
	redislock "github.com/RichardKnop/machinery/v2/locks/redis"
	"github.com/RichardKnop/machinery/v2/tasks"
	"github.com/robwittman/launchbox/api"
	v1 "k8s.io/api/core/v1"
	v12 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"os"
	"path/filepath"
	"strconv"
)

var nullArgs = []tasks.Arg{}

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
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		panic(err)
	}
	clientset, err := kubernetes.NewForConfig(config)

	return map[string]interface{}{
		"namespace.create": func(applicationId uint) error {
			app, err := apiClient.Apps().Find(applicationId)
			if err != nil {
				return err
			}
			_, err = clientset.CoreV1().Namespaces().Create(context.TODO(), &v1.Namespace{
				ObjectMeta: v12.ObjectMeta{
					Name: app.Namespace,
					Labels: map[string]string{
						"launchbox.io/application.id": strconv.Itoa(int(app.ID)),
					},
				},
			}, v12.CreateOptions{})
			return err
		},
		"service.create": func(applicationId uint, projectId uint) error {
			// Create a service account
			app, err := apiClient.Apps().Find(applicationId)
			if err != nil {
				return err
			}
			project, err := apiClient.Projects().Find(projectId)
			if err != nil {
				return err
			}
			_, err = clientset.CoreV1().ServiceAccounts(app.Namespace).Create(context.TODO(), &v1.ServiceAccount{
				ObjectMeta: v12.ObjectMeta{
					Name: project.GetFriendlyName(),
					Labels: map[string]string{
						"launchbox.io/application.id": strconv.Itoa(int(app.ID)),
						"launchbox.io/project.id":     strconv.Itoa(int(project.ID)),
					},
				},
			}, v12.CreateOptions{})
			return err
		},
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