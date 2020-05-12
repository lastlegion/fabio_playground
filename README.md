# Fabio Playground

A simple setup to play around with http://fabiolb.net/


1. Starting dummy services

```
go run books/main.go
go run grocery/main.go
```

You should have books service on `8080` and grocery service on `8081`


2. Register services to consul 
```
consul agent -dev -config-dir=<root_dir>/fabio_example/consul
```

3. Run Fabio
```
fabio
```

4. Testing the services
```
~/dev/fabio_example/consul on  master ⌚ 15:42:08
$ curl localhost:9999/book
hello from books service

~/dev/fabio_example/consul on  master ⌚ 15:42:13
$ curl localhost:9999/grocery
hello from grocery service
```

You can also see routes at localhost:9998
