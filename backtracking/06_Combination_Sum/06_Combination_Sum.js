const testCases = [
  {
    candidates: [2, 3, 6, 7],
    target: 7,
    expected: [[2, 2, 3], [7]],
  },
  {
    candidates: [2, 3, 5],
    target: 8,
    expected: [
      [2, 2, 2, 2],
      [2, 3, 3],
      [3, 5],
    ],
  },
  {
    candidates: [2],
    target: 1,
    expected: [],
  },
  {
    candidates: [1],
    target: 1,
    expected: [[1]],
  },
  {
    candidates: [1, 2],
    target: 4,
    expected: [
      [1, 1, 1, 1],
      [1, 1, 2],
      [2, 2],
    ],
  },
];

(function () {
  testCases.forEach((testCase, i) => {
    console.log(`Test Case ${i + 1}:`);
    console.log(
      `Input: candidates = ${JSON.stringify(testCase.candidates)}, target = ${
        testCase.target
      }`
    );
    const result = combinationSum(testCase.candidates, testCase.target);

    // Sort inner arrays for comparison
    const sortedResult = result
      .map((inner) => [...inner].sort((a, b) => a - b))
      .sort((a, b) => JSON.stringify(a).localeCompare(JSON.stringify(b)));
    const sortedExpected = testCase.expected
      .map((inner) => [...inner].sort((a, b) => a - b))
      .sort((a, b) => JSON.stringify(a).localeCompare(JSON.stringify(b)));

    console.log(`Result: ${JSON.stringify(sortedResult)}`);
    console.log(`Expected: ${JSON.stringify(sortedExpected)}`);

    if (JSON.stringify(sortedResult) === JSON.stringify(sortedExpected)) {
      console.log(`Test case ${i + 1} passed`);
    } else {
      console.log(`Test case ${i + 1} failed`);
    }
    console.log();
  });
})();

function combinationSum(candidates, target) {
  let result = [];

  function combinationSumRecursive(start, current, currentSum) {
    // Base cases
    if (currentSum > target) return;

    if (currentSum == target) {
      result.push([...current]);
      return;
    }

    // Recursive case
    for (let i = start; i < candidates.length; i++) {
      current.push[candidates[i]];
      combinationSumRecursive(i, current, currentSum + candidates[i]);
      current.pop();
    }
  }

  combinationSumRecursive(0, [], 0);
  return result;
}
