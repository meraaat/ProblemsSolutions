// Recursive
function lengthOfLISRecursive(nums) {
  function lis(i) {
    if (i === 0) {
      return 1;
    }
    let maxLen = 1;
    for (let j = 0; j < i; j++) {
      if (nums[j] < nums[i]) {
        maxLen = Math.max(maxLen, 1 + lis(j));
      }
    }
    return maxLen;
  }

  let maxLength = 0;
  for (let i = 0; i < nums.length; i++) {
    maxLength = Math.max(maxLength, lis(i));
  }
  return maxLength;
}

// Memoization
function lengthOfLISMemo(nums) {
  const memo = new Array(nums.length).fill(-1);

  function lis(i) {
    if (i === 0) {
      return 1;
    }
    if (memo[i] !== -1) {
      return memo[i];
    }
    let maxLen = 1;
    for (let j = 0; j < i; j++) {
      if (nums[j] < nums[i]) {
        maxLen = Math.max(maxLen, 1 + lis(j));
      }
    }
    memo[i] = maxLen;
    return maxLen;
  }

  let maxLength = 0;
  for (let i = 0; i < nums.length; i++) {
    maxLength = Math.max(maxLength, lis(i));
  }
  return maxLength;
}

// Tabulation
function lengthOfLISTab(nums) {
  if (nums.length === 0) {
    return 0;
  }
  const dp = new Array(nums.length).fill(1);

  for (let i = 1; i < nums.length; i++) {
    for (let j = 0; j < i; j++) {
      if (nums[i] > nums[j]) {
        dp[i] = Math.max(dp[i], dp[j] + 1);
      }
    }
  }

  let maxLength = 0;
  for (const length of dp) {
    maxLength = Math.max(maxLength, length);
  }
  return maxLength;
}

// Patience Sorting with Binary Search
function lengthOfLISPatience(nums) {
  const piles = [];

  for (const num of nums) {
    let l = 0;
    let r = piles.length;

    while (l < r) {
      const mid = Math.floor((l + r) / 2);
      if (piles[mid] < num) {
        l = mid + 1;
      } else {
        r = mid;
      }
    }

    if (l === piles.length) {
      piles.push(num);
    } else {
      piles[l] = num;
    }
  }

  return piles.length;
}

const testCases = [
  { nums: [10, 9, 2, 5, 3, 7, 101, 18], expect: 4 },
  { nums: [0, 1, 0, 3, 2, 3], expect: 4 },
  { nums: [7, 7, 7, 7, 7, 7, 7], expect: 1 },
  { nums: [1, 2, 1, 4, 3, 4, 5], expect: 5 },
  { nums: [], expect: 0 },
  { nums: [1], expect: 1 },
  { nums: [1, 2], expect: 2 },
  { nums: [2, 1], expect: 1 },
  { nums: [1, 3, 2, 4, 5], expect: 4 },
  { nums: [1, 3, 6, 7, 9, 4, 10, 5, 6], expect: 6 },
];

const lisFuncs = {
  Recursive: lengthOfLISRecursive,
  Memoization: lengthOfLISMemo,
  Tabulation: lengthOfLISTab,
  PatienceSorting: lengthOfLISPatience,
};

for (const name in lisFuncs) {
  console.log(`Testing ${name}:`);
  testCases.forEach((tc, i) => {
    console.log(`Test Case ${i + 1}: nums=${JSON.stringify(tc.nums)}`); // Use JSON.stringify for array output
    const res = lisFuncs[name](tc.nums);
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
