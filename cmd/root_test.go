package cmd_test

import (
	"fmt"
	"strings"
	"testing"
)

func Test_GetName(t *testing.T) {
	CertificatePath := "wildcard-undangtemanrc.com_fullchain.pem"
	res := strings.Split(CertificatePath, "_")
	fmt.Println(res[0])
}
