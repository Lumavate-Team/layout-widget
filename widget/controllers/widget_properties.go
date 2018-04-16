package controllers

import (
  properties "github.com/Lumavate-Team/lumavate-go-common/properties"
  components "github.com/Lumavate-Team/lumavate-go-common/components"
  _"os"
  _"fmt"
)

type LumavateProperties struct {
}

func (lp *LumavateProperties) GetLayoutProperties() [] properties.PropertyType {
  props := [] properties.PropertyType {}

  props = append(props, &properties.PropertyText{
		&properties.PropertyBase{"templateRowStart", "", "", "Grid Row Start", ""}, "", properties.PropertyOptionsText{Rows: 3}})
  props = append(props, &properties.PropertyText{
		&properties.PropertyBase{"templateRowEnd", "", "", "Grid Row End", ""}, "", properties.PropertyOptionsText{ Rows: 3 }})
  props = append(props, &properties.PropertyText{
		&properties.PropertyBase{"templateColumnStart", "", "", "Grid Column Start", ""}, "", properties.PropertyOptionsText{Rows: 3}})
  props = append(props, &properties.PropertyText{
		&properties.PropertyBase{"templateColumnEnd", "", "", "Grid Column End", ""}, "", properties.PropertyOptionsText{Rows: 3}})
	return props
}

func (lp *LumavateProperties) GetGridItemsProperty() *properties.PropertyComponents {
  return &properties.PropertyComponents {
    &properties.PropertyBase{"gridItems", "Grid", "Grid Items", "Grid Items", ""},
    [] *properties.Component{}, properties.PropertyOptionsComponent{[] string {"navigation", "video", "quote"}, [] *properties.Component {lp.GetNavigationComponent(), lp.GetVideoComponent(), lp.GetQuoteComponent()} },
  }
}

func (lp *LumavateProperties) GetNavigationComponent() *properties.Component {
  props := [] properties.PropertyType {}
  props = append(props, &properties.PropertyImage{
    &properties.PropertyBase{"image", "", "", "Background Image", ""}})
  props = append(props, &properties.PropertyTranslatedText{
    &properties.PropertyBase{"title", "", "", "Title", ""}, "", properties.PropertyOptionsText{}})
  props = append(props, &properties.PropertyPageLink{
    &properties.PropertyBase{"pageLink", "", "", "Page URL", ""}})
		//Image Scaling
		//Fill, Fit, Stretch, Tile
	props = append(props, lp.GetLayoutProperties()...)
  return &properties.Component{"navigation", "", "navigation", "Navigation", "x", "Navigation", props}
}

func (lp *LumavateProperties) GetVideoComponent() *properties.Component {
  props := [] properties.PropertyType {}
  props = append(props, &properties.PropertyTranslatedText{
    &properties.PropertyBase{"title", "", "", "Title", ""}, "", properties.PropertyOptionsText{}})
  props = append(props, &properties.PropertyText{
		&properties.PropertyBase{"video", "", "", "Video URL", ""}, "https://www.youtube.com/embed/VIDEO_ID", properties.PropertyOptionsText{}})
	props = append(props, lp.GetLayoutProperties()...)
  return &properties.Component{"video", "", "video", "Video", "x", "Video", props}
}

func (lp *LumavateProperties) GetQuoteComponent() *properties.Component {
	comp :=properties.LoadComponent("https://experience.john.labelnexusdev.com", "1.0.0", "quote")
	comp.Properties = append(comp.Properties, lp.GetLayoutProperties()...)
	comp.Category = "quote"
	comp.Section = ""
	return comp
}

/*
 * Returns all properties for the widget
 */
func (lp *LumavateProperties) GetAllProperties() [] properties.PropertyType {
  return [] properties.PropertyType {
    components.GetNavBarProperty(),
    components.GetNavBarItemsProperty(),
    &properties.PropertyColor{
      &properties.PropertyBase{"backgroundColor", "General", "Settings", "Background Color", ""},
      "#ffffff"},
		&properties.PropertyNumeric{
			&properties.PropertyBase{"padding", "Grid", "Grid Layout", "Padding", ""}, 0, properties.PropertyOptionsNumeric{ Min: 0, Max: 32}},
		&properties.PropertyText{
			&properties.PropertyBase{"gridTemplateRows", "Grid", "Grid Layout", "Grid Row Template", ""}, "", properties.PropertyOptionsText{Rows: 3}},
		&properties.PropertyText{
			&properties.PropertyBase{"gridTemplateColumns", "Grid", "Grid Layout", "Grid Column Template", ""}, "", properties.PropertyOptionsText{Rows: 3}},
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
		lp.GetQuoteComponent(),
  }
}
