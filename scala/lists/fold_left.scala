val a = List(1, 3, 5, 7)

// 和reduce一样，不过foldLeft有个start value
println(a.foldLeft(0)(_ + _))
println(a.foldLeft(0)(_ * _))
println(a.foldLeft(10)(_ + _))
println(a.foldLeft(10)(_ * _))

