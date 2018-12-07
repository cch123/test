use rand::Rng;
use std::cmp::Ordering;
use std::io;

fn main() {
    let sec_num = rand::thread_rng().gen_range(1, 100);
    loop {
        println!("Guess the number");
        println!("Please insert your guess");
        let mut guess = String::new();
        io::stdin()
            .read_line(&mut guess)
            .expect("Failed to readline");

        println!("You guessed: {}", guess);
        let guess: u32 = guess.trim().parse().expect("Please type a number");
        match guess.cmp(&sec_num) {
            Ordering::Less => println!("Too small"),
            Ordering::Greater => println!("Too great"),
            Ordering::Equal => println!("equal"),
        }
    }
}
