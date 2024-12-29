const testCases = [
  {
    input: [1, 2, 2],
    expected: [[], [1], [2], [1, 2], [2, 2], [1, 2, 2]],
  },
  {
    input: [0],
    expected: [[], [0]],
  },
  {
    input: [],
    expected: [[]],
  },
];

(function () {
  testCases.forEach((testCase, i) => {
    console.log(`Test Case ${i + 1}:`);
    console.log(`Input: ${JSON.stringify(testCase.input)}`);
    const result = subsetsWithDup(testCase.input);
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

function subsetsWithDup(nums) {
  let result = [];
  nums.sort((a, b) => a - b);

  function helper(i, subsets) {
    // Base case
    if (i === nums.length) {
      result.push([...subsets]);
      return;
    }

    // Recursive case
    // Include
    subsets.push(nums[i]);
    helper(i + 1, subsets);
    subsets.pop(); // Backtracking

    while (i < nums.length - 1 && nums[i] == nums[i + 1]) i++;
    helper(i + 1, subsets);
  }

  helper(0, []);
  return result;
}
