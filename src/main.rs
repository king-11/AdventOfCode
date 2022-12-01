use std::fs;

mod day1;

fn main() {
    let content =  fs::read_to_string("data/day1.txt").unwrap();
    println!("Answer {}", day1::part2(&content));
}
