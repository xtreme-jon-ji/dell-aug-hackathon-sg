## Pre-Requisites
### Install the PKS CLI
Download and install the Pivotal Container Service Command Line Interface (PKS CLI):

https://network.pivotal.io/products/pivotal-container-service

**Note**: You will be redirected to sign in page for registration before you can download the PKS CLI

Move the pks cli to the /usr/local/bin (rename to pks)
```
$ mv pks-linux-amd64-1.6.1-build.20 /usr/local/bin/pks
```

### Install the kubectl CLI
Download and install the Kubernetes Command Line Tool (kubectl):

https://kubernetes.io/docs/tasks/tools/install-kubectl/


## Accessing the Platform
First, you'll need to login to the PKS platform:
```
$ pks login -a <pks_api_host> -u <username> -p <password> -k
```

Next, run the following command in order to set the kubeconfig file to your home folder for the first time.
```
$ pks get-credentials <cluster-name>
```
The `pks_api_host`, `username`, `password`, `cluster-name` will only be provided upon request. You may request for the trial access via the slack channel **before** Hackathon Weekend . The access to the actual environment will only be given during the Hackathon Weekend (7th & 8th March), which can also request via the same slack channel

Once kubeconfig file is set, get-credentials is no longer needed. Please use get-kubeconfig command for authentication.
Example:

```

<clsuter-name> - tracy-pks-cluster-1
<pks_api_host> - https://api.pks.tracy.cf-app.com
<username> - team_1

$ pks get-kubeconfig tracy-pks-cluster-1 -u team_1 -a https://api.pks.tracy.cf-app.com -k              
Password: ******
Fetching kubeconfig for cluster tracy-pks-cluster-1 and user team_1.
You can now use the kubeconfig for user team_1:

$kubectl config use-context tracy-pks-cluster-1

$kubectl config set-context --current --namespace=team-1                                             
Context "tracy-pks-cluster-1" modified.

$kubectl get pods                                                                                         
No resources found in team-1 namespace.

$kubectl get pods -n default                                                                               
Error from server (Forbidden): pods is forbidden: User "user_team_1" cannot list resource "pods" in API group "" in the namespace "default"
```
As you may notice, permission segregation is done by **namespaces**. Please remember to switch to your own team namespace before running.


## Deploying the Sample Application
### Sample Docker Image
A sample docker image is available in Docker Hub with the repository path of **docker.io/djsoon/docker101**
There are TWO (2) ways to deploy the sample application as POD and SERVICE: **Imperative & Declarative**

### Imperative

1. Create a POD
```
kubectl run <pod-name> --image=<docker-image> --restart=Never --port=<port-num> 
```
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;For Eg.
```
$ kubectl run docker101-pod --image=docker.io/djsoon/docker101 --restart=Never --port=8080
```

2. Create a SERVICE
```
kubectl expose pod <pod-name> --name=<service-name> --port=<port-num> --type=NodePort --dry-run -o yaml > <output-yaml-file>.yaml
```
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;For Eg.
```
$ kubectl expose pod docker101-pod --name=docker101-svc --port=8080 --target-port=8080 --type=NodePort --dry-run -o yaml > docker101-svc.yaml
```
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Edit the YAML file to include the specific Node Port. **Note:** You will need to request for the specific node port range via slack channel

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;In the example below: the port number 31080 is assigned as to the **nodePort**
```
apiVersion: v1
kind: Service
metadata:
  labels:
    run: docker101-pod
  name: docker101-svc
spec:
  ports:
  - port: 8080
    protocol: TCP
    targetPort: 8080
    nodePort: 31080
  selector:
    run: docker101-pod
  type: NodePort
```
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Once the NodePort is added, deploy the service with the command below:
```
$ kubectl apply -f docker101-svc.yaml
```

### Declarative

Create a YAML file and specify the POD & SERVICE resources in the file:

```
apiVersion: v1
kind: Pod
metadata:
  labels:
    run: docker101-pod
  name: docker101-pod
spec:
  containers:
  - image: docker.io/djsoon/docker101
    name: docker101-pod
    ports:
    - containerPort: 8080
  restartPolicy: Never
---
apiVersion: v1
kind: Service
metadata:
  labels:
    run: docker101-pod
  name: docker101-svc
spec:
  ports:
  - port: 8080
    protocol: TCP
    targetPort: 8080
    nodePort: 31080
  selector:
    run: docker101-pod
  type: NodePort
```
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Deploy the YAML file with the command below:
```
$ kubectl apply -f <YAML file>
```

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;For Eg.

```
$ kubectl apply -f docker101-pod.yaml
```

### POD & SERVICE validation

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Check if the docker101 POD is in running state
```
$ kubectl get pod
NAME            READY   STATUS    RESTARTS   AGE
docker101-pod   1/1     Running   0          1m
```
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Check if the docker101 SERVICE is in running state and listening to port 31080
```
$ kubectl get service
NAME            TYPE        CLUSTER-IP       EXTERNAL-IP   PORT(S)          AGE
docker101-svc   NodePort    10.100.200.161   <none>        8080:31080/TCP   1m
kubernetes      ClusterIP   10.100.200.1     <none>        443/TCP          5d9h
```

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Access to the docker101 service (HTML Page) by keying in the URL on the browser address bar

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;For Eg.
```
http://<cluster-host>:31080
```

**Note:** Please also request for the **cluster-host** via the slack channel
  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;
You should see the HTML page as below:
![docker101](/docker101.png)



## Links

Topics to explore:

Docker Overview
https://docs.docker.com/engine/docker-overview

Kubernetes Concept
https://kubernetes.io/docs/concepts

Pivotal Container Service Concept & Features
https://pivotal.io/platform/pivotal-container-service


