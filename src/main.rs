mod day1;
mod day2;
mod day3;
mod day4;
mod day5;
mod day6;
mod day7;
mod day8;
mod day9;
mod day10;
mod day12;

fn main() {
    let content =  std::fs::read_to_string("./data/day7.txt").unwrap();
    let t0 = std::time::Instant::now();
    let result = day7::part2(&content);
    let t1 = std::time::Instant::now().duration_since(t0);
    println!("Answer {}", result);
    println!("Time used: ({:?})", t1);
}
