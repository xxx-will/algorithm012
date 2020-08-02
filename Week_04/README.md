学习笔记


贪心套路：
    循环 or 递归
        单层逻辑-- 枚举每种case情况。使用最有解。 不需要保存档次最优解结果


二分思路
    前提单调有序
    确定好上下界
    一次迭代掉一半
    

半有序数组
    上下界，中点。
    若左界点小于中点 且 右界小于左界 则 迭代各取左右半中点
    若左界不小于中点 则，在左半边，新左右界为左界和中点
    若右界不大于中点 则，在右半边，新左右界为 中点和右界
    迭代查找，若左界大于右届，且右界小于中点，则返回，即为找到结果。

DFS
    递归，或者手动维护一个栈

    模版代码：
    递归

    set visted
    dfs(node, visted) {
        if visted
            return

        visited.add(node)
        process(node)

        next == generate()

        dfs(next, visted)
    }

    栈
        visted  set
        stack
    dfs(node, stack, visited) {
        stack.push(node)

        for stack.empty() == false {
            stack.pop(node)
            visited.add(node)
            
            process(node)
            next = generate()
            stack.push(next)
        }

    }

BFS
    不用栈，用队列

    queue
    visited set
func bfs(node, queue, visited) {
    if node visited {
        return
    }

    queue.push(node)

    for queue.empty() == false {
        queue.pop()
        visited.add(node)

        process(node)
        new = generate()
        queue.push(new)
    }

}


递归，bfs，dfs的范型模版心中有了基本的模子。在完成作业过程中，还是困难多多。
1，重复子问题的分析，归纳
2，基于基础模版的具体题目，具体实现上
3，复杂度优化
三方面的问题，离掌握还有很长的路。。。
