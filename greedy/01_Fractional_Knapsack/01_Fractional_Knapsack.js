function fractionalKnapsackSimple(W, arr, n) {
  // [profit, weight]
  arr.sort((a, b) => b[0] / b[1] - a[0] / a[1]);

  let remainingWeight = W;
  let value = 0;

  for (let i = 0; i < n; i++) {
    if (remainingWeight === 0) break;
    let weight = Math.min(remainingWeight, arr[i][1]);

    remainingWeight -= weight;
    value += (arr[i][0] / arr[i][1]) * weight;
  }

  return weight;
}

function fractionalKnapsack(W, items) {
  const n = items.length;
  if (n === 0 || W === 0) {
    return 0;
  }

  const itemStructs = items.map(([profit, weight]) => ({
    profit,
    weight,
    ratio: profit / weight,
  }));

  itemStructs.sort((a, b) => b.ratio - a.ratio);

  let totalProfit = 0;
  let remainingWeight = W;

  for (const item of itemStructs) {
    if (item.weight <= remainingWeight) {
      totalProfit += item.profit;
      remainingWeight -= item.weight;
    } else {
      totalProfit += (remainingWeight / item.weight) * item.profit;
      remainingWeight = 0;
      break;
    }
  }

  return totalProfit;
}

const testCases = [
  {
    W: 25,
    items: [
      [70, 10],
      [90, 20],
      [150, 30],
    ],
    expect: 137.5,
  },
  {
    W: 45,
    items: [
      [70, 10],
      [90, 20],
      [150, 30],
    ],
    expect: 242.5,
  },
  {
    W: 10,
    items: [
      [10, 2],
      [5, 3],
      [15, 5],
      [7, 7],
      [6, 1],
      [18, 6],
    ],
    expect: 43,
  },
  { W: 10, items: [], expect: 0 },
  { W: 0, items: [[10, 5]], expect: 0 },
  {
    W: 50,
    items: [
      [60, 10],
      [100, 20],
      [120, 30],
    ],
    expect: 240,
  },
  {
    W: 10,
    items: [
      [10, 5],
      [1, 1],
    ],
    expect: 11,
  },
  {
    W: 30,
    items: [
      [60, 10],
      [100, 20],
      [120, 30],
    ],
    expect: 160,
  },
  {
    W: 5,
    items: [
      [60, 10],
      [100, 20],
      [120, 30],
    ],
    expect: 30,
  },
];

for (let i = 0; i < testCases.length; i++) {
  const tc = testCases[i];
  console.log(
    `Test Case ${i + 1}: W=${tc.W}, items=${JSON.stringify(tc.items)}`
  );
  const res = fractionalKnapsack(tc.W, tc.items);
  console.log(`Got: ${res}`);
  console.log(`Expected: ${tc.expect}`);
  if (Math.abs(res - tc.expect) < 0.0001) {
    // Compare floating points with tolerance
    console.log("PASS");
  } else {
    console.log("FAIL");
  }
  console.log();
}
