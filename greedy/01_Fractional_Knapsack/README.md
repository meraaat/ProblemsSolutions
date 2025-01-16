# Problem: Fractional Knapsack

## Problem Statement

Determine how to optimally fill a knapsack with a capacity of `W` kilograms using a list of `N` items, where each item is represented by a pair `[profit, weight]`. In the Fractional Knapsack problem, you can take fractions of items to maximize the total profit in the knapsack. (`N` will be greater than or equal to 1)

## Constraints

*   `1 <= N <= 10^5` (Number of items)
*   `1 <= W <= 10^5` (Knapsack capacity)
*   `1 <= profit_i, weight_i <= 10^4` (Profit and weight of each item)

## Function Signatures

### JavaScript

```javascript
// JavaScript
function fractionalKnapsack(W, items)
```

### Go

```go
// Go
func fractionalKnapsack(W int, items [][]int) float64
```

## Example

```
Example 1:
items = [[70, 10], [90, 20], [150, 30]]
W = 25
Output: 145
Explanation: Take the entire first item (profit 70, weight 10).
Take 15 units of the second item (90 * (15/20) = 67.5). Total profit is 70 + 67.5 = 137.5.
Take 0 units of the last item. Total profit is 137.5.

Example 2:
items = [[70, 10], [90, 20], [150, 30]]
W = 45
Output: 242.5
Explanation: Take the entire first item (profit 70, weight 10).
Take the entire second item (profit 90, weight 20).
Take 15 units of the third item (150 * (15/30) = 75). Total profit is 70 + 90 + 75 = 235.

Example 3:
items = [[10,2],[5,3],[15,5],[7,7],[6,1],[18,6]]
W=10
Output: 39.666666666666664
Explanation: Take the entire 6 unit weight item with profit 18, then take the entire 1 unit weight item with profit 6, then take 3 units of the 5 unit weight item with profit 15 that is 9 profit, then take the entire 2 unit weight item with profit 10. 18+6+9+10= 43.

Example 4:
items = []
W = 10
Output: 0
Explanation: No items to take.

Example 5:
items = [[10,5]]
W = 0
Output: 0
Explanation: No capacity to take anything.

```
## Approaches

The Fractional Knapsack problem is optimally solved using a greedy approach.

### Greedy Approach

1.  **Calculate the profit-to-weight ratio** for each item.
2.  **Sort the items** in descending order of their profit-to-weight ratio.
3.  Iterate through the sorted items:
    *   If the current item's weight is less than or equal to the remaining knapsack capacity, take the whole item and update the remaining capacity and total profit.
    *   Otherwise, take a fraction of the item that fills the remaining capacity. Calculate the fractional profit and update the total profit.
4.  Return the total profit.

## Complexity Analysis

| Approach | Time Complexity | Space Complexity |
|----------|-----------------|-----------------|
| Greedy   | O(N log N)      | O(1) or O(N) if sorting is considered |

*   **Greedy:** O(N log N) due to sorting. The space complexity is O(1) if sorting is done in place or O(N) if a new array is created while sorting.
