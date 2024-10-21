package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

var logLevels = []string{"INFO", "WARN", "ERROR", "DEBUG"}

func main() {
	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// Simulate log generation
	for i := 0; i < 100; i++ {
		logMessage := generateLog()
		fmt.Println(logMessage)
		time.Sleep(50 * time.Millisecond) // Wait for 250 milliseconds before generating the next log
	}
}

func generateLog() string {
	level := getRandomLogLevel()
	timestamp := time.Now().Format("2006-01-02T15:04:05.999999") // RFC 3339 format with microsecond precision
	pid := os.Getpid()                                           // Get the process ID
	message := randomMessage()
	return fmt.Sprintf("%s, [%s #%d] %s -- Main: %s", levelPrefix(level), timestamp, pid, level, message)
}

func levelPrefix(level string) string {
	switch level {
	case "INFO":
		return "I"
	case "WARN":
		return "W"
	case "ERROR":
		return "E"
	case "DEBUG":
		return "D"
	default:
		return "I" // Default to INFO if not found
	}
}

func getRandomLogLevel() string {
	return logLevels[rand.Intn(len(logLevels))]
}

func randomMessage() string {
	messages := []string{
		"Check out this info!",
		"User logged in successfully.",
		"Started GET \"/posts\" for 127.0.0.1 at",
		"Processing by PostsController#index as HTML",
		"Completed 200 OK in 45ms (Views: 30.0ms | ActiveRecord: 15.0ms)",
		"Started POST \"/posts\" for 127.0.0.1 at",
		"Rendered posts/index.html.erb within layouts/application",
		"Unpermitted parameters: :secret",
		"Started GET \"/users/1\" for 127.0.0.1 at",
		"Completed 404 Not Found in 12ms (Views: 10.0ms | ActiveRecord: 2.0ms)",
		"Processing by UsersController#show as HTML",
		"SQL (0.5ms)  SELECT \"users\".* FROM \"users\" WHERE \"users\".\"id\" = $1 LIMIT $2  [[\"id\", 1], [\"LIMIT\", 1]]",
		"Started PATCH \"/users/1\" for 127.0.0.1 at",
		"Completed 204 No Content in 35ms",
		"Processing by CommentsController#create as HTML",
		"Redirected to /posts/1",
		"Completed 302 Found in 20ms (ActiveRecord: 10.0ms)",
		"Started DELETE \"/posts/1\" for 127.0.0.1 at",
		"Completed 500 Internal Server Error in 50ms (Views: 0.0ms | ActiveRecord: 25.0ms)",
		"Error rendering post: Post not found",
	}
	return messages[rand.Intn(len(messages))]
}
