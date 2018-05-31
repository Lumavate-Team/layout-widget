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

	var displayHelp string = `Denotes when this item should be displayed:
* Both: Display during _optimal_ & _degraded_ rendering (default)
* Optimal: Display during _optimal_ rendering on newer browsers supporting CSS Grid
* Degraded: Display only during _degraded_ rendering (browsers that do **not** support CSS Grid)`

  props = append(props, &properties.PropertyText{
		&properties.PropertyBase{"cssClass", "", "", "CSS Class", "Denotes the class (as defined in the Layout CSS) that will be added to the styling of this item."}, "", properties.PropertyOptionsText{}})
  props = append(props, &properties.PropertyDropdown{
		&properties.PropertyBase{"displayMode", "", "", "Display Mode", displayHelp}, "both", displayOptions})
  props = append(props, &properties.PropertyText{
		&properties.PropertyBase{"templateRowStart", "", "", "Grid Row Start", "This is Row at which this grid item will start"}, "", properties.PropertyOptionsText{Rows: 3}})
  props = append(props, &properties.PropertyText{
		&properties.PropertyBase{"templateRowEnd", "", "", "Grid Row End", "This is the Row at which this grid item will end"}, "", properties.PropertyOptionsText{ Rows: 3 }})
  props = append(props, &properties.PropertyText{
		&properties.PropertyBase{"templateColumnStart", "", "", "Grid Column Start", "This is the Column at which the grid item will start"}, "", properties.PropertyOptionsText{Rows: 3}})
  props = append(props, &properties.PropertyText{
		&properties.PropertyBase{"templateColumnEnd", "", "", "Grid Column End", "This is the Column at which the grid item will end"}, "", properties.PropertyOptionsText{Rows: 3}})
	return props
}

func (lp *LumavateProperties) GetGridItemsProperty() *properties.PropertyComponents {
  return &properties.PropertyComponents {
    &properties.PropertyBase{"gridItems", "Grid", "Grid Items", "Grid Items", ""},
    [] *properties.Component{}, properties.PropertyOptionsComponent{[] string {"navigation", "video", "text", "app", "form"}, [] *properties.Component {lp.GetNavigationComponent(), lp.GetVideoComponent(), lp.GetTextComponent(), lp.GetAppComponent(), lp.GetGridFormComponent()} },
  }
}

func (lp *LumavateProperties) GetAppComponent() *properties.Component {
  props := [] properties.PropertyType {}

  props = append(props, &properties.PropertyTranslatedText{
    &properties.PropertyBase{"title", "", "", "Title", ""}, "", properties.PropertyOptionsText{}})

  props = append(props, &properties.PropertyToggle{
    &properties.PropertyBase{"openNewWindow", "", "", "Open in New Window", ""}, false})

  props = append(props, &properties.PropertyText{
		&properties.PropertyBase{"appLink", "", "", "App Link", "Defines the **_deep link_** into the application"}, "", properties.PropertyOptionsText{}})

  props = append(props, &properties.PropertyText{
		&properties.PropertyBase{"apple", "", "", "iTunes Store", "Link to the application in iTunes"}, "", properties.PropertyOptionsText{}})

  props = append(props, &properties.PropertyText{
		&properties.PropertyBase{"google", "", "", "Google Play", "Link to the application in Google Play"}, "", properties.PropertyOptionsText{}})

  props = append(props, &properties.PropertyText{
		&properties.PropertyBase{"microsoft", "", "", "Windows Store", "Link to the application in the Windows Store"}, "", properties.PropertyOptionsText{}})

  props = append(props, &properties.PropertyToggle{
    &properties.PropertyBase{"useBackgroundImage", "", "", "Use Background Image", ""}, false})

  props = append(props, &properties.PropertyImage{
    &properties.PropertyBase{"image", "", "", "Background Image", ""}})

	// Background Image Scaling Options
	options := make(map[string]string)
	options["fill"] = "Fill"
	options["fit"] = "Fit"
	options["stretch"] = "Stretch"
	options["repeat"] = "Repeat"

	var scaleHelp string = `Denotes how the image will appear as the grid item scales.

* Fill: Sets _background-size_ to "cover" & will fill the entire width/height even if the image cannot be fully displayed

* Fit: Sets _background-size_ to "contain" & will fit the image inside the item maintaining aspect ratio so the entire image may be displayed

* Stretch: Sets _background-size_ to "100% 100%", stretching the image to the exact size of the item disregarding aspect ratios

* Repeat: Sets _background-repeat_ to repeat the image starting from the center of the item`

  props = append(props, &properties.PropertyDropdown{
		&properties.PropertyBase{"imageScaling", "", "", "Background Image Scaling", scaleHelp}, "fill",options})

  props = append(props, &properties.PropertyToggle{
    &properties.PropertyBase{"useBackgroundColor", "", "", "Use Background Color", ""}, false})

  props = append(props, &properties.PropertyColor{
    &properties.PropertyBase{"backgroundColor", "", "", "Background Color", ""}, "#ffffff"})

	props = append(props, lp.GetLayoutProperties()...)
	image := fmt.Sprintf("%v%v/static/images/application.svg", os.Getenv("WIDGET_URL_PREFIX"), os.Getenv("PUBLIC_KEY"))
  return &properties.Component{"app", "", "app", "App", image, "App", props}
}

func (lp *LumavateProperties) GetNavigationComponent() *properties.Component {
  props := [] properties.PropertyType {}

  props = append(props, &properties.PropertyTranslatedText{
    &properties.PropertyBase{"title", "", "", "Title", ""}, "", properties.PropertyOptionsText{}})

  props = append(props, &properties.PropertyToggle{
    &properties.PropertyBase{"useBackgroundImage", "", "", "Use Background Image", ""}, false})

  props = append(props, &properties.PropertyImage{
    &properties.PropertyBase{"image", "", "", "Background Image", ""}})

	// Background Image Scaling Options
	options := make(map[string]string)
	options["fill"] = "Fill"
	options["fit"] = "Fit"
	options["stretch"] = "Stretch"
	options["repeat"] = "Repeat"

	var scaleHelp string = `Denotes how the image will appear as the grid item scales.

* Fill: Sets _background-size_ to "cover" & will fill the entire width/height even if the image cannot be fully displayed

* Fit: Sets _background-size_ to "contain" & will fit the image inside the item maintaining aspect ratio so the entire image may be displayed

* Stretch: Sets _background-size_ to "100% 100%", stretching the image to the exact size of the item disregarding aspect ratios

* Repeat: Sets _background-repeat_ to repeat the image starting from the center of the item`

  props = append(props, &properties.PropertyDropdown{
		&properties.PropertyBase{"imageScaling", "", "", "Background Image Scaling", scaleHelp}, "fill",options})

  props = append(props, &properties.PropertyToggle{
    &properties.PropertyBase{"useBackgroundColor", "", "", "Use Background Color", ""}, false})

  props = append(props, &properties.PropertyColor{
    &properties.PropertyBase{"backgroundColor", "", "", "Background Color", ""}, "#ffffff"})

  props = append(props, &properties.PropertyPageLink{
    &properties.PropertyBase{"pageLink", "", "", "Page URL", ""}})

	props = append(props, lp.GetLayoutProperties()...)
	image := fmt.Sprintf("%v%v/static/images/navigation.svg", os.Getenv("WIDGET_URL_PREFIX"), os.Getenv("PUBLIC_KEY"))
  return &properties.Component{"navigation", "", "navigation", "Navigation", image, "Navigation", props}
}

func (lp *LumavateProperties) GetVideoComponent() *properties.Component {
	var videoHelp = `This should be the URL to the desired YouTube video, using the "embed" URL provided by YouTube & containing any relevant querystring parameters

### Example

https://www.youtube.com/embed/[VIDEOID]`
  props := [] properties.PropertyType {}
  props = append(props, &properties.PropertyTranslatedText{
    &properties.PropertyBase{"title", "", "", "Title", ""}, "", properties.PropertyOptionsText{}})
  props = append(props, &properties.PropertyText{
		&properties.PropertyBase{"video", "", "", "Video URL", videoHelp}, "https://www.youtube.com/embed/VIDEO_ID", properties.PropertyOptionsText{}})
	props = append(props, lp.GetLayoutProperties()...)
	image := fmt.Sprintf("%v%v/static/images/video.svg", os.Getenv("WIDGET_URL_PREFIX"),os.Getenv("PUBLIC_KEY"))
  return &properties.Component{"video", "", "video", "Video", image, "Video", props}
}

func (lp *LumavateProperties) GetTextComponent() *properties.Component {
  props := [] properties.PropertyType {}
  props = append(props, &properties.PropertyTranslatedText{
    &properties.PropertyBase{"title", "", "", "Title", ""}, "", properties.PropertyOptionsText{}})
  props = append(props, &properties.PropertyText{
		&properties.PropertyBase{"text", "", "", "Text", ""}, "", properties.PropertyOptionsText{}})
	props = append(props, lp.GetLayoutProperties()...)
	image := fmt.Sprintf("%v%v/static/images/text.svg", os.Getenv("WIDGET_URL_PREFIX"),os.Getenv("PUBLIC_KEY"))
  return &properties.Component{"text", "", "text", "Text", image, "Text", props}
}

func (lp *LumavateProperties) GetGridFormComponent() *properties.Component {
  props := [] properties.PropertyType {}
  props = append(props, &properties.PropertyTranslatedText{
    &properties.PropertyBase{"title", "", "", "Title", ""}, "", properties.PropertyOptionsText{}})
  props = append(props, lp.GetLayoutProperties()...)
	image := fmt.Sprintf("%v%v/static/images/form.svg", os.Getenv("WIDGET_URL_PREFIX"),os.Getenv("PUBLIC_KEY"))
  return &properties.Component{"form", "", "form", "Form", image, "Form", props}
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

	var cssHelp string = `Defines custom CSS used within the Layout.  Use the defined css classes here to add custom styling to each grid item.`

  return [] properties.PropertyType {
    components.GetNavBarProperty(),
    components.GetNavBarItemsProperty(),
    &properties.PropertyText{
      &properties.PropertyBase{"formAction", "Actions", "Microservices", "Registration URI", ""}, "", properties.PropertyOptionsText{}},
    &properties.PropertyColor{
      &properties.PropertyBase{"backgroundColor", "General", "Settings", "Background Color", ""}, "#ffffff"},
		&properties.PropertyToggle{
			&properties.PropertyBase{"displayBackgroundImage", "General", "Settings", "Display Background Image", ""}, false},
		&properties.PropertyImage{
			&properties.PropertyBase{"backgroundImage", "General", "Settings", "Background Image", ""}},
		&properties.PropertyNumeric{
			&properties.PropertyBase{"padding", "Grid", "Grid Layout", "Padding", "Denotes the number of pixels to be used for padding between grid items"}, 0, properties.PropertyOptionsNumeric{ Min: 0, Max: 32}},
		&properties.PropertyText{
			&properties.PropertyBase{"gridTemplateRows", "Grid", "Grid Layout", "Grid Row Template", rowhelp}, "", properties.PropertyOptionsText{Rows: 3}},
		&properties.PropertyText{
			&properties.PropertyBase{"gridTemplateColumns", "Grid", "Grid Layout", "Grid Column Template", colhelp}, "", properties.PropertyOptionsText{Rows: 3}},
		&properties.PropertyText{
			&properties.PropertyBase{"inlineCss", "CSS", "CSS", "Custom CSS", cssHelp}, "", properties.PropertyOptionsText{ReadOnly: false, Rows:5}},
    lp.GetGridItemsProperty(),
    components.GetFormItemsProperty(),
  }
}

/*
 * Returns all components for the widget
 */
func (lp *LumavateProperties) GetAllComponents() [] *properties.Component {
  return [] *properties.Component {
    components.GetNavBarComponent(),
    components.GetNavBarItemComponent(),
    lp.GetAppComponent(),
    lp.GetNavigationComponent(),
    lp.GetVideoComponent(),
		lp.GetTextComponent(),
    lp.GetGridFormComponent(),
		components.GetTextFormComponent(),
		components.GetDateFormComponent(),
		components.GetDropDownFormComponent(),
		components.GetCheckboxFormComponent(),
		components.GetAddressFormComponent(),
		components.GetEmailFormComponent(),
		components.GetHiddenFormComponent(),
  }
}
