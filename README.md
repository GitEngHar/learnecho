# Let's try Echo

```shell
go run server.go

   ____    __
  / __/___/ /  ___
 / _// __/ _ \/ _ \
/___/\__/_//_/\___/ v4.13.4
High performance, minimalist Go web framework
https://echo.labstack.com
____________________________________O/_______
                                    O\
⇨ http server started on [::]:1323
```

## Run Example

After run server , Redirect this [page](http://localhost:1323/).

[ref](https://echo.labstack.com/docs/quick-start)

### context

#### bad practice

run server
```shell
go run ./callByValue/good/request_cancel.go
```

request
```shell
curl localhost:9010/
```