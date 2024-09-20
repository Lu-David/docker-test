package main

import (
	"log"

	"k8s.io/client-go/kubernetes"
)

func main() {
	log.Printf("Starting")

	_, err := kubernetes.NewForConfig(nil)
	log.Fatal(err)
}
