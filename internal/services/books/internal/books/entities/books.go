package entities

type Book struct {
	ID           string `gorm:"primaryKey;type:uuid"       comment:"ID"`
	BranchID     string `gorm:"not null;type:varchar(255)"   comment:"Branch ID"`
	Date         string `gorm:"not null;type:date"           comment:"Date YYYY-MM-DD"`
	StartTime    string `gorm:"not null;type:varchar(5)"     comment:"Start Time HH:MM"`
	EndTime      string `gorm:"not null;type:varchar(5)"     comment:"End Time HH:MM"`
	Capacity     int    `gorm:"not null;type:int"            comment:"Capacity"`
	Booked       bool   `gorm:"not null;type:boolean"        comment:"Booked Status"`
	Amount       int    `gorm:"not null;type:int"          comment:"Amount"`
	MobileNumber string `gorm:"not null;type:varchar(255)" comment:"Mobile Number"`
	Note         string `gorm:"not null;type:text"         comment:"Note"`
	CommonCQRSHistoryModel
}

type Books []Book

type UpdateBook struct {
	ID string

	Booked *bool
	Amount *int
	Note   *string
}

func (u *UpdateBook) RemoveUnchangedFields(book Book) *UpdateBook {
	if u == nil {
		return nil
	}
	if u.ID != book.ID {
		return nil
	}
	if u.Booked != nil && *u.Booked == book.Booked {
		u.Booked = nil
	}
	if u.Amount != nil && *u.Amount == book.Amount {
		u.Amount = nil
	}
	if u.Note != nil && *u.Note == book.Note {
		u.Note = nil
	}
	return u
}

func (b *UpdateBook) ToUpdateMap() map[string]any {
	m := make(map[string]any)
	if b.Booked != nil {
		m["booked"] = *b.Booked
	}
	if b.Amount != nil {
		m["amount"] = *b.Amount
	}
	if b.Note != nil {
		m["note"] = *b.Note
	}
	return m
}
