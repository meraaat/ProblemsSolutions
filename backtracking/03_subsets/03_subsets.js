const testCases = [
  {
    input: [1, 2, 3],
    expected: [[], [3], [2], [2, 3], [1], [1, 3], [1, 2], [1, 2, 3]],
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
    const result = subsets(testCase.input);
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

function subsets(nums) {
  let result = [];

  function helper(nums, i, subsets) {
    if (i === nums.length) {
      result.push(subsets.slice());
      return;
    }
    // Don't add
    helper(nums, i + 1, subsets);

    //Add
    subsets.push(nums[i]);
    helper(nums, i + 1, subsets);

    subsets.pop();
  }

  helper(nums, 0, []);
  return result;
}
