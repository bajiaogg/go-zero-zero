



## model

```shell
goctl model mysql datasource -url="root:rootroot@tcp(127.0.0.1:3306)/zero_user" -table="*"  -dir="./model" -c --style goZero

goctl api go -api desc/user.api -dir . -style goZero --home ../../../deploy/goctl/

goctl rpc protoc pb/user.proto --go_out=./ --go-grpc_out=./ --zrpc_out=./ --style goZero --home ../../../deploy/goctl/

```



## 端口：

<table>
    <tr>
        <td>服务</td>
        <td>服务类型</td>
        <td>端口</td>
    </tr>
    <tr>
        <td rowspan="2">user</td>
        <td>api</td>
        <td>5001</td>
    </tr>
    <tr>
        <td>rpc</td>
        <td>5501</td>
    </tr>
    <tr>
        <td rowspan="2">identity</td>
        <td>api</td>
        <td>5002</td>
    </tr>
    <tr>
        <td>rpc</td>
        <td>5502</td>
    </tr>
    <tgor>
        <td rowspan="2">notification</td>
        <td>api</td>
        <td>5002</td>
    </tgor>
    <tr>
        <td>rpc</td>
        <td>5502</td>
    </tr>
</table>
