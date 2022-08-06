package interest


import (
	"net/http"
	"github.com/dingtra/rundb"
	"fmt"
)

func Http (w http.ResponseWriter, r *http.Request) {
	session, _ := rundb.Store.Get(r, "session")
	let := InterestedStruct{}

	let.VerifyInterested(r, session.Values["usersid"].(string))

	if let.Success == true {
		fmt.Fprintf(w, "%s", let.Details)
	}
}