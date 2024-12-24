const testCases = [
  [],
  [0, 5],
  [1, 3, 5],
  [-4, -1, 0, 3, 10],
  [-7, -3, 2, 3, 11],
];

(function () {
  testCases.forEach((testcase, i) => {
    console.log(`testcase ${i}:`);
    console.log(`Sorted Squared Array: ${SortedSquaredArray(testcase)}`);
    console.log(`Two Pointers: ${SquareAndSortOptimized(testcase)}`);
  });
})();

// Approach 1: Sort and Square - T: O(n log n) | S:(1)
function SortedSquaredArray(array) {
  return array
    .map((num) => {
      return num * num;
    })
    .sort((a, b) => {
      return a - b;
    });
}

// Approach2: Two Pointers - T: O(n)
function SquareAndSortOptimized(array) {
  const result = new Array(array.length).fill(0);
  let left = 0;
  let right = array.length - 1;

  for (let i = array.length - 1; i >= 0; i--) {
    let leftSquare = array[left] * array[left];
    let rightSquare = array[right] * array[right];

    if (leftSquare > rightSquare) {
      result[i] = leftSquare;
      left++;
    } else {
      result[i] = rightSquare;
      right--;
    }
  }

  return result;
}
