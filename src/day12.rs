use std::{
    cmp::Reverse,
    collections::{BinaryHeap, HashSet},
};

use itertools::Itertools;

#[derive(Debug, PartialEq, Eq, PartialOrd, Ord, Copy, Clone)]
struct Item(i32, (usize, usize));

fn convert_to_elevation(chr: char) -> u8 {
    if chr.is_lowercase() {
        chr as u8
    } else if chr == 'S' {
        'a' as u8
    } else {
        'z' as u8
    }
}

#[allow(dead_code)]
pub fn part1(content: &str) -> i32 {
  part(content, &|y: char| y == 'S')
}

#[allow(dead_code)]
pub fn part2(content: &str) -> i32 {
  part(content, &|y: char| y == 'S' || y == 'a')
}

fn part(content: &str, check: &dyn Fn(char) -> bool) -> i32 {
    let grid = content
        .lines()
        .map(|line| line.chars().collect_vec())
        .collect_vec();

    let (n, m) = (grid.len(), grid[0].len());

    let mut queue = BinaryHeap::new();
    let mut marked = HashSet::new();

    for (i, x) in grid.iter().enumerate() {
      for (j, y) in x.iter().enumerate() {
        if check(*y) {
          queue.push(Reverse(Item(0, (i, j))));
          marked.insert((i, j));
        }
      }
    }

    while queue.len() != 0 {
        let item = queue.peek().unwrap().0;
        let (i, j) = (item.1 .0, item.1 .1);
        let distance = item.0;
        // println!("{} {} {:?}", grid[i][j], convert_to_elevation(grid[i][j]), item);

        if grid[i][j] == 'E' {
            return item.0;
        }

        if i > 0 && !marked.contains(&(i - 1, j)) {
            if convert_to_elevation(grid[i - 1][j]) <= convert_to_elevation(grid[i][j]) + 1 {
                queue.push(Reverse(Item(distance + 1, (i - 1, j))));
                marked.insert((i - 1, j));
            }
        }
        if j > 0 && !marked.contains(&(i, j - 1)) {
            if convert_to_elevation(grid[i][j - 1]) <= convert_to_elevation(grid[i][j]) + 1 {
                queue.push(Reverse(Item(distance + 1, (i, j - 1))));
                marked.insert((i, j - 1));
            }
        }
        if i < n - 1 && !marked.contains(&(i + 1, j)) {
            if i == 0 && j == 0
                || convert_to_elevation(grid[i + 1][j]) <= convert_to_elevation(grid[i][j]) + 1
            {
                queue.push(Reverse(Item(distance + 1, (i + 1, j))));
                marked.insert((i + 1, j));
            }
        }
        if j < m - 1 && !marked.contains(&(i, j + 1)) {
            if (i == 0 && j == 0)
                || convert_to_elevation(grid[i][j + 1]) <= convert_to_elevation(grid[i][j]) + 1
            {
                queue.push(Reverse(Item(distance + 1, (i, j + 1))));
                marked.insert((i, j + 1));
            }
        }

        queue.pop();
    }

    0
}

#[cfg(test)]
mod tests {
    use super::*;

    const CASE: &str = "Sabqponm
abcryxxl
accszExk
acctuvwj
abdefghi";

    #[test]
    fn part1_test() {
        let result = part1(CASE);
        assert_eq!(result, 31);
    }

    #[test]
    fn part2_test() {
        let result = part2(CASE);
        assert_eq!(result, 29);
    }
}
