package buisnesslogic

import (
	"fmt"
	"testing"

	"github.com/MarcBernstein0/gamedeals/api/model"
)

func TestReturnResponse(t *testing.T) {
	testReq := model.Request{
		Title:  "string",
		System: "PC",
	}
	resp, err := FindGameDeals(testReq)
	if err != nil {
		t.Errorf("Error when calling FindGameDeals %v\n", err)
	}
	fmt.Printf("%+v\n", resp)
	if resp.Title != testReq.Title {
		t.Errorf("Not returning right game title: exptected %v got %v\n", testReq.Title, resp.Title)
	}

}
