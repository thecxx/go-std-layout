# Standard golang project layout

### 目录结构
```
/* 常规目录 */

/api
协议文件，Swagger/thrift/protobuf 等

/build
编译脚本、生成逻辑等等

/cmd
应用入口（比如 /cmd/demo/）目录，这个目录下面，每个文件在编译之后都会生成一个可执行的文件。
不要把很多的代码放到这个目录下面，这里面的代码尽可能简单。

/configs
配置文件

/docs
设计文档

/pkg
主要的业务代码、项目中内敛的依赖包

/pkg/runtime
运行时包，可以存放一些程序运行时可能需要处理的问题，比如：error、panic、内存监控等等一系列程序运行中的问题

/test
其他测试目录，功能测试，性能测试等

/vendor
项目依赖的其他第三方库，使用 govendor/dep 工具来管理依赖

/* 可选目录 */

/examples
应用程序或者公共库使用的一些例子

/deployments
部署相关
```