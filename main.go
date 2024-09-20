package main

import (
	"log"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func main() {
	log.Printf("Starting")

	clusterConfig, _ := rest.InClusterConfig()
	clientset, _ := kubernetes.NewForConfig(clusterConfig)
	log.Printf(clientset.DiscoveryClient.LegacyPrefix)
	log.Printf("Done")
}
