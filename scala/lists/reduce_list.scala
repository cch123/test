val a = List(1,2,3, 10)

println(a.reduceLeft(_ + _))
println(a.reduceLeft(_* _))
