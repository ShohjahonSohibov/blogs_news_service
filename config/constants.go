package config

import "time"

const (
	DatabaseQueryTimeLayout = `'YYYY-MM-DD"T"HH24:MI:SS"."MS"Z"TZ'`
	// DatabaseTimeLayout
	DatabaseTimeLayout string = time.RFC3339
	// AccessTokenExpiresInTime ...
	AccessTokenExpiresInTime time.Duration = 30 * 24 * 60 * time.Minute
	// RefreshTokenExpiresInTime ...
	RefreshTokenExpiresInTime time.Duration = 30 * 24 * 60 * time.Minute

	// Corp
	AccessTokenExpiresInTimeCorp time.Duration = 06 * 60 * time.Minute
	// OtpExpiresInTime ...
	OtpExpiresInTime time.Duration = 5 * time.Minute
	// OtpTemplateText ...
	OtpTemplateTextRu string = "Используйте код %s для авторизации mediapark.uz\n%s"
	OtpTemplateTextUz string = "Avtorizatsiya uchun %s kodidan foydalaning mediapark.uz\n%s"
	OtpTemplateTextKr string = "Aвторизация учун %s кодидан фойдаланинг mediapark.uz\n%s"

	SlugCashbackNewUserPrice = "new_user_price"

	SourceWeb    = "web"
	SourceMobile = "mobile"
)

const (
	RoleIDCorperativeUser = "67896169-951d-43b4-ad54-bf1a3513878c"
	RoleIDUser            = "06d63125-e7a2-4616-afa4-cd50ee3ac33d"
	OperatorUserRoleID    = "e0085103-fe31-4f8e-9e73-177c9aa1821d"
)
