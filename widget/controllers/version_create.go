package controllers

import (
  "fmt"
  common_controller "github.com/Lumavate-Team/lumavate-go-common/controllers"
  "github.com/bitly/go-simplejson"
  "encoding/json"
  "strings"
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
  for index, _ := range body_items {
		if id, ok := body.Get("bodyItems").GetIndex(index).Get("componentData").CheckGet("Id"); ok {
			id_val, _ := id.String()
			components = append(components, ComponentStruct{id_val})
		}
  }

	resource_body, _ := this.LumavateGet("/pwa/v1/resources")
	resources := ResourcesStruct{}
  json.Unmarshal(resource_body, &resources)
	fmt.Println(resources)

  pre_script := `
    var lp = document.querySelector('luma-core-context');
    lp.componentOnReady().then(function() {
      promises = [];
      promises.push(lp.getToken());
%s
      console.log(promises);
      Promise.all(promises).then( (values) => {
        token        = values.shift();
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
  script, _ := body.Get("script").String()

  parts := strings.Split(script, begin_delim)
  script = parts[len(parts)-1]

  parts = strings.Split(script, end_delim)
  script = parts[0]


  promises_push := ""
  for _, comp := range components {
    promises_push += fmt.Sprintf("      promises.push(lp.getComponent('%s'));\n", comp.Id)
  }
  for _, page := range resources.Payload.Data.Pages {
    promises_push += fmt.Sprintf("      promises.push(lp.getComponent('%s'));\n", page.Id)
  }
  for _, microservice := range resources.Payload.Data.Microservices {
    promises_push += fmt.Sprintf("      promises.push(lp.getComponent('%s'));\n", microservice.Id)
  }

  assignment := ""
  for _, comp := range components {
    assignment += fmt.Sprintf("        c_%-10s = values.shift(); \n", comp.Id)
  }
  for _, page := range resources.Payload.Data.Pages {
    assignment += fmt.Sprintf("        p_%-10s = values.shift(); /* %-20s */\n", page.Id, page.Url)
  }
  for _, microservice := range resources.Payload.Data.Microservices {
    assignment += fmt.Sprintf("        m_%-10s = values.shift(); /* %-20s */\n", microservice.Id, microservice.Url)
  }

  script = fmt.Sprintf(pre_script, promises_push, assignment, begin_delim, script, end_delim)

  body.Set("script", script)
  this.ServeJSON()
}

