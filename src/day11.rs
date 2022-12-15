use std::{fmt::Debug, str::FromStr, string::ParseError};

use itertools::Itertools;
use regex::Regex;

enum Operator {
    Add,
    Multiply,
}

impl Debug for Operator {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        match self {
            Operator::Add => write!(f, "+"),
            Operator::Multiply => write!(f, "*"),
        }
    }
}

impl FromStr for Operator {
    type Err = ParseError;
    fn from_str(s: &str) -> Result<Self, Self::Err> {
        match s {
            "+" => Ok(Operator::Add),
            "*" => Ok(Operator::Multiply),
            _ => unreachable!(),
        }
    }
}

#[derive(Debug)]
struct Monkey {
    items: Vec<i64>,
    op: Operator,
    op_val: Option<i64>,
    test: i64,
    throw: (usize, usize),
    inspected: i64,
}

impl FromStr for Monkey {
    type Err = ParseError;
    fn from_str(s: &str) -> Result<Self, Self::Err> {
        let items_regex = Regex::new(r"\d+").unwrap();
        let op_regex = Regex::new(r"\s+Operation: new = old (.) (\d+|old)").unwrap();
        let test_regex = Regex::new(r"\s+Test: divisible by (\d+)").unwrap();
        let throw_regex = Regex::new(r"\s+If (true|false): throw to monkey (\d+)").unwrap();

        let lines = s.lines().collect_vec();

        let items = items_regex
            .find_iter(lines[1])
            .map(|item| item.as_str().parse::<i64>().unwrap())
            .collect_vec();

        let caps = op_regex.captures(lines[2]).unwrap();
        let (op, op_val) = (
            caps.get(1).unwrap().as_str().parse::<Operator>().unwrap(),
            caps.get(2).unwrap().as_str().parse::<i64>(),
        );

        let op_val = match op_val {
            Ok(val) => Some(val),
            Err(_) => None,
        };

        let test = test_regex
            .captures(lines[3])
            .unwrap()
            .get(1)
            .unwrap()
            .as_str()
            .parse::<i64>()
            .unwrap();

        let true_throw = throw_regex
            .captures(lines[4])
            .unwrap()
            .get(2)
            .unwrap()
            .as_str()
            .parse::<usize>()
            .unwrap();
        let false_throw = throw_regex
            .captures(lines[5])
            .unwrap()
            .get(2)
            .unwrap()
            .as_str()
            .parse::<usize>()
            .unwrap();

        Ok({
            Monkey {
                items,
                op,
                op_val,
                test,
                throw: (true_throw, false_throw),
                inspected: 0,
            }
        })
    }
}

impl Monkey {
    fn operate(&self, val: i64) -> i64 {
        let op_use_val = match self.op_val {
            Some(const_val) => {
                const_val
            },
            None => val
        };

        match self.op {
            Operator::Add => val + op_use_val,
            Operator::Multiply => val * op_use_val
        }
    }
}

#[allow(dead_code)]
pub fn part1(content: &str) -> i64 {
    part(content, 20, 3)
}

#[allow(dead_code)]
pub fn part2(content: &str) -> i64 {
    part(content, 10000, 1)
}

fn part(content: &str, count: i64, worry_reduce: i64) -> i64 {
    let mut monkeys = content.split("\n\n").map(|line| line.parse::<Monkey>().unwrap()).collect_vec();

    let modulo = monkeys.iter().fold(1, |acc, mk| acc * mk.test);

    for _ in 0..count {
        let mut monkey_items = monkeys.iter().map(|mk| mk.items.clone()).collect_vec();
        for (idx, monkey) in monkeys.iter_mut().enumerate() {
            let items = monkey_items[idx].clone();
            for item in items.iter() {
                let worry = (monkey.operate(*item) / worry_reduce) % modulo;
                if worry % monkey.test == 0 {
                    monkey_items[monkey.throw.0].push(worry);
                } else {
                    monkey_items[monkey.throw.1].push(worry);
                }
            }
            monkey.inspected += items.len() as i64;
            monkey_items[idx].clear();
        }
        for (idx, monkey) in monkeys.iter_mut().enumerate() {
            monkey.items = monkey_items[idx].clone();
        }
    }

    monkeys.iter().sorted_by_key(|mk| -mk.inspected).take(2).fold(1, |acc, mk| acc * mk.inspected)
}

#[cfg(test)]
mod tests {
    use super::*;

    const CASE: &str = "Monkey 0:
  Starting items: 79, 98
  Operation: new = old * 19
  Test: divisible by 23
    If true: throw to monkey 2
    If false: throw to monkey 3

Monkey 1:
  Starting items: 54, 65, 75, 74
  Operation: new = old + 6
  Test: divisible by 19
    If true: throw to monkey 2
    If false: throw to monkey 0

Monkey 2:
  Starting items: 79, 60, 97
  Operation: new = old * old
  Test: divisible by 13
    If true: throw to monkey 1
    If false: throw to monkey 3

Monkey 3:
  Starting items: 74
  Operation: new = old + 3
  Test: divisible by 17
    If true: throw to monkey 0
    If false: throw to monkey 1";

    #[test]
    fn part1_test() {
        let result = part1(CASE);
        assert_eq!(result, 10605)
    }

    #[test]
    fn part2_test() {
        let result = part2(CASE);
        assert_eq!(result, 2713310158)
    }
}
