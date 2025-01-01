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
