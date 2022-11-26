package example_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/0x00b/wx-openapi/api/client/miniprogram"
)

func TestMain(m *testing.T) {
	cfg := miniprogram.NewConfiguration()
	client := miniprogram.NewAPIClient(cfg)

	req := client.MiniProgramAuthApi.SnsJscode2sessionGet(context.Background())
	rsp, hr, e := req.Appid("q").GrantType("x").JsCode("x").Secret("x").Execute()

	fmt.Println(rsp, hr, e)
}
