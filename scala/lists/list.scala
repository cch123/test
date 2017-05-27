// 注意，和 java 是反着的
//eq tests identity (same object):
val a = List(1, 2, 3)
val b = List(1, 2, 3)
(a eq b)

//== tests equality (same content):
val a = List(1, 2, 3)
val b = List(1, 2, 3)
(a == b)
