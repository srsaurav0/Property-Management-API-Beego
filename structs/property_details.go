package structs

type PropertyDetailsResponse struct {
	ID        string `json:"ID"`
	Feed      int    `json:"Feed"`
	Published bool   `json:"Published"`
	GeoInfo   struct {
		Categories []struct {
			Name       string   `json:"Name"`
			Slug       string   `json:"Slug"`
			Type       string   `json:"Type"`
			Display    []string `json:"Display"`
			LocationID string   `json:"LocationID"`
		} `json:"Categories"`
		City        string `json:"City"`
		Country     string `json:"Country"`
		CountryCode string `json:"CountryCode"`
		Display     string `json:"Display"`
		LocationID  string `json:"LocationID"`
		StateAbbr   string `json:"StateAbbr"`
		Lat         string `json:"Lat"`
		Lng         string `json:"Lng"`
	} `json:"GeoInfo"`
	Property struct {
		Amenities map[string]string `json:"Amenities"`
		Counts    struct {
			Bedroom   int `json:"Bedroom"`
			Bathroom  int `json:"Bathroom"`
			Reviews   int `json:"Reviews"`
			Occupancy int `json:"Occupancy"`
		} `json:"Counts"`
		EcoFriendly  bool   `json:"EcoFriendly"`
		FeatureImage string `json:"FeatureImage"`
		Image        *struct {
			Count  int      `json:"Count,omitempty"`
			Images []string `json:"Images,omitempty"`
		} `json:"Image,omitempty"`
		Price                  int                `json:"Price"`
		PropertyName           string             `json:"PropertyName"`
		PropertySlug           string             `json:"PropertySlug"`
		PropertyType           string             `json:"PropertyType"`
		PropertyTypeCategoryId string             `json:"PropertyTypeCategoryId"`
		ReviewScore            int                `json:"ReviewScore"`
		ReviewScores           map[string]float64 `json:"ReviewScores,omitempty"`
		RoomSize               float64            `json:"RoomSize"`
		MinStay                int                `json:"MinStay"`
		UpdatedAt              string             `json:"UpdatedAt"`
	} `json:"Property"`
	Partner struct {
		ID         string   `json:"ID"`
		Archived   []string `json:"Archived"`
		OwnerID    string   `json:"OwnerID"`
		HcomID     string   `json:"HcomID"`
		BrandId    string   `json:"BrandId"`
		URL        string   `json:"URL"`
		UnitNumber string   `json:"UnitNumber"`
		EpCluster  string   `json:"EpCluster"`
	} `json:"Partner"`
}
