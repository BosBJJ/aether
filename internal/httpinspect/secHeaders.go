package httpinspect

type HeaderCheck struct {
	Name           string `json:"name"`
	DisplayName    string `json:"display_name"`
	GoodLabel      string `json:"good_label"`
	Risk           string `json:"risk"`
	Recommendation string `json:"recommendation"`
}




var SecurityHeaders = []HeaderCheck{
	{
		Name:           "Strict-Transport-Security",
		DisplayName:    "HSTS",
		GoodLabel:      "Enabled",
		Risk:           "Traffic can be downgraded to insecure HTTP",
		Recommendation: "Add: Strict-Transport-Security: max-age=31536000; includeSubDomains",
	},
	{
		Name:           "Content-Security-Policy",
		DisplayName:    "CSP",
		GoodLabel:      "Present",
		Risk:           "Cross-site scripting (XSS) attacks are possible",
		Recommendation: "Add a strong Content-Security-Policy header",
	},
	{
		Name:           "X-Frame-Options",
		DisplayName:    "X-Frame-Options",
		GoodLabel:      "Present",
		Risk:           "Clickjacking attacks are possible",
		Recommendation: "Add: X-Frame-Options: DENY or SAMEORIGIN",
	},
	{
		Name:           "X-Content-Type-Options",
		DisplayName:    "X-Content-Type-Options",
		GoodLabel:      "nosniff",
		Risk:           "MIME-type sniffing attacks possible",
		Recommendation: "Add: X-Content-Type-Options: nosniff",
	},
	{
		Name:           "Referrer-Policy",
		DisplayName:    "Referrer-Policy",
		GoodLabel:      "Present",
		Risk:           "Sensitive referrer information may leak to other sites",
		Recommendation: "Add: Referrer-Policy: strict-origin-when-cross-origin",
	},
	{
		Name:           "Permissions-Policy",
		DisplayName:    "Permissions-Policy",
		GoodLabel:      "Present",
		Risk:           "Browser features (camera, mic, geolocation) are unrestricted",
		Recommendation: "Add Permissions-Policy to control browser features",
	},
	{
		Name:           "Cross-Origin-Opener-Policy",
		DisplayName:    "COOP",
		GoodLabel:      "Present",
		Risk:           "Cross-origin attacks (e.g. Spectre) are more likely",
		Recommendation: "Add: Cross-Origin-Opener-Policy: same-origin",
	},
	{
		Name:           "Server",
		DisplayName:    "Server",
		GoodLabel:      "Hidden",
		Risk:           "Leaking server software version helps attackers",
		Recommendation: "Hide or obscure the Server header",
	},
	{
		Name:           "X-Powered-By",
		DisplayName:    "X-Powered-By",
		GoodLabel:      "Hidden",
		Risk:           "Reveals backend technology and version",
		Recommendation: "Remove X-Powered-By header",
	},
	{
		Name:           "X-XSS-Protection",
		DisplayName:    "X-XSS-Protection",
		GoodLabel:      "Enabled",
		Risk:           "Older browsers have reduced XSS protection",
		Recommendation: "Add: X-XSS-Protection: 1; mode=block (if needed)",
	},
	{
		Name:           "X-AspNet-Version",
		DisplayName:    "X-AspNet-Version",
		GoodLabel:      "Hidden",
		Risk:           "Reveals ASP.NET version to attackers",
		Recommendation: "Remove X-AspNet-Version header",
	},
}