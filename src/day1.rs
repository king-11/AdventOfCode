use itertools::Itertools;

#[allow(dead_code)]
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

#[allow(dead_code)]
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

#[cfg(test)]
mod tests {
    use super::*;

    const CASE: &str = "1000
2000
3000

4000

5000
6000

7000
8000
9000

10000";

    #[test]
    fn part2_test() {
        let result = part2(CASE);
        assert_eq!(result, 45000);
    }

    #[test]
    fn part1_test() {
        let result = part1(CASE);
        assert_eq!(result, 24000);
    }
}
