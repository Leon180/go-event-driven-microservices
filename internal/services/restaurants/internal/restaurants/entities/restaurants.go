package entities

type Restaurant struct {
	ID          string `gorm:"primaryKey;type:uuid" comment:"ID"`
	Name        string `gorm:"not null;type:varchar(255)" comment:"Name"`
	Description string `gorm:"not null;type:text" comment:"Description"`
	CommonCQRSHistoryModel
}

type Restaurants []Restaurant

type UpdateRestaurant struct {
	ID           string
	Name         *string
	Description  *string
	ActiveStatus *bool
}

func (u *UpdateRestaurant) RemoveUnchangedFields(restaurant Restaurant) {
	if u.ID != restaurant.ID {
		return
	}
	if u.Name != nil && *u.Name == restaurant.Name {
		u.Name = nil
	}
	if u.Description != nil && *u.Description == restaurant.Description {
		u.Description = nil
	}
	if u.ActiveStatus != nil && *u.ActiveStatus == restaurant.ActiveStatus {
		u.ActiveStatus = nil
	}
}

func (u *UpdateRestaurant) ToUpdateMap() map[string]interface{} {
	updateMap := make(map[string]interface{})
	if u.Name != nil {
		updateMap["name"] = *u.Name
	}
	if u.Description != nil {
		updateMap["description"] = *u.Description
	}
	if u.ActiveStatus != nil {
		updateMap["active_status"] = *u.ActiveStatus
	}
	return updateMap
}

type UpdateRestaurants []UpdateRestaurant
