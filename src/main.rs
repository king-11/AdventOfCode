use std::fs;

mod day2;

fn main() {
    let content =  fs::read_to_string("data/day2.txt").unwrap();
    let t0 = std::time::Instant::now();
    let result = day2::part2(&content);
    let t1 = std::time::Instant::now().duration_since(t0);
    println!("Answer {}", result);
    println!("Time used: ({:?})", t1);
}
