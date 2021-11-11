package model

// 站点信息
type Site struct {
	Avatar string
	Slogan string
	Name   string
	Domain string
	Notice string
	Desc   string
}

//社交信息
type Social struct {
	Id    int
	Title string
	Icon  string
	Color string
	Href  string
}
