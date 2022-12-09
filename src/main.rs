mod day1;
mod day2;
mod day3;
mod day4;
mod day5;

fn main() {
    let content =  include_str!("../data/day5.txt");
    let t0 = std::time::Instant::now();
    let result = day5::part2(&content);
    let t1 = std::time::Instant::now().duration_since(t0);
    println!("Answer {}", result);
    println!("Time used: ({:?})", t1);
}
