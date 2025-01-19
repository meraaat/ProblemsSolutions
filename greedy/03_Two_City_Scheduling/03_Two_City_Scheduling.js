function twoCitySchedCost(costs) {
  const n = costs.length / 2;

  // Sort by the difference between cost to A and cost to B
  costs.sort((a, b) => a[0] - a[1] - (b[0] - b[1]));

  let minCost = 0;
  for (let i = 0; i < n; i++) {
    minCost += costs[i][0]; // Cost to city A for the first n people
    minCost += costs[i + n][1]; // Cost to city B for the remaining n people
  }

  return minCost;
}

const testCases = [
  {
    costs: [
      [10, 20],
      [30, 200],
      [400, 50],
      [30, 20],
    ],
    expect: 110,
  },
  {
    costs: [
      [5, 20],
      [30, 100],
      [400, 30],
      [50, 10],
    ],
    expect: 75,
  },
  {
    costs: [
      [20, 700],
      [400, 50],
      [900, 600],
      [200, 150],
      [800, 100],
      [500, 450],
    ],
    expect: 1470,
  },
  {
    costs: [
      [259, 770],
      [448, 54],
      [926, 667],
      [184, 139],
      [840, 118],
      [577, 469],
    ],
    expect: 1859,
  },
  {
    costs: [
      [10, 90],
      [80, 20],
      [30, 70],
      [60, 40],
    ],
    expect: 110,
  },
  {
    costs: [
      [10, 10],
      [10, 10],
    ],
    expect: 20,
  },
  {
    costs: [
      [0, 2],
      [1, 1],
      [2, 0],
    ],
    expect: 2,
  },
];

for (let i = 0; i < testCases.length; i++) {
  const tc = testCases[i];
  console.log(`Test Case ${i + 1}: costs=${JSON.stringify(tc.costs)}`);
  const res = twoCitySchedCost(tc.costs);
  console.log(`Got: ${res}`);
  console.log(`Expected: ${tc.expect}`);
  if (res !== tc.expect) {
    console.log("FAIL");
  } else {
    console.log("PASS");
  }
  console.log();
}
