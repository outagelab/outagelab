package datapageservice

type GenerateDatapageRequest struct {
	Application string `json:"application"`
	Environment string `json:"environment"`
}

type GenerateDatapageResponse struct {
	OutageRules []*OutageRule `json:"rules"`
}

type OutageRule struct {
	Type                string               `json:"type"`
	HttpClientRequestV1 *HttpClientRequestV1 `json:"httpClientRequestV1"`
}

type HttpClientRequestV1 struct {
	Host     string `json:"host"`
	Status   *int   `json:"status"`
	Duration *int   `json:"duration"`
}
