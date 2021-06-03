package domain

type Comment struct {
	Id        uint    `json:"id"`
	Text      string  `json:"text"`
	ArticleId uint    `json:"article_id"`
	Article   Article `gorm:"foreignKey:ArticleId"`
	UserId    uint    `json:"user_id"`
	User      User    `gorm:"foreginKey:UserId"`
}
