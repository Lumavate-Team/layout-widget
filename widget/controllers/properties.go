package controllers

import (
	"encoding/json"
	"fmt"
	_ "github.com/Lumavate-Team/lumavate-go-common/components"
	common_controller "github.com/Lumavate-Team/lumavate-go-common/controllers"
	properties "github.com/Lumavate-Team/lumavate-go-common/properties"
	_ "os"
	"widget/models"
)

type PropertyController struct {
	common_controller.LumavateController
}

func (this *PropertyController) Get() {
	this.LoadAllComponentSets()
	this.Data["json"] = this.GetAllProperties()
	this.ServeJSON()
}
func (this *PropertyController) GetAllProperties() []properties.PropertyType {

	props := []properties.PropertyType{}

	props = append(props,
		&properties.PropertyCodeEditor{
			&properties.PropertyBase{"viewTemplate", "View", "Settings", "Template", ""}, ""})

	props = append(props,
		&properties.PropertyCodeEditor{
			&properties.PropertyBase{"viewModel", "View Model", "Settings", "View Model", ""}, ""})

	props = append(props, this.GetSecurityProperties())

	return props
}

func (this *PropertyController) GetSecurityProperties() *properties.PropertyComponent {
	token_data := this.ParseToken()
	propertyGroups := models.AuthGroupRequest{}
	if token_data.AuthUrl != "" {
		body, _ := this.LumavateGet(fmt.Sprintf("%vdiscover/auth-groups", token_data.AuthUrl))
		json.Unmarshal(body, &propertyGroups)
	}

	// defaults for auth group multiselect
	defaults := make([]string, 1)
	defaults = append(defaults, "")

	// loop through groups from auth-groups call and append them to multiselect property
	groups := []properties.MultiselectOption{}
	for _, element := range propertyGroups.Payload.Data {
		groups = append(groups, properties.MultiselectOption{element.Group, element.Group})
	}

	props := []properties.PropertyType{}
	c := &properties.Component{"securityType", "", "securityNone", "<None>", "a", "<None>", props, ""}

	props1 := []properties.PropertyType{}
	c1 := &properties.Component{"securityType", "", "securityAll", "All logged in users", "a", "All logged in users", props1, ""}

	props2 := []properties.PropertyType{}
	props2 = append(props2, &properties.PropertyPageLink{&properties.PropertyBase{"noAuthRedirect", "", "Security Properties (Specific)", "No Auth Redirect", ""}})
	props2 = append(props2, &properties.PropertyMultiselect{&properties.PropertyBase{"specificGroup", "", "Security Properties (Specific)", "Specific Group(s)", ""}, defaults, groups})
	c2 := &properties.Component{"securityType", "", "securitySpecific", "Specific user group(s)", "a", "Specific user group(s)", props2, ""}

	p := &properties.PropertyComponent{
		&properties.PropertyBase{"securityProperties", "Widget", "Security Settings", "Authentication", ""},
		c, &properties.PropertyOptionsComponent{[]string{"securityType"}, []*properties.Component{c, c1, c2}},
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
