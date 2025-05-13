package entity

type SaveSecurityImageReq struct {
	Email  string
	Image  string
	Phrase string
}

type GetSecurityImageReq struct {
	Email string
}

type GetSecurityImageResp struct {
	Image  string
	Phrase string
}
