val d = Nil

val c = 3 :: d
val b = 2 :: c
val a = 1 ::b
println(a)
println(b)
println(c)
println(d)

println(a.tail)
println(b.tail)
println(c.tail)
//println(d.tail) exception
