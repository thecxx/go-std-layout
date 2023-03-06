# Standard golang project layout

> From work experience and accumulation of open source projects.

### 目录结构
```
.
├── api         /* 协议文件，Swagger/thrift/protobuf 等 */
├── build       /* 编译脚本、生成逻辑等等 */
├── cmd         /* 应用入口（比如 /cmd/demo/）目录，这个目录下面，每个文件在编译之后都会生成一个可执行的文件 */
│   └── demo
├── conf        /* 配置文件 */
├── deployment  /* 部署相关 */
├── docs        /* 设计文档 */
├── examples    /* 应用程序或者公共库使用的一些例子 */
├── pkg         /* 主要的业务代码、项目中内敛的依赖包 */
│   ├── common              /* 通用目录，一般存放静态配置、错误码之类的内容 */
│   ├── config              /* 配置文件结构 */
│   ├── console             /* 命令入口 */
│   │   └── demo
│   │       └── handler
│   ├── dao                 /* 数据访问层 */
│   │   ├── cache           /* 缓存包 */
│   │   ├── db              /* 数据库包 */
│   │   ├── embed           /* 内嵌资源包 */
│   │   ├── fs              /* 文件资源包 */
│   │   └── rpc             /* Remote Procedure Call */
│   ├── port                /* 在线服务入口 */
│   │   ├── grpc
│   │   │   └── handler
│   │   └── http
│   │       ├── handler
│   │       └── middleware
│   └── service             /* 服务层，console、port中的handler用来适配不同的协议，最终由service处理 */
├── proto       /* 数据原型 */
├── tests       /* 测试目录，功能测试，性能测试等 */
└── vendor      /* 项目依赖的其他第三方库 */
```

### Directory Structure
```
.
├── api         /* Interface definition, like swagger, thrift, protobuf */
├── build       /* Build tools, scripts */
├── cmd         /* Entries (like /cmd/demo) */
│   └── demo
├── conf        /* Configuration files */
├── deployment  /* Deployment */
├── docs        /* Documentation */
├── examples    /* Some examples */
├── pkg
│   ├── common              /* General logic such as static configuration and error codes */
│   ├── config              /* Configuration definition */
│   ├── console             /* Console commands */
│   │   └── demo
│   │       └── handler
│   ├── dao                 /* Data Access Layer */
│   │   ├── cache           /* Cache */
│   │   ├── db              /* Database */
│   │   ├── embed           /* Embedded resources */
│   │   ├── fs              /* Files */
│   │   └── rpc             /* Remote Procedure Call */
│   ├── port                /* Online services */
│   │   ├── grpc
│   │   │   └── handler
│   │   └── http
│   │       ├── handler
│   │       └── middleware
│   └── service             /* Service */
├── proto       /* Prototype definition */
├── tests       /* Tests */
└── vendor
```

### Architecture
![Architecture](https://github.com/thecxx/go-std-layout/blob/master/arch.jpg)

### Sequence Diagram
![Sequence Diagram](https://github.com/thecxx/go-std-layout/blob/master/flowchart.jpg)

### Useful Tools
- glay
```
go install github.com/thecxx/go-std-layout/tools/glay/cmd/glay@latest
```
