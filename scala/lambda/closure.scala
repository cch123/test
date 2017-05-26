// A closure is a function which maintains a reference to one or more variables outside of the function scope (it "closes over" the variables). Scala will detect that you are using variables outside of scope and create an object instance to hold the shared variables.
var incrementer = 1

def closure = { x: Int ⇒
  x + incrementer
}

val result1 = closure(10)

incrementer = 2

val result2 = closure(10)

