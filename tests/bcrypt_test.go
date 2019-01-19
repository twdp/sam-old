package tests

import (
	"fmt"
	"testing"
	"tianwei.pro/business"
)

func TestBcrypt(t *testing.T)  {
	fmt.Println(business.GenerateCrypto("123"))
}