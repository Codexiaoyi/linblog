package model

type Article struct {
	Id            int
	IsTop         bool
	Banner        string
	IsHot         bool
	PubTime       string
	Title         string
	Summary       string
	Category      string
	Publisher     string
	ViewsCount    int
	CommentsCount int
}
