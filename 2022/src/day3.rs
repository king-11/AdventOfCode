use std::collections::HashSet;

use itertools::Itertools;

fn change_to_priority(ch: char) -> u32 {
    if ch.is_ascii_lowercase() {
        return (ch as u32 - 'a' as u32) + 1;
    } else {
        return (ch as u32 - 'A' as u32) + 27;
    }
}

#[allow(dead_code)]
pub fn part1(content: &str) -> i32 {
    content
        .lines()
        .map(|item| {
            let (s1, s2) = item.split_at(item.len() / 2);
            let std_set = s1.chars().collect::<HashSet<char>>();
            s2.chars().find(|ch| std_set.contains(ch)).unwrap()
        })
        .map(change_to_priority)
        .sum::<u32>() as i32
}

#[allow(dead_code)]
pub fn part2(content: &str) -> i32 {
    content
        .lines()
        .chunks(3)
        .into_iter()
        .map(|group| {
            let (s1, s2, s3) = group.collect_tuple().unwrap();
            let st1 = s1.chars().collect::<HashSet<char>>();
            let st2 = s2.chars().collect::<HashSet<char>>();
            s3.chars().find(|ch| st1.contains(&ch) && st2.contains(&ch)).unwrap()
        })
        .map(change_to_priority)
        .sum::<u32>() as i32
}

#[cfg(test)]
mod tests {
    use super::*;

    const CASE: &str = "vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw";

    #[test]
    fn part1_test() {
        let result = part1(CASE);
        assert_eq!(result, 157);
    }

    #[test]
    fn part2_test() {
        let result = part2(CASE);
        assert_eq!(result, 70);
    }
}
