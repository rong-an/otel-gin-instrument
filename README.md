## 后端 

```
# enter the project directory
cd gin-app

# install dependency
go mod tidy 

# build
go build 

# run 
SERVICE_NAME=GinApp INSECURE_MODE=true OTEL_EXPORTER_OTLP_ENDPOINT=localhost:4317 ./gin-app
```



## 前端 

```
# enter the project directory
cd ui

# install dependency
npm install

# develop
npm run dev

# build for production environment
npm run build:prod
```

