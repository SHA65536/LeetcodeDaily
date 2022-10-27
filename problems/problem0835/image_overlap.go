package problem0835

/*
You are given two images, img1 and img2, represented as binary, square matrices of size n x n. A binary matrix has only 0s and 1s as values.
We translate one image however we choose by sliding all the 1 bits left, right, up, and/or down any number of units.
We then place it on top of the other image. We can then calculate the overlap by counting the number of positions that have a 1 in both images.
Note also that a translation does not include any kind of rotation. Any 1 bits that are translated outside of the matrix borders are erased.
Return the largest possible overlap.
*/

func largestOverlap(img1 [][]int, img2 [][]int) int {
	var res, cur int
	var n int = len(img1)
	var ni, nj int
	for oi := -(n - 1); oi <= (n-1)*2; oi++ {
		for oj := -(n - 1); oj <= (n-1)*2; oj++ {
			cur = 0
			for i := 0; i < n; i++ {
				ni = i + oi
				if ni >= n || ni < 0 {
					continue
				}
				for j := 0; j < n; j++ {
					nj = j + oj
					if nj >= n || nj < 0 {
						continue
					}
					if img1[ni][nj] == 1 && img2[i][j] == 1 {
						cur++
					}
				}
			}
			if cur > res {
				res = cur
			}
		}
	}
	return res
}

func largestOverlapOpt(img1 [][]int, img2 [][]int) int {
	var res int
	var n = len(img1)
	var LA, LB []int = make([]int, 0, n), make([]int, 0, n)
	var count = map[int]int{}
	for i := 0; i < n*n; i++ {
		if img1[i/n][i%n] == 1 {
			LA = append(LA, i/n*100+i%n)
		}
		if img2[i/n][i%n] == 1 {
			LB = append(LB, i/n*100+i%n)
		}
	}
	for _, i := range LA {
		for _, j := range LB {
			count[i-j]++
		}
	}
	for _, c := range count {
		if c > res {
			res = c
		}
	}
	return res
}

/*
 int largestOverlap(vector<vector<int>>& A, vector<vector<int>>& B) {
        vector<int> LA, LB;
        int N = A.size();
        unordered_map<int, int> count;
        for (int i = 0; i < N * N; ++i)
            if (A[i / N][i % N] == 1)
                LA.push_back(i / N * 100 + i % N);
        for (int i = 0; i < N * N; ++i)
            if (B[i / N][i % N] == 1)
                LB.push_back(i / N * 100 + i % N);
        for (int i : LA) for (int j : LB) count[i - j]++;
        int res = 0;
        for (auto it : count) res = max(res, it.second);
        return res;
*/
