package main
import "fmt"

import registry "github.com/apache/dubbo-go/registry"
import zk "github.com/apache/dubbo-go/registry/zookeeper"
import xds "mosn.io/mosn/pkg/xds"

func main() {
 var r registry.Registry
 fmt.Println(r)
 var z zk.Option
 fmt.Println(z)
 var x xds.Client
 fmt.Println(x)
}
