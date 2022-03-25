# scheduler

Toy scheduler for use in Kubernetes demos.

## Usage

Annotate each node using the annotator command:

```
kubectl proxy
```
output:
```
Starting to serve on 127.0.0.1:8001
```

### Annotate all the nodes with randomly assigned colors
```
go run annotator/main.go
```
output:
```
aks-agentpool-25166730-vmss000000 #000000
aks-agentpool-25166730-vmss000001 #ffffff
aks-agentpool-25166730-vmss000002 #000000
aks-agentpool-25166730-vmss000003 #ffffff
```

### Create a deployment

```
kubectl create namespace colors
kubectl apply -f deployments/nginx.yaml --namespace=colors
```
```
deployment.apps/nginx created
```

The nginx pod(s) should be in a pending state:

```
kubectl get pods -n colors
```
output:
```
NAME                     READY     STATUS    RESTARTS   AGE
nginx-1431970305-mwghf   0/1       Pending   0          27s
```

## Build the Scheduler
```
sh ./build
```

## Run the Scheduler

As a reminder, list the nodes and note the randomly assigned color of each node.

```
go run annotator/main.go -l
```
output:
```
aks-agentpool-25166730-vmss000000 #000000
aks-agentpool-25166730-vmss000001 #ffffff
aks-agentpool-25166730-vmss000002 #000000
aks-agentpool-25166730-vmss000003 #ffffff
```

Now, randomly assign colors to the pods that are pending:
```
go run annotator/main.go -p
```
output:
```
nginx-d957c978b-5sn22 #feffef
nginx-d957c978b-qttfk #feffef
nginx-d957c978b-r6488 #604010
```


Now, run the color sort scheduler:

```
./scheduler
```

output
```
2022/03/25 04:30:50 Starting custom scheduler...
Processing pod with string color: #927a0c
Match found
2022/03/25 04:30:52 Successfully assigned nginx-d957c978b-d2bsn to aks-agentpool-25166730-vmss000000
Processing pod with string color: #00ff00
Match found
2022/03/25 04:30:54 Successfully assigned nginx-d957c978b-nrddx to aks-agentpool-25166730-vmss000001
Processing pod with string color: #aaaaaa
Match found
2022/03/25 04:30:56 Successfully assigned nginx-d957c978b-sqdsq to aks-agentpool-25166730-vmss000001
```
Hit control C.

See the scheduled pods:
```
kubectl get pods -owide --sort-by="{.spec.nodeName}" -n colors
```


Now, tear it down for the next run!
```
kubectl delete -f deployments/nginx.yaml --namespace colors
```


# TODO: not tested below here:

## Run the Scheduler on Kubernetes

```
kubectl create -f deployments/scheduler.yaml
```
```
deployment "scheduler" created
```
