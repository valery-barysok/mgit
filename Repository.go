package main

// Repository represents a git repository.
type Repository struct {
	Path             string
	LastCommitHash   string
	LastCommitTagged bool
}