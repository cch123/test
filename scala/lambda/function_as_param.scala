//Hint: a map method applies the function to each element of a list.

def makeUpper(xs: List[String]) = xs map {
    _.toUpperCase
}

def makeWhatEverYouLike(xs: List[String], sideEffect: String ⇒ String) =
xs map sideEffect

makeUpper(List("abc", "xyz", "123"))

makeWhatEverYouLike(List("ABC", "XYZ", "123"), { x ⇒
    x.toLowerCase
})

//using it inline
List("Scala", "Erlang", "Clojure") map {
    _.length
}
