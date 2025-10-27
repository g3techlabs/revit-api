package validation

var customMessages = map[string]string{
	"required":         "Field %s must be filled",
	"email":            "Invalid email address for field %s",
	"notanemail":       "Field %s is not a email, even when said so",
	"min":              "Field %s must have a minimum length of %s characters",
	"max":              "Field %s must have a maximum length of %s characters",
	"len":              "Field %s must be exactly %s characters long",
	"number":           "Field %s must be a number",
	"positive":         "Field %s must be a positive number",
	"alphanum":         "Field %s must contain only alphanumeric characters",
	"oneof":            "Invalid value for field %s",
	"password":         "Field %s must contain at least 8 characteres and one upper and lower letter",
	"datetime":         "Field %s must be a date of ISO 8601 format (YYYY-MM-DD)",
	"profilepic":       "Type of image not supported",
	"lowercase":        "Field %s must be all lowercase characters",
	"nicknametooshort": "Field %s, when set as nickname, must be at least 3 characters long",
	"notanickname":     "Field %s is not a nickname, even when said so",
	"gte":              "Field %s is not greater or equal to the specified value",
	"lte":              "Field %s must be less than or equal to the specified value",
	"gt":               "Field %s must be greater than the specified value",
	"lt":               "Field %s must be less than the specified value",
}
