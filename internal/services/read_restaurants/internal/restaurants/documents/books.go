package documents

import "go.mongodb.org/mongo-driver/bson"

type Book struct {
	ID           string `bson:"_id"`
	TableID      string `bson:"table_id"`
	AvailableID  string `bson:"available_id"`
	Amount       int    `bson:"amount"`
	MobileNumber string `bson:"mobile_number"`
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

func (u *UpdateBook) ToFilterBson() bson.M {
	return bson.M{
		"_id":          u.ID,
		"table_id":     u.TableID,
		"available_id": u.AvailableID,
	}
}

func (u *UpdateBook) ToSetBson() bson.M {
	bsonMap := bson.M{}
	if u.Amount != nil {
		bsonMap["amount"] = *u.Amount
	}
	if u.MobileNumber != nil {
		bsonMap["mobile_number"] = *u.MobileNumber
	}
	if u.ActiveStatus != nil {
		bsonMap["active_status"] = *u.ActiveStatus
	}
	return bson.M{
		"$set": bsonMap,
	}
}

type UpdateBooks []UpdateBook
