package main

import (
	twin "github.com/TwiN/go-color"
	fatih "github.com/fatih/color"
	"go.uber.org/zap"
	"time"
)

var Reset = "\033[0m"
var Red = "\033[31m"
var Green = "\033[32m"
var Yellow = "\033[33m"
var Blue = "\033[34m"
var Purple = "\033[35m"
var Cyan = "\033[36m"
var Gray = "\033[37m"
var White = "\033[97m"

const url = "http://www.google.com"

func main() {
	colorLog()
	nicePerformance()
	criticalPerformance()
}

func colorLog() {
	// Create a new color object
	c := fatih.New(fatih.FgCyan).Add(fatih.Underline)
	c.Println("Prints cyan text with an underline.")

	// Or just add them to New()
	d := fatih.New(fatih.FgCyan, fatih.Bold)
	d.Printf("This prints bold cyan %s\n", "too!.")

	// Mix up foreground and background colors, create new mixes!
	red := fatih.New(fatih.FgRed)

	boldRed := red.Add(fatih.Bold)
	boldRed.Println("This will print text in bold red.")

	whiteBackground := red.Add(fatih.BgWhite)
	whiteBackground.Println("Red text with white background.")

	info := fatih.New(fatih.FgHiWhite, fatih.BgYellow)

	info.Println("Initiating use color log")
	info.Println("Using colorize log")
}

func criticalPerformance() {
	logger, _ := zap.NewDevelopment()
	defer logger.Sync()
	logger.Info("CRITICAL PERF - failed to fetch URL",
		// Structured context as strongly typed Field values.
		zap.String("url", url),
		zap.Int("attempt", 3),
		zap.Duration("backoff", time.Second),
	)

	logger.Warn(Yellow + "WARNING - ...." + Reset)
	logger.Error(twin.Ize(twin.Red, "Something went wrong"))
	logger.Info(fatih.CyanString("My log with cyan"))
}

func nicePerformance() {
	logger, _ := zap.NewDevelopment()
	defer logger.Sync() // flushes buffer, if any
	sugar := logger.Sugar()
	sugar.Infow("NICE PERF - failed to fetch URL",
		// Structured context as loosely typed key-value pairs.
		"url", url,
		"attempt", 3,
		"backoff", time.Second,
	)
	sugar.Infof("Failed to fetch URL: %s", url)
}
