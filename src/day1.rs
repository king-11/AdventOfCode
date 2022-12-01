pub fn part1(content: &str) -> i32 {
  let mut current_count = 0;
  let mut max_count = 0;
  for line in content.lines() {
    if line == "" {
      max_count = std::cmp::max(current_count, max_count);
      current_count = 0;
      continue;
    }

    current_count += line.parse::<i32>().unwrap();
  }

  max_count
}

pub fn part2(content: &str) -> i32 {
  let mut calories: Vec<i32> = vec![];
  let mut current_count = 0;
  for line in content.lines() {
    if line == "" {
      calories.push(current_count);
      current_count = 0;
      continue;
    }

    current_count += line.parse::<i32>().unwrap();
  }

  calories.sort_by_key(|a| -a);
  calories[..3].iter().sum()
}
