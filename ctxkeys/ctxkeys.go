package ctxkeys

import (
	"context"
)

type String string

const UserNameKey String = "user"
const RequestIDKey String = "request_id"
const DBConnKey String = "db"
const UserIDKey String = "userid"
const TokenKey String = "token"
const EnvelopeKey String = "envelope"
const LoggerKey String = "logger"
const ByteBodyKey String = "bytebody"

func GetValue(ctx context.Context, key String) any {
	return ctx.Value(key)
}
