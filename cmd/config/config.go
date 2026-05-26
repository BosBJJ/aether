package config

import "database/sql"


var (
    OutputFormat string
    SaveToDB     bool
    Threads      int
    Timeout      int
    Quiet        bool
    Verbose      bool
    Pretty       bool
    DB           *sql.DB
)

func GetOutputFormat() string {
    return OutputFormat
}

func ShouldSave() bool {
    return SaveToDB
}

func GetThreads() int {
    if Threads <= 0 {
        return 50
    }
    return Threads
}

func GetTimeout() int {
    if Timeout <= 0 {
        return 5
    }
    return Timeout
}

func IsQuiet() bool {
    return Quiet
}

func IsVerbose() bool {
    return Verbose
}

func GetPretty() bool {
    return Pretty
}