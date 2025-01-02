// Recursive
function minCostClimbingStairsRecursive(cost) {
  const n = cost.length;
  function minCost(i) {
    if (i < 0) return 0;
    if (i <= 1) return cost[i];
    return cost[i] + Math.min(minCost(i - 1), minCost(i - 2));
  }
  return Math.min(minCost(n - 1), minCost(n - 2));
}

// Memoization
function minCostClimbingStairsMemo(cost) {
  const n = cost.length;
  const memo = {};
  function minCost(i) {
    if (i < 0) return 0;
    if (i <= 1) return cost[i];
    if (memo[i] !== undefined) return memo[i];
    return (memo[i] = cost[i] + Math.min(minCost(i - 1), minCost(i - 2)));
  }
  return Math.min(minCost(n - 1), minCost(n - 2));
}

// Tabulation
function minCostClimbingStairsTab(cost) {
  const n = cost.length;
  const dp = new Array(n);
  dp[0] = cost[0];
  dp[1] = cost[1];
  for (let i = 2; i < n; i++) {
    dp[i] = cost[i] + Math.min(dp[i - 1], dp[i - 2]);
  }
  return Math.min(dp[n - 1], dp[n - 2]);
}

// Space Optimized
function minCostClimbingStairsSpace(cost) {
  const n = cost.length;
  if (n <= 1) return 0;
  let one = cost[0];
  let two = cost[1];
  for (let i = 2; i < n; i++) {
    const temp = cost[i] + Math.min(one, two);
    one = two;
    two = temp;
  }
  return Math.min(one, two);
}

const testCases = [
  { cost: [10, 15, 20], expect: 15 },
  { cost: [1, 100, 1, 1, 1, 100, 1, 1, 100, 1], expect: 6 },
  { cost: [1, 2, 3], expect: 2 },
  { cost: [0, 2, 2, 1], expect: 2 },
  { cost: [0, 1, 2, 3, 4, 5], expect: 6 },
];

const minCostFuncs = {
  Recursive: minCostClimbingStairsRecursive,
  Memoization: minCostClimbingStairsMemo,
  Tabulation: minCostClimbingStairsTab,
  SpaceOptimized: minCostClimbingStairsSpace,
};

for (const name in minCostFuncs) {
  console.log(`Testing: ${name}`);
  testCases.forEach((tc, i) => {
    const res = minCostFuncs[name](tc.cost);
    console.log(`Test ${i + 1}: minCostClimbingStairs(${tc.cost})`);
    if (res !== tc.expect) {
      console.log(`FAIL: Got ${res}, Expected ${tc.expect}`);
    } else {
      console.log("PASS");
    }
  });
  console.log();
}
