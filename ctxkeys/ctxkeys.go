package ctxkeys

import (
	"context"
)

type String string

const (
	UserNameKey   String = "user"
	RequestIDKey  String = "requestid"
	DBConnKey     String = "db"
	UserIDKey     String = "userid"
	TokenKey      String = "token"
	EnvelopeKey   String = "envelope"
	LoggerKey     String = "logger"
	ByteBodyKey   String = "bytebody"
	PaginationKey String = "pagination"
	AppConfigKey  String = "appconfig"
)

func GetValue(ctx context.Context, key String) any {
	return ctx.Value(key)
}
