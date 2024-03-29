package cognitoidentity

import "github.com/unicloud-uos/uos-sdk-go/aws/request"

func init() {
	initRequest = func(r *request.Request) {
		switch r.Operation.Name {
		case opGetOpenIdToken, opGetId, opGetCredentialsForIdentity:
			r.Handlers.Sign.Clear() // these operations are unsigned
		}
	}
}
