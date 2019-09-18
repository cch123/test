struct Container(i32, i32);

trait Contains<A, B> {
    fn contains(&self,_:&A, _:&B) -> bool;
    fn first(&self) -> i32;
    fn last(&self) -> i32;
}

impl Contains<i32, i32> for Container {
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

fn difference<A, B, C>(c : &C) -> i32 
where
    C : Contains<A, B>
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
