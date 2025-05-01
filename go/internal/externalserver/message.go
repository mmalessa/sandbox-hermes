package externalserver

type StatusType string

const (
	StatusSuccess StatusType = "SUCCESS"
	StatusFail    StatusType = "FAIL"
)

type InternalRequest struct {
	Id      string            `json:"id" msgpack:"id"`
	Headers map[string]string `json:"headers" msgpack:"headers"`
	Body    map[string]string `json:"body" msgpack:"body"`
}

type InternalResponse struct {
	Id         string     `json:"id" msgpack:"id"`
	Status     StatusType `json:"status" msgpack:"status"`
	StatusCode uint16     `json:"statusCode" msgpack:"statusCode"`
	Message    string     `json:"message" msgpack:"message"`
}
