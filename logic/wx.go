package logic

import (
	"crypto/sha1"
    "fmt"
    "io"
    "log"
	"sort"
    "strings"
)

const (
	token = "testServerWeiXin"
)

//本地计算signature
func makeSignature(timestamp, nonce string) string { 
    si := []string{token, timestamp, nonce}
    sort.Strings(si)            //字典序排序
    str := strings.Join(si, "") //组合字符串
    s := sha1.New()             //返回一个新的使用SHA1校验的hash.Hash接口
    io.WriteString(s, str)      //WriteString函数将字符串数组str中的内容写入到s中
    return fmt.Sprintf("%x", s.Sum(nil))
}

func Sign () string {
	return ""
}