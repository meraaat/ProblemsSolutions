const testCases = [{ n: 1 }, { n: 2 }, { n: 3 }, { n: 4 }];

(function () {
  testCases.forEach((testcase, i) => {
    console.log(`n= ${testcase.n}`);
    console.log(`testcase ${i}:  ${toh(testcase.n, 1, 3, 2)}`);
  });
})();

// Approach : T O(n^2) | S O(n)
function toh(N, from, to, aux) {
  let count = 0;

  function helper(N, from, to, aux) {
    if (N === 1) {
      count += 1;
      console.log("move disk " + N + " from rod " + from + " to rod " + to);
      return;
    }
    helper(N - 1, from, aux, to);
    count += 1;
    console.log("move disk " + N + " from rod " + from + " to rod " + to);
    helper(N - 1, aux, to, from);
  }

  helper(N, from, to, aux);
  return count;
}
