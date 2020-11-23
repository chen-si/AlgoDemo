package DynamicProgramming

/**
 * 动态规划算法(跳跃点)求解0-1背包问题
 * @param n 物品数量
 * @param c 背包容量
 * @param v 物品价值
 * @param w 物品重量
 * @param p 跳跃点集 其中p[][0]代表重量 p[][1]代表价值
 * @param x 输出的0-1序列
 * @return 最大总价值
 */
func Knapsack(n int, c int, v []int, w []int, p [][]int, x []int) int{
	// 初始化一些变量
	head := make([]int, n + 2)
	head[n + 1] = 0
	p[0][0] = 0
	p[0][1] = 0
	left, right, next := 0, 0, 1
	head[n] = 1

	// 从最后一个物品往前追溯
	for i := n - 1; i >= 0; i --{
		k := left
		//fmt.Println(left, right)
		for j := left; j <= right ;j++{
			//fmt.Println(i, k, next)
			// 如果重量0超过背包容量
			if p[j][0] + w[i] > c{
				break
			}

			// y,m 用于指代跳跃点集 q
			y, m := p[j][0] + w[i], p[j][1] + v[i]
			// fmt.Println(i, "+", y, m)

			// 重量更小的跳跃点集-保留
			for k <= right && p[k][0] < y{
				p[next][0] = p[k][0]
				p[next][1] = p[k][1]
				next ++
				k ++
			}
			// 如果总重量相同但是价值更少，覆盖
			if k <= right && p[k][0] == y {
				if m < p[k][1]{
					m = p[k][1]
				}
				k ++
			}
			// 如果价值比之前的最大的还大 加到最后面
			if m >= p[next - 1][1]{
				p[next][0] = y
				p[next][1] = m
				next ++
			}
			// 清除上一个跳跃点集的无效部分（重量更大但是价值更少）
			for k <= right && p[k][1] <= p[next - 1][1]{
				k++
			}
		}
		// 填写相同的部分
		for k <= right {
			p[next][0] = p[k][0]
			p[next][1] = p[k][1]
			next ++
			k ++
		}

		left = right + 1
		right = next - 1
		// 记录起始点
		head[i] = next
	}
	Traceback(n, w, v, p, head, x)
	return p[next -1][1]
}

/**
 * 动态规划算法(跳跃点)求解0-1背包问题
 * @param n 物品数量
 * @param w 物品重量
 * @param v 物品价值
 * @param p 跳跃点集 其中p[][0]代表重量 p[][1]代表价值
 * @param head 记录跳跃点开始地点
 * @param x 输出的0-1序列
 */
func Traceback(n int, w []int, v []int, p [][]int, head []int, x []int){
	j, m := p[head[0] - 1][0], p[head[0] - 1][1]
	for i := 1; i <= n; i++{
		x[i - 1] = 0
		for k := head[i + 1]; k <= head[i] - 1;k ++{
			if p[k][0] + w[i - 1] == j && p[k][1] + v[i - 1] == m{
				// 将第i个物品放入背包中
				x[i - 1] = 1
				j = p[k][0]
				m = p[k][1]
				break
			}
		}
	}
}