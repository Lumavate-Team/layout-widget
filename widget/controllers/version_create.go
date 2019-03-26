package controllers

import (
  "fmt"
  common_controller "github.com/Lumavate-Team/lumavate-go-common/controllers"
  "github.com/bitly/go-simplejson"
  "encoding/json"
  "strings"
  "regexp"
	"os"
)

type VersionCreateController struct {
  common_controller.LumavateController
}

type ComponentStruct struct {
  Id string
}

type ResourcesStruct struct {
  Payload struct {
    Data struct {
      Pages [] struct {
        Id string
        Url string
      }
      Microservices [] struct {
        Id string
        Url string
      }
    }
  }
}

func (this *VersionCreateController) Post() {
  components := []ComponentStruct{}

  body, _ := simplejson.NewJson(this.Ctx.Input.RequestBody)
  body_items, _ := body.Get("bodyItems").Array()
  modal_items, _ := body.Get("modalItems").Array()
  logic_items, _ := body.Get("logicItems").Array()

  reg, _ := regexp.Compile("[^a-zA-Z0-9]+")

  for index, _ := range body_items {
    if id, ok := body.Get("bodyItems").GetIndex(index).Get("componentData").CheckGet("Id"); ok {
      id_val, _ := id.String()
      id_val = reg.ReplaceAllString(id_val, "_")
      components = append(components, ComponentStruct{id_val})
    }

    if id, ok := body.Get("bodyItems").GetIndex(index).Get("componentData").CheckGet("id"); ok {
      id_val, _ := id.String()
      id_val = reg.ReplaceAllString(id_val, "_")
      components = append(components, ComponentStruct{id_val})
    }
  }

  for index, _ := range modal_items {
    if id, ok := body.Get("modalItems").GetIndex(index).Get("componentData").CheckGet("Id"); ok {
      id_val, _ := id.String()
      id_val = reg.ReplaceAllString(id_val, "_")
      components = append(components, ComponentStruct{id_val})
    }

    if id, ok := body.Get("modalItems").GetIndex(index).Get("componentData").CheckGet("id"); ok {
      id_val, _ := id.String()
      id_val = reg.ReplaceAllString(id_val, "_")
      components = append(components, ComponentStruct{id_val})
    }
  }

  for index, _ := range logic_items {
    if id, ok := body.Get("logicItems").GetIndex(index).Get("componentData").CheckGet("Id"); ok {
      id_val, _ := id.String()
      id_val = reg.ReplaceAllString(id_val, "_")
      components = append(components, ComponentStruct{id_val})
    }

    if id, ok := body.Get("logicItems").GetIndex(index).Get("componentData").CheckGet("id"); ok {
      id_val, _ := id.String()
      id_val = reg.ReplaceAllString(id_val, "_")
      components = append(components, ComponentStruct{id_val})
    }
  }

  resource_body, _ := this.LumavateGet("/pwa/v1/resources")
  resources := ResourcesStruct{}
  json.Unmarshal(resource_body, &resources)
  fmt.Println(resources)

  pre_script := `
    var context = document.querySelector('luma-core-context');
    context.componentOnReady().then(function() {
      promises = [];
      promises.push(context.getToken());
%s
      Promise.all(promises).then( (values) => {
        token           = values.shift();
        auth_data       = context.authData;
        activation_data = context.activationData;
        domain_data     = context.domainData;
%s%s%s%s

      });
    });
  `

  begin_delim := `
        /*
        Please place your code beneath this comment
        */`

  end_delim := `
        /*
        Please place your code above this comment
        */`

  this.Data["json"] =body
	script := ""

	if os.Getenv("MODE") == "CSSGRID" {
		script, _ = body.Get("script").String()
	}
	if os.Getenv("MODE") == "KNOCKOUT" {
		script, _ = body.Get("viewModel").String()
	}

  parts := strings.Split(script, begin_delim)
  script = parts[len(parts)-1]

  parts = strings.Split(script, end_delim)
  script = parts[0]


  promises_push := ""
  for _, comp := range components {
    if comp.Id != "" {
      promises_push += fmt.Sprintf("      promises.push(context.getComponent('%s'));\n", comp.Id)
    }
  }
  for _, page := range resources.Payload.Data.Pages {
    promises_push += fmt.Sprintf("      promises.push(context.getWidget('%s'));\n", page.Id)
  }
  for _, microservice := range resources.Payload.Data.Microservices {
    promises_push += fmt.Sprintf("      promises.push(context.getMicroservice('%s'));\n", microservice.Id)
  }

  assignment := ""
  for _, comp := range components {
    if comp.Id != "" {
      assignment += fmt.Sprintf("        c_%-13s = values.shift(); \n", comp.Id)
    }
  }
  for _, page := range resources.Payload.Data.Pages {
    assignment += fmt.Sprintf("        p_%-13s = values.shift(); /* %-20s */\n", reg.ReplaceAllString(page.Id, "_"), page.Url)
  }
  for _, microservice := range resources.Payload.Data.Microservices {
    assignment += fmt.Sprintf("        m_%-13s = values.shift(); /* %-20s */\n", reg.ReplaceAllString(microservice.Id, "_"), microservice.Url)
  }

  script = fmt.Sprintf(pre_script, promises_push, assignment, begin_delim, script, end_delim)

	if os.Getenv("MODE") == "CSSGRID" {
		body.Set("script", script)
	}
	if os.Getenv("MODE") == "KNOCKOUT" {
		body.Set("viewModel", script)
	}
  this.ServeJSON()
}
