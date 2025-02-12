package constants

const (
	AllowCredential = "true"
	AllowedHeader   = "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Origin, Cache-Control, X-Requested-With, User-Agent, Accept, Postman-Token, Host, Connection, Accept-Language, Referer, Sec-Fetch-Dest, Sec-Fetch-Mode, Sec-Fetch-Site, sec-ch-ua, Sec-Ch-Ua, sec-ch-ua-mobile, Sec-Ch-Ua-Mobile, sec-ch-ua-platform, Sec-Ch-Ua-Platform" // separate with ", "
	AllowedMethods  = "POST, GET, PUT, DELETE, PATCH, OPTIONS"
	MaxAge          = "43200" // for 12 hour
)
