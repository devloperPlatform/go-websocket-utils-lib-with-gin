package websocket

import (
	"fmt"
	"testing"
)

func TestWebSocket(t *testing.T) {

	//engine := NewWebSocketServer("/dev")
	//engine.Handle("hello", func(ctx *conn.Context) error {
	//	type tmpStrut struct {
	//		Hello string `socket:"hello"`
	//	}
	//	tmpData := &tmpStrut{}
	//	//tmpData := ""
	//	_ = ctx.Unmarshal(&tmpData)
	//	fmt.Println(tmpData)
	//	andReturn, err := ctx.SendMsgAndReturn("/hello", conn.ModTypeFile, "😄😄")
	//	if err != nil {
	//		return err
	//	}
	//
	//	tmpStr := ""
	//	err = andReturn.Unmarshal(&tmpStr)
	//	if err != nil {
	//		return err
	//	}
	//	fmt.Println(tmpStr)
	//	//ctx.ReturnData(2)
	//	return nil
	//})
	////engine.GET("/", func(context *gin.Context) {
	////	context.String(200, "1sdf")
	////})
	//if err := engine.Run(":65528"); err != nil {
	//	panic(err)
	//}

	fmt.Println("完成")
}
