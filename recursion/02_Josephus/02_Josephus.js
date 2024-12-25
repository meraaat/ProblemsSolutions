const testCases = [
  { n: 1, k: 1 },
  { n: 2, k: 1 },
  { n: 2, k: 2 },
  { n: 3, k: 1 },
  { n: 3, k: 2 },
  { n: 3, k: 3 },
  { n: 3, k: 4 },
  { n: 4, k: 5 },
  { n: 4, k: 8 },
];

(function () {
  testCases.forEach((testcase, i) => {
    console.log(`n= ${testcase.n}, k= ${testcase.k}`);
    console.log(`testcase ${i}:  ${findTheWinner(testcase.n, testcase.k)}`);
    console.log(`testcase ${i}:  ${findTheWinner2(testcase.n, testcase.k)}`);
    console.log(`testcase ${i}:  ${findTheWinner3(testcase.n, testcase.k)}`);
  });
})();

// Approach 0 : T O(n^2) | S O(n)
function findTheWinner(n, k) {
  let arr = Array.from({ length: n }, (_, i) => i + 1);

  function helper(arr, startIndex) {
    // Base case
    if (arr.length === 1) return arr[0];

    // Recursion case
    let indexToRemove = (startIndex + k - 1) % arr.length;
    arr.splice(indexToRemove, 1);
    helper(arr, indexToRemove);
  }

  return helper(arr, 0);
}

// Approach 1: T O(n) | S O(n)
function findTheWinner2(n, k) {
  function josephus(n) {
    // Base case
    if (n === 1) return 0;

    // Recursive case
    return josephus(n - 1 + k) % n;
  }
  return josephus(n) + 1;
}

// Approach 2: T O(n) | S O(1) / iterative
function findTheWinner3(n, k) {
  let survivor = 0;

  for (let i = 2; i <= n; i++) {
    survivor = (survivor + k) % i;
  }
  return survivor + 1;
}
