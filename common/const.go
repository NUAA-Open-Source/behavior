package common

// Cross-sites resource sharing settings
var CORS_ALLOW_ORIGINS = []string{
	"https://safeu.a2os.club",
	"https://test.safeu.a2os.club",
	"http://safeu.a2os.club",
	"http://test.safeu.a2os.club",
}

var CORS_ALLOW_DEBUG_ORIGINS = []string{
	"http://*",
	"https://*",
}

var CORS_ALLOW_HEADERS = []string{
	"Origin",
	"Content-Length",
	"Content-Type",
	"Token",
	"X-CSRF-TOKEN",
	"withCredentials",
}

var CORS_ALLOW_METHODS = []string{
	"GET",
	"POST",
	"PUT",
	"PATCH",
	"DELETE",
	"HEAD",
}

var CORS_EXPOSE_HEADERS = []string{
	"X-CSRF-TOKEN",
	"Token",
}

var CSRF_COOKIE_SECRET = []byte("csrf-secret")

const (
	CSRF_SESSION_NAME string = "safeu-session"
	CSRF_SECRET       string = "safeu-secret"
)
