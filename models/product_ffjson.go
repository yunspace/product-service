// DO NOT EDIT!
// Code generated by ffjson <https://github.com/pquerna/ffjson>
// source: product.go
// DO NOT EDIT!

package models

import (
	"bytes"
	"errors"
	"fmt"
	fflib "github.com/pquerna/ffjson/fflib/v1"
)

func (mj *Product) MarshalJSON() ([]byte, error) {
	var buf fflib.Buffer
	if mj == nil {
		buf.WriteString("null")
		return buf.Bytes(), nil
	}
	err := mj.MarshalJSONBuf(&buf)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
func (mj *Product) MarshalJSONBuf(buf fflib.EncodingBuffer) error {
	if mj == nil {
		buf.WriteString("null")
		return nil
	}
	var err error
	var obj []byte
	_ = obj
	_ = err
	buf.WriteString(`{"Code":`)
	fflib.WriteJsonString(buf, string(mj.Code))
	buf.WriteString(`,"Name":`)
	fflib.WriteJsonString(buf, string(mj.Name))
	buf.WriteString(`,"Description":`)
	fflib.WriteJsonString(buf, string(mj.Description))
	buf.WriteString(`,"Price":`)
	fflib.AppendFloat(buf, float64(mj.Price), 'g', -1, 64)
	buf.WriteString(`,"Expiry":`)
	fflib.FormatBits2(buf, uint64(mj.Expiry), 10, mj.Expiry < 0)
	if mj.IsPlan {
		buf.WriteString(`,"IsPlan":true`)
	} else {
		buf.WriteString(`,"IsPlan":false`)
	}
	if mj.IsUnlimited {
		buf.WriteString(`,"IsUnlimited":true`)
	} else {
		buf.WriteString(`,"IsUnlimited":false`)
	}
	buf.WriteString(`,"SizeMb":`)
	fflib.FormatBits2(buf, uint64(mj.SizeMb), 10, mj.SizeMb < 0)
	if mj.Is4g {
		buf.WriteString(`,"Is4g":true`)
	} else {
		buf.WriteString(`,"Is4g":false`)
	}
	if mj.AutoRenew {
		buf.WriteString(`,"AutoRenew":true`)
	} else {
		buf.WriteString(`,"AutoRenew":false`)
	}
	buf.WriteString(`,"TermsUrl":`)
	fflib.WriteJsonString(buf, string(mj.TermsUrl))
	buf.WriteByte('}')
	return nil
}

const (
	ffj_t_Productbase = iota
	ffj_t_Productno_such_key

	ffj_t_Product_Code

	ffj_t_Product_Name

	ffj_t_Product_Description

	ffj_t_Product_Price

	ffj_t_Product_Expiry

	ffj_t_Product_IsPlan

	ffj_t_Product_IsUnlimited

	ffj_t_Product_SizeMb

	ffj_t_Product_Is4g

	ffj_t_Product_AutoRenew

	ffj_t_Product_TermsUrl
)

var ffj_key_Product_Code = []byte("Code")

var ffj_key_Product_Name = []byte("Name")

var ffj_key_Product_Description = []byte("Description")

var ffj_key_Product_Price = []byte("Price")

var ffj_key_Product_Expiry = []byte("Expiry")

var ffj_key_Product_IsPlan = []byte("IsPlan")

var ffj_key_Product_IsUnlimited = []byte("IsUnlimited")

var ffj_key_Product_SizeMb = []byte("SizeMb")

var ffj_key_Product_Is4g = []byte("Is4g")

var ffj_key_Product_AutoRenew = []byte("AutoRenew")

var ffj_key_Product_TermsUrl = []byte("TermsUrl")

func (uj *Product) UnmarshalJSON(input []byte) error {
	fs := fflib.NewFFLexer(input)
	return uj.UnmarshalJSONFFLexer(fs, fflib.FFParse_map_start)
}

func (uj *Product) UnmarshalJSONFFLexer(fs *fflib.FFLexer, state fflib.FFParseState) error {
	var err error = nil
	currentKey := ffj_t_Productbase
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
				currentKey = ffj_t_Productno_such_key
				state = fflib.FFParse_want_colon
				goto mainparse
			} else {
				switch kn[0] {

				case 'A':

					if bytes.Equal(ffj_key_Product_AutoRenew, kn) {
						currentKey = ffj_t_Product_AutoRenew
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'C':

					if bytes.Equal(ffj_key_Product_Code, kn) {
						currentKey = ffj_t_Product_Code
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'D':

					if bytes.Equal(ffj_key_Product_Description, kn) {
						currentKey = ffj_t_Product_Description
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'E':

					if bytes.Equal(ffj_key_Product_Expiry, kn) {
						currentKey = ffj_t_Product_Expiry
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'I':

					if bytes.Equal(ffj_key_Product_IsPlan, kn) {
						currentKey = ffj_t_Product_IsPlan
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffj_key_Product_IsUnlimited, kn) {
						currentKey = ffj_t_Product_IsUnlimited
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffj_key_Product_Is4g, kn) {
						currentKey = ffj_t_Product_Is4g
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'N':

					if bytes.Equal(ffj_key_Product_Name, kn) {
						currentKey = ffj_t_Product_Name
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'P':

					if bytes.Equal(ffj_key_Product_Price, kn) {
						currentKey = ffj_t_Product_Price
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'S':

					if bytes.Equal(ffj_key_Product_SizeMb, kn) {
						currentKey = ffj_t_Product_SizeMb
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'T':

					if bytes.Equal(ffj_key_Product_TermsUrl, kn) {
						currentKey = ffj_t_Product_TermsUrl
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				}

				if fflib.EqualFoldRight(ffj_key_Product_TermsUrl, kn) {
					currentKey = ffj_t_Product_TermsUrl
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffj_key_Product_AutoRenew, kn) {
					currentKey = ffj_t_Product_AutoRenew
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffj_key_Product_Is4g, kn) {
					currentKey = ffj_t_Product_Is4g
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffj_key_Product_SizeMb, kn) {
					currentKey = ffj_t_Product_SizeMb
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffj_key_Product_IsUnlimited, kn) {
					currentKey = ffj_t_Product_IsUnlimited
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffj_key_Product_IsPlan, kn) {
					currentKey = ffj_t_Product_IsPlan
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffj_key_Product_Expiry, kn) {
					currentKey = ffj_t_Product_Expiry
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffj_key_Product_Price, kn) {
					currentKey = ffj_t_Product_Price
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffj_key_Product_Description, kn) {
					currentKey = ffj_t_Product_Description
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffj_key_Product_Name, kn) {
					currentKey = ffj_t_Product_Name
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffj_key_Product_Code, kn) {
					currentKey = ffj_t_Product_Code
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				currentKey = ffj_t_Productno_such_key
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

				case ffj_t_Product_Code:
					goto handle_Code

				case ffj_t_Product_Name:
					goto handle_Name

				case ffj_t_Product_Description:
					goto handle_Description

				case ffj_t_Product_Price:
					goto handle_Price

				case ffj_t_Product_Expiry:
					goto handle_Expiry

				case ffj_t_Product_IsPlan:
					goto handle_IsPlan

				case ffj_t_Product_IsUnlimited:
					goto handle_IsUnlimited

				case ffj_t_Product_SizeMb:
					goto handle_SizeMb

				case ffj_t_Product_Is4g:
					goto handle_Is4g

				case ffj_t_Product_AutoRenew:
					goto handle_AutoRenew

				case ffj_t_Product_TermsUrl:
					goto handle_TermsUrl

				case ffj_t_Productno_such_key:
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

handle_Code:

	/* handler: uj.Code type=string kind=string quoted=false*/

	{

		{
			if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
			}
		}

		if tok == fflib.FFTok_null {

		} else {

			outBuf := fs.Output.Bytes()

			uj.Code = string(string(outBuf))

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Name:

	/* handler: uj.Name type=string kind=string quoted=false*/

	{

		{
			if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
			}
		}

		if tok == fflib.FFTok_null {

		} else {

			outBuf := fs.Output.Bytes()

			uj.Name = string(string(outBuf))

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Description:

	/* handler: uj.Description type=string kind=string quoted=false*/

	{

		{
			if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
			}
		}

		if tok == fflib.FFTok_null {

		} else {

			outBuf := fs.Output.Bytes()

			uj.Description = string(string(outBuf))

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Price:

	/* handler: uj.Price type=float64 kind=float64 quoted=false*/

	{
		if tok != fflib.FFTok_double && tok != fflib.FFTok_integer && tok != fflib.FFTok_null {
			return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for float64", tok))
		}
	}

	{

		if tok == fflib.FFTok_null {

		} else {

			tval, err := fflib.ParseFloat(fs.Output.Bytes(), 64)

			if err != nil {
				return fs.WrapErr(err)
			}

			uj.Price = float64(tval)

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Expiry:

	/* handler: uj.Expiry type=int32 kind=int32 quoted=false*/

	{
		if tok != fflib.FFTok_integer && tok != fflib.FFTok_null {
			return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for int32", tok))
		}
	}

	{

		if tok == fflib.FFTok_null {

		} else {

			tval, err := fflib.ParseInt(fs.Output.Bytes(), 10, 32)

			if err != nil {
				return fs.WrapErr(err)
			}

			uj.Expiry = int32(tval)

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_IsPlan:

	/* handler: uj.IsPlan type=bool kind=bool quoted=false*/

	{
		if tok != fflib.FFTok_bool && tok != fflib.FFTok_null {
			return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for bool", tok))
		}
	}

	{
		if tok == fflib.FFTok_null {

		} else {
			tmpb := fs.Output.Bytes()

			if bytes.Compare([]byte{'t', 'r', 'u', 'e'}, tmpb) == 0 {

				uj.IsPlan = true

			} else if bytes.Compare([]byte{'f', 'a', 'l', 's', 'e'}, tmpb) == 0 {

				uj.IsPlan = false

			} else {
				err = errors.New("unexpected bytes for true/false value")
				return fs.WrapErr(err)
			}

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_IsUnlimited:

	/* handler: uj.IsUnlimited type=bool kind=bool quoted=false*/

	{
		if tok != fflib.FFTok_bool && tok != fflib.FFTok_null {
			return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for bool", tok))
		}
	}

	{
		if tok == fflib.FFTok_null {

		} else {
			tmpb := fs.Output.Bytes()

			if bytes.Compare([]byte{'t', 'r', 'u', 'e'}, tmpb) == 0 {

				uj.IsUnlimited = true

			} else if bytes.Compare([]byte{'f', 'a', 'l', 's', 'e'}, tmpb) == 0 {

				uj.IsUnlimited = false

			} else {
				err = errors.New("unexpected bytes for true/false value")
				return fs.WrapErr(err)
			}

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_SizeMb:

	/* handler: uj.SizeMb type=int32 kind=int32 quoted=false*/

	{
		if tok != fflib.FFTok_integer && tok != fflib.FFTok_null {
			return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for int32", tok))
		}
	}

	{

		if tok == fflib.FFTok_null {

		} else {

			tval, err := fflib.ParseInt(fs.Output.Bytes(), 10, 32)

			if err != nil {
				return fs.WrapErr(err)
			}

			uj.SizeMb = int32(tval)

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Is4g:

	/* handler: uj.Is4g type=bool kind=bool quoted=false*/

	{
		if tok != fflib.FFTok_bool && tok != fflib.FFTok_null {
			return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for bool", tok))
		}
	}

	{
		if tok == fflib.FFTok_null {

		} else {
			tmpb := fs.Output.Bytes()

			if bytes.Compare([]byte{'t', 'r', 'u', 'e'}, tmpb) == 0 {

				uj.Is4g = true

			} else if bytes.Compare([]byte{'f', 'a', 'l', 's', 'e'}, tmpb) == 0 {

				uj.Is4g = false

			} else {
				err = errors.New("unexpected bytes for true/false value")
				return fs.WrapErr(err)
			}

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_AutoRenew:

	/* handler: uj.AutoRenew type=bool kind=bool quoted=false*/

	{
		if tok != fflib.FFTok_bool && tok != fflib.FFTok_null {
			return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for bool", tok))
		}
	}

	{
		if tok == fflib.FFTok_null {

		} else {
			tmpb := fs.Output.Bytes()

			if bytes.Compare([]byte{'t', 'r', 'u', 'e'}, tmpb) == 0 {

				uj.AutoRenew = true

			} else if bytes.Compare([]byte{'f', 'a', 'l', 's', 'e'}, tmpb) == 0 {

				uj.AutoRenew = false

			} else {
				err = errors.New("unexpected bytes for true/false value")
				return fs.WrapErr(err)
			}

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_TermsUrl:

	/* handler: uj.TermsUrl type=string kind=string quoted=false*/

	{

		{
			if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
			}
		}

		if tok == fflib.FFTok_null {

		} else {

			outBuf := fs.Output.Bytes()

			uj.TermsUrl = string(string(outBuf))

		}
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
