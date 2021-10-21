# Perfect

### 目录结构
```
.
├── api         /* 协议文件，Swagger/thrift/protobuf 等 */
├── build       /* 编译脚本、生成逻辑等等 */
├── cmd         /* 应用入口（比如 /cmd/demo/）目录，这个目录下面，每个文件在编译之后都会生成一个可执行的文件 */
│   └── demo
├── conf        /* 配置文件 */
├── deployment  /* 部署相关 */
├── docs        /* 设计文档 */
├── examples    /* 应用程序或者公共库使用的一些例子 */
├── pkg         /* 主要的业务代码、项目中内敛的依赖包 */
│   ├── asset               /* 程序逻辑以外的数据资源 */
│   │   ├── config          /* 配置文件结构 */
│   │   ├── embed           /* 内嵌资源 */
│   │   └── option          /* 执行选项 */
│   ├── cmd                 /* 命令入口 */
│   │   ├── handler         /* 命令行服务入口 */
│   ├── common              /* 通用目录，一般存放静态配置、错误码之类的内容 */
│   ├── dao                 /* 数据访问层(目录名或改为dal，Data Access Layer) */
│   │   ├── cache           /* 缓存逻辑 */
│   │   └── database        /* 数据库逻辑 */
│   │   └── rpc             /* Remote Procedure Call */
│   ├── portal
│   │   └── restful
│   │       ├── handler     /* 在线服务入口 */
│   │       └── middleware  /* 中间件包 */
│   └── service             /* 服务层，portal中的handler用来适配不同的协议，最终统一由service处理 */
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
│   └── demo
├── conf        /* Configuration files */
├── deployment  /* Deployment */
├── docs        /* Documentation */
├── examples    /* Some examples */
├── pkg
│   ├── asset               /* Static resources */
│   │   ├── config          /* Configuration definition */
│   │   ├── embed           /* Embedded resources */
│   │   └── option          /* Option parameters during execution */
│   ├── cmd                 /* Command entry */
│   │   ├── handler         /* Handler */
│   ├── common              /* General logic such as static configuration and error codes */
│   ├── dao                 /* Data Access Layer(or name it 'dal') */
│   │   ├── cache           /* Cache */
│   │   └── database        /* Database*/
│   │   └── rpc             /* Remote Procedure Call */
│   ├── portal
│   │   └── restful
│   │       ├── handler     /* Handler */
│   │       └── middleware  /* Middleware */
│   └── service             /* Service */
├── proto       /* Prototype definition */
├── tests       /* Tests */
└── vendor
```