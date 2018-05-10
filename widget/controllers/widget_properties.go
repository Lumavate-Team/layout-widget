package controllers

import (
  properties "github.com/Lumavate-Team/lumavate-go-common/properties"
  components "github.com/Lumavate-Team/lumavate-go-common/components"
  "os"
  "fmt"
)

type LumavateProperties struct {
}

func (lp *LumavateProperties) GetLayoutProperties() [] properties.PropertyType {
  props := [] properties.PropertyType {}

	// Background Image Scaling Options
	displayOptions := make(map[string]string)
	displayOptions["both"] = "Both"
	displayOptions["optimal"] = "Optimal"
	displayOptions["degraded"] = "Degraded"

  props = append(props, &properties.PropertyDropdown{
		&properties.PropertyBase{"displayMode", "", "", "Display Mode", "When should this tiem be displayed: Only upon a degraded experience (due to old browsers), Only when fully optimzied, or display at all times (Both)"}, "both", displayOptions})
  props = append(props, &properties.PropertyText{
		&properties.PropertyBase{"templateRowStart", "", "", "Grid Row Start", "This is Row at which this grid item will start"}, "", properties.PropertyOptionsText{Rows: 3}})
  props = append(props, &properties.PropertyText{
		&properties.PropertyBase{"templateRowEnd", "", "", "Grid Row End", ""}, "This is the Row at which this grid item will end", properties.PropertyOptionsText{ Rows: 3 }})
  props = append(props, &properties.PropertyText{
		&properties.PropertyBase{"templateColumnStart", "", "", "Grid Column Start", "This is the Column at which the grid item will start"}, "", properties.PropertyOptionsText{Rows: 3}})
  props = append(props, &properties.PropertyText{
		&properties.PropertyBase{"templateColumnEnd", "", "", "Grid Column End", ""}, "This is the Column at which the grid item will end", properties.PropertyOptionsText{Rows: 3}})
	return props
}

func (lp *LumavateProperties) GetGridItemsProperty() *properties.PropertyComponents {
  return &properties.PropertyComponents {
    &properties.PropertyBase{"gridItems", "Grid", "Grid Items", "Grid Items", ""},
    [] *properties.Component{}, properties.PropertyOptionsComponent{[] string {"navigation", "video", "text"}, [] *properties.Component {lp.GetNavigationComponent(), lp.GetVideoComponent(), lp.GetTextComponent()} },
  }
}

func (lp *LumavateProperties) GetNavigationComponent() *properties.Component {
  props := [] properties.PropertyType {}

  props = append(props, &properties.PropertyImage{
    &properties.PropertyBase{"image", "", "", "Background Image", ""}})

	// Background Image Scaling Options
	options := make(map[string]string)
	options["fill"] = "Fill"
	options["fit"] = "Fit"
	options["stretch"] = "Stretch"
	options["repeat"] = "Repeat"
  props = append(props, &properties.PropertyDropdown{
		&properties.PropertyBase{"imageScaling", "", "", "Background Image Scaling", "Denotes how the image will appear as the grid item background"}, "fill",options})

  props = append(props, &properties.PropertyColor{
    &properties.PropertyBase{"backgroundColor", "", "", "Background Color", ""}, "#ffffff"})

  props = append(props, &properties.PropertyTranslatedText{
    &properties.PropertyBase{"title", "", "", "Title", ""}, "", properties.PropertyOptionsText{}})

  props = append(props, &properties.PropertyPageLink{
    &properties.PropertyBase{"pageLink", "", "", "Page URL", ""}})

	props = append(props, lp.GetLayoutProperties()...)
	image := fmt.Sprintf("%v%vstatic/images/navigation.svg", os.Getenv("BASE_URL"), os.Getenv("WIDGET_URL_PREFIX"))
  return &properties.Component{"navigation", "", "navigation", "Navigation", image, "Navigation", props}
}

func (lp *LumavateProperties) GetVideoComponent() *properties.Component {
  props := [] properties.PropertyType {}
  props = append(props, &properties.PropertyTranslatedText{
    &properties.PropertyBase{"title", "", "", "Title", ""}, "", properties.PropertyOptionsText{}})
  props = append(props, &properties.PropertyText{
		&properties.PropertyBase{"video", "", "", "Video URL", "This should be in the format: https://www.youtube.com/embed/[VIDEOID]"}, "https://www.youtube.com/embed/VIDEO_ID", properties.PropertyOptionsText{}})
	props = append(props, lp.GetLayoutProperties()...)
	image := fmt.Sprintf("%v%vstatic/images/video.svg", os.Getenv("BASE_URL"), os.Getenv("WIDGET_URL_PREFIX"))
  return &properties.Component{"video", "", "video", "Video", image, "Video", props}
}

func (lp *LumavateProperties) GetTextComponent() *properties.Component {
  props := [] properties.PropertyType {}
  props = append(props, &properties.PropertyTranslatedText{
    &properties.PropertyBase{"title", "", "", "Title", ""}, "", properties.PropertyOptionsText{}})
  props = append(props, &properties.PropertyText{
		&properties.PropertyBase{"text", "", "", "Text", ""}, "", properties.PropertyOptionsText{}})
	props = append(props, lp.GetLayoutProperties()...)
	image := fmt.Sprintf("%v%vstatic/images/video.svg", os.Getenv("BASE_URL"), os.Getenv("WIDGET_URL_PREFIX"))
  return &properties.Component{"text", "", "text", "Text", image, "Text", props}
}

/*
 * Returns all properties for the widget
 */
func (lp *LumavateProperties) GetAllProperties() [] properties.PropertyType {
	var rowhelp string = `Denotes the number of Rows in the grid.  This can be denoted by the following:

	- Pixels(px) - Defines the row in static pixel amount
	- Percentage(%) - Defines the row in terms of percentage of screen height
	- Fractional Units(fr) - Defines the row in terms of fractional units of the screen height

	###Example
	Define 4 rows:
	The first is 25px tall, row 2 is 10% of the total screen height, and rows 3 & 4 use the remaining height and creates the third row that is twice the height of the fourth row
	25px 10% 2fr 1fr`
	var colhelp string = `Denotes the number of Columns in the grid.  This can be denoted by the following:

	- Pixels(px) - Defines the column in static pixel amount
	- Percentage(%) - Defines the column in terms of percentage of screen height
	- Fractional Units(fr) - Defines the column in terms of fractional units of the screen height

	###Example
	Define 5 columns:
	The first & fifth columns are 25px wide, columns 2,3, & 4 use the remaining width and creates the third column that is twice the width of the second and fourth column, which are equivalent in size
	25px 1fr 2fr 1fr 25px`
  return [] properties.PropertyType {
    components.GetNavBarProperty(),
    components.GetNavBarItemsProperty(),
    &properties.PropertyColor{
      &properties.PropertyBase{"backgroundColor", "General", "Settings", "Background Color", ""},
      "#ffffff"},
		&properties.PropertyNumeric{
			&properties.PropertyBase{"padding", "Grid", "Grid Layout", "Padding", "Denotes the number of pixels to be used for padding between grid items"}, 0, properties.PropertyOptionsNumeric{ Min: 0, Max: 32}},
		&properties.PropertyText{
			&properties.PropertyBase{"gridTemplateRows", "Grid", "Grid Layout", "Grid Row Template", rowhelp}, "", properties.PropertyOptionsText{Rows: 3}},
		&properties.PropertyText{
			&properties.PropertyBase{"gridTemplateColumns", "Grid", "Grid Layout", "Grid Column Template", colhelp}, "", properties.PropertyOptionsText{Rows: 3}},
    lp.GetGridItemsProperty(),
  }
}

/*
 * Returns all components for the widget
 */
func (lp *LumavateProperties) GetAllComponents() [] *properties.Component {
  return [] *properties.Component {
    components.GetNavBarComponent(),
    components.GetNavBarItemComponent(),
    lp.GetNavigationComponent(),
    lp.GetVideoComponent(),
		lp.GetTextComponent(),
  }
}
