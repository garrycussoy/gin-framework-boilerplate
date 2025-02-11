// This package provides custom error which will act as a base error template
// used accross the code
package custom_errors

type CustomError interface {
	Error() CustomErrorAttributes
}

type CustomErrorAttributes struct {
	Status  int         `json:"status"` // HTTP status code
	Code    string      `json:"code"`   // RC code
	Message string      `json:"message"`
	Detail  interface{} `json:"detail,omitempty"`
}

func (akasiaErr CustomErrorAttributes) Error() CustomErrorAttributes {
	return akasiaErr
}

func CreateCustomError(status int, rcCode string, message string, detail interface{}) CustomError {
	return CustomErrorAttributes{
		Status:  status,
		Code:    rcCode,
		Message: message,
		Detail:  detail,
	}
}
