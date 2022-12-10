mod day1;
mod day2;
mod day3;
mod day4;
mod day5;
mod day6;

fn main() {
    let content =  std::fs::read_to_string("./data/day6.txt").unwrap();
    let t0 = std::time::Instant::now();
    let result = day6::part2(&content);
    let t1 = std::time::Instant::now().duration_since(t0);
    println!("Answer {}", result);
    println!("Time used: ({:?})", t1);
}
