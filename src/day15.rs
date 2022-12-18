use std::collections::BTreeSet;

use nom::{
    bytes::complete::tag,
    character::complete::{self, newline},
    multi::separated_list1,
    sequence::{preceded, separated_pair},
    IResult,
};

#[derive(Debug, PartialEq, Eq, PartialOrd, Ord, Clone, Copy)]
struct Position(i64, i64);

impl Position {
    fn distance(&self, other: &Position) -> u64 {
        self.0.abs_diff(other.0) + self.1.abs_diff(other.1)
    }
}

fn parse_sensor(input: &str) -> IResult<&str, Position> {
    let (input, sx) = preceded(tag("Sensor at x="), complete::i64)(input)?;
    let (input, sy) = preceded(tag(", y="), complete::i64)(input)?;

    Ok((input, Position(sx, sy)))
}

fn parse_beacon(input: &str) -> IResult<&str, Position> {
    let (input, sx) = preceded(tag("closest beacon is at x="), complete::i64)(input)?;
    let (input, sy) = preceded(tag(", y="), complete::i64)(input)?;

    Ok((input, Position(sx, sy)))
}

fn parse_line(input: &str) -> IResult<&str, (Position, Position)> {
    let (input, (sensor, beacon)) = separated_pair(parse_sensor, tag(": "), parse_beacon)(input)?;

    Ok((input, (sensor, beacon)))
}

fn parse_lines(input: &str) -> IResult<&str, Vec<(Position, Position)>> {
    let (input, result) = separated_list1(newline, parse_line)(input)?;

    Ok((input, result))
}

#[allow(dead_code)]
pub fn part1(content: &str) -> i64 {
    line_by_line(content, 2000000)
}

fn line_by_line(content: &str, y: i64) -> i64 {
    let sb_vec = parse_lines(content).unwrap().1;

    let mut blocked = BTreeSet::new();
    for (s, b) in sb_vec.iter() {
        let distance = s.distance(b) as i64;
        let dis_y = s.1.abs_diff(y) as i64;

        if dis_y > distance {
            continue;
        }

        let dis_x = distance as i64 - dis_y;
        for x in (s.0 - dis_x)..=(s.0 + dis_x) {
            blocked.insert(x);
        }

        if b.1 == y && blocked.contains(&b.0) {
            blocked.remove(&b.0);
        }
    }

    blocked.len() as i64
}

#[allow(dead_code)]
pub fn part2(content: &str) -> i64 {
    bounding_box(content, 4000000)
}

fn bounding_box(content: &str, max_range: i64) -> i64 {
    let sb_vec = parse_lines(content).unwrap().1;

    for (s, b) in sb_vec.iter() {
        let distance = s.distance(b) as i64 + 1;
        for x in (s.0 - distance)..(s.0 + distance) {
            let rem_dist = distance - x.abs_diff(s.0) as i64;
            let positions = vec![Position(x, s.1 + rem_dist), Position(x, s.1 - rem_dist)];

            let result = positions
                .iter()
                .filter(|&pos| pos.0 >= 0 && pos.1 >= 0 && pos.0 <= max_range && pos.1 <= max_range)
                .find(|&pos| {
                    sb_vec
                        .iter()
                        .all(|(si, bi)| si.distance(pos) > si.distance(bi))
                });

            match result {
                None => continue,
                Some(res) => {
                    return res.0 * 4000000 + res.1;
                }
            }
        }
    }

    0
}

#[cfg(test)]
mod tests {
    use super::*;

    const CASE: &str = "Sensor at x=2, y=18: closest beacon is at x=-2, y=15
Sensor at x=9, y=16: closest beacon is at x=10, y=16
Sensor at x=13, y=2: closest beacon is at x=15, y=3
Sensor at x=12, y=14: closest beacon is at x=10, y=16
Sensor at x=10, y=20: closest beacon is at x=10, y=16
Sensor at x=14, y=17: closest beacon is at x=10, y=16
Sensor at x=8, y=7: closest beacon is at x=2, y=10
Sensor at x=2, y=0: closest beacon is at x=2, y=10
Sensor at x=0, y=11: closest beacon is at x=2, y=10
Sensor at x=20, y=14: closest beacon is at x=25, y=17
Sensor at x=17, y=20: closest beacon is at x=21, y=22
Sensor at x=16, y=7: closest beacon is at x=15, y=3
Sensor at x=14, y=3: closest beacon is at x=15, y=3
Sensor at x=20, y=1: closest beacon is at x=15, y=3";

    #[test]
    fn test_part1() {
        let res = line_by_line(CASE, 10);
        assert_eq!(res, 26);
    }

    #[test]
    fn test_part2() {
        let res = bounding_box(CASE, 20);
        assert_eq!(res, 56000011);
    }
}
