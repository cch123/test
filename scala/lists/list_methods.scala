val a = List(1, 3, 5, 7, 9)

println(a.length)

println(a.reverse)
println(a)

a.map {
    v => v * 100
}

println(a)

a.filter {
    v => v % 3 == 0
}

println(a)

// 上面的操作都会生成新的列表，并不会影响到原始的对象
