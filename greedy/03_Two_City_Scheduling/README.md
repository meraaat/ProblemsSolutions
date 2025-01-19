# Problem: Two City Scheduling

## Problem Statement

A company is planning to interview `2n` people. You are given the array `costs` where `costs[i] = [aCosti, bCosti]`. The cost of flying the `i`th person to city A is `aCosti`, and the cost of flying the `i`th person to city B is `bCosti`.

Return the minimum cost to fly every person to a city such that exactly `n` people arrive in each city. (`costs` has an even number of elements, and the number of elements is greater than or equal to 2).

## Constraints

*   `2 <= costs.length <= 100`
*   `costs.length` is even.
*   `10 <= aCosti, bCosti <= 1000`

## Function Signatures

### JavaScript

```javascript
// JavaScript
function twoCitySchedCost(costs)
```

### Go

```go
// Go
func twoCitySchedCost(costs [][]int) int
```

## Example

```
Example 1:
Input: costs = [[10,20],[30,200],[400,50],[30,20]]
Output: 110
Explanation:
The first person goes to city A for a cost of 10.
The second person goes to city A for a cost of 30.
The third person goes to city B for a cost of 50.
The fourth person goes to city A for a cost of 30.
The total minimum cost is 10 + 30 + 50 + 30 = 110.

Example 2:
Input: costs = [[5,20],[30,100],[400,30],[50,10]]
Output: 75
Explanation:
The first person goes to city A for a cost of 5.
The second person goes to city A for a cost of 30.
The third person goes to city B for a cost of 30.
The fourth person goes to city B for a cost of 10.
The total minimum cost is 5 + 30 + 30 + 10 = 75.

Example 3:
Input: costs = [[20,700], [400,50], [900,600], [200,150], [800,100],[500,450]]
Output: 1470

Example 4:
Input: costs = [[259,770],[448,54],[926,667],[184,139],[840,118],[577,469]]
Output: 1859
```
## Approaches

This problem can be efficiently solved using sorting and a greedy approach.

### Greedy Approach (with Sorting)

1.  **Calculate the difference** between the cost of going to city A and the cost of going to city B for each person: `diff[i] = aCosti - bCosti`.

2.  **Sort the `costs` array** based on these differences in ascending order. This means that people with the largest negative differences (meaning it's much cheaper to send them to B) will be at the beginning of the sorted array, and people with the largest positive differences (meaning it's much cheaper to send them to A) will be at the end.

3.  **Send the first `n` people** in the sorted array to city A.

4.  **Send the remaining `n` people** to city B.

5.  **Calculate the total cost.**

## Complexity Analysis

| Approach         | Time Complexity | Space Complexity |
|------------------|-----------------|-----------------|
| Greedy (Sorting) | O(n log n)      | O(n) or O(1) if sorting is in place |

*   **Greedy (Sorting):** O(n log n) due to the sorting step. The space complexity is O(n) if the sorting creates a copy of the input array, and O(1) if the sorting is done in-place.
