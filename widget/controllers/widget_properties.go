package controllers

import (
  properties "github.com/Lumavate-Team/go-properties"
  ims_components "github.com/Lumavate-Team/ims-go-components"
  _ "os"
)

type LumavateProperties struct {
}

/*
 * Returns parking property for the widget
 */
func (lp *LumavateProperties) GetParkingProperty() *properties.PropertyComponents {
  return &properties.PropertyComponents {
    &properties.PropertyBase{"parking", "Parking", "Alernate Settings", "", ""},
    [] *properties.Component{}, properties.PropertyOptionsComponent{[] string {"parking"}, [] *properties.Component {lp.GetParkingComponent()} },
  }
}

/*
 * Returns parking component for the widget
 */
func (lp *LumavateProperties) GetParkingComponent() *properties.Component {
  props := [] properties.PropertyType {}
  props = append(props, &properties.PropertyText{
    &properties.PropertyBase{"altDate", "", "", "Alternate Date", ""}, "Alternate Date", properties.PropertyOptionsText{}})

  props = append(props, &properties.PropertyImage{
      &properties.PropertyBase{"altImage", "", "", "Alternate Parking Image", ""}})

  return &properties.Component{"parking", "", "parking", "Parking", "", "Parking", props }
}

/*
 * Returns all properties for the widget
 */
func (lp *LumavateProperties) GetAllProperties() [] properties.PropertyType {
  return [] properties.PropertyType {
    ims_components.GetTitleProperty(),
    ims_components.GetNavBarProperty(),
    ims_components.GetNavBarItemsProperty(),
    &properties.PropertyColor{
      &properties.PropertyBase{"backgroundColor", "General", "Properties", "Background Color", ""},
      "#ffffff"},
  	&properties.PropertyImage{
		  &properties.PropertyBase{"parkingImage", "Parking", "Main Settings", "Main Parking Image", ""}},
    lp.GetParkingProperty(),
  }
}

/*
 * Returns all components for the widget
 */
func (lp *LumavateProperties) GetAllComponents() [] *properties.Component {
  return [] *properties.Component {
    ims_components.GetTitleComponent(),
    ims_components.GetNavBarComponent(),
    ims_components.GetNavBarItemComponent(),
    lp.GetParkingComponent(),
  }
}