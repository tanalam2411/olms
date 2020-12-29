##### Installing Strimzi kafka operator

1. Install OLM

```bash
$ kubectl apply -f https://github.com/operator-framework/operator-lifecycle-manager/releases/download/v0.17.0/crds.yaml

customresourcedefinition.apiextensions.k8s.io/catalogsources.operators.coreos.com created
customresourcedefinition.apiextensions.k8s.io/clusterserviceversions.operators.coreos.com created
customresourcedefinition.apiextensions.k8s.io/installplans.operators.coreos.com created
customresourcedefinition.apiextensions.k8s.io/operatorgroups.operators.coreos.com created
customresourcedefinition.apiextensions.k8s.io/operators.operators.coreos.com created
customresourcedefinition.apiextensions.k8s.io/subscriptions.operators.coreos.com created
```
```bash
$ kubectl apply -f https://github.com/operator-framework/operator-lifecycle-manager/releases/download/v0.17.0/olm.yaml
namespace/olm created
namespace/operators created
serviceaccount/olm-operator-serviceaccount created
clusterrole.rbac.authorization.k8s.io/system:controller:operator-lifecycle-manager created
clusterrolebinding.rbac.authorization.k8s.io/olm-operator-binding-olm created
deployment.apps/olm-operator created
deployment.apps/catalog-operator created
clusterrole.rbac.authorization.k8s.io/aggregate-olm-edit created
clusterrole.rbac.authorization.k8s.io/aggregate-olm-view created
operatorgroup.operators.coreos.com/global-operators created
operatorgroup.operators.coreos.com/olm-operators created
clusterserviceversion.operators.coreos.com/packageserver created
catalogsource.operators.coreos.com/operatorhubio-catalog created

```


2. Installing strimzi operator

```bash
$ kubectl create -f https://operatorhub.io/install/strimzi-kafka-operator.yaml
subscription.operators.coreos.com/my-strimzi-kafka-operator created
```
```bash
$ kubectl get csv -n operators
NAME                               DISPLAY   VERSION   REPLACES                           PHASE
strimzi-cluster-operator.v0.20.1   Strimzi   0.20.1    strimzi-cluster-operator.v0.20.0   Pending
```
```bash
$ kubectl get all -n operators
NAME                                                   READY   STATUS              RESTARTS   AGE
pod/strimzi-cluster-operator-v0.20.1-5dc7cb547-d8nqr   0/1     ContainerCreating   0          61s

NAME                                               READY   UP-TO-DATE   AVAILABLE   AGE
deployment.apps/strimzi-cluster-operator-v0.20.1   0/1     1            0           61s

NAME                                                         DESIRED   CURRENT   READY   AGE
replicaset.apps/strimzi-cluster-operator-v0.20.1-5dc7cb547   1         1         0       61s
```
```bash
$ kubectl get csv -n operators
NAME                               DISPLAY   VERSION   REPLACES                           PHASE
strimzi-cluster-operator.v0.20.1   Strimzi   0.20.1    strimzi-cluster-operator.v0.20.0   Succeeded
```
```bash
$ kubectl get ns
NAME                 STATUS   AGE
default              Active   8h
kube-node-lease      Active   8h
kube-public          Active   8h
kube-system          Active   8h
local-path-storage   Active   8h
olm                  Active   2m1s
operators            Active   2m1s
```

3. Installing kafka - https://operatorhub.io/operator/strimzi-kafka-operator - Custom Resource Definitions

```yaml
apiVersion: kafka.strimzi.io/v1beta1
kind: Kafka
metadata:
  name: my-cluster
spec:
  kafka:
    version: 2.6.0
    replicas: 3
    listeners:
      - name: plain
        port: 9092
        type: internal
        tls: false
      - name: tls
        port: 9093
        type: internal
        tls: true
    config:
      offsets.topic.replication.factor: 3
      transaction.state.log.replication.factor: 3
      transaction.state.log.min.isr: 2
      log.message.format.version: '2.6'
      inter.broker.protocol.version: '2.6'
    storage:
      type: ephemeral
  zookeeper:
    replicas: 3
    storage:
      type: ephemeral
  entityOperator:
    topicOperator: {}
    userOperator: {}
```

```bash
$ kubectl create -f kafka.yaml 
kafka.kafka.strimzi.io/my-cluster created
```

```bash
$ kubectl get all
NAME                                              READY   STATUS    RESTARTS   AGE
pod/my-cluster-entity-operator-6b7f77b756-2m2wn   3/3     Running   0          71s
pod/my-cluster-kafka-0                            1/1     Running   0          100s
pod/my-cluster-kafka-1                            1/1     Running   0          100s
pod/my-cluster-kafka-2                            1/1     Running   0          100s
pod/my-cluster-zookeeper-0                        1/1     Running   2          4m7s
pod/my-cluster-zookeeper-1                        1/1     Running   3          4m7s
pod/my-cluster-zookeeper-2                        1/1     Running   0          4m7s

NAME                                  TYPE        CLUSTER-IP      EXTERNAL-IP   PORT(S)                      AGE
service/kubernetes                    ClusterIP   10.96.0.1       <none>        443/TCP                      8h
service/my-cluster-kafka-bootstrap    ClusterIP   10.103.134.60   <none>        9091/TCP,9092/TCP,9093/TCP   101s
service/my-cluster-kafka-brokers      ClusterIP   None            <none>        9091/TCP,9092/TCP,9093/TCP   101s
service/my-cluster-zookeeper-client   ClusterIP   10.96.65.154    <none>        2181/TCP                     4m9s
service/my-cluster-zookeeper-nodes    ClusterIP   None            <none>        2181/TCP,2888/TCP,3888/TCP   4m8s

NAME                                         READY   UP-TO-DATE   AVAILABLE   AGE
deployment.apps/my-cluster-entity-operator   1/1     1            1           71s

NAME                                                    DESIRED   CURRENT   READY   AGE
replicaset.apps/my-cluster-entity-operator-6b7f77b756   1         1         1       71s

NAME                                    READY   AGE
statefulset.apps/my-cluster-kafka       3/3     100s
statefulset.apps/my-cluster-zookeeper   3/3     4m8s

```