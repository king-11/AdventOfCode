use itertools::Itertools;
fn increment(val: i32, cycle: i32) -> i32 {
  if [20, 60, 100, 140, 180, 220].contains(&cycle) {
    return cycle*val;
  }

  0
}

#[allow(dead_code)]
pub fn part1(content: &str) -> i32 {
  let mut cycle = 0;
  let mut prod = 0;
  let mut val = 1;
  content.lines().for_each(|line| {
    if line.starts_with("noop") {
      cycle += 1;
      prod += increment(val, cycle);
    } else {
      cycle += 1;
      prod += increment(val, cycle);
      cycle += 1;
      prod += increment(val, cycle);
      val += line.split(" ").collect_vec()[1].parse::<i32>().unwrap();
    }
  });

  prod
}

#[allow(dead_code)]
pub fn part2(content: &str) -> i32 {
  let mut cycle = 0;
  let mut val = 1;

  let print_or_not = |c: i32, v: i32| {
    if (c % 40 - v).abs() <= 1 {
      print!("#");
    } else {
      print!(".")
    }

    if (c + 1) % 40 == 0 {
      println!("");
    }
  };

  content.lines().for_each(|line| {
    if line.starts_with("noop") {
      print_or_not(cycle, val);
      cycle += 1;
    } else {
      print_or_not(cycle, val);
      cycle += 1;

      print_or_not(cycle, val);
      cycle += 1;

      val += line.split(" ").collect_vec()[1].parse::<i32>().unwrap();
    }
  });

  0
}

#[cfg(test)]
mod tests {
  use super::*;

  #[test]
  fn part1_test() {
    let result = part1(CASE);
    assert_eq!(result, 13140)
  }

  #[test]
  fn part2_test() {
    let result = part2(CASE);
    assert_eq!(result, 0)
  }
  const CASE: &str = "addx 15
addx -11
addx 6
addx -3
addx 5
addx -1
addx -8
addx 13
addx 4
noop
addx -1
addx 5
addx -1
addx 5
addx -1
addx 5
addx -1
addx 5
addx -1
addx -35
addx 1
addx 24
addx -19
addx 1
addx 16
addx -11
noop
noop
addx 21
addx -15
noop
noop
addx -3
addx 9
addx 1
addx -3
addx 8
addx 1
addx 5
noop
noop
noop
noop
noop
addx -36
noop
addx 1
addx 7
noop
noop
noop
addx 2
addx 6
noop
noop
noop
noop
noop
addx 1
noop
noop
addx 7
addx 1
noop
addx -13
addx 13
addx 7
noop
addx 1
addx -33
noop
noop
noop
addx 2
noop
noop
noop
addx 8
noop
addx -1
addx 2
addx 1
noop
addx 17
addx -9
addx 1
addx 1
addx -3
addx 11
noop
noop
addx 1
noop
addx 1
noop
noop
addx -13
addx -19
addx 1
addx 3
addx 26
addx -30
addx 12
addx -1
addx 3
addx 1
noop
noop
noop
addx -9
addx 18
addx 1
addx 2
noop
noop
addx 9
noop
noop
noop
addx -1
addx 2
addx -37
addx 1
addx 3
noop
addx 15
addx -21
addx 22
addx -6
addx 1
noop
addx 2
addx 1
noop
addx -10
noop
noop
addx 20
addx 1
addx 2
addx 2
addx -6
addx -11
noop
noop
noop";
}
