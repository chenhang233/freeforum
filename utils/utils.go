package utils

import (
	"freeforum/utils/logs"
	"math/rand"
	"strings"
	"time"
)

const words = "qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM"

func RandomWordsIndex(words string) int {
	nR := rand.New(rand.NewSource(time.Now().UnixNano()))
	wLen := len(words)
	return nR.Intn(wLen - 1)
}

func RandomUUID(addr string) string {
	sb := strings.Builder{}
	for i := 0; i < 10; i++ {
		addrList := strings.Split(addr, ":")
		if len(addrList) < 2 {
			logs.LOG.Error.Println("len(addrList) < 2")
			addrList = append(addrList, "99", "88")
		}
		for i, v := range addrList {
			sb.WriteByte(byte(i))
			sb.WriteByte(v[RandomWordsIndex(v)])
		}

	}
	for i := 0; i < 20; i++ {
		sb.WriteByte(words[RandomWordsIndex(words)])
	}
	return sb.String()
}
