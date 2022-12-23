use std::collections::BTreeMap;

use itertools::Itertools;
use lazy_static::lazy_static;
use nom::{branch::alt, character::complete, multi::many1, IResult, Parser};

enum Move {
    Left,
    Right,
}

fn parse_air(input: &str) -> IResult<&str, Vec<Move>> {
    let (input, moves) = many1(alt((
        complete::char('>').map(|_| Move::Right),
        complete::char('<').map(|_| Move::Left),
    )))(input)?;

    Ok((input, moves))
}

#[derive(Clone)]
struct RockFormation {
    rocks: Vec<(i32, i32)>,
}

const X_MIN: i32 = 1;
const X_MAX: i32 = 7;

lazy_static! {
    static ref ROCKS: Vec<RockFormation> = vec![
        // minus
        RockFormation {
            rocks: vec![(1, 0), (2, 0), (3, 0), (4, 0)]
        },
        // plus
        RockFormation {
            rocks: vec![(2, 2), (1, 1), (2, 1), (3, 1), (2, 0)]
        },
        // L shape
        RockFormation {
            rocks: vec![(3, 2), (3, 1), (1, 0), (2, 0), (3, 0)]
        },
        // Line
        RockFormation {
            rocks: vec![(1, 3), (1, 2), (1, 1), (1, 0)]
        },
        // square
        RockFormation {
            rocks: vec![(1, 1), (2, 1), (1, 0), (2, 0)]
        }
    ];
}

impl RockFormation {
    fn air_displace(&mut self, air_move: &Move, blocked: &BTreeMap<i32, Vec<i32>>) {
        let x_mov = match air_move {
            Move::Left => -1,
            Move::Right => 1,
        };

        if self.rocks.iter().any(|(x, y)| {
            blocked.get(y).unwrap_or(&vec![]).contains(&(x + x_mov))
                || x + x_mov < X_MIN
                || x + x_mov > X_MAX
        }) {
            return;
        }

        self.rocks.iter_mut().for_each(|(x, _)| *x = *x + x_mov);
    }
    fn height_displace(&mut self, y_mov: i32) {
        self.rocks.iter_mut().for_each(|(_, y)| *y = *y + y_mov);
    }
    fn bottom_found(&mut self, blocked: &BTreeMap<i32, Vec<i32>>) -> bool {
        self.rocks.iter().any(|(x, y)| blocked.get(&(y-1)).unwrap_or(&vec![]).contains(x))
    }
}

pub fn part1(content: &str) -> i32 {
    let (_, air) = parse_air(content).unwrap();
    let mut air_moves = air.iter().cycle();

    let mut rock_moves = ROCKS.iter().cycle();
    let mut blocked = BTreeMap::new();

    blocked.insert(0, (1..7).map(|i| i).collect_vec());
    let mut y_max = 0;
    for _ in 0..2022 {
        let mut rock = rock_moves.next().unwrap().clone();
        rock.height_displace(y_max + 3 + 1);

        rock.air_displace(&Move::Right, &blocked);
        rock.air_displace(&Move::Right, &blocked);

        loop {
            match air_moves.next() {
                None => unreachable!(),
                Some(mov) => rock.air_displace(mov, &blocked),
            };

            if rock.bottom_found(&blocked) {
                for (x, y) in rock.rocks.iter() {
                    y_max = y_max.max(*y);
                    match blocked.get_mut(&y) {
                        Some(val) => val.push(*x),
                        None => {
                            blocked.insert(*y, vec![*x]);
                        }
                    }
                }
                break;
            }

            rock.height_displace(-1);
        }
    }

    y_max
}

#[cfg(test)]
mod tests {
    use super::*;

    const CASE: &str = ">>><<><>><<<>><>>><<<>>><<<><<<>><>><<>>";

    #[test]
    fn test_part1() {
        let res = part1(CASE);
        assert_eq!(res, 3068)
    }
}
