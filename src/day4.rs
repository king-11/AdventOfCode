use itertools::Itertools;

#[allow(dead_code)]
pub fn part1(content: &str) -> i32 {
  content
  .lines()
  .filter_map(|s| {
    s.split(|c: char| !c.is_numeric())
    .map(|n| n.parse::<i32>().unwrap())
    .collect_tuple()
  })
  .filter(|(a1, a2, b1, b2)| a1 <= b1 && b2 <= a2 || b1 <= a1 && a2 <= b2)
  .count() as i32
}

#[allow(dead_code)]
pub fn part2(content: &str) -> i32 {
  content
  .lines()
        .filter_map(|s| {
            s.split(|c: char| !c.is_numeric())
                .map(|n| n.parse::<i32>().unwrap())
                .next_tuple()
        })
        .filter(|(a1, a2, b1, b2)| a2 >= b1 && b2 >= a1)
        .count() as i32
}

#[cfg(test)]
mod tests {
    use super::*;

    const CASE: &str = "2-4,6-8
2-3,4-5
5-7,7-9
2-8,3-7
6-6,4-6
2-6,4-8";

    #[test]
    fn part1_test() {
        let result = part1(CASE);
        assert_eq!(result, 2);
    }

    #[test]
    fn part2_test() {
        let result = part2(CASE);
        assert_eq!(result, 4);
    }
}
