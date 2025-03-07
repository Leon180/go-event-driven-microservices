package entities

type Book struct {
	ID           string `gorm:"primaryKey;type:uuid" comment:"ID"`
	TableID      string `gorm:"not null;type:uuid" comment:"Table ID"`
	AvailableID  string `gorm:"not null;type:uuid" comment:"Available ID"`
	Amount       int    `gorm:"not null;type:int" comment:"Amount"`
	MobileNumber string `gorm:"not null;type:varchar(255)" comment:"Mobile Number"`
	CommonCQRSHistoryModel
}

type Books []Book

type UpdateBook struct {
	ID           string
	TableID      string
	AvailableID  string
	Amount       *int
	MobileNumber *string
	ActiveStatus *bool
}

func (u *UpdateBook) RemoveUnchangedFields(book Book) *UpdateBook {
	if u.ID != book.ID {
		return nil
	}
	if u.Amount != nil && *u.Amount == book.Amount {
		u.Amount = nil
	}
	if u.MobileNumber != nil && *u.MobileNumber == book.MobileNumber {
		u.MobileNumber = nil
	}
	if u.ActiveStatus != nil && *u.ActiveStatus == book.ActiveStatus {
		u.ActiveStatus = nil
	}
	return u
}

func (u *UpdateBook) ToUpdateMap() map[string]any {
	updateMap := make(map[string]any)
	if u.Amount != nil {
		updateMap["amount"] = *u.Amount
	}
	if u.MobileNumber != nil {
		updateMap["mobile_number"] = *u.MobileNumber
	}
	if u.ActiveStatus != nil {
		updateMap["active_status"] = *u.ActiveStatus
	}
	return updateMap
}

type UpdateBooks []UpdateBook
