package structs

type Bedrock struct {
	ID   int    `gorm:"column:id; primaryKey; autoIncrement" json:"id"`
	FUID string `gorm:"column:fuid; type:char(36); unique;not null" json:"fuid"`
	XUID string `gorm:"column:xuid; type:char(36); unique;not null" json:"xuid"`
}

func (Bedrock) TableName() string {
	return "bedrock"
}
