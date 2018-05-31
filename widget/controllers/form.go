package controllers

import (
	common_controller "github.com/Lumavate-Team/lumavate-go-common"
  	_"github.com/Lumavate-Team/lumavate-go-common/properties/component_data"
	"widget/models"
	"encoding/json"
	"fmt"
	"strconv"
)


type FormController struct {
	common_controller.LumavateController
}

func (this *FormController) Post() {
	luma_response := models.LumavateRequest{}
	json.Unmarshal(this.LumavateGetData(), &luma_response)

	var register map[string]interface{}

	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &register); err != nil {
		fmt.Println(err)
	}

	fmt.Println(register)

	b, _ := json.Marshal(&register)
	q := fmt.Sprintf("%v/%v",
		luma_response.Payload.Data.FormAction,
		"persons")
	resp, status := this.LumavatePost(q, b, true)

	fmt.Println(q)

	fmt.Println(status)
	if status != "200" {
		var json_resp map[string]interface{}
		json.Unmarshal(resp, &json_resp)

		if err := json.Unmarshal(resp, &json_resp); err != nil {
			fmt.Println(err)
		}

		this.Data["json"] = json_resp
		code, _ := strconv.Atoi(status)
		this.Ctx.Output.SetStatus(code)
	} else {
		this.Ctx.Output.SetStatus(204)
	}
	this.ServeJSON()

}
