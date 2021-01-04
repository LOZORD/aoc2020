use log::trace;
use std::collections::vec_deque::VecDeque;
use std::collections::HashSet;
use std::env;
use std::io::{stdin, BufRead};
const BUFF_SIZE: usize = 25;
fn main() {
    env_logger::init();

    let bad_num: i64 = match env::var("BAD_NUM") {
        Ok(v) => v.parse().unwrap(),
        Err(_) => 0,
    };

    if bad_num == 0 {
        part1();
    } else {
        part2(bad_num);
    }
}

fn part1() {
    let mut nums: VecDeque<i64> = VecDeque::with_capacity(BUFF_SIZE);

    for line in stdin().lock().lines() {
        let t = line.unwrap();
        trace!("got line = {}", t);
        let n: i64 = t.parse().unwrap();
        trace!("got = {}, len = {}", n, nums.len());

        if nums.iter().count() < BUFF_SIZE {
            nums.push_back(n);
             // We still need to digest the complete preamble.
            continue;
        }

        trace!("preamble complete! {:?}", nums);

        let mut found = false;
        let mut set: HashSet<i64> = HashSet::new();
        for m in &nums {
            if set.contains(&(n - m)) {
                found = true;
                break;
            }

            set.insert(*m);
        }

        if !found {
            // Success!
            println!("GOT_BAD_NUMBER = {}", n);
            return;
        }

        nums.pop_front();
        nums.push_back(n);
    }
}

fn part2(bad_num: i64) {
    // Read the entire list of numbers.
    let mut all_nums: Vec<i64> = Vec::new();
    for line in stdin().lock().lines() {
        let t = line.unwrap();
        let n: i64 = t.parse().unwrap();
        all_nums.push(n);
    }

    for i in 0..all_nums.len() {
        let mut sum = 0;
        let mut v = Vec::new();

        if all_nums[i] == bad_num {
            // Region size must be > 1, and 0 isn't in our list.
            continue;
        }

        for j in i..all_nums.len() {
            let n = all_nums[j];
            sum += n;
            v.push(n);

            if sum == bad_num {
                // Success!
                let min = v.iter().min().unwrap();
                let max = v.iter().max().unwrap();
                println!("FOUND_REGION = ({} + {}) = {}", min, max, min + max);
                return;
            }

            if sum > bad_num {
                // We've gone over - onto the next index for a possible region.
                break;
            }
        }
    }
}
