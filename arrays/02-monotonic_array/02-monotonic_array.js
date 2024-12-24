const testCases = [
  [],
  [1],
  [3, 2, 1],
  [1, 2, 1],
  [1, 3, 2],
  [1, 2, 4, 2],
  [1, 2, 3, 5, 13],
  [1, 1, 2, 3, 3, 4, 4],
  [6, 5, 4, 4, 3, 2, 1],
];

(function () {
  testCases.forEach((testcase, i) => {
    console.log(`testcase ${i}:  ${isMonotonic(testcase)}`);
  });
})();

// isMonotonic - T: O(n) | S: O(1)
function isMonotonic(array) {
  let isIncreasing = true;
  let isDecreasing = true;

  for (let i = 1; i < array.length; i++) {
    if (array[i] < array[i - 1]) isIncreasing = false;
    if (array[i] > array[i - 1]) isDecreasing = false;
  }

  return isIncreasing || isDecreasing;
}
