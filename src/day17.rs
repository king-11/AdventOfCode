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
    rocks: Vec<u8>,
    position: usize
}

lazy_static! {
    static ref ROCKS: Vec<RockFormation> = vec![
        // minus
        RockFormation {
            rocks: vec![0b_1111],
            position: 0
        },
        // plus
        RockFormation {
            rocks: vec![0b_010, 0b_111, 0b_010],
            position: 0
        },
        // L shape
        RockFormation {
            rocks: vec![0b_100, 0b_100, 0b_111],
            position: 0
        },
        // Line
        RockFormation {
            rocks: vec![0b_1, 0b_1, 0b_1, 0b_1],
            position: 0
        },
        // square
        RockFormation {
            rocks: vec![0b_11, 0b_11],
            position: 0
        }
    ];
}

impl RockFormation {
    fn air_displace(&mut self, air_move: &Move, blocked: &BTreeMap<usize, u8>) {
        let new_vals = match air_move {
            Move::Left => self.rocks.iter().map(|val| *val >> 1).collect_vec(),
            Move::Right => self.rocks.iter().map(|val| *val << 1).collect_vec(),
        };

        if new_vals
            .iter()
            .zip(self.rocks.iter())
            .rev()
            .enumerate()
            .any(|(idx, (&val, &old_val))| *blocked.get(&(self.position + idx)).unwrap_or(&0) & val != 0 || val.count_ones() != old_val.count_ones() || val & 1 << 7 != 0)
        {
            return;
        }

        self.rocks = new_vals;
    }
    fn move_down(&mut self) {
        self.position -= 1;
    }
    fn bottom_found(&mut self, blocked: &BTreeMap<usize, u8>) -> bool {
        self.rocks
            .iter()
            .rev()
            .enumerate()
            .any(|(idx, &val)| {
                *blocked.get(&(idx + self.position - 1)).unwrap_or(&0) & val != 0
            })
    }
}

pub fn part1(content: &str) -> usize {
    let (_, air) = parse_air(content).unwrap();
    let mut air_moves = air.iter().cycle();

    let mut rock_moves = ROCKS.iter().cycle();
    let mut blocked: BTreeMap<usize, u8> = BTreeMap::new();

    blocked.insert(0, 0b1111111);
    let mut y_max = 0;
    for _ in 0..2022 {
        let mut rock = rock_moves.next().unwrap().clone();

        rock.position = y_max + 3 + 1;
        rock.rocks.iter_mut().for_each(|val| *val <<= 2);

        loop {
            match air_moves.next() {
                None => unreachable!(),
                Some(mov) => rock.air_displace(mov, &blocked),
            };

            if rock.bottom_found(&blocked) {
                y_max = y_max.max(rock.position + (rock.rocks.len() - 1));
                for (idx, val) in rock.rocks.iter().rev().enumerate() {
                    match blocked.get_mut(&(idx + rock.position)) {
                        Some(block) => {*block |= *val;},
                        None => {
                            blocked.insert(idx + rock.position, *val);
                        }
                    }
                }
                break;
            }

            rock.move_down();
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
