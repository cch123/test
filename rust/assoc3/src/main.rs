struct Container(i32, i32);

trait Contains {
    type A;
    type B;

    fn contains(&self,_:&Self::A, _:&Self::B) -> bool;
    fn first(&self) -> i32;
    fn last(&self) -> i32;
}

impl Contains for Container {
    type A = i32;
    type B = i32;

    fn contains(&self, n1:&i32, n2:&i32) -> bool {
        &self.0 == &self.1
    }

    fn first(&self) -> i32 {
        self.0
    }

    fn last(&self) -> i32 {
        self.1
    }
}

fn difference<C: Contains>(c : &C) -> i32 
{
    c.last() - c.first()
}

fn main() {
    let c = Container(10,99);
    println!("{}", c.contains(&1,&2));
    println!("{}", c.first());
    println!("{}", c.last());
    println!("{}", difference(&c));
}
