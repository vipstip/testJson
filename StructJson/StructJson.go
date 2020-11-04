package StructJson

type BidRequest struct {
	Device DataDevice `json:"device"`
	Site   DataSite   `json:"site"`
	Gdpr   DataGdpr   `json:"gdpr"`
	Schain DataSchain `json:"schain"`
	Bids   []DataBid  `json:"bids"`
}

type DataDevice struct {
	Ua       string `json:"ua"`
	Height   int    `json:"height"`
	Width    int    `json:"width"`
	Dnt      int    `json:"dnt"`
	Language string `json:"language"`
}

type DataSite struct {
	Id       string `json:"id"`
	Page     string `json:"page"`
	Referrer string `json:"referrer"`
	Hostname string `json:"hostname"`
}

type DataGdpr struct {
}

type DataSchain struct {
	Ver      string    `json:"ver"`
	Complete int       `json:"complete"`
	Nodes    DataNodes `json:"nodes"`
}

type DataBid struct {
	Bidder              string         `json:"bidder"`
	Params              DataParams     `json:"params"`
	Crumbs              DataCrumbs     `json:"crumbs"`
	MediaTypes          DataMediaTypes `json:"mediaTypes"`
	AdUnitCode          string         `json:"adUnitCode"`
	TransactionId       string         `json:"transaction_id"`
	Sizes               [][]int        `json:"sizes"`
	BidId               string         `json:"bidId"`
	BidderRequestId     string         `json:"bidderRequestId"`
	AuctionId           string         `json:"auctionId"`
	Src                 string         `json:"src"`
	BidRequestsCount    int            `json:"bidRequestsCount"`
	BidderRequestsCount int            `json:"bidderRequestsCount"`
	BidderWinsCount     int            `json:"bidderWinsCount"`
	Schain              DataSchain     `json:"schain"`
	TargetKey           int            `json:"targetKey"`
}

type DataNodes struct {
	Asi string `json:"asi"`
	Sid string `json:"sid"`
	Hp  int    `json:"hp"`
}

type DataParams struct {
	SiteId     string  `json:"site_id"`
	FloorPrice float64 `json:"floor_price"`
}

type DataCrumbs struct {
	Pubcid string `json:"pubcid"`
}

type DataMediaTypes struct {
	Banner DataBanner `json:"banner"`
}

type DataBanner struct {
	Sizes [][]int `json:"sizes"`
}
