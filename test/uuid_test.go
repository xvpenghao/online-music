package test

import (
	"fmt"
	"github.com/satori/go.uuid"
	"strings"
	"testing"
)

func Test_UUID(t *testing.T) {
	uid, _ := uuid.NewV4()
	res := strings.Replace(uid.String(), "-", "", -1)
	fmt.Println(len(res))
	fmt.Println(res)
}
