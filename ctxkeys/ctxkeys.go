package ctxkeys

import (
	"context"
)

type String string

const UserNameKey String = "user"
const RequestIDKey String = "requestid"
const DBConnKey String = "db"
const UserIDKey String = "userid"
const TokenKey String = "token"
const EnvelopeKey String = "envelope"
const LoggerKey String = "logger"
const ByteBodyKey String = "bytebody"
const QueryParamsKey String = "queryparams"
const AppConfigKey String = "appconfig"

func GetValue(ctx context.Context, key String) any {
	return ctx.Value(key)
}
