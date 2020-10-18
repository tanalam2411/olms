

```bash
olms$ operator-sdk init --domain=olms.com --repo=github.com/tanalam2411/olms
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

```

```bash

```
