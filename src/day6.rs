use itertools::Itertools;

#[allow(dead_code)]
pub fn part1(content: &str) -> i32 {
  part(content, 4)
}

#[allow(dead_code)]
pub fn part2(content: &str) -> i32 {
  part(content, 14)
}

fn part(content: &str, size_pac: usize) -> i32 {
  content
    .chars()
    .fold((String::new(), 0), |mut acc, chr| {
      if acc.0.len() < size_pac {
        acc.0.push(chr);
      }

      if acc.0.chars().all_unique() {
        return acc;
      } else {
        acc.0.remove(0);
        return (acc.0, acc.1 + 1);
      }
    })
    .1 + size_pac as i32
}

#[cfg(test)]
mod tests {
  use super::*;

  const CASES: [&str; 4] = ["bvwbjplbgvbhsrlpgdmjqwftvncz", "nppdvjthqldpwncqszvftbrmjlhg", "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw"];

  const RESULT_1: [i32; 4] = [5, 6, 10, 11];
  const RESULT_2: [i32; 4] = [23, 23, 29, 26];

  #[test]
  fn part1_test() {
    let result = CASES.iter().map(|case| part1(case)).collect_vec();
    assert_eq!(result, RESULT_1);
  }

  #[test]
  fn part2_test() {
    let result = CASES.iter().map(|case| part2(case)).collect_vec();
    assert_eq!(result, RESULT_2);
  }
}
