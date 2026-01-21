# Developed by @Viggo Van
[Email](mailto:wayne3van@gmail.com)


### 多仓库go语言微服务-apiLink
### github.com/wegge0857/PolyrepoGoMicros-ApiLink


### 生成全部.pb.go文件
```bash
go run gen.go
```


### 建立本地关联
### 1. 初始化工作区（如果还没初始化过），在userMicros的外层执行
```bash
go work init
```


### 2. 将 API 仓库和微服务仓库都加入工作区，在userMicros的外层执行
```bash
go work use ./apiLink
go work use ./userMicros
go work use ./etfMicros
go work use ./bbfMicros
```

### 使用wire命令，go work有个限制，不能直接 go generate ./...
```bash
go install github.com/google/wire/cmd/wire@latest
```
### 在userMicros/cmd/userMicros中执行生wire_gen.go
```bash
wire
```

# 端口备注
```yaml
微服务              http            grpc
bffMicros:          8600            9600
etfMicros:          8601            9602
userMicros:         8602            9602
```