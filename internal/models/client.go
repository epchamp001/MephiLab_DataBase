package models

type Client struct {
	ID         uint   `gorm:"primaryKey"`
	Role       Role   `gorm:"not null"` // Role: sender, receiver
	FirstName  string `gorm:"not null"`
	LastName   string `gorm:"not null"`
	MiddleName string
	Phone      string `gorm:"not null; unique"`
	Email      string `gorm:"not null; unique"`
	Address    string
	Chats      []Chat      `gorm:"foreignKey:ParticipantID"`
	OrdersSent []Order     `gorm:"foreignKey:SenderID"`
	OrdersRecv []Order     `gorm:"foreignKey:RecipientID"`
	PromoCodes []PromoCode `gorm:"foreignKey:ClientID"`
}
