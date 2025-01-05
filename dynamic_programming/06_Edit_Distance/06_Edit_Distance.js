// Recursive
function minDistanceRecursive(word1, word2) {
  function minDistance(i, j) {
    if (i < 0) {
      return j + 1;
    }
    if (j < 0) {
      return i + 1;
    }
    if (word1[i] === word2[j]) {
      return minDistance(i - 1, j - 1);
    }
    return (
      1 +
      Math.min(
        minDistance(i - 1, j),
        minDistance(i, j - 1),
        minDistance(i - 1, j - 1)
      )
    );
  }
  return minDistance(word1.length - 1, word2.length - 1);
}

// Memoization
function minDistanceMemo(word1, word2) {
  const memo = Array(word1.length + 1)
    .fill(null)
    .map(() => Array(word2.length + 1).fill(-1));

  function minDistance(i, j) {
    if (i < 0) {
      return j + 1;
    }
    if (j < 0) {
      return i + 1;
    }
    if (memo[i + 1][j + 1] !== -1) {
      return memo[i + 1][j + 1];
    }
    if (word1[i] === word2[j]) {
      memo[i + 1][j + 1] = minDistance(i - 1, j - 1);
    } else {
      memo[i + 1][j + 1] =
        1 +
        Math.min(
          minDistance(i - 1, j),
          minDistance(i, j - 1),
          minDistance(i - 1, j - 1)
        );
    }
    return memo[i + 1][j + 1];
  }
  return minDistance(word1.length - 1, word2.length - 1);
}

// Tabulation
function minDistanceTab(word1, word2) {
  const m = word1.length;
  const n = word2.length;
  const dp = Array(m + 1)
    .fill(null)
    .map(() => Array(n + 1).fill(0));

  for (let i = 0; i <= m; i++) {
    dp[i][0] = i;
  }
  for (let j = 0; j <= n; j++) {
    dp[0][j] = j;
  }

  for (let i = 1; i <= m; i++) {
    for (let j = 1; j <= n; j++) {
      if (word1[i - 1] === word2[j - 1]) {
        dp[i][j] = dp[i - 1][j - 1];
      } else {
        dp[i][j] = 1 + Math.min(dp[i - 1][j], dp[i][j - 1], dp[i - 1][j - 1]);
      }
    }
  }
  return dp[m][n];
}

// Space Optimized
function minDistanceSpace(word1, word2) {
  const m = word1.length;
  const n = word2.length;

  if (m < n) {
    [word1, word2] = [word2, word1];
    [m, n] = [n, m];
  }

  let prev = Array(n + 1).fill(0);
  let curr = Array(n + 1).fill(0);

  for (let j = 0; j <= n; j++) {
    prev[j] = j;
  }

  for (let i = 1; i <= m; i++) {
    curr[0] = i;
    for (let j = 1; j <= n; j++) {
      if (word1[i - 1] === word2[j - 1]) {
        curr[j] = prev[j - 1];
      } else {
        curr[j] = 1 + Math.min(prev[j], curr[j - 1], prev[j - 1]);
      }
    }
    [prev, curr] = [curr, prev];
  }
  return prev[n];
}

const testCases = [
  { word1: "horse", word2: "ros", expect: 3 },
  { word1: "intention", word2: "execution", expect: 5 },
  { word1: "", word2: "", expect: 0 },
  { word1: "a", word2: "", expect: 1 },
  { word1: "", word2: "b", expect: 1 },
  { word1: "kitten", word2: "sitting", expect: 3 },
  { word1: "apple", word2: "aple", expect: 1 },
  { word1: "intention", word2: "intention", expect: 0 },
  { word1: "algorithm", word2: "algoritm", expect: 1 },
  { word1: "park", word2: "spake", expect: 3 },
];

const minDistanceFuncs = {
  Recursive: minDistanceRecursive,
  Memoization: minDistanceMemo,
  Tabulation: minDistanceTab,
  SpaceOptimized: minDistanceSpace,
};

for (const name in minDistanceFuncs) {
  console.log(`Testing ${name}:`);
  testCases.forEach((tc, i) => {
    console.log(`Test Case ${i + 1}: word1="${tc.word1}", word2="${tc.word2}"`);
    const res = minDistanceFuncs[name](tc.word1, tc.word2);
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
