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

func (lp *LumavateProperties) GetTilesProperty() *properties.PropertyComponents {
  return &properties.PropertyComponents {
    &properties.PropertyBase{"gridItems", "Tiles", "Tile Settings", "Tile Settings", ""},
    [] *properties.Component{}, properties.PropertyOptionsComponent{[] string {"tile", "quote"}, [] *properties.Component {lp.GetTileComponent(), lp.GetQuoteComponent()} },
  }
}

func (lp *LumavateProperties) GetTileComponent() *properties.Component {
  props := [] properties.PropertyType {}
  props = append(props, &properties.PropertyTranslatedText{
    &properties.PropertyBase{"title", "", "", "Title", ""}, "Tile", properties.PropertyOptionsText{}})
	props = append(props, lp.GetLayoutProperties()...)
  return &properties.Component{"tile", "", "tile", "Tiles", "x", "Tile", props}
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
			&properties.PropertyBase{"padding", "Tiles", "Tile Layout", "Padding", ""}, 0, properties.PropertyOptionsNumeric{ Min: 0, Max: 32}},
		&properties.PropertyText{
			&properties.PropertyBase{"gridTemplateRows", "Tiles", "Tile Layout", "Grid Row Template", ""}, "", properties.PropertyOptionsText{Rows: 3}},
		&properties.PropertyText{
			&properties.PropertyBase{"gridTemplateColumns", "Tiles", "Tile Layout", "Grid Column Template", ""}, "", properties.PropertyOptionsText{Rows: 3}},
    lp.GetTilesProperty(),
  }
}

/*
 * Returns all components for the widget
 */
func (lp *LumavateProperties) GetAllComponents() [] *properties.Component {
  return [] *properties.Component {
    components.GetNavBarComponent(),
    components.GetNavBarItemComponent(),
    lp.GetTileComponent(),
		lp.GetQuoteComponent(),
  }
}
