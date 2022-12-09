use std::fs;

mod day1;
mod day2;
mod day3;
mod day4;

fn main() {
    let content =  fs::read_to_string("data/day3.txt").unwrap();
    let t0 = std::time::Instant::now();
    let result = day4::part2(&content);
    let t1 = std::time::Instant::now().duration_since(t0);
    println!("Answer {}", result);
    println!("Time used: ({:?})", t1);
}
