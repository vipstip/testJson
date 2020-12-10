package StructJson

import "github.com/bsm/openrtb/v3"

//
type BidRequest struct {
	openrtb.BidRequest
}

//	ID  string `json:"id"`
//	Imp []struct {
//		ID     string `json:"id"`
//		Banner struct {
//			W     int      `json:"w"`
//			H     int      `json:"h"`
//			Mimes []string `json:"mimes"`
//		} `json:"banner"`
//		Bidfloor    float64 `json:"bidfloor"`
//		Bidfloorcur string  `json:"bidfloorcur"`
//		Secure      int     `json:"secure"`
//	} `json:"imp"`
//	Site struct {
//		ID        string `json:"id"`
//		Domain    string `json:"domain"`
//		Page      string `json:"page"`
//		Publisher struct {
//			ID string `json:"id"`
//		} `json:"publisher"`
//	} `json:"site"`
//	Device struct {
//		Ua  string `json:"ua"`
//		IP  string `json:"ip"`
//		Geo struct {
//			Lat     float64 `json:"lat"`
//			Lon     float64 `json:"lon"`
//			Country string  `json:"country"`
//		} `json:"geo"`
//		Os         string `json:"os"`
//		Osv        string `json:"osv"`
//		Devicetype int    `json:"devicetype"`
//	} `json:"device"`
//	Regs struct {
//		Ext struct {
//			Gdpr int `json:"gdpr"`
//		} `json:"ext"`
//	} `json:"regs"`
//	User struct {
//		ID       string `json:"id"`
//		Buyeruid string `json:"buyeruid"`
//	} `json:"user"`
//	At   int      `json:"at"`
//	Tmax int      `json:"tmax"`
//	Bcat []string `json:"bcat"`
//	Ext  struct {
//		Payid  string `json:"payid"`
//		Schain struct {
//			Complete int `json:"complete"`
//			Nodes    []struct {
//				Asi string `json:"asi"`
//				Sid string `json:"sid"`
//				Rid string `json:"rid"`
//				Hp  int    `json:"hp"`
//			} `json:"nodes"`
//			Ver string `json:"ver"`
//		} `json:"schain"`
//	} `json:"ext"`
//}
