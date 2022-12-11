use itertools::Itertools;

#[allow(dead_code)]
pub fn part1(content: &str) -> i32 {
    let grid = content
        .lines()
        .map(|line| line.as_bytes().iter().map(|c| c - b'0').collect_vec())
        .collect_vec();

        let mut around_grid = (0..grid.len())
        .map(|_| (0..grid[0].len()).map(|_| (0, 0, 0, 0)).collect_vec())
        .collect_vec();

        let mut marked = (0..grid.len())
        .map(|_| (0..grid[0].len()).map(|_| false).collect_vec())
        .collect_vec();

        for (i, row) in grid.as_slice().iter().enumerate() {
            for (j, &x) in row.as_slice().iter().enumerate() {
                if i > 0 {
                around_grid[i][j].0 = grid[i - 1][j].max(around_grid[i - 1][j].0);
            }
            if j > 0 {
                around_grid[i][j].3 = grid[i][j - 1].max(around_grid[i][j - 1].3);
            }

            if i == 0
            || j == 0
            || i == grid.len() - 1
            || j == row.len() - 1
            || x > around_grid[i][j].0
            || x > around_grid[i][j].3
            {
                marked[i][j] = true;
            }
        }

        for (j, &x) in row.as_slice().iter().enumerate().rev() {
            if j + 1 < row.len() {
                around_grid[i][j].1 = grid[i][j + 1].max(around_grid[i][j + 1].1);
            }

            if x > around_grid[i][j].1 {
                marked[i][j] = true;
            }
        }
    }

    for (i, row) in grid.as_slice().iter().enumerate().rev() {
        for (j, &x) in row.as_slice().iter().enumerate() {
            if i + 1 < grid.len() {
                around_grid[i][j].2 = grid[i + 1][j].max(around_grid[i + 1][j].2);
            }

            if x > around_grid[i][j].2 {
                marked[i][j] = true
            }
        }
    }

    marked
    .as_slice()
    .iter()
        .map(|row| row.as_slice().iter().filter(|&x| *x).count())
        .sum::<usize>() as i32
}

#[allow(dead_code)]
pub fn part2(content: &str) -> i32 {
    let grid = content
    .lines()
    .map(|line| line.as_bytes().iter().map(|c| c - b'0').collect_vec())
        .collect_vec();

    let mut max_area = 0;
    for (i, row) in grid.as_slice().iter().enumerate() {
        for (j, &x) in row.as_slice().iter().enumerate() {
            let mut area = 1;

            let mut count = 0;
            let mut idx = j as i32 - 1;
            loop {
                if idx < 0 {
                    break;
                }
                count += 1;
                if grid[i][idx as usize] < x {
                    idx -= 1;
                } else {
                    break;
                }
            }
            area *= count;

            let mut count = 0;
            let mut idx = i as i32 - 1;
            loop {
                if idx < 0 {
                    break;
                }
                count += 1;
                if grid[idx as usize][j] < x {
                    idx -= 1;
                } else {
                    break;
                }
            }
            area *= count;

            let mut count = 0;
            let mut idx = (j + 1) as i32;
            loop {
                if idx >= row.len() as i32 {
                    break;
                }
                count += 1;
                if grid[i][idx as usize] < x {
                    idx += 1;
                } else {
                    break;
                }
            }
            area *= count;

            let mut count = 0;
            let mut idx = (i + 1) as i32;
            loop {
                if idx >= grid.len() as i32 {
                    break;
                }
                count += 1;
                if grid[idx as usize][j] < x {
                    idx += 1;
                } else {
                    break;
                }
            }
            area *= count;
            max_area = max_area.max(area);
        }
    }

    max_area
}

#[cfg(test)]
mod tests {
    use super::*;

    const CASE: &str = "30373
25512
65332
33549
35390";

    #[test]
    fn part1_test() {
        let result = part1(CASE);
        assert_eq!(result, 21);
    }

    #[test]
    fn part2_test() {
        let result = part2(CASE);
        assert_eq!(result, 8);
    }
}
