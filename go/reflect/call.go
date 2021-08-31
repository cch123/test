package main
  
import "fmt"
import "reflect"
  
type T struct {}
  
func (t *T) Speak(x int) {
    fmt.Println("Alex", x)
}
  
func main() {
    var t T
      
    // use of Call() method
    val := reflect.ValueOf(&t).MethodByName("Speak").Call([]reflect.Value{reflect.ValueOf(1)})
      
    fmt.Println(val)
}      
