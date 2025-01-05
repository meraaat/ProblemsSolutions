// Recursive
function longestCommonSubsequenceRecursive(text1, text2) {
  function lcs(i, j) {
    if (i < 0 || j < 0) {
      return 0;
    }
    if (text1[i] === text2[j]) {
      return 1 + lcs(i - 1, j - 1);
    }
    return Math.max(lcs(i - 1, j), lcs(i, j - 1));
  }
  return lcs(text1.length - 1, text2.length - 1);
}

// Memoization
function longestCommonSubsequenceMemo(text1, text2) {
  const memo = Array(text1.length + 1)
    .fill(null)
    .map(() => Array(text2.length + 1).fill(-1));

  function lcs(i, j) {
    if (i < 0 || j < 0) {
      return 0;
    }
    if (memo[i + 1][j + 1] !== -1) {
      return memo[i + 1][j + 1];
    }
    if (text1[i] === text2[j]) {
      memo[i + 1][j + 1] = 1 + lcs(i - 1, j - 1);
    } else {
      memo[i + 1][j + 1] = Math.max(lcs(i - 1, j), lcs(i, j - 1));
    }
    return memo[i + 1][j + 1];
  }
  return lcs(text1.length - 1, text2.length - 1);
}

// Tabulation
function longestCommonSubsequenceTab(text1, text2) {
  const m = text1.length;
  const n = text2.length;
  const dp = Array(m + 1)
    .fill(null)
    .map(() => Array(n + 1).fill(0));

  for (let i = 1; i <= m; i++) {
    for (let j = 1; j <= n; j++) {
      if (text1[i - 1] === text2[j - 1]) {
        dp[i][j] = dp[i - 1][j - 1] + 1;
      } else {
        dp[i][j] = Math.max(dp[i - 1][j], dp[i][j - 1]);
      }
    }
  }
  return dp[m][n];
}

// Space Optimized
function longestCommonSubsequenceSpace(text1, text2) {
  const m = text1.length;
  const n = text2.length;

  if (m < n) {
    [text1, text2] = [text2, text1]; // Swap strings if text1 is shorter
    [m, n] = [n, m];
  }

  let prev = Array(n + 1).fill(0);
  let curr = Array(n + 1).fill(0);

  for (let i = 1; i <= m; i++) {
    for (let j = 1; j <= n; j++) {
      if (text1[i - 1] === text2[j - 1]) {
        curr[j] = prev[j - 1] + 1;
      } else {
        curr[j] = Math.max(prev[j], curr[j - 1]);
      }
    }
    [prev, curr] = [curr, prev]; // Swap arrays
  }
  return prev[n];
}

const testCases = [
  { text1: "abcde", text2: "ace", expect: 3 },
  { text1: "abc", text2: "abc", expect: 3 },
  { text1: "abc", text2: "def", expect: 0 },
  { text1: "abcdef", text2: "acef", expect: 4 },
  { text1: "", text2: "", expect: 0 },
  { text1: "abc", text2: "", expect: 0 },
  { text1: "", text2: "def", expect: 0 },
  { text1: "fish", text2: "fosh", expect: 3 },
];

const lcsFuncs = {
  Recursive: longestCommonSubsequenceRecursive,
  Memoization: longestCommonSubsequenceMemo,
  Tabulation: longestCommonSubsequenceTab,
  SpaceOptimized: longestCommonSubsequenceSpace,
};

for (const name in lcsFuncs) {
  console.log(`Testing ${name}:`);
  testCases.forEach((tc, i) => {
    console.log(`Test Case ${i + 1}: text1="${tc.text1}", text2="${tc.text2}"`);
    const res = lcsFuncs[name](tc.text1, tc.text2);
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
