package httpinspect


type TechSignature struct {
    Name        string
    Category    string
    Confidence  int
    Headers     map[string]string
    Cookies     []string
    BodyPattern []string
}
// Patterns sourced and refined from Wappalyzer and manual research.
var FingerprintDB = []TechSignature{
	// === Web Servers ===
	{Name: "Nginx", Category: "Web Server", Confidence: 9, Headers: map[string]string{"Server": "nginx"}},
	{Name: "Apache", Category: "Web Server", Confidence: 8, Headers: map[string]string{"Server": "apache"}},
	{Name: "OpenResty", Category: "Web Server", Confidence: 8, Headers: map[string]string{"Server": "openresty"}},

	// === CDNs ===
	{Name: "Cloudflare", Category: "CDN", Confidence: 9, Headers: map[string]string{"Server": "cloudflare"}, Cookies: []string{"__cfduid", "__cfruid"}},
	{Name: "Akamai", Category: "CDN", Confidence: 7, Headers: map[string]string{"Server": "akamai"}},

	// === Backend ===
	{Name: "Go (Golang)", Category: "Backend", Confidence: 7, Headers: map[string]string{"Server": "go"}},
	{Name: "Laravel", Category: "Backend", Confidence: 8, Headers: map[string]string{"X-Powered-By": "laravel"}, Cookies: []string{"laravel_session"}},
	{Name: "Django", Category: "Backend", Confidence: 8, Cookies: []string{"csrftoken"}, BodyPattern: []string{"csrfmiddlewaretoken"}},
	{Name: "PHP", Category: "Backend", Confidence: 6, Headers: map[string]string{"X-Powered-By": "php"}, Cookies: []string{"PHPSESSID"}},
	{Name: "Node.js", Category: "Backend", Confidence: 6, Headers: map[string]string{"X-Powered-By": "express"}},

	// === Frontend ===
	{Name: "React", Category: "Frontend", Confidence: 8, BodyPattern: []string{"react-dom", "__REACT_DEVTOOLS_GLOBAL_HOOK__"}},
	{Name: "Vue.js", Category: "Frontend", Confidence: 8, BodyPattern: []string{"data-v-"}},
	{Name: "Angular", Category: "Frontend", Confidence: 7, BodyPattern: []string{"ng-version"}},

	// === CMS ===
	{Name: "WordPress", Category: "CMS", Confidence: 9, BodyPattern: []string{"wp-content", "wp-includes", "wp-emoji"}},
	{Name: "Shopify", Category: "CMS", Confidence: 8, BodyPattern: []string{"shopify.com", "cdn.shopify"}},
	{Name: "Wix", Category: "CMS", Confidence: 7, BodyPattern: []string{"_wix"}},

	// === Analytics ===
	{Name: "Google Analytics", Category: "Analytics", Confidence: 8, BodyPattern: []string{"google-analytics.com", "gtag/js"}},
	{Name: "Hotjar", Category: "Analytics", Confidence: 7, BodyPattern: []string{"hotjar"}},

	// === Error Tracking ===
	{Name: "Sentry", Category: "Error Tracking", Confidence: 7, BodyPattern: []string{"_sentry", "sentry.io"}},
}