
//此源码被清华学神尹成大魔王专业翻译分析并修改
//尹成QQ77025077
//尹成微信18510341407
//尹成所在QQ群721929980
//尹成邮箱 yinc13@mails.tsinghua.edu.cn
//尹成毕业于清华大学,微软区块链领域全球最有价值专家
//https://mvp.microsoft.com/zh-cn/PublicProfile/4033620
package goDataStructure

import (
	"fmt"
	"io/ioutil"
	"math"
	"math/rand"
	"sort"
	"strings"
	"time"
)

func TestQueue(q Queue, op int) time.Duration {
	startTime := time.Now()
	for i := 0; i < op; i++ {
		fmt.Println(i)
		q.EnQueue(Student{"test", rand.Intn(10000)})
	}
	for i := 0; i < op-1; i++ {
		fmt.Println(i)
		q.DeQueue()
	}

	return time.Since(startTime)
}
func TestStack(s Stack, op int) time.Duration {
	startTime := time.Now()
	for i := 0; i < op; i++ {
		s.Push(Student{"test", rand.Intn(10000)})
	}
	for i := 0; i < op-1; i++ {
		s.Pop()
	}

	return time.Since(startTime)
}

func TestSet(s Set) time.Duration {
	startTime := time.Now()

	file := "./acm/pride-and-prejudice.txt"
//文件：=“/acm/测试文本”

	word := divideWord(file)
	sort.Strings(word)

	for _, value := range word {

		ls := strings.ToLower(value)

		s.Add(Stringer(ls))
	}
	fmt.Println(s.GetSize())

	return time.Since(startTime)
}
func TestBigData(op int, m Map) time.Duration {
	startTime := time.Now()
	rand.Seed(time.Now().Unix())

	a := make([]Integer, op)
	for i := 0; i < op; i++ {
		a[i] = Integer(rand.Intn(math.MaxInt32))
	}

	redBlackTreeMap := CreateRedBlackTreeMap()
	for _, value := range a {
		redBlackTreeMap.Add(value, nil)
	}

	return time.Since(startTime)
}
func TestMap(m Map) time.Duration {
	startTime := time.Now()

	file := "./acm/pride-and-prejudice.txt"
//文件：=“/acm/测试文本”

	word := divideWord(file)
//排序字符串（字）
	for _, value := range word {
		ls := strings.ToLower(value)
		s := Stringer(ls)
		num := m.Get(s)
		if num == nil {
			m.Add(s, 1)
		} else {
			m.Set(s, num.(int)+1)
		}
	}

	fmt.Print("单词的个数是：")
	fmt.Println(m.GetSize())
	fmt.Print("the 的个数是：")
	fmt.Println(m.Get(Stringer("the")))
	fmt.Print("pride 的个数是：")
	fmt.Println(m.Get(Stringer("pride")))
	fmt.Print("prejudice 的个数是：")
	fmt.Println(m.Get(Stringer("prejudice")))
	for _, value := range word {
		m.Contains(Stringer(value))
	}
	return time.Since(startTime)
}
func divideWord(path string) []string {
	var wordMap []string
	file, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Printf("读文件错误%v", err)
	}
	text := string(file)
//按字母读取，除26个字母（大小写）之外的所有字符均认为是分隔符
	var size = 0
	letterStart := false
	var str = ""
	for _, value := range text {
		if (value >= 65 && value <= 90) || (value >= 97 && value <= 122) {
			letterStart = true
			str += fmt.Sprintf("%c", value)
		} else {
			if letterStart {
				size++
				letterStart = false
				wordMap = append(wordMap, str)
//fmt.println（str，wordmap）
				str = ""
			}
		}
	}
	return wordMap
}
func TestHeap(h Heap, op int) time.Duration {
	startTime := time.Now()
	rand.Seed(time.Now().Unix())
//T：=[]利息78、73、47、72、47、39、27、69、62、46、3、7、31、8、4、33、37、18、2、44
	for i := 0; i < op; i++ {
		h.Add(Integer(rand.Intn(10000)))
	}
//对于u，值：=范围t
//h.加法（整数（值））
//}
	var r []int
	for i := 0; i < op; i++ {
		n := int(h.Remove().(Integer))
		r = append(r, n)
	}
	for i := 1; i < op; i++ {
		if r[i-1] < r[i] {
			panic("err")
		}
	}
	return time.Since(startTime)
}
func TestUnionFind(uf UnionFind, op int) time.Duration {
	startTime := time.Now()
	rand.Seed(time.Now().Unix())
	size := uf.GetSize()

	for i := 0; i < op; i++ {
		a := rand.Intn(size)
		b := rand.Intn(size)
		uf.Union(a, b)
	}

	for i := 0; i < op; i++ {
		a := rand.Intn(size)
		b := rand.Intn(size)
		uf.IsConnected(a, b)
	}

	return time.Since(startTime)
}
