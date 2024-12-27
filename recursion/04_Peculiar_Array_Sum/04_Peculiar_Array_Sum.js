const testCases = [
  { input: [2, 3, [4, 1, 2]] },
  { input: [1, 2, [7, [3, 4], 2]] },
  { input: [1] },
  { input: [[1]] },
  { input: [[1, 2], 3] },
  { input: [[[1]]] },
  { input: [1, [2, [3]]] },
];

(function () {
  testCases.forEach((testCase, i) => {
    console.log(`Test Case ${i + 1}:`);
    console.log(`Input: ${JSON.stringify(testCase.input)}`); // Use JSON.stringify for better output
    const result = peculiarArraySum(testCase.input);
    console.log(`Result: ${result}`);
    console.log(); // Add an empty line for readability
  });
})();

// Approach 0 : T O(n) | S O(d)
function peculiarArraySum(arr, power = 1) {
  let sum = 0;

  arr.forEach((item) => {
    if (Array.isArray(item)) {
      sum += peculiarArraySum(item, power + 1);
    } else {
      sum += item;
    }
  });

  return Math.pow(sum, power);
}
