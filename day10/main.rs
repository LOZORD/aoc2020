// use cached::proc_macro::cached;
use log::{info, trace};
use std::collections::BTreeSet;
// use std::collections::{HashMap};
use std::env;
use std::io::{stdin, BufRead};
// use std::vec::Vec;
// extern crate cached;
// #[macro_use]
// extern crate cached;

fn main() {
    env_logger::init();
    trace!("starting up...");

    let nums: BTreeSet<i64> = stdin()
        .lock()
        .lines()
        .map(|line| line.unwrap().parse().unwrap())
        .collect();

    info!("got nums = {:?}", nums);

    let is_part2 = !env::var("PART2")
        .ok()
        .unwrap_or(String::from(""))
        .is_empty();

    if !is_part2 {
        part1(&nums);
    } else {
       println!("Part 2 is solved in another castle... (see Python code)")
    }
}

fn part1(nums: &BTreeSet<i64>) {
    let mut diff1 = 0;
    let mut diff3 = 0;
    let max = *nums.iter().max().unwrap();
    let mut cur: i64 = max;
    for n in nums.iter().rev() {
        if *n == max {
            continue;
        }
        let diff = cur - n;
        info!("{} - {} = {}", cur, n, diff);
        match diff {
            1 => diff1 += 1,
            3 => diff3 += 1,
            _ => (),
        }
        cur = *n;
    }

    // Adaptor chain -> device.
    match cur {
        1 => diff1 += 1,
        3 => diff3 += 1,
        _ => (),
    }

    // Device -> outlet.
    diff3 += 1;

    println!("DIFF_1 = {}", diff1);
    println!("DIFF_3 = {}", diff3);
    println!("DIFF_1 * DIFF_3 = {}", diff1 * diff3);
}

/*

NOTE: I couldn't solve Part 2 in Rust because some sort of global memoiztion/caching needs to happen.
Rust makes that challenging (for my noob brain), so I practiced my Python for Part 2 instead. 

fn part2(nums: &BTreeSet<i64>) {
    let s: Vec<i64> = nums.iter().rev().map(|n| *n).collect();
    println!("COUNT = {}", do_count(&s, 0, *nums.iter().max().unwrap()));
}

// use cached::SizedCache;
// use std::thread::sleep;
// use std::time::Duration;

// cached_key! {
//     Key = { format!("{}-{}", idx, cur); };

// }

// #[cached()]
fn do_count(nums: &[i64], idx: usize, cur: i64) -> i64 {
    return do_count2(nums, idx, cur);
}

struct P2Solver {
    // nums: 
    nums: Vec<i64>,
    memo:  HashMap<(usize, i64), i64>
}

impl P2Solver {
    fn slice(&self) -> &[i64] {
        return &self.nums;
    }

    // fn new2(nums: BTreeSet<i64>) -> P2Solver {
    //     return P2Solver::new();
    // }

    fn count(&mut self, idx: usize, cur: i64) -> i64 {
        if idx == self.slice().len() {
            if cur <= 3 {
                return 1;
            } else {
                return 0;
            }
        }

        let mut sum = 0;
        for (i, n) in self.slice()[idx..].iter().enumerate() {
            let diff = cur - n;
            let within_joltage = (0 <= diff && diff <= 3);
            if !within_joltage {
                continue;
            }

            let mut c = 0;
            if self.memo.contains_key(&(i, *n)) {
               c= *self.memo.get(&(i,*n)).unwrap();
            } else {
                c = self.count(i, *n);
                // let x = self.memo.get_mut
                // self.memo.insert((i, *n), c);
            }
        }

        return sum;
    }
}

fn do_count2(nums: &[i64], idx: usize, cur: i64) -> i64 {
    trace!("examining cur = {}, idx = {}", cur, idx);

    let ns = &nums[idx..];

    if cur == 0 && ns.len() == 0 {
        trace!("reached base case 1");
        return 1;
    }

    if ns.len() == 0  {
        trace!("ran out of things to check");
        return 0;
    }


   
    let mut sum = 0;
    for (i, n) in ns.iter().enumerate() {
        let diff = cur - n;
        if 1 <= diff && diff <= 3 {
            let next = i + 1;
            trace!(
                "recursing on cur={} from diff={} with num length={}",
                *n,
                diff,
                ns.len() - 1,
            );
            sum += do_count(&nums, next, *n);
        }
        if diff > 3 {
            break;
        }
    }

    return sum;
}
*/