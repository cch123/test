// A companion object can also see private values and variables of the corresponding classes' instantiated objects:
class Person(val name: String, private val superheroName: String) //The superhero name is private!

object Person {
  def showMeInnerSecret(x: Person) = x.superheroName
}

val clark = new Person("Clark Kent", "Superman")
val peter = new Person("Peter Parker", "Spiderman")

println(Person.showMeInnerSecret(clark))
