const testCases = [
  {
    n: 4,
    k: 2,
    expected: [
      [1, 2],
      [1, 3],
      [1, 4],
      [2, 3],
      [2, 4],
      [3, 4],
    ],
  },
  {
    n: 1,
    k: 1,
    expected: [[1]],
  },
  {
    n: 5,
    k: 3,
    expected: [
      [1, 2, 3],
      [1, 2, 4],
      [1, 2, 5],
      [1, 3, 4],
      [1, 3, 5],
      [1, 4, 5],
      [2, 3, 4],
      [2, 3, 5],
      [2, 4, 5],
      [3, 4, 5],
    ],
  },
  {
    n: 2,
    k: 1,
    expected: [[1], [2]],
  },
  {
    n: 2,
    k: 2,
    expected: [[1, 2]],
  },
];

(function () {
  testCases.forEach((testCase, i) => {
    console.log(`Test Case ${i + 1}:`);
    console.log(`Input: n = ${testCase.n}, k = ${testCase.k}`);
    const result = combine(testCase.n, testCase.k);

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

function combine(n, k) {
  let result = [];
  function helper(start, current) {
    // Base case
    if (current.length === k) {
      result.push([...current]);
      return;
    }

    // Recursive case
    let need = k - current.length;
    for (let i = start; i <= n - need + 1; i++) {
      current.push(i);
      helper(i + 1, current);
      current.pop();
    }
  }

  helper(1, []);
  return result;
}
