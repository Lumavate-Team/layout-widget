package controllers

import (
  common_controller "github.com/Lumavate-Team/lumavate-go-common/controllers"
  properties "github.com/Lumavate-Team/lumavate-go-common/properties"
  components "github.com/Lumavate-Team/lumavate-go-common/components"
  "encoding/json"
  "widget/models"
  "fmt"
  "os"
)

type PropertyController struct {
  common_controller.LumavateController
}

func (this *PropertyController) Get() {
  this.LoadAllComponentSets()
  this.Data["json"] = this.GetAllProperties()
  this.ServeJSON()
}

func (this *PropertyController) GetLogicProperties() []properties.PropertyType {
  placementOptions := make(map[string]string)
  placementOptions["top"] = "Top"
  placementOptions["bottom"] = "Bottom"

  props := []properties.PropertyType{}

  props = append(props, &properties.PropertyDropdown{
    &properties.PropertyBase{"placement", "", "Placement Settings", "Component Placement", help_component_placement}, "top", placementOptions})

  return props
}

func (this *PropertyController) GetLayoutProperties() []properties.PropertyType {
  props := []properties.PropertyType{}

  // Background Image Scaling Options
  displayOptions := make(map[string]string)
  displayOptions["both"] = "Both"
  displayOptions["optimal"] = "Optimal"
  displayOptions["degraded"] = "Degraded"

  vjustifyOptions := make(map[string]string)
  vjustifyOptions["start"] = "Top"
  vjustifyOptions["end"] = "Bottom"
  vjustifyOptions["center"] = "Center"
  vjustifyOptions["stretch"] = "Stretch"

  hjustifyOptions := make(map[string]string)
  hjustifyOptions["start"] = "Left"
  hjustifyOptions["end"] = "Right"
  hjustifyOptions["center"] = "Center"
  hjustifyOptions["stretch"] = "Stretch"

  props = append(props, &properties.PropertyDropdown{
    &properties.PropertyBase{"displayMode", "", "Placement Settings", "Display Mode", help_mode}, "both", displayOptions})
  props = append(props, &properties.PropertyNumeric{
    &properties.PropertyBase{"templateRowStart", "", "Placement Settings", "Body Row Start", "This is Row at which this grid item will start"}, 1, properties.PropertyOptionsNumeric{Min: 1, Max: 100}})
  props = append(props, &properties.PropertyNumeric{
    &properties.PropertyBase{"templateRowSpan", "", "Placement Settings", "Number of Rows to Span", "This is the Row at which this grid item will end"}, 1, properties.PropertyOptionsNumeric{Min: 1, Max: 100}})
  props = append(props, &properties.PropertyNumeric{
    &properties.PropertyBase{"templateColumnStart", "", "Placement Settings", "Body Column Start", "This is the Column at which the grid item will start"}, 1, properties.PropertyOptionsNumeric{Min: 1, Max: 100}})
  props = append(props, &properties.PropertyNumeric{
    &properties.PropertyBase{"templateColumnSpan", "", "Placement Settings", "Number of Columns to Span", "This is the Column at which the grid item will end"}, 1, properties.PropertyOptionsNumeric{Min: 1, Max: 100}})
  props = append(props, &properties.PropertyDropdown{
		&properties.PropertyBase{"horizontalAlignment", "", "Placement Settings", "Horizontal Alignment", ""}, "stretch", hjustifyOptions})
  props = append(props, &properties.PropertyDropdown{
		&properties.PropertyBase{"verticalAlignment", "", "Placement Settings", "Vertical Alignment", ""}, "stretch", vjustifyOptions})
  props = append(props, &properties.PropertyNumeric{
    &properties.PropertyBase{"paddingLeft", "", "Placement Settings", "Padding Left (px)", ""}, 0, properties.PropertyOptionsNumeric{Min: 0, Max: 200}})
  props = append(props, &properties.PropertyNumeric{
    &properties.PropertyBase{"paddingRight", "", "Placement Settings", "Padding Right (px)", ""}, 0, properties.PropertyOptionsNumeric{Min: 0, Max: 200}})
  props = append(props, &properties.PropertyNumeric{
    &properties.PropertyBase{"paddingTop", "", "Placement Settings", "Padding Top (px)", ""}, 0, properties.PropertyOptionsNumeric{Min: 0, Max: 200}})
  props = append(props, &properties.PropertyNumeric{
    &properties.PropertyBase{"paddingBottom", "", "Placement Settings", "Padding Bottom (px)", ""}, 0, properties.PropertyOptionsNumeric{Min: 0, Max: 200}})
  return props
}

func (this *PropertyController) GetAllProperties() []properties.PropertyType {

    props := []properties.PropertyType{}

    if os.Getenv("MODE") == "CSSGRID" {
      props = append(props,
        &properties.PropertyToggle{
          &properties.PropertyBase{"displayHeader", "Header", "Header Settings", "Display Header", ""}, false},
        this.GetDynamicComponentProperty("header", "header", "Header", "Header Settings", "Header Data", ""),
        &properties.PropertyToggle{
          &properties.PropertyBase{"displayFooter", "Footer", "Footer Settings", "Display Footer", ""}, false},
        this.GetDynamicComponentProperty("footer", "footer", "Footer", "Footer Settings", "Footer Data", ""),
      )

      props = append(props,
        &properties.PropertyColor{
          &properties.PropertyBase{"backgroundColor", "Body", "Body Settings", "Background Color", ""}, "#ffffff"})

      props = append(props,
        &properties.PropertyToggle{
          &properties.PropertyBase{"displayBackgroundImage", "Body", "Body Settings", "Display Background Image", ""}, false})

      props = append(props,
        &properties.PropertyImage{
          &properties.PropertyBase{"backgroundImage", "Body", "Body Settings", "Background Image", ""}})

      props = append(props, this.GetBodyProperties())

      props = append(props, this.GetBodyItems())

      props = append(props,
        &properties.PropertyCodeEditor{
          &properties.PropertyBase{"script", "Script", "Javascript", "On Pageload", ""}, ""})
    }

    if os.Getenv("MODE") == "KNOCKOUT" {
      props = append(props,
        &properties.PropertyCodeEditor{
          &properties.PropertyBase{"viewTemplate", "View", "Settings", "Template", ""}, ""})

      props = append(props, this.GetTemplateProperties())

      props = append(props,
        &properties.PropertyCodeEditor{
          &properties.PropertyBase{"viewModel", "View Model", "Settings", "View Model", ""}, ""})

      props = append(props, this.GetTranslationProperties())
      props = append(props, this.GetVariableProperties())
    }

    props = append(props, this.GetSecurityProperties())

    props = append(props, this.GetLogicItems())

    props = append(props, this.GetDynamicComponentsProperty("modal", "modalItems", "Modal", "Modal Items", "", ""))

    for _, element := range components.GetAddToHomeProperties() {
      props = append(props, element)
    }

  return props
}

func (this *PropertyController) GetSecurityProperties() *properties.PropertyComponent {
  token_data := this.ParseToken()
  propertyGroups := models.AuthGroupRequest {}
  if token_data.AuthUrl != "" {
    body, _ := this.LumavateGet(fmt.Sprintf("%vdiscover/auth-groups",token_data.AuthUrl))
    json.Unmarshal(body, &propertyGroups)
  }

  // defaults for auth group multiselect
  defaults := make([]string, 1)
  defaults = append(defaults, "")

  // loop through groups from auth-groups call and append them to multiselect property
  groups := [] properties.MultiselectOption {}
  for _, element := range propertyGroups.Payload.Data {
    groups = append(groups, properties.MultiselectOption{element.Group, element.Group})
  }

  props := [] properties.PropertyType {}
  c := &properties.Component{"securityType", "", "securityNone", "<None>", "a", "<None>", props, ""}

  props1 := [] properties.PropertyType {}
  c1 := &properties.Component{"securityType", "", "securityAll", "All logged in users", "a", "All logged in users", props1, ""}

  props2 := [] properties.PropertyType {}
  props2 = append(props2, &properties.PropertyPageLink{ &properties.PropertyBase{"noAuthRedirect", "", "Security Properties (Specific)", "No Auth Redirect", ""}})
  props2 = append(props2, &properties.PropertyMultiselect{ &properties.PropertyBase{"specificGroup", "", "Security Properties (Specific)", "Specific Group(s)", ""}, defaults, groups})
  c2 := &properties.Component{"securityType", "", "securitySpecific", "Specific user group(s)", "a", "Specific user group(s)", props2, ""}

  p := &properties.PropertyComponent{
    &properties.PropertyBase{"securityProperties", "Widget", "Security Settings", "Authentication", ""},
    c, &properties.PropertyOptionsComponent{[] string {"securityType"}, [] *properties.Component {c, c1, c2} },
  }

  return p
}

func (this *PropertyController) GetTemplateProperties() *properties.PropertyComponents {
  props := [] properties.PropertyType {}
  props = append(props, &properties.PropertyText{&properties.PropertyBase{"templateId", "", "", "Template Id", ""}, "", properties.PropertyOptionsText{}})
  props = append(props, &properties.PropertyCodeEditor{&properties.PropertyBase{"template", "", "", "Template", ""}, ""})
  c := &properties.Component{"templateType", "", "templateType", "Template", "a", "Template", props, "{{ componentData.templateId }} "}

  p := &properties.PropertyComponents{
    &properties.PropertyBase{"templates", "Templates", "Template Settings", "Templates", ""},
    []*properties.Component {}, &properties.PropertyOptionsComponent{[] string {"templateProperties"}, [] *properties.Component {c} },
  }

  return p
}

func (this *PropertyController) GetTranslationProperties() *properties.PropertyComponents {
  props := [] properties.PropertyType {}
  props = append(props, &properties.PropertyText{&properties.PropertyBase{"stringId", "", "", "String Id", ""}, "", properties.PropertyOptionsText{}})
  props = append(props, &properties.PropertyTranslatedText{&properties.PropertyBase{"string", "", "", "String", ""}, "", properties.PropertyOptionsText{}})
  c := &properties.Component{"translationType", "", "translationType", "Translation", "a", "Translation", props, "{{ componentData.stringId }} - {{ componentData.string['en-us'] | truncate(40) }} "}

  p := &properties.PropertyComponents{
    &properties.PropertyBase{"translations", "Translations", "Translation Settings", "Translations", ""},
    []*properties.Component {}, &properties.PropertyOptionsComponent{[] string {"translationProperties"}, [] *properties.Component {c} },
  }

  return p
}

func (this *PropertyController) GetVariableProperties() *properties.PropertyComponents {
  textprops := [] properties.PropertyType {}
  textprops = append(textprops, &properties.PropertyText{&properties.PropertyBase{"variableId", "", "", "Variable Id", ""}, "", properties.PropertyOptionsText{}})
  textprops = append(textprops, &properties.PropertyText{&properties.PropertyBase{"stringValue", "", "", "Variable", ""}, "", properties.PropertyOptionsText{}})
  tc := &properties.Component{"variableType", "", "stringVariableType", "String", "a", "String", textprops, "{{ componentData.variableId }} - {{ componentData.stringValue | truncate(40) }} "}

  intprops := [] properties.PropertyType {}
  intprops = append(intprops, &properties.PropertyText{&properties.PropertyBase{"variableId", "", "", "Variable Id", ""}, "", properties.PropertyOptionsText{}})
  intprops = append(intprops, &properties.PropertyNumeric{&properties.PropertyBase{"intValue", "", "", "Variable", ""}, 0, properties.PropertyOptionsNumeric{-2147483647, 2147483647}})

  ic := &properties.Component{"variableType", "", "intVariableType", "Integer", "a", "Integer", intprops, "{{ componentData.variableId }} - {{ componentData.intValue }} "}

  colorprops := [] properties.PropertyType {}
  colorprops = append(colorprops, &properties.PropertyText{&properties.PropertyBase{"variableId", "", "", "Variable Id", ""}, "", properties.PropertyOptionsText{}})
  colorprops = append(colorprops, &properties.PropertyColor{&properties.PropertyBase{"colorValue", "", "", "Variable", ""}, "#cccccc"})
  cc := &properties.Component{"variableType", "", "colorVariableType", "Color", "a", "Color", colorprops, "{{ componentData.variableId }} - {{ componentData.colorValue }} "}

  imageprops := [] properties.PropertyType {}
  imageprops = append(imageprops, &properties.PropertyText{&properties.PropertyBase{"variableId", "", "", "Variable Id", ""}, "", properties.PropertyOptionsText{}})
  imageprops = append(imageprops, &properties.PropertyImage{&properties.PropertyBase{"imageValue", "", "", "Variable", ""}})
  uc := &properties.Component{"variableType", "", "imageVariableType", "Image", "a", "Image", imageprops, "{{ componentData.variableId }}"}

  p := &properties.PropertyComponents{
    &properties.PropertyBase{"variables", "Variables", "Variable Settings", "Variables", ""},
    []*properties.Component {}, &properties.PropertyOptionsComponent{[] string {"variableProperties"}, [] *properties.Component {tc, ic, cc, uc} },
  }

  return p
}

func (this *PropertyController) GetLogicItems() *properties.PropertyComponents {
  p := this.GetDynamicComponentsProperty("logic", "logicItems", "Logic", "Logic Components", "", "")

  for _, component := range p.Options.Components {
    for _, property := range component.Properties {
      p := property.(map[string]interface{})
      if p["section"] ==  nil || p["section"] == ""  {
        p["section"] = component.Label + " Settings"
      }
    }
    component.Properties = append(component.Properties, this.GetLogicProperties()...)
  }
  return p
}

func (this *PropertyController) GetBodyItems() *properties.PropertyComponents {
  p := this.GetDynamicComponentsProperty("body", "bodyItems", "Body", "Body Items", "", "")
  for _, component := range p.Options.Components {
    for _, property := range component.Properties {
      p := property.(map[string]interface{})
      if p["section"] ==  nil || p["section"] == ""  {
        p["section"] = component.Label + " Settings"
      }
    }
    component.Properties = append(component.Properties, this.GetLayoutProperties()...)
  }
  return p
}
func (this *PropertyController) GetBodyProperties() *properties.PropertyComponent {
  props := [] properties.PropertyType {}

  props = append(props, &properties.PropertyNumeric{
		&properties.PropertyBase{"bodyNumRows", "", "Body Properties (Basic)", "Number Of Rows", ""}, 5, properties.PropertyOptionsNumeric{Min: 1, Max: 20}})
  props = append(props, &properties.PropertyNumeric{
		&properties.PropertyBase{"bodyNumColumns", "", "Body Properties (Basic)", "Number Of Columns", ""}, 1, properties.PropertyOptionsNumeric{Min: 1, Max: 20}})
  props = append(props, &properties.PropertyNumeric{
		&properties.PropertyBase{"bodyMaxWidth", "", "Body Properties (Basic)", "Max Width (pixels)", ""}, 0, properties.PropertyOptionsNumeric{Min: 0, Max: 10000}})
  props = append(props, &properties.PropertyNumeric{
    &properties.PropertyBase{"paddingLeft", "", "Body Properties (Basic)", "Padding Left (px)", ""}, 0, properties.PropertyOptionsNumeric{Min: 0, Max: 200}})
  props = append(props, &properties.PropertyNumeric{
    &properties.PropertyBase{"paddingRight", "", "Body Properties (Basic)", "Padding Right (px)", ""}, 0, properties.PropertyOptionsNumeric{Min: 0, Max: 200}})
  props = append(props, &properties.PropertyNumeric{
    &properties.PropertyBase{"paddingTop", "", "Body Properties (Basic)", "Padding Top (px)", ""}, 0, properties.PropertyOptionsNumeric{Min: 0, Max: 200}})
  props = append(props, &properties.PropertyNumeric{
    &properties.PropertyBase{"paddingBottom", "", "Body Properties (Basic)", "Padding Bottom (px)", ""}, 0, properties.PropertyOptionsNumeric{Min: 0, Max: 200}})

  c := &properties.Component{"body-items", "", "body-items-basic", "Basic", "a", "Basic", props, ""}

  props2 := [] properties.PropertyType {}

  props2 = append(props2, &properties.PropertyText{
		&properties.PropertyBase{"bodyTemplateRows", "", "Body Properties (Advanced)", "Body Row Template", help_row_template}, "", properties.PropertyOptionsText{}})
  props2 = append(props2, &properties.PropertyText{
		&properties.PropertyBase{"bodyTemplateColumns", "", "Body Properties (Advanced)", "Body Column Template", help_column_template}, "", properties.PropertyOptionsText{}})
  props2 = append(props2, &properties.PropertyText{
		&properties.PropertyBase{"bodyRowGap", "", "Body Properties (Advanced)", "Body Row Gap", "This sets the size of the gap (gutter) between the grid rows"}, "", properties.PropertyOptionsText{}})
  props2 = append(props2, &properties.PropertyText{
		&properties.PropertyBase{"bodyColumnGap", "Body Properties (Advanced)", "Body Properties (Advanced)", "Body Column Gap", "This sets the size of the gap (gutter) between the grid columns"}, "", properties.PropertyOptionsText{}})
  props2 = append(props2, &properties.PropertyNumeric{
    &properties.PropertyBase{"paddingLeft", "", "Body Properties (Advanced)", "Padding Left (px)", ""}, 0, properties.PropertyOptionsNumeric{Min: 0, Max: 200}})
  props2 = append(props2, &properties.PropertyNumeric{
    &properties.PropertyBase{"paddingRight", "", "Body Properties (Advanced)", "Padding Right (px)", ""}, 0, properties.PropertyOptionsNumeric{Min: 0, Max: 200}})
  props2 = append(props2, &properties.PropertyNumeric{
    &properties.PropertyBase{"paddingTop", "", "Body Properties (Advanced)", "Padding Top (px)", ""}, 0, properties.PropertyOptionsNumeric{Min: 0, Max: 200}})
  props2 = append(props2, &properties.PropertyNumeric{
    &properties.PropertyBase{"paddingBottom", "", "Body Properties (Advanced)", "Padding Bottom (px)", ""}, 0, properties.PropertyOptionsNumeric{Min: 0, Max: 200}})

  c2 := &properties.Component{"body-items", "", "body-items-advanced", "Advanced", "a", "Advanced", props2, ""}

  p := &properties.PropertyComponent{
    &properties.PropertyBase{"bodyProperties", "Body", "Body Settings", "Body Style", ""},
    c, &properties.PropertyOptionsComponent{[] string {"body-items"}, [] *properties.Component {c, c2} },
  }

  return p
}

/*
* Help Text Globals
 */

var help_component_placement string = `Indicates whether this logic component should be placed above or below the layout grid`

var help_mode string = `Denotes when this item should be displayed: * Both:
Display during _optimal_ & _degraded_ rendering (default) * Optimal: Display
during _optimal_ rendering on newer browsers supporting CSS Grid * Degraded:
Display only during _degraded_ rendering (browsers that do **not** support
CSS Grid)`

var help_row_template string = `Denotes the number of Rows in the grid.  This can be denoted by the following:
- Pixels(px): Defines the row in static pixel amount
- Percentage(%): Defines the row in terms of percentage of screen height
- Fractional Units(fr): Defines the row in terms of fractional units of the screen height

**_Scroll down for links to additional resources_**
### Example

#### Define 4 rows:
The first is 25px tall, row 2 is 10% of the total screen height, and rows 3 & 4 use the remaining height and creates the third row that is twice the height of the fourth row

#### Correct Setting:
25px 10% 2fr 1fr

#### Layout is based on CSS Grid.
Learn more about CSS Grid here: <a href="https://www.w3schools.com/css/css_grid.asp" target="_blank">W3Schools</a> <a href="https://cssgridgarden.com/" target="_blank">CSS Grid Garden</a>`

var help_column_template string = `Denotes the number of Columns in the grid.  This can be denoted by the following:
* Pixels(px): Defines the column in static pixel amount
* Percentage(%): Defines the column in terms of percentage of screen height
* Fractional Units(fr): Defines the column in terms of fractional units of the screen width

**_Scroll down for links to additional resources_**
### Example

#### Define 5 columns:
The first & fifth columns are 25px wide, columns 2,3, & 4 use the remaining width and creates the third column that is twice the width of the second and fourth column, which are equivalent in size

#### Correct Setting:
25px 1fr 2fr 1fr 25px

### Layout is based on CSS Grid.
Learn more about CSS Grid here: <a href="https://www.w3schools.com/css/css_grid.asp" target="_blank">W3Schools</a> <a href="https://cssgridgarden.com/" target="_blank">CSS Grid Garden</a>`
