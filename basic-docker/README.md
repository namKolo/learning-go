### Demo Docker with Go App

#### prerequisite

```
  go get https://github.com/namKolo/learning-go
```

#### RUN 
```
docker run -it --rm --name go-basic-instance -p 8080:8080 \
   -v ${GOPATH}/github.com/namKolo/learning-go/basic-docker:/go/src/BasicGoApp -w /go/src/BasicGoApp go-basic
```