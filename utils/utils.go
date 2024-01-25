package utils

import (
	"math/rand"
	"strconv"
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
		time.Sleep(50)
		addrList := strings.Split(addr, ":")
		if len(addrList) < 2 {
			addrList = make([]string, 0)
			nR := rand.New(rand.NewSource(time.Now().UnixNano()))
			t1 := strconv.Itoa(nR.Int())
			t2 := strconv.Itoa(nR.Int())
			addrList = append(addrList, t1, t2)
		}
		for i, v := range addrList {
			sb.WriteByte(byte(i))
			sb.WriteByte(v[RandomWordsIndex(v)])
		}

	}
	for i := 0; i < 20; i++ {
		time.Sleep(50)
		sb.WriteByte(words[RandomWordsIndex(words)])
	}
	return sb.String()
}
