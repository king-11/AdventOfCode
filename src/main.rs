use std::fs;

mod day1;

fn main() {
    let content =  fs::read_to_string("data/day1.txt").unwrap();
    let t0 = std::time::Instant::now();
    let result = day1::part2(&content);
    let t1 = std::time::Instant::now().duration_since(t0);
    println!("Answer {}", result);
    println!("Time used: ({:?})", t1);
}
