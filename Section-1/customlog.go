package main

import "fmt"

type logLevel int

const (
	LogTrace logLevel = iota
	LogInfo
	LogDebug
	LogWarning
	LogError
)

var levelNames = []string{"Trace", "Info", "Debug", "Warning", "Error"}

func (l logLevel) String() string {
	if l < LogTrace || l > LogError {
		return "Unknown"
	}
	return levelNames[l]
}

func printLogLevel(level logLevel) {
	fmt.Printf("Log Level { %d : %s }\n", level, level.String())
}
func main() {

	printLogLevel(LogTrace)
	printLogLevel(LogInfo)
	printLogLevel(LogDebug)
	printLogLevel(LogWarning)
	printLogLevel(LogError)
}
