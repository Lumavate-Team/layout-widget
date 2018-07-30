package controllers

import (
  common_controller "github.com/Lumavate-Team/lumavate-go-common"
)

type SingleUseTokenStruct struct {
  Token string `json:"token"`
}

type SingleUseTokenController struct {
  common_controller.LumavateController
}

func (this *SingleUseTokenController) Post() {
  lr := this.GetRequest()
  token_obj, code := lr.GetSingleUseToken()

  if code == 200 {
    this.Ctx.Output.SetStatus(200)
    result := SingleUseTokenStruct{ token_obj.Payload.Data.Token }
    this.Data["json"] = &result
    this.ServeJSON()
  } else if code == 401 {
    this.Ctx.Output.SetStatus(401)
    result := SingleUseTokenStruct{ "" }
    this.Data["json"] = &result
    this.ServeJSON()
  } else if code == 403 {
    this.Ctx.Output.SetStatus(403)
    result := SingleUseTokenStruct{ "" }
    this.Data["json"] = &result
    this.ServeJSON()
  } else {
    this.Ctx.Output.SetStatus(500)
    result := SingleUseTokenStruct{ "" }
    this.Data["json"] = &result
    this.ServeJSON()
  }
}
