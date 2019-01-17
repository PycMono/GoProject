package impl

import (
	"fmt"
	"moqikaka.com/LoginServer/src/api"
	"moqikaka.com/LoginServer/src/bll/partnerBll"
	"testing"
)

var (
	tt = `{"token":"1d218734689c83bde922055087a2ccca","ExtraData":"{\"access_token\":\"c580531b8382bea88ef402f3aa991e0a\"}"}`
)

func TestName(t *testing.T) {
	partnerDB := partnerBll.GetItem(1005)
	fmt.Println(partnerDB.PartnerID)
	//_ = NewLoginHandler_mqhy(partnerDB, tt)
	api.CallOne(partnerDB, tt)
}
