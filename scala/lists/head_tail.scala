// head headOption tail
val a = List(1,2,3)
// 注意这两个的区别，scala里的Option是支持为空的head的，但不带option的可能会抛异常
a.head
a.headOption
// tail没有tailOption之类的东西
a.tail
