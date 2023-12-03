use itertools::Itertools;
use regex::Regex;
use std::collections::HashMap;

#[derive(Debug)]
enum FileType {
    Directory,
    File(usize),
}

#[derive(Debug)]
struct LinuxFile {
    name: String,
    file_type: FileType,
    children: HashMap<String, usize>,
    parent: Option<usize>,
}

impl LinuxFile {
    fn new(name: String, size: Option<usize>) -> Self {
        let children = HashMap::new();
        match size {
            None => LinuxFile {
                name,
                file_type: FileType::Directory,
                children,
                parent: None,
            },
            Some(val) => LinuxFile {
                name,
                file_type: FileType::File(val),
                children,
                parent: None,
            },
        }
    }
    fn file_size(&self, buf: &Vec<LinuxFile>) -> usize {
        match self.file_type {
            FileType::File(sz) => sz,
            FileType::Directory => self
                .children
                .iter()
                .map(|(_, idx)| buf.get(*idx).unwrap().file_size(buf))
                .sum(),
        }
    }
}

#[allow(dead_code)]
pub fn part1(content: &str) -> i32 {
    let max_size = 100000;
    let pointers = file_tree(content);
    pointers
        .iter()
        .filter_map(|file| {
            let sz = file.file_size(&pointers);
            if sz <= max_size {
                return match file.file_type {
                    FileType::File(_) => None,
                    FileType::Directory => Some(sz),
                };
            }
            None
        })
        .sum::<usize>() as i32
}

#[allow(dead_code)]
pub fn part2(content: &str) -> i32 {
    let pointers = file_tree(content);
    let mut root_size = 0;
    pointers
        .iter()
        .filter_map(|file| {
            let sz = file.file_size(&pointers);
            if file.name == "/".to_string() {
                root_size = sz;
            }
            return match file.file_type {
                FileType::File(_) => None,
                FileType::Directory => Some(sz)
            };
        })
        .sorted()
        .find(|dir_size| 70000000 - root_size + dir_size >= 30000000)
        .unwrap() as i32
}

fn file_tree(content: &str) -> Vec<LinuxFile> {
    let cd_regex = Regex::new(r"\$ cd (.+)").unwrap();
    let ls_regex = Regex::new(r"\$ ls").unwrap();
    let file_regex = Regex::new("(\\d+) (.+)").unwrap();
    let dir_regex = Regex::new("dir (.+)").unwrap();

    let root = LinuxFile::new("/".to_string(), None);
    let mut current_idx = 0usize;

    let mut pointers = vec![];
    pointers.push(root);

    for line in content.lines() {
        let n = pointers.len();
        let pointer = pointers.get_mut(current_idx).unwrap();

        if cd_regex.is_match(line) {
            let target = cd_regex.captures(line).unwrap().get(1).unwrap().as_str();

            if pointer.children.contains_key(target) {
                let child = *pointer.children.get(target).unwrap();
                current_idx = child;
            } else if target == ".." {
                let parent = pointer.parent.unwrap();
                current_idx = parent;
            } else {
                current_idx = 0;
            }
        } else if ls_regex.is_match(line) {
            continue;
        } else if file_regex.is_match(line) {
            let caps = file_regex.captures(line).unwrap();
            let (szf, name) = (
                caps.get(1).unwrap().as_str().parse::<usize>().unwrap(),
                caps.get(2).unwrap().as_str(),
            );

            let mut file = LinuxFile::new(name.to_string(), Some(szf));
            file.parent = Some(current_idx);

            pointer.children.insert(name.to_string(), n);

            pointers.push(file);
        } else if dir_regex.is_match(line) {
            let caps = dir_regex.captures(line).unwrap();
            let name = caps.get(1).unwrap().as_str();
            let mut file = LinuxFile::new(name.to_string(), None);

            file.parent = Some(current_idx);
            pointer.children.insert(name.to_string(), n);
            pointers.push(file);
        }
    }

    pointers
}

#[cfg(test)]
mod tests {
    use super::*;

    const CASE: &str = "$ cd /
$ ls
dir a
14848514 b.txt
8504156 c.dat
dir d
$ cd a
$ ls
dir e
29116 f
2557 g
62596 h.lst
$ cd e
$ ls
584 i
$ cd ..
$ cd ..
$ cd d
$ ls
4060174 j
8033020 d.log
5626152 d.ext
7214296 k";

    #[test]
    fn part1_test() {
        let result = part1(CASE);
        assert_eq!(95437, result);
    }

    #[test]
    fn part2_test() {
        let result = part2(CASE);
        assert_eq!(24933642, result);
    }
}
