// Tabulation (Bottom-Up DP)
function unboundedKnapsack(W, wt, val, n) {
  const dp = new Array(W + 1).fill(0);

  for (let w = 0; w <= W; w++) {
    for (let i = 0; i < n; i++) {
      if (wt[i] <= w) {
        dp[w] = Math.max(dp[w], val[i] + dp[w - wt[i]]);
      }
    }
  }
  return dp[W];
}

const testCases = [
  { W: 6, wt: [2, 2], val: [5, 10], n: 2, expect: 30 },
  { W: 8, wt: [1, 3, 4], val: [1, 4, 5], n: 3, expect: 11 },
  { W: 4, wt: [4, 5, 1], val: [1, 2, 3], n: 3, expect: 12 },
  { W: 10, wt: [2, 5, 7], val: [1, 6, 18], n: 3, expect: 36 },
  { W: 5, wt: [1, 2, 3], val: [10, 15, 20], n: 3, expect: 50 },
  { W: 10, wt: [1, 2, 3, 4, 5], val: [1, 2, 3, 4, 5], n: 5, expect: 10 },
];

testCases.forEach((tc, i) => {
  console.log(
    `Test Case ${i + 1}: W=${tc.W}, wt=${tc.wt}, val=${tc.val}, n=${tc.n}`
  );
  const res = unboundedKnapsack(tc.W, tc.wt, tc.val, tc.n);
  console.log(`Got: ${res}`);
  console.log(`Expected: ${tc.expect}`);
  if (res !== tc.expect) {
    console.log("FAIL");
  } else {
    console.log("PASS");
  }
  console.log();
});
