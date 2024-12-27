const testCases = [
  {
    input: [1, 2, 3],
    expected: [
      [1, 2, 3],
      [1, 3, 2],
      [2, 1, 3],
      [2, 3, 1],
      [3, 2, 1],
      [3, 1, 2],
    ],
  },
  {
    input: [0, 1],
    expected: [
      [0, 1],
      [1, 0],
    ],
  },
  { input: [1], expected: [[1]] },
  { input: [], expected: [] },
  {
    input: [1, 2],
    expected: [
      [1, 2],
      [2, 1],
    ],
  },
];

(function () {
  testCases.forEach((testCase, i) => {
    console.log(`Test Case ${i + 1}:`);
    console.log(`Input: ${JSON.stringify(testCase.input)}`);
    const result = permute(testCase.input);
    console.log(`Result: ${JSON.stringify(result)}`); // Use JSON.stringify for nested arrays
    console.log(`Expected: ${JSON.stringify(testCase.expected)}`);
    if (JSON.stringify(result) === JSON.stringify(testCase.expected)) {
      console.log(`Test case ${i + 1} passed`);
    } else {
      console.log(`Test case ${i + 1} failed`);
    }
    console.log();
  });
})();

// Approach : T O(n!*n) | S O(n)
function permute(nums) {
  let result = [];
  let n = nums.length;

  function helper(i) {
    // Base case
    if (i === n - 1) {
      result.push([...nums]);
      return;
    }

    // Recursive case
    for (let j = i; j < n; j++) {
      [nums[i], nums[j]] = [nums[j], nums[i]];
      helper(i + 1);
      // Backtracking
      [nums[i], nums[j]] = [nums[j], nums[i]]; // Reverting changes
    }
  }
  helper(0);
  return result;
}
