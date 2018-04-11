package main

import (
	"flag"
	"fmt"

	ecsdeploy "github.com/mohuk/ecs-deploy"
)

var (
	cluster string
	image   string
	region  string
	service string
)

func main() {
	flag.StringVar(&cluster, "c", "", "ECS Cluster name")
	flag.StringVar(&image, "i", "", "Image name")
	flag.StringVar(&region, "r", "us-east-1", "Region")
	flag.StringVar(&service, "s", "", "ECS Service to deploy")

	flag.Parse()

	if cluster == "" || image == "" || service == "" {
		fmt.Println("Missing parameters")
	}

	ecsdeploy.Init(region)
	serviceOutput, err := ecsdeploy.GetServiceDefinition(cluster, service)
	if err != nil {
		fmt.Println(err)
	}

	regTaskDefOutput, err := ecsdeploy.RegisterTaskDefinition(*serviceOutput.TaskDefinition, image)
	if err != nil {
		fmt.Println(err)
	}

	updateServiceOutput, err := ecsdeploy.UpdateService(cluster, service, *regTaskDefOutput.TaskDefinitionArn)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Updated Cluster: %s,\n Service: %s,\n Image: %s\n", *updateServiceOutput.ClusterArn, *updateServiceOutput.ServiceName, image)
}
