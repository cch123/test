//An object that has the same name as a class is called a companion object of the class, and it is often used to contain factory methods for the class that it complements:

class Movie(val name: String, val year: Short)

object Movie {
    def academyAwardBestMoviesForYear(x: Short) = {
        x match {
            case 1930 => Some(new Movie("oh fuc", 1930))
            case 1931 => Some(new Movie("oh no", 1931))
            case 1932 => Some(new Movie("ooooo", 1932))
            case _ => None
        }
    }
}

print(Movie.academyAwardBestMoviesForYear(1932).get.name)
