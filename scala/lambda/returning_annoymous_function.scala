def addWithSyntaxSugar(x: Int) = (y: Int) ⇒ x + y

addWithSyntaxSugar(1).isInstanceOf[Function1[Int, Int]]

addWithSyntaxSugar(2)(3)

def fiveAdder = addWithSyntaxSugar(5)

fiveAdder(5)

addWithSyntaxSugar(1).isInstanceOf[Function1[_, _]]
