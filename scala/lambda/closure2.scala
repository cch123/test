//We can take that closure and throw it into a method and it will still hold the environment:
def summation(x: Int, y: Int ⇒ Int) = y(x)

var incrementer = 3
def closure = (x: Int) ⇒ x + incrementer

val result = summation(10, closure)

incrementer = 4
val result2 = summation(10, closure)

