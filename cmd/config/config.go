package config


var (
    OutputFormat string
    SaveToDB     bool
    Threads      int
    Timeout      int
    Quiet        bool
    Verbose      bool
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