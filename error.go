package zgo

const (
	TooManyRequest      = -32
	InvalidAccessToken  = -216
	InvalidRefreshToken = -14014
	ExpiredRefreshToken = -14020
)

type ErrorResp struct {
	ErrorName        string `json:"error_name"`
	ErrorReason      string `json:"error_reason"`
	RefDoc           string `json:"ref_doc"`
	ErrorDescription string `json:"error_description"`
	Error            int    `json:"error"`
}
