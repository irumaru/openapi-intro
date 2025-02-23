# OASからコード生成

```
oapi-codegen -package generated ./api.yaml > ./generated/api.go
```

# 依存関係の解決

```
go mod tidy
```

# Curlコマンド

```
curl http://localhost:8080/users
```
```
curl -X POST -H "Content-Type: application/json" -d '{"id":2, "name":"Hanako Yamada"}' -i http://localhost:1323/users
```

# 参考

https://qiita.com/TomohiroSakai/items/1c259750352b49e96a18
