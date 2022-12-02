#[allow(dead_code)]
pub fn part1(content: &str) -> i32 {
    content
        .lines()
        .map(|item| match item {
            "B Z" => 1 + 6,
            "A Y" => 1 + 6,
            "C X" => 1 + 6,

            "A X" => 2 + 3,
            "B Y" => 2 + 3,
            "C Z" => 2 + 3,

            "A Z" => 3 + 0,
            "C Y" => 3 + 0,
            "B X" => 3 + 0,

            _ => 0,
        })
        .sum()
}

pub fn part2(content: &str) -> i32 {
    content
        .lines()
        .map(|item| match item {
            "C Z" => 1 + 6,
            "A Y" => 1 + 3,
            "B X" => 1 + 0,

            "A Z" => 2 + 6,
            "B Y" => 2 + 3,
            "C X" => 2 + 0,

            "B Z" => 3 + 6,
            "C Y" => 3 + 3,
            "A X" => 3 + 0,

            _ => 0,
        })
        .sum()
}
