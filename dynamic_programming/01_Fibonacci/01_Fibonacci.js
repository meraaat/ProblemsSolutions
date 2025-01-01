const testCases = [
  { n: 0, expect: 0 },
  { n: 1, expect: 1 },
  { n: 2, expect: 1 },
  { n: 3, expect: 2 },
  { n: 4, expect: 3 },
  { n: 5, expect: 5 },
  { n: 10, expect: 55 },
  { n: 20, expect: 6765 },
  { n: 30, expect: 832040 },
];

const fibFunctions = {
  Recursive: fibRecursive,
  Memoization: fibMemoization,
  Tabulation: fibTabulation,
  SpaceOptimized: fibSpaceOptimized,
};

function runTests() {
  // Encapsulate tests in a function
  for (const name in fibFunctions) {
    console.log(`Testing: ${name}`);
    testCases.forEach((testCase, i) => {
      const result = fibFunctions[name](testCase.n);
      console.log(
        `Test Case ${i + 1}: fib(${testCase.n}) = ${result}, Expected: ${
          testCase.expect
        }`
      );
      if (result === testCase.expect) {
        console.log("PASS");
      } else {
        console.log("FAIL");
      }
    });
    console.log();
  }
}

function main() {
  const n = 10;
  console.log(`fibRecursive(${n}) = ${fibRecursive(n)}`);
  console.log(`fibMemoization(${n}) = ${fibMemoization(n)}`);
  console.log(`fibTabulation(${n}) = ${fibTabulation(n)}`);
  console.log(`fibSpaceOptimized(${n}) = ${fibSpaceOptimized(n)}`);
  runTests(); // Call the tests
}

main(); // Call the main function

// Recursive approach
function fibRecursive(n) {
  if (n <= 1) return n;
  return fibRecursive(n - 1) + fibRecursive(n - 2);
}

// Memoization (Top-Down DP)
function fibMemo(n) {
  const memo = {};
  function fib(n) {
    if (n <= 1) return n;
    if (memo[n] !== undefined) return memo[n];
    return (memo[n] = fib(n - 1) + fib(n - 2));
  }
  return fib(n);
}

// Tabulation (Bottom-Up DP)
function fibTab(n) {
  if (n <= 1) return n;
  const dp = [0, 1];
  for (let i = 2; i <= n; i++) {
    dp[i] = dp[i - 1] + dp[i - 2];
  }
  return dp[n];
}

// Space Optimized
function fibSpace(n) {
  if (n <= 1) return n;
  let a = 0,
    b = 1;
  for (let i = 2; i <= n; i++) {
    [a, b] = [b, a + b];
  }
  return b;
}
