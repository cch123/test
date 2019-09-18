trait P {
    fn println(&self, x: i32);
}

impl P for i32 {
    fn println(&self, x: i32) {
        println!("{}", x);
    }
}

fn getP() -> impl P {
    return 1;
}

fn main() {
    let mut x = 1;
    let mut y = getP();
    y.println(10);
}
