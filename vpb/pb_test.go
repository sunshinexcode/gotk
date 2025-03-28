package vpb_test

import (
	"testing"

	"github.com/sunshinexcode/gotk/vbase64"
	"github.com/sunshinexcode/gotk/vpb"
	"github.com/sunshinexcode/gotk/vpb/test"
	"github.com/sunshinexcode/gotk/vtest"
)

func TestMarshal(t *testing.T) {
	request := &test.Request{
		AppId:     "970CA35de60c44645bbae8a215061b33",
		RequestId: "test",
	}

	b, err := vpb.Marshal(vpb.Message(request))

	vtest.Nil(t, err)
	vtest.Equal(t, "CiA5NzBDQTM1ZGU2MGM0NDY0NWJiYWU4YTIxNTA2MWIzMxIEdGVzdA==", vbase64.EncodeToStr(b))
}

func TestUnmarshal(t *testing.T) {
	reply := &test.Reply{
		Code: "0",
		Msg:  "test",
	}
	b, err := vpb.Marshal(vpb.Message(reply))

	vtest.Nil(t, err)
	vtest.Equal(t, "CgEwEgR0ZXN0", vbase64.EncodeToStr(b))

	resp := &test.Reply{}
	err = vpb.Unmarshal(b, resp)

	vtest.Nil(t, err)
	vtest.Equal(t, "0", resp.Code)
	vtest.Equal(t, "test", resp.Msg)
}
