package controllers

import (
  common_controller "github.com/Lumavate-Team/lumavate-go-common"
	component_data "github.com/Lumavate-Team/lumavate-go-common/properties/component_data"
  "encoding/json"
  "widget/models"
	"os"
	"fmt"
	"strings"
  _"github.com/bitly/go-simplejson"
  "reflect"
	"strconv"
)

type MainController struct {
  common_controller.LumavateController
}

func (this *MainController) Post() {
	luma_response := models.LumavateRequest{}
	json.Unmarshal(this.LumavateGetData(), &luma_response)

	var register map[string]interface{}

	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &register); err != nil {
		fmt.Println(err)
	}

	b, _ := json.Marshal(&register)
	q := fmt.Sprintf("%v",
		luma_response.Payload.Data.FormAction)
	resp, status := this.LumavatePost(q, b, true)

	if status != "200" {
		fmt.Println("Post to person failed")
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

func (this *MainController) Get() {
  luma_response := models.LumavateRequest {}
  err := json.Unmarshal(this.LumavateGetData(), &luma_response)

  if err != nil {
    this.Abort("500")
  }

  luma_response.Payload.Data.NavBar.ComponentData.NavBarItems = luma_response.Payload.Data.NavBarItems

  this.LayoutSections["FormScript"] = ""

  for i, element := range luma_response.Payload.Data.GridItems {
      if reflect.TypeOf(element.Component).Elem().Name() == "FormStruct" {
				var tmpForm component_data.FormStruct
				tmpForm.FormItems = luma_response.Payload.Data.FormItems
				luma_response.Payload.Data.GridItems[i].Component = tmpForm
        this.LayoutSections["FormScript"] = "register_script.tpl"
      }
    }

  this.Data["data"] = luma_response.Payload.Data
	this.Data["dnsInfo"] = fmt.Sprintf("%s%s", os.Getenv("PROTO"), this.Ctx.Input.Host())

	this.Layout = "layout/layout.tpl"
	mode := this.GetString("mode")
	if strings.ToLower(mode) != "degraded" {
		this.TplName = "index.tpl"
	} else {
		this.TplName = "degraded.tpl"
	}

  this.LayoutSections["HtmlHead"] = "html_head.tpl"
  this.LayoutSections["Scripts"] = "script.tpl"
  this.LayoutSections["FooterContent"] = "home_footer.tpl"
}
