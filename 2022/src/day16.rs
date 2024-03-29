use std::{
    collections::{BTreeMap, BTreeSet, BinaryHeap, HashMap},
    fmt::Display,
};

use itertools::Itertools;
use rayon::{*, prelude::{IntoParallelRefMutIterator, IndexedParallelIterator, ParallelIterator, IntoParallelRefIterator}};
use nom::{
    branch::alt,
    bytes::complete::tag,
    character::complete::{self, alpha1, newline},
    multi::{separated_list0, separated_list1},
    sequence::{preceded, separated_pair, terminated},
    IResult,
};

#[derive(Clone, Debug)]
struct Node {
    name: String,
    flow: u32,
    connections: Vec<u32>,
}

impl Node {
    fn new(name: &str, flow: u32, connections: Vec<u32>) -> Self {
        Self {
            name: name.to_string(),
            flow: flow,
            connections: connections,
        }
    }
}

fn parse_node(input: &str) -> IResult<&str, (&str, u32, Vec<&str>)> {
    let (input, (name, flow)) = terminated(
        separated_pair(
            preceded(tag("Valve "), alpha1),
            tag(" has flow rate="),
            complete::u32,
        ),
        tag("; "),
    )(input)?;
    let (input, _) = alt((
        tag("tunnels lead to valves "),
        tag("tunnel leads to valve "),
    ))(input)?;
    let (input, connections) = separated_list0(tag(", "), alpha1)(input)?;

    Ok((input, (name, flow, connections)))
}

fn parse_lines(input: &str) -> IResult<&str, Vec<Node>> {
    let (input, nodes_connections) = separated_list1(newline, parse_node)(input)?;

    let mut n: u32 = 0;
    let mut order_name = HashMap::new();
    for (name, _, _) in &nodes_connections {
        order_name.insert(name.to_string(), n);
        n += 1;
    }

    let mut nodes = vec![];
    for (name, flow, connections) in nodes_connections.iter() {
        let connect_u32 = connections
            .iter()
            .map(|name| *order_name.get(&name.to_string()).unwrap())
            .collect_vec();

        nodes.push(Node::new(name, *flow, connect_u32));
    }

    Ok((input, nodes))
}

#[derive(PartialEq, Eq, PartialOrd, Ord)]
struct SearchSpace(i32, i32, usize, BTreeSet<usize>);

impl Display for SearchSpace {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        write!(f, "score: {}, time: {}, idx: {}, opened: {:?}", self.0, self.1, self.2, self.3)
    }
}

#[allow(dead_code)]
pub fn part1(content: &str) -> i32 {
    let (_, nodes) = parse_lines(content).unwrap();
    let sign_nodes: HashMap<usize, usize> = HashMap::from_iter(
        nodes
            .iter()
            .enumerate()
            .filter_map(|(i, node)| if node.flow > 0 { Some(i) } else { None })
            .enumerate()
            .map(|(new_idx, i)| (i, new_idx + 1))
    );

    let node_count = sign_nodes.len();
    let mut node_mask = 0;
    for i in 1..=node_count {
        node_mask |= 1 << i;
    }

    part(&nodes, 30, node_mask, &sign_nodes)
}

#[allow(dead_code)]
pub fn part2(content: &str) -> i32 {
    let (_, nodes) = parse_lines(content).unwrap();
    let sign_nodes: HashMap<usize, usize> = HashMap::from_iter(
        nodes
            .iter()
            .enumerate()
            .filter_map(|(i, node)| if node.flow > 0 { Some(i) } else { None })
            .enumerate()
            .map(|(mask_id, i)| (i, mask_id + 1))
    );
    let node_count = sign_nodes.len();

    let mut node_mask = 0;
    for i in 1..=node_count {
        node_mask |= 1 << i;
    }

    let mut pressure = vec![0; node_mask+1];

    pressure.par_iter_mut().enumerate().for_each(|(i , p)| {
        *p = part(&nodes, 26, i as i32, &sign_nodes);
    });

    pressure.par_iter().enumerate().map(|(i, p)| {
        let mask_e = !i & node_mask;
        *p + pressure[mask_e]
    }).max().unwrap() as i32
}

fn part(nodes: &Vec<Node>, max_time: i32, consider_nodes: i32, sign_nodes: &HashMap<usize, usize>) -> i32 {
    let mut prune_queue = BinaryHeap::new();

    let mut time_idx_set: BTreeMap<(i32, usize), i32> = BTreeMap::new();

    let start = nodes.iter().position(|node| node.name == "AA").unwrap();

    // score, time, usize, opened
    prune_queue.push(SearchSpace(0, 1, start, BTreeSet::new()));
    let mut max_pressure = 0;

    while prune_queue.len() != 0 {
        let search = prune_queue.pop().unwrap();

        let (score, time, idx, mut opened) = (search.0, search.1, search.2, search.3);

        if *time_idx_set.get(&(time, idx)).unwrap_or(&-1) >= score {
            continue;
        }

        let mask_idx = *sign_nodes.get(&idx).unwrap_or(&0) as i32;
        if (consider_nodes & (1 << mask_idx) == 0) && mask_idx != 0 {
            continue;
        }

        time_idx_set.insert((time, idx), score);
        if time == max_time {
            max_pressure = max_pressure.max(score);
            continue;
        }

        if !opened.contains(&idx) && nodes[idx].flow > 0 {
            opened.insert(idx);

            let new_score = score + opened.iter().map(|&i| nodes[i].flow).sum::<u32>() as i32;
            let new_state = SearchSpace(new_score, time + 1, idx, opened.clone());

            prune_queue.push(new_state);
            opened.remove(&idx);
        }

        let new_score = score + opened.iter().map(|&i| nodes[i].flow).sum::<u32>() as i32;
        for connect in nodes[idx].connections.iter() {
            let new_state = SearchSpace(new_score, time + 1, *connect as usize, opened.clone());
            prune_queue.push(new_state);
        }
    }

    max_pressure as i32
}

#[cfg(test)]
mod tests {
    use super::*;

    const CASE: &str = "Valve AA has flow rate=0; tunnels lead to valves DD, II, BB
Valve BB has flow rate=13; tunnels lead to valves CC, AA
Valve CC has flow rate=2; tunnels lead to valves DD, BB
Valve DD has flow rate=20; tunnels lead to valves CC, AA, EE
Valve EE has flow rate=3; tunnels lead to valves FF, DD
Valve FF has flow rate=0; tunnels lead to valves EE, GG
Valve GG has flow rate=0; tunnels lead to valves FF, HH
Valve HH has flow rate=22; tunnel leads to valve GG
Valve II has flow rate=0; tunnels lead to valves AA, JJ
Valve JJ has flow rate=21; tunnel leads to valve II";

    #[test]
    fn test_part1() {
        let res = part1(CASE);
        assert_eq!(res, 1651);
    }

    #[test]
    fn test_part2() {
        let res = part2(CASE);
        assert_eq!(res, 1707);
    }
}
