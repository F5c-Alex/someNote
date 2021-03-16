// 你这个学期必须选修 numCourses 门课程，记为 0 到 numCourses - 1 。

// 在选修某些课程之前需要一些先修课程。 先修课程按数组 prerequisites 给出，其中 prerequisites[i] = [ai, bi] ，表示如果要学习课程 ai 则 必须 先学习课程  bi 。

// 例如，先修课程对 [0, 1] 表示：想要学习课程 0 ，你需要先完成课程 1 。
// 请你判断是否可能完成所有课程的学习？如果可以，返回 true ；否则，返回 false 。

// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/course-schedule
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。


func canFinish(numCourses int, prerequisites [][]int) bool {
	visited := make([]int,numCourses)
	table := make([][]int,numCourses)
	valid := true
	for i:=0;i<numCourses;i++{
			table[i]=[]int{}
	}
	for _,v := range prerequisites{
			table[v[0]] = append(table[v[0]],v[1])
	}
	var dfs func(int)
	dfs = func(cur int){
			if !valid{
					return
			}
			visited[cur] = 1
			for _,v := range table[cur]{
					if visited[v] == 0{
							dfs(v)
					}else if visited[v]==1{
							valid = false
							return
					}
			}
			visited[cur] = 2
			return
	}
	for k,v := range visited{
			if v == 0{
					dfs(k)
			} 
			if !valid{
					return false
			}
	}
	return true
}

