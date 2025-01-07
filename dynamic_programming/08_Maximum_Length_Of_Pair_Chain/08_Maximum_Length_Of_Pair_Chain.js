// Recursive
function findLongestChainRecursive(pairs) {
  pairs.sort((a, b) => a[0] - b[0]); // Sort by left values
  const n = pairs.length;

  function findChain(index) {
    if (index === n) {
      return 0;
    }

    let maxLength = 1;
    for (let i = index + 1; i < n; i++) {
      if (pairs[index][1] < pairs[i][0]) {
        maxLength = Math.max(maxLength, 1 + findChain(i));
      }
    }
    return maxLength;
  }

  let maxLength = 0;
  for (let i = 0; i < n; i++) {
    maxLength = Math.max(maxLength, findChain(i));
  }
  return maxLength;
}

// Memoization
function findLongestChainMemo(pairs) {
  pairs.sort((a, b) => a[0] - b[0]);
  const n = pairs.length;
  const memo = new Array(n).fill(-1);

  function findChain(index) {
    if (index === n) {
      return 0;
    }
    if (memo[index] !== -1) {
      return memo[index];
    }

    let maxLength = 1;
    for (let i = index + 1; i < n; i++) {
      if (pairs[index][1] < pairs[i][0]) {
        maxLength = Math.max(maxLength, 1 + findChain(i));
      }
    }
    memo[index] = maxLength;
    return maxLength;
  }
  let maxLength = 0;
  for (let i = 0; i < n; i++) {
    maxLength = Math.max(maxLength, findChain(i));
  }
  return maxLength;
}

// Tabulation
function findLongestChainTab(pairs) {
  pairs.sort((a, b) => a[0] - b[0]);
  const n = pairs.length;
  if (n === 0) {
    return 0;
  }
  const dp = new Array(n).fill(1);

  for (let i = 1; i < n; i++) {
    for (let j = 0; j < i; j++) {
      if (pairs[j][1] < pairs[i][0]) {
        dp[i] = Math.max(dp[i], dp[j] + 1);
      }
    }
  }

  let maxLength = 0;
  for (const len of dp) {
    maxLength = Math.max(maxLength, len);
  }
  return maxLength;
}

// Greedy
function findLongestChainGreedy(pairs) {
  pairs.sort((a, b) => a[1] - b[1]); // Sort by right values
  let currentEnd = -Infinity;
  let chainLength = 0;

  for (const pair of pairs) {
    if (pair[0] > currentEnd) {
      currentEnd = pair[1];
      chainLength++;
    }
  }
  return chainLength;
}

const testCases = [
  {
    pairs: [
      [1, 2],
      [2, 3],
      [3, 4],
    ],
    expect: 2,
  },
  {
    pairs: [
      [1, 2],
      [7, 8],
      [4, 5],
    ],
    expect: 3,
  },
  {
    pairs: [
      [1, 2],
      [2, 9],
      [4, 5],
    ],
    expect: 2,
  },
  { pairs: [[1, 2]], expect: 1 },
  { pairs: [], expect: 0 },
  {
    pairs: [
      [1, 2],
      [1, 3],
    ],
    expect: 1,
  },
  {
    pairs: [
      [1, 2],
      [3, 4],
      [2, 8],
    ],
    expect: 2,
  },
  {
    pairs: [
      [1, 2],
      [2, 3],
      [4, 5],
      [1, 9],
    ],
    expect: 3,
  },
  {
    pairs: [
      [-10, -8],
      [8, 9],
      [-5, 0],
      [6, 10],
      [-6, -4],
      [1, 7],
      [9, 10],
      [-4, 7],
    ],
    expect: 4,
  },
];

const findLongestChainFuncs = {
  Recursive: findLongestChainRecursive,
  Memoization: findLongestChainMemo,
  Tabulation: findLongestChainTab,
  Greedy: findLongestChainGreedy,
};

for (const name in findLongestChainFuncs) {
  console.log(`Testing ${name}:`);
  testCases.forEach((tc, i) => {
    console.log(`Test Case ${i + 1}: pairs=${JSON.stringify(tc.pairs)}`);
    const res = findLongestChainFuncs[name](tc.pairs);
    console.log(`Got: ${res}`);
    console.log(`Expected: ${tc.expect}`);
    if (res !== tc.expect) {
      console.log("FAIL");
    } else {
      console.log("PASS");
    }
    console.log();
  });
  console.log("--------------------");
}
