package validation

var customMessages = map[string]string{
	"required":   "Field %s must be filled",
	"email":      "Invalid email address for field %s",
	"min":        "Field %s must have a minimum length of %s characters",
	"max":        "Field %s must have a maximum length of %s characters",
	"len":        "Field %s must be exactly %s characters long",
	"number":     "Field %s must be a number",
	"positive":   "Field %s must be a positive number",
	"alphanum":   "Field %s must contain only alphanumeric characters",
	"oneof":      "Invalid value for field %s",
	"password":   "Field %s must contain at least 8 characteres and one upper and lower letter",
	"datetime":   "Field %s must be a date of ISO 8601 format (YYYY-MM-DD)",
	"profilepic": "Type of image not supported",
}
