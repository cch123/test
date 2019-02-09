#[derive(Debug)]
struct Ref<'a, T>(&'a T);

#[derive(Debug)]
struct Solution {
    i: i32,
}

#[derive(Debug)]
enum Expr<'a> {
    AddExpr(Box<Expr<'a>>, Box<Expr<'a>>),
    Lit(&'a str),
}

// input same with out put, can be ignored
fn get_expr(i: &str) -> Expr {
    Expr::AddExpr(Box::new(Expr::Lit(i)), Box::new(Expr::Lit("234")))
}

trait Animal{
    fn speak(&self);
}

#[derive(Debug)]
struct Human <'a>{
    category : &'a str,
    population: i32,
}

impl <'a>Animal for Human<'a> {
    fn speak(&self) {
        println!("category: {}, population: {}", self.category, self.population);
    }
}

fn gen_human<'a>(s: &'a str) -> Box<Animal+'a> {
    Box::new(Human{category:s, population : 12})
}

fn main() {
    let i = Solution { i: 123 };
    let x = Ref(&i); // or Ref::<Solution>(&i);
    println!("{:?}", x);

    let expr = get_expr("333");
    println!("{:?}", expr);

    let human = gen_human("eat me bite me");
    human.speak();
}
