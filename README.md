
#Instructions  
This plugin is on golang base code, if golang is not installed then you have to go through the installation instructions here https://go.dev/doc/install

Or do  
```bash
sudo apt update

sudo apt install -y golang

go version

#Information   


kubectl-top-plugin

A custom kubectl plugin to display resource usage for pods and nodes.



Clone the repository:
```bash
git clone https://github.com/ragatgen/kubectl-top-plugin.git


cd kubectl-top-plugin

Build the binary:

go build -o kubectl-top kubectl-top.go

Make the binary executable:

chmod +x kubectl-top

Move it to a directory in your PATH:

sudo mv kubectl-top /usr/local/bin/

Check the installation

command -v kubectl-top

ouput should be 

/usr/local/bin/kubectl-top

Usage

#Run the plugin using:  
```bash
kubectl-top

It will prompt you for a namespace and display resource usage for the pods in that namespace, followed by node resource usage.


The working plugin should do the following successfully

randra@Lenovo-WorkRG:/mnt/c/Users/ragatgen$ kubectl-top

Enter the namespace to check: gorilla

Pods resource usage in namespace: gorilla

NAME                          CPU(cores)   MEMORY(bytes) 

pod-lister-5cb9fbcddb-xr2dm   4m           1Mi

Nodes resource usage:
NAME                                CPU(cores)   CPU%   MEMORY(bytes)   MEMORY% 

aks-nodepool1-29206393-vmss000006   157m         8%     5265Mi          104%

aks-nodepool1-29206393-vmss000007   190m         10%    1776Mi          35%

