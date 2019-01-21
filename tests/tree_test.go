package tests

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"testing"
)

func TestTreeValid(t *testing.T) {
	tr := beego.NewTree()
	tr.AddRouter("/v1/shop/:branchId([0-9]+)", "sam")
	ctx := context.NewContext()
	ctx.Input.SetData("branchId", 2222)

	obj := tr.Match("/v1/shop/123", ctx)

	if obj == nil || obj.(string) != "sam" {
		t.Fatal("fff")
	}
	if vv := ctx.Input.Param(":branchId"); vv != "123" {
		t.Fatal("...")
	}

}