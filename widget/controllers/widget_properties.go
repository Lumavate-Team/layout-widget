package controllers

import (
  properties "github.com/Lumavate-Team/go-properties"
  _ "os"
)

type LumavateProperties struct {
}

/*
 * Defines a simple, text property
 */
func (lp *LumavateProperties) GetTextProperty() *properties.PropertyText {
  return &properties.PropertyText{
    &properties.PropertyBase{"sampleText", "Quote", "Simple Settings", "Hello World Text", ""},
    "Hello, World!", properties.PropertyOptionsText{}}
}

/*
 * Gets a property that allows entry for 'Quote' data
 */
func (lp *LumavateProperties) GetComponentProperty() *properties.PropertyComponent {
  return &properties.PropertyComponent{
    &properties.PropertyBase{"quote", "Quote", "Quote Settings", "Quote Data", ""},
    lp.GetQuoteComponent(), properties.PropertyOptionsComponent{[] string {"all"}, [] *properties.Component {lp.GetQuoteComponent()} },
  }
}

/*
 * Gets a description for the 'Quote' component.  This is defined in a central place
 */
func (lp *LumavateProperties) GetQuoteComponent() *properties.Component {
  //return properties.LoadComponent(os.Getenv("BASE_URL"), "1.0.0", "quote")
  return properties.LoadComponent("https://experience.john.labelnexusdev.com", "1.0.0", "quote")
}

/*
 * Returns all properties for the widget
 */
func (lp *LumavateProperties) GetAllProperties() [] properties.PropertyType {
  return [] properties.PropertyType {
    lp.GetTextProperty(),
    lp.GetComponentProperty(),
  }
}

/*
 * Returns all components for the widget
 */
func (lp *LumavateProperties) GetAllComponents() [] *properties.Component {
  return [] *properties.Component {
    lp.GetQuoteComponent(),
  }
}
