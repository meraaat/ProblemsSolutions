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
    console.log(`testcase ${i}:  ${kthGrammar(testcase.n, testcase.k)}`);
  });
})();

// Recursion Approach : T O(n) | S O(n)
function kthGrammar(n, k) {
  // Base case
  if (n === 1) return 0;

  // Recursion Base
  const length = Math.pow(2, n - 1);
  const mid = length / 2;

  if (k <= mid) return kthGrammar(n - 1, k);
  else return 1 - kthGrammar(n - 1, k - mid);
}
