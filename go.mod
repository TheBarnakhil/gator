module github.com/TheBarnakhil/gator

go 1.24.1

require (
	github.com/TheBarnakhil/gator/internal/config v0.0.0
	github.com/google/uuid v1.6.0
	github.com/lib/pq v1.10.9
)

replace github.com/TheBarnakhil/gator/internal/config => ./internal/config
