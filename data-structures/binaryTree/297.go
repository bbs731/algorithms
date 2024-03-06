package binaryTree

import (
	"strconv"
	"strings"
)

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	var buf strings.Builder
	var dfs func(*TreeNode)
	dfs = func(r *TreeNode) {
		if r == nil {
			buf.WriteString("X,")
			return
		}
		buf.WriteString(strconv.Itoa(r.Val))
		buf.WriteString(",")
		dfs(r.Left)
		dfs(r.Right)
	}
	dfs(root)
	return buf.String()
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	p := strings.Split(data, ",")
	var build func() *TreeNode
	build = func() *TreeNode {
		// 这里是真的tmd 的神奇啊， 不用管 p 的长度。因为 build() 返回 nil 之后， nil 不会在去call build.
		// 最后，结束的刚刚好。所有的node 都按照之前序列化的顺序，build 好了。
		//if len(p) == 0 {
		//	return nil
		//}
		n := p[0]
		p = p[1:]
		if n == "X" {
			return nil
		}

		val, _ := strconv.Atoi(n)
		return &TreeNode{val, build(), build()}
	}

	return build()
}
