use std::collections::HashSet;

use itertools::Itertools;

use if_chain::if_chain;

fn change_to_int(ch: char) -> u32 {
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
            let leng = item.len();
            let mut std_set = HashSet::new();
            let mut chr = item.chars().nth(0).unwrap();
            for (i, ch) in item.char_indices() {
                if i * 2 >= leng {
                    if std_set.contains(&ch) {
                        chr = ch;
                    }
                } else {
                    std_set.insert(ch);
                }
            }
            chr
        })
        .map(change_to_int)
        .sum::<u32>() as i32
}

#[allow(dead_code)]
pub fn part2(content: &str) -> i32 {
    content
        .lines()
        .batching(|it| {
            if_chain! {
                if let Some(s1) = it.next();
                if let Some(s2) = it.next();
                if let Some(s3) = it.next();
                then {
                    return Some((s1, s2, s3));
                }
            }

            None
        })
        .map(|group| {
            let st1 = group.0.chars().collect::<HashSet<char>>();
            let st2 = group.1.chars().collect::<HashSet<char>>();

            let mut chr = group.0.chars().nth(0).unwrap();

            for ch in group.2.chars() {
                if st1.contains(&ch) && st2.contains(&ch) {
                    chr = ch;
                }
            }

            chr
        })
        .map(change_to_int)
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
