# Problem: Task Scheduler

## Problem Statement

You are given an array of CPU tasks, `tasks`, each represented by letters A to Z, and a cooling time, `n`. Each cycle or interval allows the completion of one task. Tasks can be completed in any order, but there's a constraint: identical tasks must be separated by at least `n` intervals due to cooling time.

Return the minimum number of intervals required to complete all tasks.

## Constraints

* `1 <= tasks.length <= 10^4`
* `0 <= n <= 100`
* `tasks[i]` is an uppercase English letter.

## Function Signatures

### JavaScript

```javascript
// JavaScript
function leastInterval(tasks, n)
```

### Go

```go
// Go
func leastInterval(tasks []byte, n int) int
```

## Example

```
Example 1:
Input: tasks = ['A','A','A','B','B','B'], n = 2
Output: 8
Explanation: A -> B -> idle -> A -> B -> idle -> A -> B.
There are three of each task. After each task A, you must wait 2 cycles to perform another task A. Same with task B.

Example 2:
Input: tasks = ['A','A','A','B','B','B'], n = 0
Output: 6
Explanation: On a n = 0 case, no cooling time is required, so all tasks can be done consecutively.

Example 3:
Input: tasks = ['A','A','A','A','A','A','B','C','D','E','F','G'], n = 2
Output: 16
Explanation: One possible schedule is: A -> B -> C -> A -> D -> E -> A -> F -> G -> A -> idle -> idle -> A -> idle -> idle -> A.

Example 4:
Input: tasks = ['A','A','A','B','B','B','C','C'], n = 2
Output: 8
Explanation: One possible schedule is: A -> B -> C -> A -> B -> C -> A -> B.

Example 5:
Input: tasks = ['A','A','B','B'], n = 3
Output: 6
Explanation: A -> B -> idle -> idle -> A -> B

```

## Approaches

This problem can be solved using a greedy approach based on frequency counting.

### Greedy Approach (Frequency Counting)

1. **Count the frequency** of each task.

2. **Find the maximum frequency** (`maxFreq`).

3. **Count the number of tasks with maximum frequency** (`maxFreqCount`).

4. The minimum number of intervals is given by:

    `max(tasks.length, (maxFreq - 1) * (n + 1) + maxFreqCount)`

    * `(maxFreq - 1) * (n + 1)` represents the intervals needed to schedule all but the last set of the most frequent tasks, including the cooling periods.
    * `maxFreqCount` represents the number of tasks that have the maximum frequency. We add this because they can be scheduled together at the end without further cooling.
    * `tasks.length` is needed in case the schedule can be made without any idle times (like when n=0 or n is small compared to the diversity of tasks).

## Complexity Analysis

| Approach         | Time Complexity | Space Complexity |
|------------------|-----------------|-----------------|
| Greedy (Counting)| O(N)            | O(1)            |

* **Greedy (Counting):** O(N) time because we iterate through the tasks once to count frequencies. The space complexity is O(1) because we use a fixed-size array (26 for uppercase English letters) to store the frequencies.
