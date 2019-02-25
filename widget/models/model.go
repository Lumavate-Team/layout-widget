package models

/* TBD - what do we do with the stuff below? */

type LumavateDomain struct {
  Payload struct {
    Data struct {
      Domain string
      RuntimeData map[string]interface{}
    }
  }
}


// structs used for getting designer defined user groups
type AuthGroupRequest struct {
  Payload struct {
    Data []GroupStruct
  }
}
type GroupStruct struct {
  Group string `json:"name"`
}

// struct used to get login status of user
type GroupRequest struct {
  Payload struct {
    Data struct{
      Roles       []string `json:"roles"`
      Status      string `json:"status"`
    }
  }
}

