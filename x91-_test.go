package pckrAuth

import (
	"fmt"
	"testing"
)

func TestPremier (t *testing.T) {
	data := `
{
        "figureMngr": {
                "ntwrkId": {
                        "part1": "",
                        "part2": "",
                        "idAthntcVrfctnFile": ""
                },
                "accessInfo": {
                        "userName": "",
                        "passWord": ""
                },
                "db1": ""
        },
        "userAthrztInfo": {
                "user-x", {
                        "athrztId": "flap",
                        "athrztKey": "can"
                }
        }
}
`
	_x5100, _x5200 := RsblVldtAthrzt ("x", []string {"flap", "can", data})
	fmt.Println (_x5100, _x5200)
	
	_x6100, _x6200 := RsblVldtAthrzt ("y", []string {"flap", "can", data})
	fmt.Println (_x6100, _x6200)
	
	_x7100, _x7200 := RsblVldtAthrzt ("x", []string {"xlap", "can", data})
	fmt.Println (_x7100, _x7200)
	
	_x8100, _x8200 := RsblVldtAthrzt ("x", []string {"flap", "xan", data})
	fmt.Println (_x8100, _x8200)

	_x9100, _x9200 := RsblVldtAthrzt ("", []string {"", "", data})
	fmt.Println (_x9100, _x9200)
}
