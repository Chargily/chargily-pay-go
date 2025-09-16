package models

//================= GENERIC RESPONSES FOR AN ARRAY OF DATA OF THE SAME TYPE =========================//
///////////////////////////////////////////////////////////////////////////////////////////////////////
// Generic response for retrieving all entries of any type
type RetrieveAll[T any] struct {
	Livemode                bool                         `json:"livemode"`
	CurrentPage             int                          `json:"current_page"`
	Data                    []T                   		 `json:"data"` 
	FirstPageURL            string                       `json:"first_page_url"`
	LastPage                int                          `json:"last_page"`
	LastPageURL             string                       `json:"last_page_url"`
	NextPageURL             *string                      `json:"next_page_url"` 
	Path                    string                       `json:"path"`
	PerPage                 int                          `json:"per_page"`
	PrevPageURL             *string                      `json:"prev_page_url"` 
	Total                   int                          `json:"total"`
}
/////////////////////////////////////////////////////////////////////////////////////////////////////