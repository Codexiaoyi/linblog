package model

type Article struct {
	Id            int
	IsTop         bool
	Banner        string
	IsHot         bool
	PubTime       int
	Title         string
	Summary       string
	Category      string
	ViewsCount    int
	CommentsCount int
}
