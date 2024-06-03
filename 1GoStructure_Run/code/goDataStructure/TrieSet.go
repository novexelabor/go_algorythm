
//此源码被清华学神尹成大魔王专业翻译分析并修改
//尹成QQ77025077
//尹成微信18510341407
//尹成所在QQ群721929980
//尹成邮箱 yinc13@mails.tsinghua.edu.cn
//尹成毕业于清华大学,微软区块链领域全球最有价值专家
//https://mvp.microsoft.com/zh-cn/PublicProfile/4033620
package goDataStructure

type TrieSet struct {
	root *trieSetNode
	size int
}

type trieSetNode struct {
	isWord bool
	next   map[rune]*trieSetNode
}

func CreateTrieSet() *TrieSet {
	return &TrieSet{
		root: &trieSetNode{
			isWord: false,
			next:   make(map[rune]*trieSetNode),
		},
		size: 0,
	}
}

func (mt *TrieSet) GetSize() int {
	return mt.size
}

func (mt *TrieSet) Add(word Comparable) { //非递减写入法todo完成add的递减写入法
	cur := mt.root
	char := []rune(string(word.(Stringer)))

	for _, value := range char {
		if _, ok := cur.next[value]; !ok {
//newnode：=createtrienode（）。
//fmt.println（newnode）
			cur.next[value] = &trieSetNode{
				isWord: false,
				next:   map[rune]*trieSetNode{},
			}
		}
		cur = cur.next[value]
	}
	if !cur.isWord {
		cur.isWord = true
		mt.size++
	}
}
func (mt *TrieSet) Contains(word Comparable) bool {
	cur := mt.root
	char := []rune(string(word.(Stringer)))

	for _, value := range char {
		if _, ok := cur.next[value]; !ok {
			return false
		}
		cur = cur.next[value]
	}
	return cur.isWord
}
func (mt *TrieSet) IsPrefix(prefix Comparable) bool {
	cur := mt.root
	char := []rune(string(prefix.(Stringer)))

	for _, value := range char {
		if _, ok := cur.next[value]; !ok {
			return false
		}
		cur = cur.next[value]
	}
	return true
}

func (mt *TrieSet) Remove(value Comparable) {
	panic("implement me")
}

func (mt *TrieSet) IsEmpty() bool {
	return mt.size == 0
}

func (mt *TrieSet) String() string {
	panic("implement me")
}
