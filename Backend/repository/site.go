package repository

import "linblog/model"

type SiteRepository struct {
}

func (site *SiteRepository) Get() interface{} {
	siteInfo := &model.Site{
		Avatar: "https://s2.ax1x.com/2020/01/17/1SCadg.png",
		Slogan: "怎就让这不停燃烧的心,就这样耗尽消失在平庸里",
		Name:   "aaayi′blog",
		Domain: "https://www.fengziy.cn",
		Notice: "",
		Desc:   "一个It技术的探索者",
	}
	return siteInfo
}

func (site *SiteRepository) GetSocials() []*model.Social {
	socials := []*model.Social{
		{
			Id:    1,
			Title: "QQ",
			Icon:  "iconqq",
			Color: "#1AB6FF ",
			Href:  "http://wpa.qq.com/msgrd?v=3&uin=874183200&site=qq&menu=yes",
		},
		{
			Id:    2,
			Title: "GitHub",
			Icon:  "icongithub",
			Color: "",
			Href:  "https://github.com/Codexiaoyi",
		},
	}
	return socials
}
