package controllers

import (
  properties "github.com/Lumavate-Team/go-properties"
  ims_components "github.com/Lumavate-Team/ims-go-components"
  _"os"
  _"fmt"
)

type LumavateProperties struct {
}

func (lp *LumavateProperties) GetPrimaryContactProperty() *properties.PropertyComponents {
  return &properties.PropertyComponents {
    &properties.PropertyBase{"primaryContacts", "Contacts", "Primary Contact Settings", "Contact Settings", ""},
    [] *properties.Component{}, properties.PropertyOptionsComponent{[] string {"contact"}, [] *properties.Component {lp.GetContactComponent()} },
  }
}

func (lp *LumavateProperties) GetSecondaryContactProperty() *properties.PropertyComponents {
  return &properties.PropertyComponents {
    &properties.PropertyBase{"secondaryContacts", "Contacts", "Secondary Contact Settings", "Contact Settings", ""},
    [] *properties.Component{}, properties.PropertyOptionsComponent{[] string {"contact"}, [] *properties.Component {lp.GetContactComponent()} },
  }
}

func (lp *LumavateProperties) GetContactComponent() *properties.Component {
  props := [] properties.PropertyType {}
  props = append(props, &properties.PropertyText{
    &properties.PropertyBase{"firstName", "", "", "First Name", ""}, "First Name", properties.PropertyOptionsText{}})

  props = append(props, &properties.PropertyText{
    &properties.PropertyBase{"lastName", "", "", "Last Name", ""}, "Last Name", properties.PropertyOptionsText{}})

  props = append(props, &properties.PropertyText{
    &properties.PropertyBase{"jobTitle", "", "", "Title", ""}, "Title", properties.PropertyOptionsText{}})

  props = append(props, &properties.PropertyText{
    &properties.PropertyBase{"phoneNumber", "", "", "Phone Number", ""}, "Phone Number", properties.PropertyOptionsText{}})

  props = append(props, &properties.PropertyText{
    &properties.PropertyBase{"email", "", "", "Email", ""}, "Email", properties.PropertyOptionsText{}})

  return &properties.Component{"contact", "", "contact-component", "Contact", "x", "Contact", props}
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
    lp.GetPrimaryContactProperty(),
    lp.GetSecondaryContactProperty(),
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
    lp.GetContactComponent(),
  }
}
