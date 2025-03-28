package vmiddleware_test

type (
	Request struct {
		Age  int    `form:"age" json:"age"`
		Name string `form:"name" json:"name"`
		Sign string `form:"sign'" json:"sign"`
	}

	RequestExpire struct {
		Age     int    `form:"age" json:"age"`
		Expired int64  `form:"expired,omitempty'" json:"expired,omitempty"`
		Name    string `form:"name" json:"name"`
		Ts      int64  `form:"ts'" json:"ts"`
	}

	RequestNoTraceId struct {
		Age  int    `form:"age" json:"age"`
		Name string `form:"name" json:"name"`
	}

	RequestOmitempty struct {
		Age  int    `form:"age,omitempty" json:"age,omitempty"`
		Name string `form:"name,omitempty" json:"name,omitempty"`
		Sign string `form:"sign,omitempty'" json:"sign,omitempty"`
	}

	RequestRequestId struct {
		Age       int    `form:"age" json:"age"`
		Name      string `form:"name" json:"name"`
		RequestId string `form:"requestId" json:"requestId"`
	}

	RequestTraceId struct {
		Age     int    `form:"age" json:"age"`
		Name    string `form:"name" json:"name"`
		Sign    string `form:"sign'" json:"sign"`
		TraceId string `form:"traceId" json:"traceId"`
	}

	RequestValid struct {
		Age  int    `form:"age" json:"age" v:"required|min:18|max:30"`
		Name string `form:"name" json:"name" v:"required|min-length:2|max-length:5#name is required||name length must be less than 5"`
	}
)

const (
	secret = "test"
)
