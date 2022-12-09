use std::collections::VecDeque;

use itertools::Itertools;
use regex::Regex;

#[allow(dead_code)]
pub fn part2(content: &str) -> i32 {
  let mut lines = content.lines();

  let stacks = lines
    .by_ref()
    .take_while(|line| !line.is_empty())
    .map(|line| {
      line.chars()
        .chunks(4)
        .into_iter()
        .flat_map(|mut chars| chars.nth(1))
        .map(|chr| if chr == ' ' { None } else { Some(chr) })
        .collect::<VecDeque<Option<char>>>()
    })
    .collect::<VecDeque<VecDeque<Option<char>>>>();

  let mut stacks = convert_stack(stacks);
  let re = Regex::new(r"move (\d+) from (\d+) to (\d+)").unwrap();

  lines
    .for_each(|line| {
      let captures = re.captures(line).unwrap();

      let (cn, src, dest) = (captures.get(1).unwrap().as_str().parse::<usize>().unwrap(), captures.get(2).unwrap().as_str().parse::<usize>().unwrap(), captures.get(3).unwrap().as_str().parse::<usize>().unwrap());

      let src_take = stacks[src-1].iter().map(|x| *x).take(cn).collect_vec();

      for x in src_take.iter().rev() {
        stacks[dest-1].push_front(*x);
      }

      (0..cn).for_each(|_| { stacks[src-1].pop_front(); });
    });


  stacks.iter().for_each(|x| print!("{}", x.front().unwrap()));
  println!("");
  0
}

fn convert_stack<T: Copy>(stacks: VecDeque<VecDeque<Option<T>>>) -> VecDeque<VecDeque<T>> {
  let max_size = stacks.iter().map(|s| s.len()).max().unwrap();
  let mut new_stacks = VecDeque::new();
  (0..max_size).for_each(|_| new_stacks.push_back(VecDeque::new()));

  for x in stacks {
    for (i, &opt_val) in x.iter().enumerate() {
      if opt_val.is_some() {
        let val = opt_val.unwrap();
        new_stacks[i].push_back(val);
      }
    }
  }

  new_stacks
}
