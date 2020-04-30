/*
item fragments
In macros, you can use item fragments to interpolate 
items into the body of traits, impls, 
and extern blocks. For example:
*/
macro_rules! mac_trait {
    ($i:item) => {
        trait T { $i }
    }
}
mac_trait! {
    fn foo() {}
}

struct P;

impl T for P {
    fn foo() {
        println!("{}", "oh no");
    }
}

fn main() {
    println!("{}, {}", u32::MAX, f32::MAX);
    println!("{}, {}", u128::MAX, f64::MAX);
    P::foo();

    // this is not 1.43 feature~
    println!("{}", env!("PATH"));
}
