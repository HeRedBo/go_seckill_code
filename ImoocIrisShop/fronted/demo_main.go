package main

import (
	"ImoocIrisShop/encrypt"
	"fmt"
)

func main() {
	text := "10" // 你要加密的数据
	//AesKey := []byte("#HvL%$o0oNNoOZnk#o2qbqCeQB1iXeIR") // 对称秘钥长度必须是16的倍数
	//fmt.Printf("明文: %s\n秘钥: %s\n", text, string(AesKey))
	encrypted, err := encrypt.EnPwdCode([]byte(text))
	if err != nil {
		panic(err)
	}
	fmt.Printf("加密后: %s\n", encrypted)
	//encrypteds, _ := base64.StdEncoding.DecodeString("xvhqp8bT0mkEcAsNK+L4fw==")
	origin, err := encrypt.DePwdCode(encrypted)
	if err != nil {
		panic(err)
	}
	fmt.Printf("解密后明文: %s\n", string(origin))
}
