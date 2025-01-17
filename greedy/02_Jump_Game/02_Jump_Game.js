function canJump(nums) {
  let reachable = 0;
  for (let i = 0; i < nums.length; i++) {
    if (i > reachable) {
      return false;
    }
    reachable = Math.max(reachable, i + nums[i]);
  }
  return reachable >= nums.length - 1;
}

const testCases = [
  { nums: [2, 3, 1, 1, 4], expect: true },
  { nums: [3, 2, 1, 0, 4], expect: false },
  { nums: [1, 3, 4, 1, 1, 2, 1], expect: true },
  { nums: [1, 3, 4, 2, 1, 1, 0, 1, 1], expect: false },
  { nums: [0], expect: true },
  { nums: [2, 0, 0], expect: true },
  { nums: [1, 0, 1, 0], expect: false },
  { nums: [1, 1, 1, 1], expect: true },
  { nums: [1, 1, 0, 1], expect: false },
  { nums: [5, 0, 0, 0, 0, 0, 1], expect: false },
  { nums: [2, 5, 0, 0], expect: true },
];

for (let i = 0; i < testCases.length; i++) {
  const tc = testCases[i];
  console.log(`Test Case ${i + 1}: nums=${JSON.stringify(tc.nums)}`);
  const res = canJump(tc.nums);
  console.log(`Got: ${res}`);
  console.log(`Expected: ${tc.expect}`);
  if (res !== tc.expect) {
    console.log("FAIL");
  } else {
    console.log("PASS");
  }
  console.log();
}
