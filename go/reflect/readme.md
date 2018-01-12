# reflect

## 基本类型

bool, int, uint, uintptr, float, complex

```go
var a = 1

// change a
reflect.ValueOf(a).SetInt(1) // panic, ValueOf(a) is not addressable
reflect.ValueOf(&a).Elem().SetInt(2) // correct
```

## map type

```go
var a map[string]int

fmt.Println(reflect.TypeOf(a)) // reflect.Type : map[string]int
fmt.Println(reflect.TypeOf(a).Kind()) // reflect.Kind : map
fmt.Println(reflect.TypeOf(a).Elem()) // reflect.Type : int
fmt.Println(reflect.TypeOf(a).Key())  // reflect.Type : string

// 使用反射初始化 map
reflect.ValueOf(a).Set(reflect.MakeMap(reflect.TypeOf(a))) // error ! ValueOf(a) is not addressable
reflect.ValueOf(&a).Elem().Set(reflect.MakeMap(reflect.TypeOf(a))) // correct

// set key value
reflect.ValueOf(a).SetMapIndex(reflect.ValueOf("abc"), reflect.ValueOf(1))

```

reflect.Type 是个 interface{} 类型，如果要按照类型做 switch，一定要按照 reflect.Kind 来做，因为 Kind 是枚举值。

当然，也可以使用 type switch，具体参考 elasticsql。

## array type

## ptr type
