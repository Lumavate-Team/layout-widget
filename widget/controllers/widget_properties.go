package controllers

import (
	properties "github.com/Lumavate-Team/lumavate-go-common/properties"
)

type lumavateProperties struct {
	*properties.LumavateProperties
}

func (self *lumavateProperties) GetLayoutProperties() []properties.PropertyType {
	props := []properties.PropertyType{}

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

	props = append(props, &properties.PropertyText{
		&properties.PropertyBase{"cssClass", "", "", "CSS Class", "Denotes the class (as defined in the Layout CSS) that will be added to the styling of this item."}, "", properties.PropertyOptionsText{}})
	props = append(props, &properties.PropertyDropdown{
		&properties.PropertyBase{"displayMode", "", "", "Display Mode", help_mode}, "both", displayOptions})
	props = append(props, &properties.PropertyNumeric{
		&properties.PropertyBase{"templateRowStart", "", "", "Body Row Start", "This is Row at which this grid item will start"}, 1, properties.PropertyOptionsNumeric{Min: 1, Max: 100}})
	props = append(props, &properties.PropertyNumeric{
		&properties.PropertyBase{"templateRowSpan", "", "", "Number of Rows to Span", "This is the Row at which this grid item will end"}, 1, properties.PropertyOptionsNumeric{Min: 1, Max: 100}})
	props = append(props, &properties.PropertyNumeric{
		&properties.PropertyBase{"templateColumnStart", "", "", "Body Column Start", "This is the Column at which the grid item will start"}, 1, properties.PropertyOptionsNumeric{Min: 1, Max: 100}})
	props = append(props, &properties.PropertyNumeric{
		&properties.PropertyBase{"templateColumnSpan", "", "", "Number of Columns to Span", "This is the Column at which the grid item will end"}, 1, properties.PropertyOptionsNumeric{Min: 1, Max: 100}})
	props = append(props, &properties.PropertyDropdown{
		&properties.PropertyBase{"alignSelf", "", "", "Row justification", "Position of Component in Grid row axis"}, "stretch", justifyOptions})
	props = append(props, &properties.PropertyDropdown{
		&properties.PropertyBase{"justifySelf", "", "", "Column Justification", "Position of Component in Grid Along column axis"}, "stretch", justifyOptions})
	return props
}

func (self *lumavateProperties) GetAllProperties() []properties.PropertyType {

	return []properties.PropertyType{
		&properties.PropertyToggle{
			&properties.PropertyBase{"displayHeader", "Header", "Settings", "Display Header", ""}, false},
		self.DynamicComponents.GetDynamicComponentProperty("header", "header", "Header", "Header Settings", "Header Data", ""),
		&properties.PropertyToggle{
			&properties.PropertyBase{"displayFooter", "Footer", "Settings", "Display Footer", ""}, false},
		self.DynamicComponents.GetDynamicComponentProperty("footer", "footer", "Footer", "Footer Settings", "Footer Data", ""),
		&properties.PropertyColor{
			&properties.PropertyBase{"backgroundColor", "General", "Settings", "Background Color", ""}, "#ffffff"},
		&properties.PropertyToggle{
			&properties.PropertyBase{"displayBackgroundImage", "General", "Settings", "Display Background Image", ""}, false},
		&properties.PropertyImage{
			&properties.PropertyBase{"backgroundImage", "General", "Settings", "Background Image", ""}},
		self.GetBodyProperties(),
		self.GetBodyItems(),
		self.DynamicComponents.GetDynamicComponentsProperty("modal", "modalItems", "Modal", "Modal Items", "Modal Items", ""),
	}
}

func (self *lumavateProperties) GetBodyItems() *properties.PropertyComponents {
	p := self.DynamicComponents.GetDynamicComponentsProperty("body", "bodyItems", "Body", "Body Items", "Body Items", "")
	for _, component := range p.Options.Components {
		component.Properties = append(component.Properties, self.GetLayoutProperties()...)
	}
	return p
}

func (self *lumavateProperties) GetBodyProperties() *properties.PropertyComponent {
	justifyOptions := make(map[string]string)
	justifyOptions["start"] = "Start"
	justifyOptions["end"] = "End"
	justifyOptions["center"] = "Center"
	justifyOptions["stretch"] = "Stretch"
	justifyOptions["space-around"] = "Space Around"
	justifyOptions["space-between"] = "Space Between"
	justifyOptions["space-evenly"] = "Space Evenly"

	props := [] properties.PropertyType {}

	props = append(props, &properties.PropertyNumeric{&properties.PropertyBase{"bodyNumRows", "", "Body Properties (Basic)", "Number Of Rows", ""}, 5, properties.PropertyOptionsNumeric{Min: 1, Max: 20}})
	props = append(props, &properties.PropertyNumeric{&properties.PropertyBase{"bodyNumColumns", "", "Body Properties (Basic)", "Number Of Columns", ""}, 5, properties.PropertyOptionsNumeric{Min: 1, Max: 20}})
	props = append(props, &properties.PropertyNumeric{&properties.PropertyBase{"bodyMaxWidth", "", "Body Properties (Basic)", "Max Width (pixels)", ""}, 0, properties.PropertyOptionsNumeric{Min: 0, Max: 10000}})

	c := &properties.Component{"body-items", "", "body-items-basic", "Basic", "a", "Basic", props}

	props2 := [] properties.PropertyType {}

	props2 = append(props2, &properties.PropertyText{&properties.PropertyBase{"bodyTemplateRows", "", "Body Properties (Advanced)", "Body Row Template", help_row_template}, "", properties.PropertyOptionsText{}})
	props2 = append(props2, &properties.PropertyText{&properties.PropertyBase{"bodyTemplateColumns", "", "Body Properties (Advanced)", "Body Column Template", help_column_template}, "", properties.PropertyOptionsText{}})
	props2 = append(props2, &properties.PropertyText{&properties.PropertyBase{"bodyRowGap", "", "Body Properties (Advanced)", "Body Row Gap", "This sets the size of the gap (gutter) between the grid rows"}, "", properties.PropertyOptionsText{}})
	props2 = append(props2, &properties.PropertyText{&properties.PropertyBase{"bodyColumnGap", "", "Body Properties (Advanced)", "Body Column Gap", "This sets the size of the gap (gutter) between the grid columns"}, "", properties.PropertyOptionsText{}})
	props2 = append(props2, &properties.PropertyDropdown{&properties.PropertyBase{"justifyContent", "", "Body Properties (Advanced)", "Body Row Alignment", "This property aligns the grid along the row axis"}, "start", justifyOptions})
	props2 = append(props2, &properties.PropertyDropdown{&properties.PropertyBase{"alignContent", "", "Body Properties (Advanced)", "Body Column Alignment", "This property aligns the grid along the column axis"}, "start", justifyOptions})

	c2 := &properties.Component{"body-items", "", "body-items-advanced", "Advanced", "a", "Advanced", props2}

	p := &properties.PropertyComponent{
		&properties.PropertyBase{"bodyProperties", "Body", "", "Body Style", ""},
		c, &properties.PropertyOptionsComponent{[] string {"body-items"}, [] *properties.Component {c, c2} },
	}

	return p
}

/*
* Help Text Globals
 */
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
