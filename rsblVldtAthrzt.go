//-- p --
package pckrAuth

//-- r --
import (
	"fmt"
	"github.com/tidwall/gjson"
)

//-- i --
func RsblVldtAthrzt (userId string, athrztDtls []string) (err error, vldty bool) {
	athrztIdPath := fmt.Sprintf ("userAthrztInfo.user-%s.athrztId", userId)
	athrztKeyPath := fmt.Sprintf ("userAthrztInfo.user-%s.athrztKey", userId)

	_x5100 := gjson.Get (athrztDtls [2], athrztIdPath).String ()
	_x5200 := gjson.Get (athrztDtls [2], athrztKeyPath).String ()

	if _x5100 != "" && _x5100 == athrztDtls [0] && _x5200 != && _x5200 == athrztDtls [1] {
		return nil, true
	} else {
		return nil, false
	}
}
