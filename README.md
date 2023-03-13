# Carbon-Aware Scheduler
Aim of this work will be to propose and develop a scheduling policy for intelligent placement of serverless workloads in different regions depending on the availability of renewable resources at given point of time. To achieve this, we will exploit  kubernetes’ container orchestration capabilities and Knative’s capability of building serverless applications.

GreenCourier is a carbon-aware Kubernetes plugin to intelligently schedule
serverless functions in regions of low carbon emission.
GreenCourier optimises delivery of serverless functions across geo-spatial multi-cluster Kubernetes environment in the cloud for carbon efficiency.
GreenCourier has production-ready tech stack and one-click away from integrating
with existing geographically distributed clusters with Liqo.
## System Architecture

![System Architecture](/Images/system-diagram.png "GreenCourier Architecture")

## Installation
We need a cluster with Knative enabled in management cluster and target clusters which are
geographically distributed to be connected to management cluster using [Liqo](https://github.com/liqotech/liqo).

### Knative Installation
Knative can be installed using this [guide](https://knative.dev/docs/install/yaml-install/serving/install-serving-with-yaml/) or using the config
YAMLs given in this repository under config/knative directory.  It is important to note, a small
change in service-core config specification as shown below:

```
apiVersion: v1
kind: ConfigMap
metadata:
  name: config-features
  namespace: knative-serving
  labels:
    app.kubernetes.io/name: knative-serving
    app.kubernetes.io/component: controller
    app.kubernetes.io/version: "1.7.1"
  annotations:
    knative.dev/example-checksum: "4d5feafc"
data:
  kubernetes.podspec-schedulername: "enabled"  # setting scheduler-name in knative spec should be enabled
```

### Multi-Cluster Setup using Liqo
Once Knative is deployed in Management cluster, it is important to setup multi-cluster topology
using Liqo, depending on kubernetes distribution and providers, Liqo can be installed
using following [guide](https://docs.liqo.io/en/v0.5.2/installation/install.html) from official documentation.

Once Liqo is successfully installed on all Kubernetes clusters, we need to setup the
topology within those clusters.  To do that, every peer cluster should be run with following generate command:

```
liqoctl generate peer-command
```

Using the join-peer command is generated, which should be copied and should be executed in 
management cluster.

Once this is done for every peer cluster, we need to enable offloading in particular namespace,
so that every pod is offloaded to corresponding peer clusters, using following command.

```
liqoctl offload namespace {$NAMESPACE} --namespace-mapping-strategy EnforceSameName --pod-offloading-strategy Remote
```

It is also important to annotate the peer clusters with corresponding region in management cluster, so that it is picked up 
by GreenCourier during scheduling.  Annotating particular cluster is possible using following command
after peering of cluster is done.

```
kubectl annotate node {$NODENAME} node.kubernetes.io/region={$REGION}
```

### Install GreenCourier
Once the cluster setup up is done, it is important for us to install metrics-collector in local cluster.

```
# Replace $USERNAME and $PASSWORD with WattTime credentials in deployment YAML
kubectl apply -f config/metrics-collector/metric-collector-deployment.yaml
```
Once metrics-collector is deployed, we can deploy the plugin code using [Helm](https://helm.sh/) using following command:
```
helm install GreenCourier carbon-scheduler/charts/
```

### Deploying Functions
When metrics-collector and plugin code is deployed, we can just start deploying functions in Knative
with just addition of one line in the function spec YAML.
```
...
spec:
  schedulerName: kube-carbon-scheduler
...
```
Example function spec with schedulerName added.
```
apiVersion: serving.knative.dev/v1
kind: Service
metadata:
  name: hello
  namespace: default
spec:
  template:
    metadata:
      annotations:
        autoscaling.knative.dev/target: "10"
    spec:
      schedulerName: kube-carbon-scheduler
      containers:
        - image: gcr.io/knative-samples/helloworld-go
          ports:
            - containerPort: 8080
          env:
            - name: TARGET
              value: "GreenCourier Demo"
```
## Evaluation
Evaluation of GreenCourier was done against other scheduling schemes based on following
three metrics.
- Pod placement efficiency
- Carbon Emission
- Scheduling and Response time

### Pod Placement Efficiency

![Pod Placement](/Images/pod_placement.png "Pod Placement Metric")

### Carbon Emission
For functions whose execution time is less than 50ms:

![Carbon Emission](/Images/carboncompare_v2_1.png "Carbon Emission Metric - Function with Execution time < 50ms")

For functions whose execution time is greater than 50ms:

![Carbon Emission](/Images/carboncompare_v2_2.png "Carbon Emission Metric - Function with Execution time > 50ms")

### Scheduling and Binding Latency
Scheduling Latency:

![Scheduling Latency](/Images/gc-default.png "Scheduling Latency")

Binding Latency:

![Scheduling Latency](/Images/liqo-kubelet.png "Binding Latency")

### Response time
![Response Time](/Images/response-time.png "Response Time")


## Contact
Thandayuthapani Subramanian [:e-mail:](thandayuthapani.subramanian@tum.de) [![Linkedin](https://i.stack.imgur.com/gVE0j.png)](https://www.linkedin.com/in/thandayuthapani/) [![GitHub](https://i.stack.imgur.com/tskMh.png)](https://github.com/thandayuthapani)  
Mohak Chadha [:e-mail:](mohak.chadha@tum.de)[![Linkedin](https://i.stack.imgur.com/gVE0j.png)](https://www.linkedin.com/in/mohak-chadha-1490707b) [![GitHub](https://i.stack.imgur.com/tskMh.png)](https://github.com/kky-fury)