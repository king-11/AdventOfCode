use std::{fmt::Debug, str::FromStr, string::ParseError};

use itertools::Itertools;
use nom::{*, bytes::complete::tag, character::complete::{self, newline, multispace1, alphanumeric1}, multi::{separated_list0, separated_list1}, branch::alt};

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

fn parse_items(input: &str) -> IResult<&str, Vec<i64>> {
    let (input, _) = multispace1(input)?;
    let (input, _) = tag("Starting items: ")(input)?;
    let (input, items) = separated_list0(tag(", "), complete::i64)(input)?;

    Ok((input, items))
}

fn parse_operation(input: &str) -> IResult<&str, Operator> {
    let (input, _) = multispace1(input)?;
    let (input, _) = tag("Operation: new = old ")(input)?;
    let (input, op) = alt((tag("+"), tag("*")))(input)?;
    let (input, _) = multispace1(input)?;

    Ok((input, op.parse().unwrap()))
}

fn parse_monkey(input: &str) -> IResult<&str, Monkey> {
    let (input, _) = tag("Monkey ")(input)?;
    let (input, _) = complete::i64(input)?;
    let (input, _) = tag(":")(input)?;
    let (input, _) = newline(input)?;

    let (input, items) = parse_items(input)?;
    let (input, _) = newline(input)?;

    let (input, op) = parse_operation(input)?;
    let (input, val) = alt((tag("old"), alphanumeric1))(input)?;
    let op_val = match val.parse::<i64>() {
        Err(_) => None,
        Ok(val) => Some(val)
    };
    let (input, _) = newline(input)?;

    let (input, _) = multispace1(input)?;
    let (input, _) = tag("Test: divisible by ")(input)?;
    let (input, test) = complete::i64(input)?;
    let (input, _) = newline(input)?;

    let (input, _) = multispace1(input)?;
    let (input, _) = tag("If true: throw to monkey ")(input)?;
    let (input, true_val) = complete::i64(input)?;
    let (input, _) = newline(input)?;

    let (input, _) = multispace1(input)?;
    let (input, _) = tag("If false: throw to monkey ")(input)?;
    let (input, false_val) = complete::i64(input)?;

    Ok((input, Monkey { items, op_val, op, test, inspected: 0, throw: (true_val as usize, false_val as usize)}))
}

fn parse_monkeys(input: &str) -> IResult<&str, Vec<Monkey>> {
    let (input, monkeys) = separated_list1(tag("\n\n"), parse_monkey)(input)?;

    Ok((input, monkeys))
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
    let mut monkeys = parse_monkeys(content).unwrap().1;
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
