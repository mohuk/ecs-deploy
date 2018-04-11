package ecsdeploy

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/client"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ecs"
)

var (
	sess client.ConfigProvider
	svc  ecs.ECS
)

// Init initializes a new AWS session to be consumed by the Amazon ECS service
func Init(region string) {
	var err error
	sess, err = session.NewSession(&aws.Config{
		Region: aws.String(region),
	})

	if err != nil {
		panic(err)
	}

	svc = *ecs.New(sess)
}

// GetServiceDefinition returns the service definition
func GetServiceDefinition(cluster string, service string) (ecs.Service, error) {
	fmt.Printf("Fetching Service Definition: cluster %s, service %s\n", cluster, service)
	servicesInput := &ecs.DescribeServicesInput{
		Cluster: aws.String(cluster),
		Services: []*string{
			aws.String(service),
		},
	}
	servicesOutput, err := svc.DescribeServices(servicesInput)
	return *servicesOutput.Services[0], err
}

// GetCurrentTaskDefinition retuns the running task definition
func GetCurrentTaskDefinition(taskDefinition string) (ecs.TaskDefinition, error) {
	fmt.Printf("Fetching Task Definition: task definition %s\n", taskDefinition)
	taskDefinitionInput := &ecs.DescribeTaskDefinitionInput{
		TaskDefinition: aws.String(taskDefinition),
	}

	taskDefinitionOutput, err := svc.DescribeTaskDefinition(taskDefinitionInput)
	return *taskDefinitionOutput.TaskDefinition, err
}

// RegisterTaskDefinition is dope
func RegisterTaskDefinition(taskDefinition string, image string) (ecs.TaskDefinition, error) {
	fmt.Printf("Registering Task Definition: task definition %s, image %s\n", taskDefinition, image)
	taskDef, err := GetCurrentTaskDefinition(taskDefinition)
	if err != nil {
		fmt.Println(err)
	}
	taskDef.ContainerDefinitions[0].SetImage(image)
	registerTaskDefInput := &ecs.RegisterTaskDefinitionInput{
		ContainerDefinitions:    taskDef.ContainerDefinitions,
		Cpu:                     taskDef.Cpu,
		ExecutionRoleArn:        taskDef.ExecutionRoleArn,
		Family:                  taskDef.Family,
		Memory:                  taskDef.Memory,
		NetworkMode:             taskDef.NetworkMode,
		PlacementConstraints:    taskDef.PlacementConstraints,
		RequiresCompatibilities: taskDef.RequiresCompatibilities,
		TaskRoleArn:             taskDef.TaskRoleArn,
		Volumes:                 taskDef.Volumes,
	}
	registerTaskDefOutput, err := svc.RegisterTaskDefinition(registerTaskDefInput)
	return *registerTaskDefOutput.TaskDefinition, err
}

// UpdateService is dope
func UpdateService(cluster string, service string, taskDefinitionARN string) (ecs.Service, error) {
	fmt.Printf("Updating Service Definition: cluster %s, service %s, taskDefARN %s\n", cluster, service, taskDefinitionARN)
	serviceDefinition, err := GetServiceDefinition(cluster, service)
	serviceDefinition.SetTaskDefinition(taskDefinitionARN)
	updateServiceDefinitionInput := &ecs.UpdateServiceInput{
		Cluster:                       &cluster,
		DeploymentConfiguration:       serviceDefinition.DeploymentConfiguration,
		DesiredCount:                  serviceDefinition.DesiredCount,
		HealthCheckGracePeriodSeconds: serviceDefinition.HealthCheckGracePeriodSeconds,
		NetworkConfiguration:          serviceDefinition.NetworkConfiguration,
		PlatformVersion:               serviceDefinition.PlatformVersion,
		Service:                       &service,
		TaskDefinition:                serviceDefinition.TaskDefinition,
	}

	updateServiceDefinitionOutput, err := svc.UpdateService(updateServiceDefinitionInput)
	return *updateServiceDefinitionOutput.Service, err
}
