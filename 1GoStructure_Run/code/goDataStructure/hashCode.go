
//此源码被清华学神尹成大魔王专业翻译分析并修改
//尹成QQ77025077
//尹成微信18510341407
//尹成所在QQ群721929980
//尹成邮箱 yinc13@mails.tsinghua.edu.cn
//尹成毕业于清华大学,微软区块链领域全球最有价值专家
//https://mvp.microsoft.com/zh-cn/PublicProfile/4033620
package goDataStructure

import (
	"crypto/md5"
	"fmt"
//“哈希”
	"io"
)

func Te() {
//A:＝42
	a := md5.New()
	b := md5.New()
	c := md5.New()
	io.WriteString(a, "The fog is getting thicker!")
	io.WriteString(b, "And Leon's getting laaarger!")
	io.WriteString(c, "The fog is getting thicker!")
	io.WriteString(c, "And Leon's getting laaarger!")

	fmt.Printf("%x\n%x\n", a.Sum(nil), b.Sum(nil))
	fmt.Printf("%x\n", c.Sum(nil))
}
