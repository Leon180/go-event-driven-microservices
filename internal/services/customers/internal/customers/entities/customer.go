package entities

type Customer struct {
	ID           string `gorm:"primaryKey;type:uuid" comment:"ID"`
	FirstName    string `gorm:"not null;type:varchar(255)" comment:"First Name"`
	LastName     string `gorm:"not null;type:varchar(255)" comment:"Last Name"`
	Email        string `gorm:"not null;type:varchar(255)" comment:"Email"`
	MobileNumber string `gorm:"not null;type:varchar(255)" comment:"Mobile Number"`
	ActiveSwitch bool   `gorm:"not null;type:boolean" comment:"Active Switch"`
	CommonHistoryModelWithUpdate
}

func (c *Customer) IsActive() bool {
	return c.ActiveSwitch
}

type Customers []Customer

type UpdateCustomer struct {
	ID           string
	FirstName    *string
	LastName     *string
	Email        *string
	MobileNumber *string
	ActiveSwitch *bool
}

func (u *UpdateCustomer) RemoveUnchangedFields(customer Customer) {
	if u.ID != customer.ID {
		return
	}
	if u.FirstName != nil && *u.FirstName == customer.FirstName {
		u.FirstName = nil
	}
	if u.LastName != nil && *u.LastName == customer.LastName {
		u.LastName = nil
	}
	if u.Email != nil && *u.Email == customer.Email {
		u.Email = nil
	}
	if u.MobileNumber != nil && *u.MobileNumber == customer.MobileNumber {
		u.MobileNumber = nil
	}
	if u.ActiveSwitch != nil && *u.ActiveSwitch == customer.ActiveSwitch {
		u.ActiveSwitch = nil
	}
}

func (u *UpdateCustomer) NoUpdates() bool {
	return u.FirstName == nil && u.LastName == nil && u.Email == nil && u.MobileNumber == nil && u.ActiveSwitch == nil
}

func (u *UpdateCustomer) ToUpdateMap() map[string]interface{} {
	updateMap := make(map[string]interface{})
	if u.FirstName != nil {
		updateMap["first_name"] = *u.FirstName
	}
	if u.LastName != nil {
		updateMap["last_name"] = *u.LastName
	}
	if u.Email != nil {
		updateMap["email"] = *u.Email
	}
	if u.MobileNumber != nil {
		updateMap["mobile_number"] = *u.MobileNumber
	}
	if u.ActiveSwitch != nil {
		updateMap["active_switch"] = *u.ActiveSwitch
	}
	return updateMap
}
