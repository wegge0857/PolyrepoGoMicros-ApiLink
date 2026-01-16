# 多仓库go语言微服务-apiLink
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