// Recursive (Naive)
function knapsackRecursive(W, wt, val, n) {
    if (n === 0 || W === 0) {
        return 0;
    }
    if (wt[n - 1] > W) {
        return knapsackRecursive(W, wt, val, n - 1);
    }
    return Math.max(
        val[n - 1] + knapsackRecursive(W - wt[n - 1], wt, val, n - 1),
        knapsackRecursive(W, wt, val, n - 1)
    );
}

// Memoization (Top-Down DP)
function knapsackMemo(W, wt, val, n) {
    const memo = Array(n + 1).fill(null).map(() => Array(W + 1).fill(-1));

    function knap(W, n) {
        if (n === 0 || W === 0) {
            return 0;
        }
        if (memo[n][W] !== -1) {
            return memo[n][W];
        }
        if (wt[n - 1] > W) {
            memo[n][W] = knap(W, n - 1);
            return memo[n][W];
        }
        memo[n][W] = Math.max(
            val[n - 1] + knap(W - wt[n - 1], n - 1),
            knap(W, n - 1)
        );
        return memo[n][W];
    }
    return knap(W, n);
}

// Tabulation (Bottom-Up DP)
function knapsackTab(W, wt, val, n) {
    const dp = Array(n + 1).fill(null).map(() => Array(W + 1).fill(0));

    for (let i = 0; i <= n; i++) {
        for (let w = 0; w <= W; w++) {
            if (i === 0 || w === 0) {
                dp[i][w] = 0;
            } else if (wt[i - 1] <= w) {
                dp[i][w] = Math.max(
                    val[i - 1] + dp[i - 1][w - wt[i - 1]],
                    dp[i - 1][w]
                );
            } else {
                dp[i][w] = dp[i - 1][w];
            }
        }
    }
    return dp[n][W];
}

const testCases = [
    { W: 8, wt: [4, 5, 2], val: [3, 9, 5], n: 3, expect: 14 },
    { W: 4, wt: [4, 5, 1], val: [1, 2, 3], n: 3, expect: 3 },
    { W: 10, wt: [10, 20], val: [60, 100], n: 2, expect: 100 },
    { W: 50, wt: [10, 20, 30], val: [60, 100, 120], n: 3, expect: 220 },
    { W: 10, wt: [1, 2, 3, 4, 5], val: [1, 2, 3, 4, 5], n: 5, expect: 10 },
];

const knapsackFuncs = {
    "Recursive": knapsackRecursive,
    "Memoization": knapsackMemo,
    "Tabulation": knapsackTab,
};

for (const name in knapsackFuncs) {
    console.log(`Testing ${name}:`);
    testCases.forEach((tc, i) => {
        console.log(`Test Case ${i + 1}: W=${tc.W}, wt=${tc.wt}, val=${tc.val}, n=${tc.n}`);
        const res = knapsackFuncs[name](tc.W, tc.wt, tc.val, tc.n);
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