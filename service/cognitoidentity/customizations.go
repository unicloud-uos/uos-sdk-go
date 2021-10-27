package cognitoidentity

import "gitlab.rd.unicloud.com/Storage-oss-service/uos-sdk-go.git/aws/request"

func init() {
	initRequest = func(r *request.Request) {
		switch r.Operation.Name {
		case opGetOpenIdToken, opGetId, opGetCredentialsForIdentity:
			r.Handlers.Sign.Clear() // these operations are unsigned
		}
	}
}
