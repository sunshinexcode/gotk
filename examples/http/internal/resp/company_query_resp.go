package resp

type CompanyQueryResp struct {
	CompanyName string `json:"companyName,omitempty"`
	Id          int64  `json:"id,omitempty"`
}
