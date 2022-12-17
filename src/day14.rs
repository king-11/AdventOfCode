use std::collections::HashSet;

use itertools::Itertools;

#[derive(PartialEq, Eq, PartialOrd, Ord, Hash, Clone, Copy, Debug)]
struct Position(i32, i32);

impl Position {
    fn new(x: i32, y: i32) -> Self {
        Position(x, y)
    }
    fn move_tup(&mut self, graph: &HashSet<Position>, y_max: i32) -> bool {
        if self.1 >= y_max {
            return false;
        }
        let down = Position(self.0, self.1 + 1);
        let down_left = Position(self.0 - 1, self.1 + 1);
        let down_right = Position(self.0 + 1, self.1 + 1);

        let movements = vec![down, down_left, down_right];

        match movements.iter().find(|pos| !graph.contains(pos)) {
            None => false,
            Some(val) => {
                *self = *val;
                true
            }
        }
    }
}

fn parse_walls(content: &str) -> HashSet<Position> {
    let mut graph = HashSet::new();
    content.lines().for_each(|line| {
        let ranges = line
            .split(" -> ")
            .map(|range| {
                range
                    .split(",")
                    .map(|num| num.parse::<i32>().unwrap())
                    .collect_tuple::<(i32, i32)>()
                    .unwrap()
            })
            .collect_vec();

        ranges
            .into_iter()
            .tuple_windows::<((i32, i32), (i32, i32))>()
            .for_each(|(tup1, tup2)| {
                if tup1.0 == tup2.0 {
                    let (start, end) = (tup1.1.min(tup2.1), tup1.1.max(tup2.1));
                    for y in start..=end {
                        graph.insert(Position::new(tup1.0, y));
                    }
                } else {
                    let (start, end) = (tup1.0.min(tup2.0), tup1.0.max(tup2.0));
                    for x in start..=end {
                        graph.insert(Position::new(x, tup1.1));
                    }
                }
            });
    });

    graph
}

#[allow(dead_code)]
pub fn part1(content: &str) -> i32 {
    let mut graph = parse_walls(content);
    let y_max = graph.iter().map(|tup| tup.1).max().unwrap();

    let mut sand_particles = 0;
    loop {
        let mut location = Position::new(500, 0);
        while location.move_tup(&graph, y_max + 1) {}
        if location.1 > y_max {
            break;
        }
        sand_particles += 1;
        graph.insert(location);
    }

    sand_particles
}

#[allow(dead_code)]
pub fn part2(content: &str) -> i32 {
  let mut graph = parse_walls(content);
    let y_max = graph.iter().map(|tup| tup.1).max().unwrap() + 1;

    let mut sand_particles = 0;
    loop {
        let mut location = Position::new(500, 0);
        while location.move_tup(&graph, y_max) {}

        sand_particles += 1;
        graph.insert(location);

        if location.1 == 0 {
            break;
        }
    }

    sand_particles
}

#[cfg(test)]
mod tests {
    use super::*;

    const CASE: &str = "498,4 -> 498,6 -> 496,6
503,4 -> 502,4 -> 502,9 -> 494,9";

    #[test]
    fn test_part1() {
        let res = part1(CASE);
        assert_eq!(res, 24);
    }

    #[test]
    fn test_part2() {
        let res = part2(CASE);
        assert_eq!(res, 93);
    }
}
