package main

import "fmt"

func main() {
	//byte又叫uint8类型，代表了ascii码的一个字符
	str := "hello"
	b := []byte(str)
	b[0] = 'a'
	fmt.Println("", string(b))

	//rune类型实际是一个int32类型 代表一个utf-8类型 汉字 日文其他符合文字都用utf-8编码
	//一个utf-8占用4个字节
	strCN := "拒绝了"
	fmt.Println("len(rune):", len(strCN))
	bCN := []rune(strCN)
	bCN[0] = 'a'
	fmt.Println("", string(bCN))
}
