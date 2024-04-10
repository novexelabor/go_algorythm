package HashTableArray

import (
	"crypto/sha512"

	"errors"
)

const (
	Deleted      = iota //数据已经被删除
	MintableSize = 100  //哈希表的大小
	legimate     = iota //已经存在的合法数据
	Empty        = iota //数据为空

)

//哈希函数自定义，获取哈希值
func MySHA(str interface{}, tableSize int) int {
	var Hashvar int = 0
	var chars []byte
	if strings, ok := str.(string); ok { //能否转换为字符串
		chars = []byte(strings) //字符串转化为字节数组
	}
	for _, v := range chars {
		Hashvar = (Hashvar<<17 | 123&1235 ^ 139) + int(v) //哈希算法,自定义的
	}
	return Hashvar % MintableSize

}

func MySHA256(str string, tableSize int) int {
	shaobj := sha512.New()    //哈希算法加密,固定长度的哈希值
	shaobj.Write([]byte(str)) //哈希
	mybytes := shaobj.Sum(nil)

	var Hashvar int = 0
	for _, v := range mybytes {
		Hashvar = (Hashvar<<3 | 123&1235 ^ 139) + int(v) //哈希算法
	}
	return Hashvar % MintableSize

}

type HashFunc func(data interface{}, tableSize int) int //函数签名,函数指针

type HashEntry struct {
	data interface{} //数据
	kind int         //类型
}

type HashTable struct {
	tableSize int          //哈希表的大小
	theCells  []*HashEntry //数组，每一个元素是指针指向哈希结构
	hashfunc  HashFunc     //调用哈希函数
}
type HashtableGO interface {
	Find(data interface{}) int      //查找数据
	Insert(data interface{})        //插入数据
	Empty()                         //为空
	GetValue(index int) interface{} //抓取value
}

func NewHashTable(size int, hash HashFunc) (*HashTable, error) {
	if size < MintableSize {
		return nil, errors.New("哈希表太小")
	}
	if hash == nil {
		return nil, errors.New("没有哈希函数")
	}
	hashtable := new(HashTable)                   //创建哈希表
	hashtable.tableSize = size                    //设置哈希表大小
	hashtable.theCells = make([]*HashEntry, size) //数组分配内存
	hashtable.hashfunc = hash                     //设置哈希函数
	for i := 0; i < hashtable.tableSize; i++ {
		hashtable.theCells[i] = new(HashEntry) //开辟元素内存
		hashtable.theCells[i].data = nil
		hashtable.theCells[i].kind = Empty //设置为空,用标志来表示元素情况
	}

	return hashtable, nil

}
func (ht *HashTable) Find(data interface{}) int {
	var collid int = 0
	curpos := ht.hashfunc(data, ht.tableSize) //计算哈希位置
	if ht.theCells[curpos].kind != Empty && ht.theCells[curpos].data != data {
		collid += 1            //哈希冲突
		curpos := 2*curpos - 1 //平方探测，处理冲突
		if curpos > ht.tableSize {
			curpos -= ht.tableSize //越界，返回
		}

	}
	return curpos

}
func (ht *HashTable) Insert(data interface{}) {
	pos := ht.Find(data)      //查找数据位置
	entry := ht.theCells[pos] //插入数据记录状态
	if entry.kind != legimate {
		entry.kind = legimate
		entry.data = data //插入数据
	}
}
func (ht *HashTable) Empty() {
	for i := 0; i < ht.tableSize; i++ {
		if ht.theCells[i] == nil {
			continue //循环清空数据
		}
		ht.theCells[i].kind = Deleted //删除数据，无法删除数组，只能人为的设置删除标志
	}
}
func (ht *HashTable) GetValue(index int) interface{} {
	if index > ht.tableSize {
		return nil //判断大小
	}
	entry := ht.theCells[index] //取出数据
	if entry.kind == legimate {
		return entry.data
	} else {
		return nil
	}

}
