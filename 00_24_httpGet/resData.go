package main

type RankingResponse struct {
	CategoryRanking struct {
		Conditions struct {
			LastModified          string      `json:"last_modified"`
			StartDate             string      `json:"start_date"`
			EndDate               string      `json:"end_date"`
			CategoryId            string      `json:"category_id"`
			Gender                interface{} `json:"gender"`
			Generation            interface{} `json:"generation"`
			Period                string      `json:"period"`
			TotalResultsAvailable int         `json:"total_results_available"`
			FirstResultsPosition  int         `json:"first_results_position"`
			TotalResultsReturned  int         `json:"total_results_returned"`
		} `json:"conditions"`
		RankingData []struct {
			Rank  int    `json:"rank"`
			Name  string `json:"name"`
			Url   string `json:"url"`
			Image struct {
				Id     string `json:"id"`
				Small  string `json:"small"`
				Medium string `json:"medium"`
			} `json:"image"`
			Review struct {
				Rate  float64 `json:"rate"`
				Count int     `json:"count"`
				Url   string  `json:"url"`
			} `json:"review"`
			Store struct {
				Id   string `json:"id"`
				Name string `json:"name"`
				Url  string `json:"url"`
			} `json:"store"`
		} `json:"ranking_data"`
		Categories struct {
			Current struct {
				Id       string `json:"id"`
				ParentId string `json:"parent_id"`
				Url      string `json:"url"`
				Title    struct {
					Short  string      `json:"short"`
					Medium interface{} `json:"medium"`
					Long   interface{} `json:"long"`
				} `json:"title"`
				Path struct {
					Category []struct {
						Depth    int    `json:"depth"`
						Id       string `json:"id"`
						ParentId string `json:"parent_id"`
						Title    struct {
							Name string `json:"name"`
						} `json:"title"`
					} `json:"category"`
				} `json:"path"`
			} `json:"current"`
			Children struct {
				Child []struct {
					SortOrder int    `json:"sort_order"`
					Id        string `json:"id"`
					Url       string `json:"url"`
					Title     struct {
						Short  string `json:"short"`
						Medium string `json:"medium"`
						Long   string `json:"long"`
					} `json:"title"`
				} `json:"child"`
			} `json:"children"`
		} `json:"categories"`
	} `json:"category_ranking"`
}
