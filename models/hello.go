package models

type GetHelloAPIResponse struct {
	Message string `json:"messageBody"`
	Name    string `json:"name"`
}

type PostHelloAPIRequestParams struct {
	Name string `json:"personName"`
	Age  int    `json:"personAge"`
}
