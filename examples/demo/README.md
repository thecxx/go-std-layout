<!--
 Copyright 2022 Kami
 
 Licensed under the Apache License, Version 2.0 (the "License");
 you may not use this file except in compliance with the License.
 You may obtain a copy of the License at
 
     http://www.apache.org/licenses/LICENSE-2.0
 
 Unless required by applicable law or agreed to in writing, software
 distributed under the License is distributed on an "AS IS" BASIS,
 WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 See the License for the specific language governing permissions and
 limitations under the License.
-->

### Examples
- build
```
go build -o demo github.com/thecxx/go-std-layout/examples/demo/cmd/demo
```

- demo --help
```
go-std-layout-demo

Usage:
  demo [flags]
  demo [command]

Available Commands:
  help        Help about any command
  port        Startup an online server
  print       Print any message

Flags:
  -h, --help      help for demo
  -v, --version   version for demo

Use "demo [command] --help" for more information about a command.
```

- demo port --help
```
Startup an online server

Usage:
  demo port [flags]

Flags:
  -h, --help       help for port
  -p, --port int   a port for server to bind (default 8090)
```
