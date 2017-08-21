# app-healthz

app-healthz is a sample app that demonstrates how to leverage the health endpoint pattern.

## Create Docker Image

```
go get github.com/kelseyhightower/app-healthz
```

Build the go binary

```
GOOS=linux bash build
```

```
docker build . -t kelseyhightower/app-healthz:1.0.0
```

```
docker push kelseyhightower/app-healthz:1.0.0
```
