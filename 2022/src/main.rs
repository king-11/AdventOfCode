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
mod day11;
mod day12;
mod day13;
mod day14;
mod day15;
mod day16;
mod day17;

fn main() {
    let content =  std::fs::read_to_string("./data/day17.txt").unwrap();
    let t0 = std::time::Instant::now();
    let result = day17::part1(&content);
    let t1 = std::time::Instant::now().duration_since(t0);
    println!("Answer {}", result);
    println!("Time used: ({:?})", t1);
}
