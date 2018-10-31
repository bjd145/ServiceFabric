package main

import (
	"net/http"
	"fmt"
	"os"
	"crypto/tls"
	sf "github.com/jjcollinge/servicefabric"
)

const (
	kindStateful  = "Stateful"
	kindStateless = "Stateless"
)

type ServiceItemExtended struct {
	sf.ServiceItem
	Application sf.ApplicationItem
	Partitions  []PartitionItemExtended
	Labels      map[string]string
}

type PartitionItemExtended struct {
	sf.PartitionItem
	Replicas  []sf.ReplicaItem
	Instances []sf.InstanceItem
}

type sfClient interface {
	GetApplications() (*sf.ApplicationItemsPage, error)
	GetServices(appName string) (*sf.ServiceItemsPage, error)
	GetPartitions(appName, serviceName string) (*sf.PartitionItemsPage, error)
	GetReplicas(appName, serviceName, partitionName string) (*sf.ReplicaItemsPage, error)
	GetInstances(appName, serviceName, partitionName string) (*sf.InstanceItemsPage, error)
	GetServiceExtensionMap(service *sf.ServiceItem, app *sf.ApplicationItem, extensionKey string) (map[string]string, error)
	GetServiceLabels(service *sf.ServiceItem, app *sf.ApplicationItem, prefix string) (map[string]string, error)
	GetProperties(name string) (bool, map[string]string, error)
}

type replicaInstance interface {
	GetReplicaData() (string, *sf.ReplicaItemBase)
}

func getValidInstances(sfClient sfClient, app sf.ApplicationItem, service sf.ServiceItem, partition sf.PartitionItem) []sf.InstanceItem {
	var validInstances []sf.InstanceItem
	if instances, err := sfClient.GetInstances(app.ID, service.ID, partition.PartitionInformation.ID); err != nil {
		fmt.Println(err)
	} else {
		for _, instance := range instances.Items {
			fmt.Println("Working on (" + instance.ID + ") - " + instance.Address )
			validInstances = append(validInstances, instance)
		}
	}
	return validInstances
}

func main() {
	
	cert, err := tls.LoadX509KeyPair("servicefabric.crt", "servicefabric.key")
	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{cert}}

	tlsConfig.BuildNameToCertificate()
	transport := &http.Transport{TLSClientConfig: tlsConfig}
	netClient := &http.Client{Transport: transport}

	c, err := sf.NewClient(netClient, os.Args[1], "3.0", tlsConfig)

	if err != nil {
		fmt.Println(err)
	}
	
	apps, err := c.GetApplications()
	if err != nil {
		fmt.Println(err)
	}

	var results []ServiceItemExtended
	for _, app := range apps.Items {
		fmt.Println("Working on " + app.ID )

		services, err := c.GetServices(app.ID)
		if err != nil {
			fmt.Println(err)
		}

		for _, service := range services.Items {
			fmt.Println("Working on " + service.ID )
			item := ServiceItemExtended{
				ServiceItem: service,
				Application: app,
			}

			if partitions, err := c.GetPartitions(app.ID, service.ID); err != nil {
				fmt.Println(err)
			} else {
				for _, partition := range partitions.Items {
					partitionExt := PartitionItemExtended{PartitionItem: partition}
					partitionExt.Instances = getValidInstances(c, app, service, partition)
					item.Partitions = append(item.Partitions, partitionExt)
				}
			}
			results = append(results, item)	
		}
	}
}

