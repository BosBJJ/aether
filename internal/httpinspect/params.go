package httpinspect

type URLAnalysis struct {
    OriginalURL string		`json:"original_url"`
    FinalURL    string		`json:"final_url"`
    Redirects   []string	`json:"redirects"`
    Findings    []Finding	`json:"findings"`
    RiskLevel   string		`json:"risk_level"`
}

type Finding struct {
    URL         string		`json:"url"`
    Description string		`json:"description"`
    Severity    string		`json:"severity"`
}


// Detection data sourced from:
// - OWASP Testing Guide (open redirects, sensitive parameters)
// - PayloadsAllTheThings (https://github.com/swisskyrepo/PayloadsAllTheThings)
// - Common bug bounty and phishing research patterns


var SuspiciousParams = map[string]string{
    // High Risk
    "redirect":    	"High",
    "redir":      	"High",
    "next":        	"High",
    "url":         	"High",
    "return":      	"High",
    "return_to":   	"High",
    "goto":        	"High",
    "destination": 	"High",
    "continue":    	"High",
    "forward":     	"High",
    "callback":    	"High",
	"sub1":			"High",
	"sub2": 		"High",
	"sub3": 		"High",
	"sub4": 		"High",
	"sub5": 		"High",

    // Medium Risk
    "link":     	"Medium",
    "path":     	"Medium",
    "file":     	"Medium",
    "download": 	"Medium",
    "source":   	"Medium",
    "ref":      	"Medium",
    "referer":  	"Medium",
    "location": 	"Medium",
    "to":      		"Medium",
    "target":  		"Medium",
    
	// Sensitive
    "token":    	"Medium",
    "auth":     	"Medium",
    "login":    	"Medium",
    "logout":   	"Medium",
    "reset":    	"Medium",
    "recover":  	"Medium",
    "password": 	"Medium",
}

// Suspicious TLDs sourced from Spamhaus TLD reputation data (spamhaus.org/statistics/tlds/)

var SuspiciousTLDs = []string{
	".tk",
	".ml",
	".ga",
	".cf",
	".gq",
	".top",
	".xyz",
	".bond",
	".xin",
	".cfd",
	".vip",
	".icu",
	".lol",
	".fun",
	".pw",
	".cc",
	".sbs",
	".club",
	".online",
	".info",
	".party",
	".click",
	".bid",
	".date",
	".stream",
	".work",
	".shop",
	".site",
	".live",
}