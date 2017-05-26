val a = List(1,2,3)

a(0)
a(2)
a(5)

// ??
intercept[IndexOutOfBoundsException] {
    println(a(5))
}
