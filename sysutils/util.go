package sysutils

import (
	"github.com/twinj/uuid"
	"strings"
)

func Token() string {
	//crutime := time.Now().Unix()
	//h := md5.New()
	//io.WriteString(h, strconv.FormatInt(crutime, 10))
	u4 := uuid.NewV4()
	token := strings.Replace(u4.String(), "-", "",-1) // fmt.Sprintf("%x", h.Sum(nil))
	return token
}
