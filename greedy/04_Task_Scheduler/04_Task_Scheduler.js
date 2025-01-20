function leastInterval(tasks, n) {
  if (n === 0) {
    return tasks.length;
  }

  const freq = new Array(26).fill(0);
  for (const task of tasks) {
    freq[task.charCodeAt(0) - "A".charCodeAt(0)]++;
  }

  let maxFreq = 0;
  let maxFreqCount = 0;
  for (const f of freq) {
    if (f > maxFreq) {
      maxFreq = f;
      maxFreqCount = 1;
    } else if (f === maxFreq) {
      maxFreqCount++;
    }
  }

  return Math.max(tasks.length, (maxFreq - 1) * (n + 1) + maxFreqCount);
}

const testCases = [
  { tasks: ["A", "A", "A", "B", "B", "B"], n: 2, expect: 8 },
  { tasks: ["A", "A", "A", "B", "B", "B"], n: 0, expect: 6 },
  {
    tasks: ["A", "A", "A", "A", "A", "A", "B", "C", "D", "E", "F", "G"],
    n: 2,
    expect: 16,
  },
  { tasks: ["A", "A", "A", "B", "B", "B", "C", "C"], n: 2, expect: 8 },
  { tasks: ["A", "A", "B", "B"], n: 3, expect: 6 },
  { tasks: ["C", "E", "C", "E", "E", "C"], n: 2, expect: 8 },
  { tasks: ["C", "E", "C", "E", "E", "C"], n: 3, expect: 10 },
  { tasks: ["A", "A", "A", "B", "B", "B", "C", "D", "E"], n: 2, expect: 9 },
  { tasks: ["A", "A", "A", "B", "B", "B", "C", "D", "E"], n: 3, expect: 11 },
  { tasks: ["A", "A", "A", "A", "B", "B", "C", "C", "D"], n: 2, expect: 11 },
  { tasks: ["A", "A", "A", "B", "B"], n: 2, expect: 7 },
];

for (let i = 0; i < testCases.length; i++) {
  const tc = testCases[i];
  console.log(
    `Test Case ${i + 1}: tasks=${JSON.stringify(tc.tasks)}, n=${tc.n}`
  );
  const res = leastInterval(tc.tasks, tc.n);
  console.log(`Got: ${res}`);
  console.log(`Expected: ${tc.expect}`);
  if (res !== tc.expect) {
    console.log("FAIL");
  } else {
    console.log("PASS");
  }
  console.log();
}
