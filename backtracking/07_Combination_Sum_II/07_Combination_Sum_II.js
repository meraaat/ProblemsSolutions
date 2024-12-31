const testCases = [
  {
    candidates: [3, 5, 2, 1, 3],
    target: 7,
    expected: [
      [1, 3, 3],
      [2, 5],
    ],
  },
  {
    candidates: [10, 1, 2, 7, 6, 1, 5],
    target: 8,
    expected: [
      [1, 1, 6],
      [1, 2, 5],
      [1, 7],
      [2, 6],
    ],
  },
  {
    candidates: [2, 5, 2, 1, 2],
    target: 5,
    expected: [[1, 2, 2], [5]],
  },
  {
    candidates: [1, 1, 1, 1, 1],
    target: 3,
    expected: [[1, 1, 1]],
  },
  {
    candidates: [1, 2],
    target: 4,
    expected: [],
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
    const result = combinationSum2(testCase.candidates, testCase.target);

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

function combinationSum2(candidates, target) {
  let result = [];
  candidates.sort((a, b) => a - b);

  function backtracks(index, currentSum, current) {
    // Base case
    if (currentSum === target) {
      result.push([...current]);
      return;
    }
    if (currentSum > target || index > candidates.length - 1) {
      return;
    }

    // Recursive case
    let seen = {};
    for (let i = index; i < candidates.length; i++) {
      if (!seen[candidates[i]]) {
        seen[candidates[i]] = true;
        current.push(candidates[i]);
        backtracks(i + 1, currentSum + candidates[i], current);
        current.pop();
      }
    }
  }

  backtracks(0, 0, []);
  return result;
}
