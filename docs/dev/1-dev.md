

```bash
olms$ operator-sdk init --project-version=2 --domain=olms.com --repo=github.com/tanalam2411/olms
Writing scaffold for you to edit...
Get controller runtime:
$ go get sigs.k8s.io/controller-runtime@v0.6.2
Update go.mod:
$ go mod tidy
Running make:
$ make
/home/afour/go/bin/controller-gen object:headerFile="hack/boilerplate.go.txt" paths="./..."
go fmt ./...
go vet ./...
go build -o bin/manager main.go
Next: define a resource with:
$ operator-sdk create api

olms$ operator-sdk create api --group=olmsg --version=v1alpha1 --kind=OLMS
Create Resource [y/n]
y
Create Controller [y/n]
y
Writing scaffold for you to edit...
api/v1alpha1/olms_types.go
controllers/olms_controller.go
Running make:
$ make
/home/afour/go/bin/controller-gen object:headerFile="hack/boilerplate.go.txt" paths="./..."
go fmt ./...
go vet ./...
go build -o bin/manager main.go
```

---

##### Verify all olm resources 

1. Namespace - `olm`, `operators`
```bash
$ k get ns | grep -E 'olm|operators'
olm                  Active   2m55s
operators            Active   2m55s
```

2. Service Account - `olm-operator-serviceaccount`
```bash
$ k get sa -n olm
NAME                          SECRETS   AGE
default                       1         3m51s
olm-operator-serviceaccount   1         3m51s
```

3. ClusterRole - `system:controller:operator-lifecycle-manager`, `aggregate-olm-edit`, `aggregate-olm-view`
```bash
$ k get clusterrole | grep -E 'aggregate-olm|lifecycle'
aggregate-olm-edit                                                     2020-11-14T19:55:19Z
aggregate-olm-view                                                     2020-11-14T19:55:19Z
system:controller:operator-lifecycle-manager                           2020-11-14T19:55:19Z
``` 

4. ClusterRoleBinding - `olm-operator-binding-olm`
```bash
$ k get clusterrolebindings | grep 'olm'
olm-operator-binding-olm                               ClusterRole/system:controller:operator-lifecycle-manager                           11m
```

5. Deployment - `olm-operator`, `catalog-operator`
```bash
$ k get deploy -n olm
NAME               READY   UP-TO-DATE   AVAILABLE   AGE
catalog-operator   1/1     1            1           12m
olm-operator       1/1     1            1           12m
packageserver      2/2     2            2           12m
```

6. OperatorGroup - `global-opertors`, `olm-operators`
```bash
$ k get og -n operators
NAME               AGE
global-operators   15m

$ k get og -n olm
NAME            AGE
olm-operators   15m
```

7. ClusterServiceVersion - `packageserver`
```bash
$ k get csv -n olm
NAME            DISPLAY          VERSION   REPLACES   PHASE
packageserver   Package Server   0.16.1               Succeeded
```

8. CatalogSource - `operatorhubio-catalog`
```bash
$ k get catsrc -n olm
NAME                    DISPLAY               TYPE   PUBLISHER        AGE
operatorhubio-catalog   Community Operators   grpc   OperatorHub.io   20m
```

9. Kafka 
```bash
$ k get packagemanifests | grep -i kafka
banzaicloud-kafka-operator                 Community Operators   24h
strimzi-kafka-operator                     Community Operators   24h
```