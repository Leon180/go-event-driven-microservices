package documents

type Book struct {
	ID           string `bson:"_id"`
	TableID      string `bson:"table_id"`
	AvailableID  string `bson:"available_id"`
	Amount       int    `bson:"amount"`
	MobileNumber string `bson:"mobile_number"`
	CommonCQRSHistoryModel

	Table     Table     `bson:"table"`
	Available Available `bson:"available"`
}

type Books []Book
