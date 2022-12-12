use itertools::Itertools;
use std::{collections::HashSet, ops::{AddAssign, SubAssign}};

#[derive(Hash, PartialEq, Eq, Clone, Copy, Debug)]
struct Point(i32, i32);

impl Point {
  fn new(tuple: (i32, i32)) -> Self {
    Point(tuple.0, tuple.1)
  }
  fn distance(&self, other: &Point) -> i32 {
    (self.0 - other.0).abs().max((self.1 - other.1).abs())
  }
  fn delta(&self, other: &Point) -> Point {
    let delta_x = if self.0 > other.0 { -1 } else if self.0 < other.0 { 1 } else { 0 };
    let delta_y = if self.1 > other.1 { -1 } else if self.1 < other.1 { 1 } else { 0 };
    Point(delta_x, delta_y)
  }
}

impl AddAssign for Point {
  fn add_assign(&mut self, rhs: Self) {
    self.0 += rhs.0;
    self.1 += rhs.1;
  }
}

impl SubAssign for Point {
  fn sub_assign(&mut self, rhs: Self) {
    self.0 -= rhs.0;
    self.1 -= rhs.1;
  }
}

#[allow(dead_code)]
pub fn part1(content: &str) -> i32 {
  part(content, 2)
}

#[allow(dead_code)]
pub fn part2(content: &str) -> i32 {
  part(content, 10)
}

fn part(content: &str, leng: i32) -> i32 {
  let mut marked_grid = HashSet::new();
  let mut points = (0..leng).map(|_| Point(0,0)).collect_vec();

  marked_grid.insert(points[leng as usize - 1]);
  for line in content.lines() {
    let (dir, val) = line.split(" ").collect_tuple().unwrap();
    let mut val = val.parse::<i32>().unwrap();

    let mov = match dir {
      "U" => (0, 1),
      "D" => (0, -1),
      "R" => (1, 0),
      "L" => (-1, 0),
      _ => unreachable!()
    };

    let mov_p = Point::new(mov);
    // println!("{} {} {:?}", dir, val, mov_p);

    while val > 0 {
      val -= 1;
      points[0] += mov_p;
      for idx in 1..leng as usize {
        let dest = points[idx-1];
        let start = &mut points[idx];
        if dest.distance(start) > 1 {
          *start += start.delta(&dest);
        }
      }
      marked_grid.insert(points[leng as usize-1]);
      // print!("({:?}, {:?}) ", points[leng as usize - 1], points[0]);
    }
    // println!("");
  }

  marked_grid.len() as i32
}

#[cfg(test)]
mod tests {
  use super::*;
  use test_case::test_case;

  const CASE1: &str = "R 4
U 4
L 3
D 1
R 4
D 1
L 5
R 2";

  const CASE2: &str = "R 5
U 8
L 8
D 3
R 17
D 10
L 25
U 20";

  #[test]
  fn part1_test() {
    let result = part1(CASE1);
    assert_eq!(result, 13);
  }

  #[test_case(CASE1, 1)]
  #[test_case(CASE2, 36)]
  fn part2_test(case: &str, res: i32) {
    let result2 = part2(case);
    assert_eq!(result2, res);
  }
}
