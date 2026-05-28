package utils

import (
	"aether/internal/database"
	"encoding/json"
	"fmt"
)


func FormatScan(scan database.ScanResult, format string) (string, error) {
    indented, err := json.MarshalIndent(json.RawMessage(scan.RawResult), "", "  ")
    if err != nil {
        return "", err
    }

    switch format {
    case "json":
        return string(indented), nil

    case "table":
        return fmt.Sprintf(`# Aether Scan Report

**Target:** %s
**Command:** %s
**Time:** %s
**Summary:** %s

## Raw Results
`+"```json\n%s\n```", scan.Target, scan.CommandType, scan.Timestamp, scan.Summary, string(indented)), nil

    case "html":
        return fmt.Sprintf(`<!DOCTYPE html>
<html>
<body>
<h1>Aether Scan Report</h1>
<p><strong>Target:</strong> %s</p>
<p><strong>Command:</strong> %s</p>
<p><strong>Time:</strong> %s</p>
<p><strong>Summary:</strong> %s</p>
<h2>Raw Results</h2>
<pre>%s</pre>
</body>
</html>`, scan.Target, scan.CommandType, scan.Timestamp, scan.Summary, string(indented)), nil

    default:
        return string(indented), nil
    }
}