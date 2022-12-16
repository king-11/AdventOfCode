use std::{cmp::Ordering, iter::zip};

use itertools::Itertools;
use json::JsonValue;

fn compare_json(left: &JsonValue, right: &JsonValue) -> Option<bool> {
    match (left, right) {
        (JsonValue::Number(_), JsonValue::Number(_)) => {
            let (i, j) = (left.as_i32().unwrap(), right.as_i32().unwrap());
            if i == j {
                return None;
            }

            Some(i < j)
        }
        (JsonValue::Number(_), JsonValue::Array(_)) => {
            let x = left.as_i32().unwrap();
            let arr_left = JsonValue::from([x].to_vec());
            compare_json(&arr_left, right)
        }
        (JsonValue::Array(_), JsonValue::Number(_)) => {
            let x = right.as_i32().unwrap();
            let arr_right = JsonValue::from([x].to_vec());
            compare_json(left, &arr_right)
        }
        (JsonValue::Array(arr_left), JsonValue::Array(arr_right)) => {
            for (x, y) in zip(arr_left, arr_right) {
                let res = compare_json(x, y);
                if res.is_some() {
                    return Some(res.unwrap());
                }
            }

            if arr_left.len() == arr_right.len() {
                return None;
            }

            Some(arr_left.len() < arr_right.len())
        }
        _ => unreachable!(),
    }
}

#[allow(dead_code)]
pub fn part1(content: &str) -> i32 {
    content
        .lines()
        .filter(|line| !line.is_empty())
        .map(|line| json::parse(line).unwrap())
        .tuples::<(_, _)>()
        .map(|(left, right)| compare_json(&left, &right).unwrap())
        .enumerate()
        .filter_map(|(idx, res)| if res { Some(idx + 1) } else { None })
        .sum::<usize>() as i32
}

#[allow(dead_code)]
pub fn part2(content: &str) -> i32 {
    let dec_2 = JsonValue::from(vec![vec![2]]);
    let dec_6 = JsonValue::from(vec![vec![6]]);
    let mut vals = content
        .lines()
        .filter(|line| !line.is_empty())
        .map(|line| json::parse(line).unwrap())
        .collect_vec();

    vals.push(dec_2);
    vals.push(dec_6);

    vals.sort_by(|a, b| match compare_json(a, b) {
        None => Ordering::Equal,
        Some(val) => {
            if val {
                Ordering::Less
            } else {
                Ordering::Greater
            }
        }
    });

    let dec_2 = JsonValue::from(vec![vec![2]]);
    let dec_6 = JsonValue::from(vec![vec![6]]);
    vals.iter()
        .positions(|val| *val == dec_2 || *val == dec_6)
        .map(|i| i + 1)
        .product::<usize>() as i32
}

#[cfg(test)]
mod tests {
    use super::*;

    const CASE: &str = "[1,1,3,1,1]
[1,1,5,1,1]

[[1],[2,3,4]]
[[1],4]

[9]
[[8,7,6]]

[[4,4],4,4]
[[4,4],4,4,4]

[7,7,7,7]
[7,7,7]

[]
[3]

[[[]]]
[[]]

[1,[2,[3,[4,[5,6,7]]]],8,9]
[1,[2,[3,[4,[5,6,0]]]],8,9]";

    #[test]
    fn part1_test() {
        let res = part1(CASE);
        assert_eq!(res, 13);
    }

    #[test]
    fn part2_test() {
        let res = part2(CASE);
        assert_eq!(res, 140);
    }
}
