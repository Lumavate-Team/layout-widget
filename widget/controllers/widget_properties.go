package controllers

import (
  properties "github.com/Lumavate-Team/go-properties"
  ims_components "github.com/Lumavate-Team/ims-go-components"
  _ "os"
)

type LumavateProperties struct {
}

/*
 * Returns all properties for the widget
 */
func (lp *LumavateProperties) GetAllProperties() [] properties.PropertyType {
  return [] properties.PropertyType {
    ims_components.GetTitleProperty(),
  }
}

/*
 * Returns all components for the widget
 */
func (lp *LumavateProperties) GetAllComponents() [] *properties.Component {
  return [] *properties.Component {
    ims_components.GetTitleComponent(),
  }
}
