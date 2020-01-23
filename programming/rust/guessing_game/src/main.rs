use std::io;
use std::cmp::Ordering;
use rand::Rng; // The Rng trait defines methods that random number generators implement, and this trait must be in scope for us to use those methods.

fn main() {
    println!("Guess the number!");

    let secret_number = rand::thread_rng().gen_range(1, 101);

    loop {
        println!("Please input your guess.");

        let mut guess = String::new(); // mutable
        io::stdin().read_line(&mut guess) // &: reference, immutable by default
            .expect("Failed to read line");

        let guess: u32 = match guess.trim().parse() {
            Ok(num) => num,
            Err(_) => continue, // The underscore, _, is a catchall value; in this example, we’re saying we want to match all Err values, no matter what information they have inside them.
        };
        
        println!("You guessed: {}", guess);

        // A match expression is made up of arms. An arm consists of a pattern and the code that should be run if the value given to the beginning of the match expression fits that arm’s pattern. 
        match guess.cmp(&secret_number) {
            Ordering::Less => println!("Too small!"),
            Ordering::Greater => println!("Too big!"),
            Ordering::Equal => {
                println!("You win!");
                break;
            }
        }
    }
}
