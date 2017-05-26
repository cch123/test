def lambda = {x: Int => x+1}
def lambda2 = (x:Int) => x+2
def lambda3 = (x:Int) => x+3

def lambda4 = new Function1[Int, Int] {
    def apply(v1: Int) : Int = v1 - 1
}

def lambda5(x: Int) = x + 1
