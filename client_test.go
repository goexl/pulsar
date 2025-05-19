package pulsar_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/goexl/pulsar"
	"github.com/goexl/pulsar/internal/message"
)

type Handler struct{}

type User struct {
	Name string `json:"name"`
}

func (h Handler) Peek() any {
	return new(User)
}

func (h Handler) Process(_ context.Context, msg any, _ *message.Extra) (err error) {
	fmt.Println(msg)

	return
}

func TestClient(t *testing.T) {
	user := new(User)
	user.Name = "test23"
	client := pulsar.New().Server().Url("http://pulsar-nzoz42zpr95v.tdmq.ap-cd.public.tencenttdmq.com:8080").
		Token("eyJrZXlJZCI6InB1bHNhci1uem96NDJ6cHI5NXYiLCJhbGciOiJIUzI1NiJ9.eyJzdWIiOiJwdWxzYXItbnpvejQyenByOTV2X2l0Y291cnNlZSJ9.1bvGreGBkB6Sj2fBNZW7a-vTJmK8T052CJTV8Ygq22k").
		Producer().Topic("persistent://pulsar-nzoz42zpr95v/itcoursee/itcoursee").Tag("local").Build().
		Consumer().Topic("persistent://pulsar-nzoz42zpr95v/itcoursee/itcoursee").Tag("test").Build().Build().Build()
	fmt.Println(client.Sender().Build().Send(context.Background(), user))
	fmt.Println(client.Handler().Build().Handle(context.Background(), new(Handler)))
	select {}
}
