package address

type ReqAddress struct {
	Consignee string `json:"consignee" binding:"required"`
	Country   string `json:"country" binding:"required"`
	Province  string `json:"province" binding:"required"`
	City      string `json:"city" binding:"required"`
	District  string `json:"district" binding:"required"`
	Address   string `json:"address" binding:"required"`
	Mobile    string `json:"mobile" binding:"required"`
	IsDefault bool   `json:"is_default"`
}

type ReqUpdate struct {
	Consignee string `json:"consignee,omitempty" binding:"required"`
	Country   string `json:"country,omitempty" binding:"required"`
	Province  string `json:"province,omitempty" binding:"required"`
	City      string `json:"city,omitempty" binding:"required"`
	District  string `json:"district,omitempty" binding:"required"`
	Address   string `json:"address,omitempty" binding:"required"`
	Mobile    string `json:"mobile,omitempty" binding:"required"`
	IsDefault bool   `json:"is_default,omitempty"`
}

type UriParams struct {
	AddressID uint `uri:"addressID"`
}
