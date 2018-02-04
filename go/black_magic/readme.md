export private function
------

1. define go:linkname function in your file
```go
//go:linkname throw runtime.throw
func throw(s string)
```

2. create empty .s file in your project folder, file name is arbitrary
```
touch a.s
```

3. go build
```
go build
```
