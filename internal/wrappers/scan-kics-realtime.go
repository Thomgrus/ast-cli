package wrappers

type KicsResultsCollection struct {
	Version     string        `json:"kics_version"`
	Count       uint          `json:"total_counter"`
	Results     []KicsQueries `json:"queries"`
	KicsSummary KicsSummary   `json:"severity_counters"`
}

type KicsQueries struct {
	QueryName   string      `json:"query_name"`
	QueryID     string      `json:"query_id"`
	Severity    string      `json:"severity"`
	Platform    string      `json:"platform"`
	Category    string      `json:"category"`
	Description string      `json:"description"`
	QueryURL    string      `json:"query_url"`
	Locations   []KicsFiles `json:"files"`
}

type KicsFiles struct {
	Filename      string `json:"file_name"`
	SimilarityID  string `json:"similarity_id"`
	Line          uint   `json:"line"`
	IssueType     string `json:"issue_type"`
	SearchKey     string `json:"search_key"`
	SearchLine    uint   `json:"search_line"`
	SearchValue   string `json:"search_value"`
	ExpectedValue string `json:"expected_value"`
	ActualValue   string `json:"actual_value"`
}

type KicsSummary struct {
	High   uint `json:"HIGH"`
	Info   uint `json:"INFO"`
	Low    uint `json:"LOW"`
	Medium uint `json:"MEDIUM"`
}
