Name: {{.serviceName}}.rpc
ListenOn: 127.0.0.1:8080
Etcd:
  Hosts:
  - 127.0.0.1:2379
  Key: {{.serviceName}}.rpc

Mysql:
  DataSource: root:rootroot@tcp(127.0.0.1:3306)/{{.serviceName}}?charset=utf8mb4&parseTime=true&loc=Asia%2FShanghai

Log:
  ServiceName: {{.serviceName}}.rpc
  Mode: console
  Encoding: plain
  Level: info