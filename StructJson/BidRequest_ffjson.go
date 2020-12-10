// Code generated by ffjson <https://github.com/pquerna/ffjson>. DO NOT EDIT.
// source: StructJson/BidRequest.go

package StructJson

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/bsm/openrtb/v3"
	fflib "github.com/pquerna/ffjson/fflib/v1"
)

// MarshalJSON marshal bytes to json - template
func (j *BidRequest) MarshalJSON() ([]byte, error) {
	var buf fflib.Buffer
	if j == nil {
		buf.WriteString("null")
		return buf.Bytes(), nil
	}
	err := j.MarshalJSONBuf(&buf)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// MarshalJSONBuf marshal buff to json - template
func (j *BidRequest) MarshalJSONBuf(buf fflib.EncodingBuffer) error {
	if j == nil {
		buf.WriteString("null")
		return nil
	}
	var err error
	var obj []byte
	_ = obj
	_ = err
	buf.WriteString(`{ "id":`)
	fflib.WriteJsonString(buf, string(j.ID))
	buf.WriteByte(',')
	if len(j.Impressions) != 0 {
		buf.WriteString(`"imp":`)
		if j.Impressions != nil {
			buf.WriteString(`[`)
			for i, v := range j.Impressions {
				if i != 0 {
					buf.WriteString(`,`)
				}
				/* Struct fall back. type=openrtb.Impression kind=struct */
				err = buf.Encode(&v)
				if err != nil {
					return err
				}
			}
			buf.WriteString(`]`)
		} else {
			buf.WriteString(`null`)
		}
		buf.WriteByte(',')
	}
	if j.Site != nil {
		if true {
			/* Struct fall back. type=openrtb.Site kind=struct */
			buf.WriteString(`"site":`)
			err = buf.Encode(j.Site)
			if err != nil {
				return err
			}
			buf.WriteByte(',')
		}
	}
	if j.App != nil {
		if true {
			/* Struct fall back. type=openrtb.App kind=struct */
			buf.WriteString(`"app":`)
			err = buf.Encode(j.App)
			if err != nil {
				return err
			}
			buf.WriteByte(',')
		}
	}
	if j.Device != nil {
		if true {
			/* Struct fall back. type=openrtb.Device kind=struct */
			buf.WriteString(`"device":`)
			err = buf.Encode(j.Device)
			if err != nil {
				return err
			}
			buf.WriteByte(',')
		}
	}
	if j.User != nil {
		if true {
			/* Struct fall back. type=openrtb.User kind=struct */
			buf.WriteString(`"user":`)
			err = buf.Encode(j.User)
			if err != nil {
				return err
			}
			buf.WriteByte(',')
		}
	}
	if j.Test != 0 {
		buf.WriteString(`"test":`)
		fflib.FormatBits2(buf, uint64(j.Test), 10, j.Test < 0)
		buf.WriteByte(',')
	}
	buf.WriteString(`"at":`)
	fflib.FormatBits2(buf, uint64(j.AuctionType), 10, j.AuctionType < 0)
	buf.WriteByte(',')
	if j.TimeMax != 0 {
		buf.WriteString(`"tmax":`)
		fflib.FormatBits2(buf, uint64(j.TimeMax), 10, j.TimeMax < 0)
		buf.WriteByte(',')
	}
	if len(j.Seats) != 0 {
		buf.WriteString(`"wseat":`)
		if j.Seats != nil {
			buf.WriteString(`[`)
			for i, v := range j.Seats {
				if i != 0 {
					buf.WriteString(`,`)
				}
				fflib.WriteJsonString(buf, string(v))
			}
			buf.WriteString(`]`)
		} else {
			buf.WriteString(`null`)
		}
		buf.WriteByte(',')
	}
	if len(j.BlockedSeats) != 0 {
		buf.WriteString(`"bseat":`)
		if j.BlockedSeats != nil {
			buf.WriteString(`[`)
			for i, v := range j.BlockedSeats {
				if i != 0 {
					buf.WriteString(`,`)
				}
				fflib.WriteJsonString(buf, string(v))
			}
			buf.WriteString(`]`)
		} else {
			buf.WriteString(`null`)
		}
		buf.WriteByte(',')
	}
	if len(j.Languages) != 0 {
		buf.WriteString(`"wlang":`)
		if j.Languages != nil {
			buf.WriteString(`[`)
			for i, v := range j.Languages {
				if i != 0 {
					buf.WriteString(`,`)
				}
				fflib.WriteJsonString(buf, string(v))
			}
			buf.WriteString(`]`)
		} else {
			buf.WriteString(`null`)
		}
		buf.WriteByte(',')
	}
	if j.AllImpressions != 0 {
		buf.WriteString(`"allimps":`)
		fflib.FormatBits2(buf, uint64(j.AllImpressions), 10, j.AllImpressions < 0)
		buf.WriteByte(',')
	}
	if len(j.Currencies) != 0 {
		buf.WriteString(`"cur":`)
		if j.Currencies != nil {
			buf.WriteString(`[`)
			for i, v := range j.Currencies {
				if i != 0 {
					buf.WriteString(`,`)
				}
				fflib.WriteJsonString(buf, string(v))
			}
			buf.WriteString(`]`)
		} else {
			buf.WriteString(`null`)
		}
		buf.WriteByte(',')
	}
	if len(j.BlockedCategories) != 0 {
		buf.WriteString(`"bcat":`)
		if j.BlockedCategories != nil {
			buf.WriteString(`[`)
			for i, v := range j.BlockedCategories {
				if i != 0 {
					buf.WriteString(`,`)
				}
				fflib.WriteJsonString(buf, string(v))
			}
			buf.WriteString(`]`)
		} else {
			buf.WriteString(`null`)
		}
		buf.WriteByte(',')
	}
	if len(j.BlockedAdvDomains) != 0 {
		buf.WriteString(`"badv":`)
		if j.BlockedAdvDomains != nil {
			buf.WriteString(`[`)
			for i, v := range j.BlockedAdvDomains {
				if i != 0 {
					buf.WriteString(`,`)
				}
				fflib.WriteJsonString(buf, string(v))
			}
			buf.WriteString(`]`)
		} else {
			buf.WriteString(`null`)
		}
		buf.WriteByte(',')
	}
	if len(j.BlockedApps) != 0 {
		buf.WriteString(`"bapp":`)
		if j.BlockedApps != nil {
			buf.WriteString(`[`)
			for i, v := range j.BlockedApps {
				if i != 0 {
					buf.WriteString(`,`)
				}
				fflib.WriteJsonString(buf, string(v))
			}
			buf.WriteString(`]`)
		} else {
			buf.WriteString(`null`)
		}
		buf.WriteByte(',')
	}
	if j.Source != nil {
		if true {
			/* Struct fall back. type=openrtb.Source kind=struct */
			buf.WriteString(`"source":`)
			err = buf.Encode(j.Source)
			if err != nil {
				return err
			}
			buf.WriteByte(',')
		}
	}
	if j.Regulations != nil {
		if true {
			/* Struct fall back. type=openrtb.Regulations kind=struct */
			buf.WriteString(`"regs":`)
			err = buf.Encode(j.Regulations)
			if err != nil {
				return err
			}
			buf.WriteByte(',')
		}
	}
	if len(j.Ext) != 0 {
		buf.WriteString(`"ext":`)

		{

			obj, err = j.Ext.MarshalJSON()
			if err != nil {
				return err
			}
			buf.Write(obj)

		}
		buf.WriteByte(',')
	}
	buf.Rewind(1)
	buf.WriteByte('}')
	return nil
}

const (
	ffjtBidRequestbase = iota
	ffjtBidRequestnosuchkey

	ffjtBidRequestID

	ffjtBidRequestImpressions

	ffjtBidRequestSite

	ffjtBidRequestApp

	ffjtBidRequestDevice

	ffjtBidRequestUser

	ffjtBidRequestTest

	ffjtBidRequestAuctionType

	ffjtBidRequestTimeMax

	ffjtBidRequestSeats

	ffjtBidRequestBlockedSeats

	ffjtBidRequestLanguages

	ffjtBidRequestAllImpressions

	ffjtBidRequestCurrencies

	ffjtBidRequestBlockedCategories

	ffjtBidRequestBlockedAdvDomains

	ffjtBidRequestBlockedApps

	ffjtBidRequestSource

	ffjtBidRequestRegulations

	ffjtBidRequestExt
)

var ffjKeyBidRequestID = []byte("id")

var ffjKeyBidRequestImpressions = []byte("imp")

var ffjKeyBidRequestSite = []byte("site")

var ffjKeyBidRequestApp = []byte("app")

var ffjKeyBidRequestDevice = []byte("device")

var ffjKeyBidRequestUser = []byte("user")

var ffjKeyBidRequestTest = []byte("test")

var ffjKeyBidRequestAuctionType = []byte("at")

var ffjKeyBidRequestTimeMax = []byte("tmax")

var ffjKeyBidRequestSeats = []byte("wseat")

var ffjKeyBidRequestBlockedSeats = []byte("bseat")

var ffjKeyBidRequestLanguages = []byte("wlang")

var ffjKeyBidRequestAllImpressions = []byte("allimps")

var ffjKeyBidRequestCurrencies = []byte("cur")

var ffjKeyBidRequestBlockedCategories = []byte("bcat")

var ffjKeyBidRequestBlockedAdvDomains = []byte("badv")

var ffjKeyBidRequestBlockedApps = []byte("bapp")

var ffjKeyBidRequestSource = []byte("source")

var ffjKeyBidRequestRegulations = []byte("regs")

var ffjKeyBidRequestExt = []byte("ext")

// UnmarshalJSON umarshall json - template of ffjson
func (j *BidRequest) UnmarshalJSON(input []byte) error {
	fs := fflib.NewFFLexer(input)
	return j.UnmarshalJSONFFLexer(fs, fflib.FFParse_map_start)
}

// UnmarshalJSONFFLexer fast json unmarshall - template ffjson
func (j *BidRequest) UnmarshalJSONFFLexer(fs *fflib.FFLexer, state fflib.FFParseState) error {
	var err error
	currentKey := ffjtBidRequestbase
	_ = currentKey
	tok := fflib.FFTok_init
	wantedTok := fflib.FFTok_init

mainparse:
	for {
		tok = fs.Scan()
		//	println(fmt.Sprintf("debug: tok: %v  state: %v", tok, state))
		if tok == fflib.FFTok_error {
			goto tokerror
		}

		switch state {

		case fflib.FFParse_map_start:
			if tok != fflib.FFTok_left_bracket {
				wantedTok = fflib.FFTok_left_bracket
				goto wrongtokenerror
			}
			state = fflib.FFParse_want_key
			continue

		case fflib.FFParse_after_value:
			if tok == fflib.FFTok_comma {
				state = fflib.FFParse_want_key
			} else if tok == fflib.FFTok_right_bracket {
				goto done
			} else {
				wantedTok = fflib.FFTok_comma
				goto wrongtokenerror
			}

		case fflib.FFParse_want_key:
			// json {} ended. goto exit. woo.
			if tok == fflib.FFTok_right_bracket {
				goto done
			}
			if tok != fflib.FFTok_string {
				wantedTok = fflib.FFTok_string
				goto wrongtokenerror
			}

			kn := fs.Output.Bytes()
			if len(kn) <= 0 {
				// "" case. hrm.
				currentKey = ffjtBidRequestnosuchkey
				state = fflib.FFParse_want_colon
				goto mainparse
			} else {
				switch kn[0] {

				case 'a':

					if bytes.Equal(ffjKeyBidRequestApp, kn) {
						currentKey = ffjtBidRequestApp
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffjKeyBidRequestAuctionType, kn) {
						currentKey = ffjtBidRequestAuctionType
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffjKeyBidRequestAllImpressions, kn) {
						currentKey = ffjtBidRequestAllImpressions
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'b':

					if bytes.Equal(ffjKeyBidRequestBlockedSeats, kn) {
						currentKey = ffjtBidRequestBlockedSeats
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffjKeyBidRequestBlockedCategories, kn) {
						currentKey = ffjtBidRequestBlockedCategories
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffjKeyBidRequestBlockedAdvDomains, kn) {
						currentKey = ffjtBidRequestBlockedAdvDomains
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffjKeyBidRequestBlockedApps, kn) {
						currentKey = ffjtBidRequestBlockedApps
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'c':

					if bytes.Equal(ffjKeyBidRequestCurrencies, kn) {
						currentKey = ffjtBidRequestCurrencies
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'd':

					if bytes.Equal(ffjKeyBidRequestDevice, kn) {
						currentKey = ffjtBidRequestDevice
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'e':

					if bytes.Equal(ffjKeyBidRequestExt, kn) {
						currentKey = ffjtBidRequestExt
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'i':

					if bytes.Equal(ffjKeyBidRequestID, kn) {
						currentKey = ffjtBidRequestID
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffjKeyBidRequestImpressions, kn) {
						currentKey = ffjtBidRequestImpressions
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'r':

					if bytes.Equal(ffjKeyBidRequestRegulations, kn) {
						currentKey = ffjtBidRequestRegulations
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 's':

					if bytes.Equal(ffjKeyBidRequestSite, kn) {
						currentKey = ffjtBidRequestSite
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffjKeyBidRequestSource, kn) {
						currentKey = ffjtBidRequestSource
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 't':

					if bytes.Equal(ffjKeyBidRequestTest, kn) {
						currentKey = ffjtBidRequestTest
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffjKeyBidRequestTimeMax, kn) {
						currentKey = ffjtBidRequestTimeMax
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'u':

					if bytes.Equal(ffjKeyBidRequestUser, kn) {
						currentKey = ffjtBidRequestUser
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'w':

					if bytes.Equal(ffjKeyBidRequestSeats, kn) {
						currentKey = ffjtBidRequestSeats
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffjKeyBidRequestLanguages, kn) {
						currentKey = ffjtBidRequestLanguages
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				}

				if fflib.SimpleLetterEqualFold(ffjKeyBidRequestExt, kn) {
					currentKey = ffjtBidRequestExt
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyBidRequestRegulations, kn) {
					currentKey = ffjtBidRequestRegulations
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyBidRequestSource, kn) {
					currentKey = ffjtBidRequestSource
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffjKeyBidRequestBlockedApps, kn) {
					currentKey = ffjtBidRequestBlockedApps
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffjKeyBidRequestBlockedAdvDomains, kn) {
					currentKey = ffjtBidRequestBlockedAdvDomains
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffjKeyBidRequestBlockedCategories, kn) {
					currentKey = ffjtBidRequestBlockedCategories
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffjKeyBidRequestCurrencies, kn) {
					currentKey = ffjtBidRequestCurrencies
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyBidRequestAllImpressions, kn) {
					currentKey = ffjtBidRequestAllImpressions
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffjKeyBidRequestLanguages, kn) {
					currentKey = ffjtBidRequestLanguages
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyBidRequestBlockedSeats, kn) {
					currentKey = ffjtBidRequestBlockedSeats
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyBidRequestSeats, kn) {
					currentKey = ffjtBidRequestSeats
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffjKeyBidRequestTimeMax, kn) {
					currentKey = ffjtBidRequestTimeMax
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffjKeyBidRequestAuctionType, kn) {
					currentKey = ffjtBidRequestAuctionType
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyBidRequestTest, kn) {
					currentKey = ffjtBidRequestTest
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyBidRequestUser, kn) {
					currentKey = ffjtBidRequestUser
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffjKeyBidRequestDevice, kn) {
					currentKey = ffjtBidRequestDevice
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffjKeyBidRequestApp, kn) {
					currentKey = ffjtBidRequestApp
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyBidRequestSite, kn) {
					currentKey = ffjtBidRequestSite
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffjKeyBidRequestImpressions, kn) {
					currentKey = ffjtBidRequestImpressions
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffjKeyBidRequestID, kn) {
					currentKey = ffjtBidRequestID
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				currentKey = ffjtBidRequestnosuchkey
				state = fflib.FFParse_want_colon
				goto mainparse
			}

		case fflib.FFParse_want_colon:
			if tok != fflib.FFTok_colon {
				wantedTok = fflib.FFTok_colon
				goto wrongtokenerror
			}
			state = fflib.FFParse_want_value
			continue
		case fflib.FFParse_want_value:

			if tok == fflib.FFTok_left_brace || tok == fflib.FFTok_left_bracket || tok == fflib.FFTok_integer || tok == fflib.FFTok_double || tok == fflib.FFTok_string || tok == fflib.FFTok_bool || tok == fflib.FFTok_null {
				switch currentKey {

				case ffjtBidRequestID:
					goto handle_ID

				case ffjtBidRequestImpressions:
					goto handle_Impressions

				case ffjtBidRequestSite:
					goto handle_Site

				case ffjtBidRequestApp:
					goto handle_App

				case ffjtBidRequestDevice:
					goto handle_Device

				case ffjtBidRequestUser:
					goto handle_User

				case ffjtBidRequestTest:
					goto handle_Test

				case ffjtBidRequestAuctionType:
					goto handle_AuctionType

				case ffjtBidRequestTimeMax:
					goto handle_TimeMax

				case ffjtBidRequestSeats:
					goto handle_Seats

				case ffjtBidRequestBlockedSeats:
					goto handle_BlockedSeats

				case ffjtBidRequestLanguages:
					goto handle_Languages

				case ffjtBidRequestAllImpressions:
					goto handle_AllImpressions

				case ffjtBidRequestCurrencies:
					goto handle_Currencies

				case ffjtBidRequestBlockedCategories:
					goto handle_BlockedCategories

				case ffjtBidRequestBlockedAdvDomains:
					goto handle_BlockedAdvDomains

				case ffjtBidRequestBlockedApps:
					goto handle_BlockedApps

				case ffjtBidRequestSource:
					goto handle_Source

				case ffjtBidRequestRegulations:
					goto handle_Regulations

				case ffjtBidRequestExt:
					goto handle_Ext

				case ffjtBidRequestnosuchkey:
					err = fs.SkipField(tok)
					if err != nil {
						return fs.WrapErr(err)
					}
					state = fflib.FFParse_after_value
					goto mainparse
				}
			} else {
				goto wantedvalue
			}
		}
	}

handle_ID:

	/* handler: j.ID type=string kind=string quoted=false*/

	{

		{
			if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
			}
		}

		if tok == fflib.FFTok_null {

		} else {

			outBuf := fs.Output.Bytes()

			j.ID = string(string(outBuf))

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Impressions:

	/* handler: j.Impressions type=[]openrtb.Impression kind=slice quoted=false*/

	{

		{
			if tok != fflib.FFTok_left_brace && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for ", tok))
			}
		}

		if tok == fflib.FFTok_null {
			j.Impressions = nil
		} else {

			j.Impressions = []openrtb.Impression{}

			wantVal := true

			for {

				var tmpJImpressions openrtb.Impression

				tok = fs.Scan()
				if tok == fflib.FFTok_error {
					goto tokerror
				}
				if tok == fflib.FFTok_right_brace {
					break
				}

				if tok == fflib.FFTok_comma {
					if wantVal == true {
						// TODO(pquerna): this isn't an ideal error message, this handles
						// things like [,,,] as an array value.
						return fs.WrapErr(fmt.Errorf("wanted value token, but got token: %v", tok))
					}
					continue
				} else {
					wantVal = true
				}

				/* handler: tmpJImpressions type=openrtb.Impression kind=struct quoted=false*/

				{
					/* Falling back. type=openrtb.Impression kind=struct */
					tbuf, err := fs.CaptureField(tok)
					if err != nil {
						return fs.WrapErr(err)
					}

					err = json.Unmarshal(tbuf, &tmpJImpressions)
					if err != nil {
						return fs.WrapErr(err)
					}
				}

				j.Impressions = append(j.Impressions, tmpJImpressions)

				wantVal = false
			}
		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Site:

	/* handler: j.Site type=openrtb.Site kind=struct quoted=false*/

	{
		/* Falling back. type=openrtb.Site kind=struct */
		tbuf, err := fs.CaptureField(tok)
		if err != nil {
			return fs.WrapErr(err)
		}

		err = json.Unmarshal(tbuf, &j.Site)
		if err != nil {
			return fs.WrapErr(err)
		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_App:

	/* handler: j.App type=openrtb.App kind=struct quoted=false*/

	{
		/* Falling back. type=openrtb.App kind=struct */
		tbuf, err := fs.CaptureField(tok)
		if err != nil {
			return fs.WrapErr(err)
		}

		err = json.Unmarshal(tbuf, &j.App)
		if err != nil {
			return fs.WrapErr(err)
		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Device:

	/* handler: j.Device type=openrtb.Device kind=struct quoted=false*/

	{
		/* Falling back. type=openrtb.Device kind=struct */
		tbuf, err := fs.CaptureField(tok)
		if err != nil {
			return fs.WrapErr(err)
		}

		err = json.Unmarshal(tbuf, &j.Device)
		if err != nil {
			return fs.WrapErr(err)
		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_User:

	/* handler: j.User type=openrtb.User kind=struct quoted=false*/

	{
		/* Falling back. type=openrtb.User kind=struct */
		tbuf, err := fs.CaptureField(tok)
		if err != nil {
			return fs.WrapErr(err)
		}

		err = json.Unmarshal(tbuf, &j.User)
		if err != nil {
			return fs.WrapErr(err)
		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Test:

	/* handler: j.Test type=int kind=int quoted=false*/

	{
		if tok != fflib.FFTok_integer && tok != fflib.FFTok_null {
			return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for int", tok))
		}
	}

	{

		if tok == fflib.FFTok_null {

		} else {

			tval, err := fflib.ParseInt(fs.Output.Bytes(), 10, 64)

			if err != nil {
				return fs.WrapErr(err)
			}

			j.Test = int(tval)

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_AuctionType:

	/* handler: j.AuctionType type=int kind=int quoted=false*/

	{
		if tok != fflib.FFTok_integer && tok != fflib.FFTok_null {
			return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for int", tok))
		}
	}

	{

		if tok == fflib.FFTok_null {

		} else {

			tval, err := fflib.ParseInt(fs.Output.Bytes(), 10, 64)

			if err != nil {
				return fs.WrapErr(err)
			}

			j.AuctionType = int(tval)

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_TimeMax:

	/* handler: j.TimeMax type=int kind=int quoted=false*/

	{
		if tok != fflib.FFTok_integer && tok != fflib.FFTok_null {
			return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for int", tok))
		}
	}

	{

		if tok == fflib.FFTok_null {

		} else {

			tval, err := fflib.ParseInt(fs.Output.Bytes(), 10, 64)

			if err != nil {
				return fs.WrapErr(err)
			}

			j.TimeMax = int(tval)

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Seats:

	/* handler: j.Seats type=[]string kind=slice quoted=false*/

	{

		{
			if tok != fflib.FFTok_left_brace && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for ", tok))
			}
		}

		if tok == fflib.FFTok_null {
			j.Seats = nil
		} else {

			j.Seats = []string{}

			wantVal := true

			for {

				var tmpJSeats string

				tok = fs.Scan()
				if tok == fflib.FFTok_error {
					goto tokerror
				}
				if tok == fflib.FFTok_right_brace {
					break
				}

				if tok == fflib.FFTok_comma {
					if wantVal == true {
						// TODO(pquerna): this isn't an ideal error message, this handles
						// things like [,,,] as an array value.
						return fs.WrapErr(fmt.Errorf("wanted value token, but got token: %v", tok))
					}
					continue
				} else {
					wantVal = true
				}

				/* handler: tmpJSeats type=string kind=string quoted=false*/

				{

					{
						if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
							return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
						}
					}

					if tok == fflib.FFTok_null {

					} else {

						outBuf := fs.Output.Bytes()

						tmpJSeats = string(string(outBuf))

					}
				}

				j.Seats = append(j.Seats, tmpJSeats)

				wantVal = false
			}
		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_BlockedSeats:

	/* handler: j.BlockedSeats type=[]string kind=slice quoted=false*/

	{

		{
			if tok != fflib.FFTok_left_brace && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for ", tok))
			}
		}

		if tok == fflib.FFTok_null {
			j.BlockedSeats = nil
		} else {

			j.BlockedSeats = []string{}

			wantVal := true

			for {

				var tmpJBlockedSeats string

				tok = fs.Scan()
				if tok == fflib.FFTok_error {
					goto tokerror
				}
				if tok == fflib.FFTok_right_brace {
					break
				}

				if tok == fflib.FFTok_comma {
					if wantVal == true {
						// TODO(pquerna): this isn't an ideal error message, this handles
						// things like [,,,] as an array value.
						return fs.WrapErr(fmt.Errorf("wanted value token, but got token: %v", tok))
					}
					continue
				} else {
					wantVal = true
				}

				/* handler: tmpJBlockedSeats type=string kind=string quoted=false*/

				{

					{
						if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
							return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
						}
					}

					if tok == fflib.FFTok_null {

					} else {

						outBuf := fs.Output.Bytes()

						tmpJBlockedSeats = string(string(outBuf))

					}
				}

				j.BlockedSeats = append(j.BlockedSeats, tmpJBlockedSeats)

				wantVal = false
			}
		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Languages:

	/* handler: j.Languages type=[]string kind=slice quoted=false*/

	{

		{
			if tok != fflib.FFTok_left_brace && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for ", tok))
			}
		}

		if tok == fflib.FFTok_null {
			j.Languages = nil
		} else {

			j.Languages = []string{}

			wantVal := true

			for {

				var tmpJLanguages string

				tok = fs.Scan()
				if tok == fflib.FFTok_error {
					goto tokerror
				}
				if tok == fflib.FFTok_right_brace {
					break
				}

				if tok == fflib.FFTok_comma {
					if wantVal == true {
						// TODO(pquerna): this isn't an ideal error message, this handles
						// things like [,,,] as an array value.
						return fs.WrapErr(fmt.Errorf("wanted value token, but got token: %v", tok))
					}
					continue
				} else {
					wantVal = true
				}

				/* handler: tmpJLanguages type=string kind=string quoted=false*/

				{

					{
						if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
							return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
						}
					}

					if tok == fflib.FFTok_null {

					} else {

						outBuf := fs.Output.Bytes()

						tmpJLanguages = string(string(outBuf))

					}
				}

				j.Languages = append(j.Languages, tmpJLanguages)

				wantVal = false
			}
		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_AllImpressions:

	/* handler: j.AllImpressions type=int kind=int quoted=false*/

	{
		if tok != fflib.FFTok_integer && tok != fflib.FFTok_null {
			return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for int", tok))
		}
	}

	{

		if tok == fflib.FFTok_null {

		} else {

			tval, err := fflib.ParseInt(fs.Output.Bytes(), 10, 64)

			if err != nil {
				return fs.WrapErr(err)
			}

			j.AllImpressions = int(tval)

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Currencies:

	/* handler: j.Currencies type=[]string kind=slice quoted=false*/

	{

		{
			if tok != fflib.FFTok_left_brace && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for ", tok))
			}
		}

		if tok == fflib.FFTok_null {
			j.Currencies = nil
		} else {

			j.Currencies = []string{}

			wantVal := true

			for {

				var tmpJCurrencies string

				tok = fs.Scan()
				if tok == fflib.FFTok_error {
					goto tokerror
				}
				if tok == fflib.FFTok_right_brace {
					break
				}

				if tok == fflib.FFTok_comma {
					if wantVal == true {
						// TODO(pquerna): this isn't an ideal error message, this handles
						// things like [,,,] as an array value.
						return fs.WrapErr(fmt.Errorf("wanted value token, but got token: %v", tok))
					}
					continue
				} else {
					wantVal = true
				}

				/* handler: tmpJCurrencies type=string kind=string quoted=false*/

				{

					{
						if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
							return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
						}
					}

					if tok == fflib.FFTok_null {

					} else {

						outBuf := fs.Output.Bytes()

						tmpJCurrencies = string(string(outBuf))

					}
				}

				j.Currencies = append(j.Currencies, tmpJCurrencies)

				wantVal = false
			}
		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_BlockedCategories:

	/* handler: j.BlockedCategories type=[]openrtb.ContentCategory kind=slice quoted=false*/

	{

		{
			if tok != fflib.FFTok_left_brace && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for ", tok))
			}
		}

		if tok == fflib.FFTok_null {
			j.BlockedCategories = nil
		} else {

			j.BlockedCategories = []openrtb.ContentCategory{}

			wantVal := true

			for {

				var tmpJBlockedCategories openrtb.ContentCategory

				tok = fs.Scan()
				if tok == fflib.FFTok_error {
					goto tokerror
				}
				if tok == fflib.FFTok_right_brace {
					break
				}

				if tok == fflib.FFTok_comma {
					if wantVal == true {
						// TODO(pquerna): this isn't an ideal error message, this handles
						// things like [,,,] as an array value.
						return fs.WrapErr(fmt.Errorf("wanted value token, but got token: %v", tok))
					}
					continue
				} else {
					wantVal = true
				}

				/* handler: tmpJBlockedCategories type=openrtb.ContentCategory kind=string quoted=false*/

				{

					{
						if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
							return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for ContentCategory", tok))
						}
					}

					if tok == fflib.FFTok_null {

					} else {

						outBuf := fs.Output.Bytes()

						tmpJBlockedCategories = openrtb.ContentCategory(string(outBuf))

					}
				}

				j.BlockedCategories = append(j.BlockedCategories, tmpJBlockedCategories)

				wantVal = false
			}
		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_BlockedAdvDomains:

	/* handler: j.BlockedAdvDomains type=[]string kind=slice quoted=false*/

	{

		{
			if tok != fflib.FFTok_left_brace && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for ", tok))
			}
		}

		if tok == fflib.FFTok_null {
			j.BlockedAdvDomains = nil
		} else {

			j.BlockedAdvDomains = []string{}

			wantVal := true

			for {

				var tmpJBlockedAdvDomains string

				tok = fs.Scan()
				if tok == fflib.FFTok_error {
					goto tokerror
				}
				if tok == fflib.FFTok_right_brace {
					break
				}

				if tok == fflib.FFTok_comma {
					if wantVal == true {
						// TODO(pquerna): this isn't an ideal error message, this handles
						// things like [,,,] as an array value.
						return fs.WrapErr(fmt.Errorf("wanted value token, but got token: %v", tok))
					}
					continue
				} else {
					wantVal = true
				}

				/* handler: tmpJBlockedAdvDomains type=string kind=string quoted=false*/

				{

					{
						if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
							return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
						}
					}

					if tok == fflib.FFTok_null {

					} else {

						outBuf := fs.Output.Bytes()

						tmpJBlockedAdvDomains = string(string(outBuf))

					}
				}

				j.BlockedAdvDomains = append(j.BlockedAdvDomains, tmpJBlockedAdvDomains)

				wantVal = false
			}
		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_BlockedApps:

	/* handler: j.BlockedApps type=[]string kind=slice quoted=false*/

	{

		{
			if tok != fflib.FFTok_left_brace && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for ", tok))
			}
		}

		if tok == fflib.FFTok_null {
			j.BlockedApps = nil
		} else {

			j.BlockedApps = []string{}

			wantVal := true

			for {

				var tmpJBlockedApps string

				tok = fs.Scan()
				if tok == fflib.FFTok_error {
					goto tokerror
				}
				if tok == fflib.FFTok_right_brace {
					break
				}

				if tok == fflib.FFTok_comma {
					if wantVal == true {
						// TODO(pquerna): this isn't an ideal error message, this handles
						// things like [,,,] as an array value.
						return fs.WrapErr(fmt.Errorf("wanted value token, but got token: %v", tok))
					}
					continue
				} else {
					wantVal = true
				}

				/* handler: tmpJBlockedApps type=string kind=string quoted=false*/

				{

					{
						if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
							return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
						}
					}

					if tok == fflib.FFTok_null {

					} else {

						outBuf := fs.Output.Bytes()

						tmpJBlockedApps = string(string(outBuf))

					}
				}

				j.BlockedApps = append(j.BlockedApps, tmpJBlockedApps)

				wantVal = false
			}
		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Source:

	/* handler: j.Source type=openrtb.Source kind=struct quoted=false*/

	{
		/* Falling back. type=openrtb.Source kind=struct */
		tbuf, err := fs.CaptureField(tok)
		if err != nil {
			return fs.WrapErr(err)
		}

		err = json.Unmarshal(tbuf, &j.Source)
		if err != nil {
			return fs.WrapErr(err)
		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Regulations:

	/* handler: j.Regulations type=openrtb.Regulations kind=struct quoted=false*/

	{
		/* Falling back. type=openrtb.Regulations kind=struct */
		tbuf, err := fs.CaptureField(tok)
		if err != nil {
			return fs.WrapErr(err)
		}

		err = json.Unmarshal(tbuf, &j.Regulations)
		if err != nil {
			return fs.WrapErr(err)
		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Ext:

	/* handler: j.Ext type=json.RawMessage kind=slice quoted=false*/

	{
		if tok == fflib.FFTok_null {

		} else {

			tbuf, err := fs.CaptureField(tok)
			if err != nil {
				return fs.WrapErr(err)
			}

			err = j.Ext.UnmarshalJSON(tbuf)
			if err != nil {
				return fs.WrapErr(err)
			}
		}
		state = fflib.FFParse_after_value
	}

	state = fflib.FFParse_after_value
	goto mainparse

wantedvalue:
	return fs.WrapErr(fmt.Errorf("wanted value token, but got token: %v", tok))
wrongtokenerror:
	return fs.WrapErr(fmt.Errorf("ffjson: wanted token: %v, but got token: %v output=%s", wantedTok, tok, fs.Output.String()))
tokerror:
	if fs.BigError != nil {
		return fs.WrapErr(fs.BigError)
	}
	err = fs.Error.ToError()
	if err != nil {
		return fs.WrapErr(err)
	}
	panic("ffjson-generated: unreachable, please report bug.")
done:

	return nil
}