// object 和 java 的静态方法类似，但实际上却是一个全局一致的单例

object Greeting {
    def english = "Hi"
    def chn = "艹"
}

println(Greeting.english)
println(Greeting.chn)
