package pckrAuth

import (
	"fmt"
	"testing"
)

func TestPremier (t *testing.T) {
	data := `{
        	"figureMngr": {
        	        "ntwrkId": {
	                        "part1": "",
                        	"part2": "",
                	        "idAthntcVrfctnFile": ""
        	        },
	                "accessInfo": {
                	        "userName" : "",
        	                "passWord" : ""
	                },
                	"db1": ""
        	},
	        "userAthrztInfo": {
                	"user-x", {
        	                "athrztId" : "aaaa0000iiiiaaaa",
	                        "athrztKey": "Dgd4gFguDTRFJ57d"
                	}
        	}
	}`
	/*--1--*/
	_x5100, _x5200 := RsblVldtAthrzt (
		"x"               , []string {"aaaa0000iiiiaaaa", "Dgd4gFguDTRFJ57d", data},
	); fmt.Println (_x5100, _x5200)
	_x9100, _x9200 := RsblVldtAthrzt (
		"chf2aubi6n5sv9j7", []string {"chla4swto9qllfr5", "_NM6i23S5[,xMZq^", data},
	); fmt.Println (_x9100, _x9200)
}
