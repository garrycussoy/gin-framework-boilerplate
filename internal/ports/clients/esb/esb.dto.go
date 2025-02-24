package esb_ports

// General DTO used accross ESB services
type GeneralResponseDTO struct {
	Message string      `json:"message"`
	Code    string      `json:"code"`
	Data    interface{} `json:"data,omitempty"`
}
