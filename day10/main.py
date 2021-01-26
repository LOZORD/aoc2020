#!/usr/bin/python3

# NOTE: This only solves Part 2 of Day 10. Part 1 is solved in main.rs.

from sys import stdin

# Memoization mapping of (nums list string representation) -> amount.
memo = {}


def main():
    nums = [0]
    for line in stdin.readlines():
        nums.append(int(line.strip()))
    nums.sort()
    print('ANSWER: %d' % do_count(nums))


def do_count(nums) -> int:
    mk = str(nums)
    if mk in memo:
        return memo[mk]

    # We've reached the end of our chain, so we know we have at least one valid chain.
    if len(nums) <= 1:
        memo[mk] = 1
        return 1

    cur = nums[0]
    s = 0
    # Start looking beyond the current joltage to avoid infinite looping!
    for idx in range(1, len(nums)):
        val = nums[idx]
        # Since our list is sorted, break if we can no longer connect plugs within
        # the joltage range.
        if val - cur > 3:
            break
        s += do_count(nums[idx:])

    memo[mk] = s
    return s


if __name__ == "__main__":
    main()
