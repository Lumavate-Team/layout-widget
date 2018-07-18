package controllers

import (
  "github.com/Lumavate-Team/lumavate-go-common/properties"
  "github.com/Lumavate-Team/lumavate-go-common"
)

func (lp *LumavateProperties) GetLayoutProperties() [] properties.PropertyType {
  props := [] properties.PropertyType {}

  // Background Image Scaling Options
  displayOptions := make(map[string]string)
  displayOptions["both"] = "Both"
  displayOptions["optimal"] = "Optimal"
  displayOptions["degraded"] = "Degraded"

  justifyOptions := make(map[string]string)
  justifyOptions["start"] = "Start"
  justifyOptions["end"] = "End"
  justifyOptions["center"] = "Center"
  justifyOptions["stretch"] = "Stretch"

  var displayHelp string = `Denotes when this item should be displayed: * Both: Display during _optimal_ & _degraded_ rendering (default) * Optimal: Display during _optimal_ rendering on newer browsers supporting CSS Grid * Degraded: Display only during _degraded_ rendering (browsers that do **not** support CSS Grid)`

  props = append(props, &properties.PropertyText{
    &properties.PropertyBase{"cssClass", "", "", "CSS Class", "Denotes the class (as defined in the Layout CSS) that will be added to the styling of this item."}, "", properties.PropertyOptionsText{}})
  props = append(props, &properties.PropertyDropdown{
    &properties.PropertyBase{"displayMode", "", "", "Display Mode", displayHelp}, "both", displayOptions})
  props = append(props, &properties.PropertyText{
    &properties.PropertyBase{"templateRowStart", "", "", "Body Row Start", "This is Row at which this grid item will start"}, "", properties.PropertyOptionsText{Rows: 3}})
  props = append(props, &properties.PropertyText{
    &properties.PropertyBase{"templateRowEnd", "", "", "Body Row End", "This is the Row at which this grid item will end"}, "", properties.PropertyOptionsText{ Rows: 3 }})
  props = append(props, &properties.PropertyText{
    &properties.PropertyBase{"templateColumnStart", "", "", "Body Column Start", "This is the Column at which the grid item will start"}, "", properties.PropertyOptionsText{Rows: 3}})
  props = append(props, &properties.PropertyText{
    &properties.PropertyBase{"templateColumnEnd", "", "", "Body Column End", "This is the Column at which the grid item will end"}, "", properties.PropertyOptionsText{Rows: 3}})
  props = append(props, &properties.PropertyDropdown{
    &properties.PropertyBase{"alignSelf", "", "", "Row justification", "Position of Component in Grid row axis"}, "stretch", justifyOptions})
  props = append(props, &properties.PropertyDropdown{
    &properties.PropertyBase{"justifySelf", "", "", "Column Justification", "Position of Component in Grid Along column axis"}, "stretch", justifyOptions})
  return props
}

/*
 * Returns all properties for the widget
 */
func (lp *LumavateProperties) GetAllProperties() [] properties.PropertyType {
  var rowhelp string = `Denotes the number of Rows in the grid.  This can be denoted by the following:
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

  var colhelp string = `Denotes the number of Columns in the grid.  This can be denoted by the following:
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

  // Background Image Scaling Options
  justifyOptions := make(map[string]string)
  justifyOptions["start"] = "Start"
  justifyOptions["end"] = "End"
  justifyOptions["center"] = "Center"
  justifyOptions["stretch"] = "Stretch"
  justifyOptions["space-around"] = "Space Around"
  justifyOptions["space-between"] = "Space Between"
  justifyOptions["space-evenly"] = "Space Evenly"


  return [] properties.PropertyType {
    &properties.PropertyToggle{
      &properties.PropertyBase{"displayHeader", "Header", "Settings", "Display Header", ""}, false},
      lp.GetHeaderProperty(),
    &properties.PropertyToggle{
      &properties.PropertyBase{"displayFooter", "Footer", "Settings", "Display Footer", ""}, false},
    lp.GetFooterProperty(),
    &properties.PropertyColor{
      &properties.PropertyBase{"backgroundColor", "General", "Settings", "Background Color", ""}, "#ffffff"},
    &properties.PropertyToggle{
      &properties.PropertyBase{"displayBackgroundImage", "General", "Settings", "Display Background Image", ""}, false},
    &properties.PropertyImage{
      &properties.PropertyBase{"backgroundImage", "General", "Settings", "Background Image", ""}},
    &properties.PropertyText{
      &properties.PropertyBase{"bodyTemplateRows", "Body", "Body Layout", "Body Row Template", rowhelp}, "", properties.PropertyOptionsText{Rows: 3}},
    &properties.PropertyText{
      &properties.PropertyBase{"bodyTemplateColumns", "Body", "Body Layout", "Body Column Template", colhelp}, "", properties.PropertyOptionsText{Rows: 3}},
    &properties.PropertyText{
      &properties.PropertyBase{"bodyRowGap", "Body", "Body Layout", "Body Row Gap", rowhelp}, "", properties.PropertyOptionsText{Rows: 3}},
    &properties.PropertyText{
      &properties.PropertyBase{"bodyColumnGap", "Body", "Body Layout", "Body Column Gap", colhelp}, "", properties.PropertyOptionsText{Rows: 3}},
    &properties.PropertyDropdown{
      &properties.PropertyBase{"justifyContent", "Body", "Body Layout", "Body Row Alignment", "This property aligns the grid along the row axis"}, "start", justifyOptions},
    &properties.PropertyDropdown{
      &properties.PropertyBase{"alignContent", "Body", "Body Layout", "Body Column Alignment", "This property aligns the grid along the column axis"}, "start", justifyOptions},
   lp.GetBodyItems(),
   lp.GetModalItems(),
  }
}

/*
* Returns all components for the widget
*/
func (lp *LumavateProperties) GetAllComponents() [] *properties.Component {
  return [] *properties.Component {
  }
}

func NewLumavateProperties(auth string) *LumavateProperties {
  lp := &LumavateProperties {auth, common.DynamicComponents{}}
  lp.DynamicComponents.LoadAllComponentSets(lp.Authorization)
  return lp
}

type LumavateProperties struct {
  Authorization string
  DynamicComponents common.DynamicComponents
}

func (self *LumavateProperties) GetModalItems() *properties.PropertyComponents {
  return self.DynamicComponents.GetDynamicComponentsProperty("modal", "modalItems", "Modal", "Modal Items", "Modal Items", "")
}


func (self *LumavateProperties) GetBodyItems() *properties.PropertyComponents {
  p := self.DynamicComponents.GetDynamicComponentsProperty("body", "bodyItems", "Body", "Body Items", "Body Items", "")
  for _, component := range p.Options.Components {
    component.Properties = append(component.Properties,self.GetLayoutProperties()...)
  }
  return p
}

func (self *LumavateProperties) GetFooterProperty() *properties.PropertyComponent {
  return self.DynamicComponents.GetDynamicComponentProperty("footer", "footer", "Footer", "Footer Settings", "Footer Data", "")
}

func (self *LumavateProperties) GetHeaderProperty() *properties.PropertyComponent {
  return self.DynamicComponents.GetDynamicComponentProperty("header", "header", "Header", "Header Settings", "Header Data", "")
}
