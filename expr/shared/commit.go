package shared

import (
	"time"
)

type CommitMetadata struct {
	// Commit message
	Message string
	// Commit author
	Author string
	// Commit creation date
	AuthorName string
	// Commit creation Author name
	Date time.Time
	// Associated tags
	Tags []string
}
