package pckrAuth

import  "errors"
import  "fmt"
import  "github.com/qb-qetell/httpMssg"
import  "github.com/tidwall/gjson"
import  "regexp"
import  "time"

func RsblVldtAthrzt (userId string, athrztDtls [ ]string) (err error, vldty bool) {
	if      userId== "" /*--xx--*/ {
		_ca01 := errors.New (fmt.Sprintf (
			 "User ID not provided." ,
		))
		return _ca01, false
	} else if len (athrztDtls) < 3 {
		_ca01 := errors.New (fmt.Sprintf (
			 "Authorization info not complete." ,
		))
		return _ca01, false
	} else if athrztDtls [0] == "" {
		_ca01 := errors.New (fmt.Sprintf (
			 "Authorization ID   not provided." ,
		))
		return _ca01, false
	} else if athrztDtls [1] == "" {
		_ca01 := errors.New (fmt.Sprintf (
			 "Authorization key  not provided." ,
		))
		return _ca01, false
	}
	/*--1--*/
	_ba01 := gjson.Get (
		athrztDtls [2], fmt.Sprintf ("userAthrztInfo.user-%s.athrztId" , userId),
	).String ()
	_ba02 := gjson.Get (
		athrztDtls [2], fmt.Sprintf ("userAthrztInfo.user-%s.athrztKey", userId),
	).String ()
	if _ba01 == athrztDtls [0] && _ba02 == athrztDtls [1] { return nil, true }
	/*--2--*/
	_bb01 := userId  ; _bb11 := athrztDtls [0]; _bb12 :=athrztDtls [1]
	for  len (_bb01  ) < 16 {_bb01 = _bb01 + "q"}
	for  len (_bb11  ) < 16 {_bb11 = _bb11 + "q"}
	_bc40 := httpMssg.Mssg_Estb  ( )
	_bc45 := fmt.Sprintf (
		 `{"cmnd": "entt_%s|athr_%s: athn", "seed": {"athrKeyy": "%s"}}`,
		  _bb01, _bb11 , _bb12 ,
	)
	_bc40.SettMssg  ([     ]byte (_bc45))
	_bc40.SettCmmnMthd ("POST")
	_bc40.SettRcpn  ("https://i1xx.amsx.octamile.com:20008")
	_bc65, _bc70 := _bc40.Send (time.Second * 16)
	if _bc65 != nil {
		_ca00 := fmt.Sprintf (
			"BB01: Could not validate authorization. [%s]",  _bc65.Error (),
		)
		_cb00 := errors.New  (_ca00)
		return _cb00, false
	}
	if _bc70.Rspn.StatusCode != 200 {
		_ca00 := fmt.Sprintf (
			"BB02: Could not validate authorization. [Response code's %d.]",
			_bc70.Rspn.StatusCode,
		)
		_cb00 := errors.New  (_ca00)
		return _cb00, false
	}
	_bc75 := string    (_bc70.ExtrRspn (  ))
	_bc80 := gjson.Get (_bc75,  "yeld.vldt").String      (     )
	if regexp.MustCompile (`^(false|true)$`).MatchString (_bc80) ==  false {
		_ca00 := fmt.Sprintf (
			"BB03: Could not validate authorization. [Response invl. [%s]]",
			_bc75,
		)
		_cb00 := errors.New  (_ca00)
		return _cb00, false
	}
	_bd00 := false
	if _bc80 == "true" { _bd00 = true }
	return nil, _bd00
}
