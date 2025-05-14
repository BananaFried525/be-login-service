package response

type ResponseStatus string

const (
	OK           ResponseStatus = "OK"
	BAD_REQUEST  ResponseStatus = "BAD_REQUEST"
	UNAUTHORIZED ResponseStatus = "UNAUTHORIZED"
	FORBIDDEN    ResponseStatus = "FORBIDDEN"
	NOT_FOUND    ResponseStatus = "NOT_FOUND"
)

type BaseResponse[T any] struct {
	Status ResponseStatus
	Data   T
}
