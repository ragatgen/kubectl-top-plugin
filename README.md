kubectl-top-plugin

A custom kubectl plugin to display resource usage for pods and nodes.

Installation

Clone the repository:

git clone https://github.com/ragatgen/kubectl-top-plugin.git


cd kubectl-top-plugin

Build the binary:

go build -o kubectl-top kubectl-top.go

Make the binary executable:

chmod +x kubectl-top

Move it to a directory in your PATH:

sudo mv kubectl-top /usr/local/bin/

Usage

Run the plugin using:

kubectl top

It will prompt you for a namespace and display resource usage for the pods in that namespace, followed by node resource usage.

Check the installation

command -v kubectl-top

ouput should be 

/usr/local/bin/kubectl-top

