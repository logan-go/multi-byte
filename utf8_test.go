package multiByte

import (
	"fmt"
	"testing"
)

func TestGetUTF8MultiBytes(t *testing.T) {
	aaa := []byte("你好你是谁啊？??????74138970514fddfglkjdfg😳😓😫😕😰😥🎋🎄💽auɔ:ɔidŋmdʒəu❹❼⑻㈨㈠¼°¼¾³∮㎏ㄕㄑㄧㄡ♫❥❤✩♬πΠαρВξ")
	result := GetUTF8MultiBytes(aaa)
	for _, v := range result {
		fmt.Println("============")
		fmt.Println(string(v.Content))
		fmt.Println("起止点：", v.StartIndex, ",", v.EndIndex)
	}
}
