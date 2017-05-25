// 单例证明
object Greeting {
    def english = "Hi"
    def chn = "艹"
}

val x = Greeting
val y = x
val z = Greeting
println(x eq y)
println(x eq z)
