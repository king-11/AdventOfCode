use itertools::Itertools;

pub fn part1(content: &str) -> i32 {
    content
        .split("\n\n")
        .map(|items| {
            items
                .split("\n")
                .fold(0i32, |sum, item| sum + item.parse::<i32>().unwrap())
        })
        .max()
        .unwrap()
}

pub fn part2(content: &str) -> i32 {
    let result_vec = content
        .split("\n\n")
        .map(|items| {
            items
                .split("\n")
                .fold(0, |sum, item| sum + item.parse::<i32>().unwrap())
        })
        .sorted_by_key(|a| -a)
        .collect::<Vec<i32>>();

    result_vec[..3].iter().sum()
}
