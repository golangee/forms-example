package nestor

import "time"

// A Tutorial contains chapters of stuff.
type Tutorial struct {
	ID        string    // ID of this resource.
	Title     string    // Title of the tutorial
	Subtitle  string    // Subtitle to tease the tutorial.
	TeaserUrl string    // TeaserUrl to an image.
	Chapters  []Chapter // Chapters provide the actual content.
}

// A Chapters consists of sections.
type Chapter struct {
	ID        string // ID of this resource.
	Title     string // Title of the chapter.
	Subtitle  string // Subtitle to tease the chapter.
	TeaserUrl string // TeaserUrl to a header section image.
	Sections  []Section
}

// A Section consistens of steps.
type Section struct {
	ID                      string        // ID of this resource.
	Objective               string        // Objective or goal to learn
	Description             string        // Description about the tutorial
	TeaserUrl               string        // TeaserUrl to a header section image.
	EstimatedDuration       time.Duration // EstimatedDuration to work through this tutorial.
	ProjectFilesDownloadUrl string        // ProjectFilesDownloadUrl to grab a zip file of ready-to-use material.
	Steps                   []Step
}

// A Step just has a left hand side (title and description) and a right hand side of multiple contents.
type Step struct {
	ID          string // ID of this resource.
	Title       string
	Description string
	Content     Content
}

// Content tuple, none or all fields may have been set.
type Content struct {
	Code     string // Code to show with line numbers.
	ImageUrl string // ImageUrl to show how it should look like.
	VideoUrl string // VideoUrl to show how it should behave or how to achieve something.
	LiveUrl  string // LiveUrl for an iframe to show compiled result.
}
