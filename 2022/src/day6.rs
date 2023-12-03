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
    .as_bytes()
    .windows(size_pac)
    .position(|window| window.iter().all_unique()).unwrap() as i32 + size_pac as i32
}

#[cfg(test)]
mod tests {
  use super::*;
  use test_case::test_case;

  #[test_case(5, "bvwbjplbgvbhsrlpgdmjqwftvncz")]
  #[test_case(6, "nppdvjthqldpwncqszvftbrmjlhg")]
  #[test_case(10, "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg")]
  #[test_case(11, "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw")]
  fn part1_test(exp_res: i32, case: &str) {
    let result = part1(case);
    assert_eq!(result, exp_res);
  }

  #[test_case(23, "bvwbjplbgvbhsrlpgdmjqwftvncz")]
  #[test_case(23, "nppdvjthqldpwncqszvftbrmjlhg")]
  #[test_case(29, "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg")]
  #[test_case(26, "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw")]
  fn part2_test(exp_res: i32, case: &str) {
    let result = part2(case);
    assert_eq!(result, exp_res);
  }
}
