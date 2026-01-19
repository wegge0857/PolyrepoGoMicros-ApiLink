# 多仓库go语言微服务-apiLink
# github.com/wegge0857/PolyrepoGoMicros-ApiLink
# apiLink结构

apiLink/
├── go.mod                // module github.com/wegge0857/PolyrepoGoMicros-ApiLink
├── user/                 // 存放 userMicros 的协议
│   └── v1/
│       ├── user.proto
│       └── user.pb.go    // 编译生成的代码
└── etf/                  // 存放 etfMicros 的协议
    └── v1/
        ├── etf.proto
        └── etf.pb.go

# 生成全部.pb.go文件
go run gen.go


# 建立本地关联
# 1. 初始化工作区（如果还没初始化过），在userMicros的外层执行
go work init

# 2. 将 API 仓库和微服务仓库都加入工作区，在userMicros的外层执行
go work use ./apiLink
go work use ./userMicros
go work use ./etfMicros
go work use ./bbfMicros

# 使用wire命令，go work有个限制，不能直接 go generate ./...
go install github.com/google/wire/cmd/wire@latest
# 在userMicros/cmd/userMicros中执行生wire_gen.go
wire



# 端口备注          http          grpc
bffMicros:          80            90
etfMicros:          8000          9000
userMicros:         8001          9001