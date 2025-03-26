package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the namespace to check: ")
	namespace, _ := reader.ReadString('\n')
	namespace = strings.TrimSpace(namespace)

	// Check if the namespace exists
	cmd := exec.Command("kubectl", "get", "ns", namespace)
	if err := cmd.Run(); err != nil {
		fmt.Printf("Namespace '%s' does not exist. Exiting.\n", namespace)
		os.Exit(1)
	}

	// Show resource usage for pods in the specified namespace
	fmt.Printf("\nPods resource usage in namespace: %s\n", namespace)
	topPods := exec.Command("kubectl", "top", "pods", "-n", namespace)
	topPods.Stdout = os.Stdout
	topPods.Stderr = os.Stderr
	topPods.Run()

	// Show resource usage for nodes
	fmt.Println("\nNodes resource usage:")
	topNodes := exec.Command("kubectl", "top", "nodes")
	topNodes.Stdout = os.Stdout
	topNodes.Stderr = os.Stderr
	topNodes.Run()
}
