module github.com/ghonzo/advent2024

go 1.23

// For generics constraints
require golang.org/x/exp v0.0.0-20241108190413-2d47ceb2692f

// Easier JSON parsing for leaderboard.go
require (
	github.com/tidwall/gjson v1.18.0
	github.com/tidwall/match v1.1.1 // indirect
	github.com/tidwall/pretty v1.2.1 // indirect
)

require github.com/ghonzo/advent2023 v0.0.0-20231222192835-fe30706b5c2d
