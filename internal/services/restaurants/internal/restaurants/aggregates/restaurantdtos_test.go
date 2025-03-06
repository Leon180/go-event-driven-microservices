package aggregates

import (
	"testing"
	"time"

	enums "github.com/Leon180/go-event-driven-microservices/internal/pkg/enums"
	mocksuuid "github.com/Leon180/go-event-driven-microservices/internal/pkg/uuid/mocks"
	"github.com/Leon180/go-event-driven-microservices/internal/services/restaurants/internal/restaurants/dtos"
	"go.uber.org/mock/gomock"
)

func TestRestaurantDTOAggregate_AddRestaurant_Aggregate(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockUUIDGenerator := mocksuuid.NewMockUUIDGenerator(ctrl)
	mockUUIDGenerator.EXPECT().GenerateUUID().Return("123e4567-e89b-12d3-a456-426614174000").AnyTimes()
	restaurantDTOAggregate := NewRestaurantDTOAggregate(mockUUIDGenerator)
	restaurant := &dtos.Restaurant{
		Name: "Test Restaurant",
		Branches: []dtos.Branch{
			{
				Name: "Test Branch1",
				Address: &dtos.Address{
					Street:  "Test Street1",
					City:    enums.CityCode(1).ToCity(),
					Country: enums.CountryCode(1).ToCountry(),
				},
				PriceRange: &dtos.PriceRange{MinPrice: 100, MaxPrice: 1000},
				Categories: []dtos.Category{
					{ID: "1", Category: enums.CategoryCode(1).ToCategory()},
					{ID: "2", Category: enums.CategoryCode(2).ToCategory()},
					{ID: "3", Category: enums.CategoryCode(3).ToCategory()},
				},
				Tables: []dtos.Table{
					{Capacity: 4}, {Capacity: 4}, {Capacity: 4}, {Capacity: 4},
					{Capacity: 2}, {Capacity: 2},
				},
				Availables: []dtos.Available{
					{Weekday: time.Monday, StartTime: "12:00", EndTime: "20:00"},
					{Weekday: time.Tuesday, StartTime: "12:00", EndTime: "20:00"},
					{Weekday: time.Wednesday, StartTime: "12:00", EndTime: "20:00"},
					{Weekday: time.Thursday, StartTime: "12:00", EndTime: "20:00"},
					{Weekday: time.Friday, StartTime: "12:00", EndTime: "20:00"},
					{Weekday: time.Saturday, StartTime: "12:00", EndTime: "22:00"},
				},
			},
			{
				Name: "Test Branch2",
				Address: &dtos.Address{
					Street:  "Test Street2",
					City:    enums.CityCode(2).ToCity(),
					Country: enums.CountryCode(1).ToCountry(),
				},
				PriceRange: &dtos.PriceRange{MinPrice: 100, MaxPrice: 1000},
				Categories: []dtos.Category{
					{ID: "1", Category: enums.CategoryCode(1).ToCategory()},
					{ID: "2", Category: enums.CategoryCode(2).ToCategory()},
					{ID: "3", Category: enums.CategoryCode(3).ToCategory()},
				},
				Tables: []dtos.Table{
					{Capacity: 8}, {Capacity: 8},
					{Capacity: 4}, {Capacity: 4}, {Capacity: 4}, {Capacity: 4}, {Capacity: 4}, {Capacity: 4},
					{Capacity: 2}, {Capacity: 2}, {Capacity: 2}, {Capacity: 2},
				},
				Availables: []dtos.Available{
					{Weekday: time.Monday, StartTime: "12:00", EndTime: "20:00"},
					{Weekday: time.Tuesday, StartTime: "12:00", EndTime: "20:00"},
					{Weekday: time.Thursday, StartTime: "12:00", EndTime: "20:00"},
					{Weekday: time.Friday, StartTime: "12:00", EndTime: "20:00"},
					{Weekday: time.Saturday, StartTime: "12:00", EndTime: "22:00"},
				},
			},
		},
	}

	t.Run("add restaurant", func(t *testing.T) {
		err := restaurantDTOAggregate.SaveRestaurant(restaurant)
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
		aggregates := restaurantDTOAggregate.GetAggregates()
		if !checkAggregates(t, aggregates[0], restaurant) {
			t.Errorf("expected aggregates to match")
		}
	})
}

func checkAggregates(t *testing.T, aggregate Restaurant, dto *dtos.Restaurant) bool {
	if aggregate.Name != dto.Name {
		t.Errorf("expected Name %v, got %v", dto.Name, aggregate.Name)
		return false
	}
	if aggregate.Description != dto.Description {
		t.Errorf("expected Description %v, got %v", dto.Description, aggregate.Description)
		return false
	}
	if len(aggregate.Branches) != len(dto.Branches) {
		t.Errorf("expected Branches %v, got %v", len(dto.Branches), len(aggregate.Branches))
		return false
	}
	for i, branch := range aggregate.Branches {
		if branch.Name != dto.Branches[i].Name {
			t.Errorf("expected Branch Name %v, got %v", dto.Branches[i].Name, branch.Name)
			return false
		}
		if branch.Description != dto.Branches[i].Description {
			t.Errorf("expected Branch Description %v, got %v", dto.Branches[i].Description, branch.Description)
			return false
		}
		if branch.Address != nil && dto.Branches[i].Address != nil {
			if branch.Address.Street != dto.Branches[i].Address.Street {
				t.Errorf("expected Branch Address Street %v, got %v", dto.Branches[i].Address.Street, branch.Address.Street)
				return false
			}
			if branch.Address.CityCode.ToCity() != dto.Branches[i].Address.City {
				t.Errorf("expected Branch Address City %v, got %v", dto.Branches[i].Address.City, branch.Address.CityCode.ToCity())
				return false
			}
			if branch.Address.PostalCode != dto.Branches[i].Address.PostalCode {
				t.Errorf("expected Branch Address Postal Code %v, got %v", dto.Branches[i].Address.PostalCode, branch.Address.PostalCode)
				return false
			}
			if branch.Address.CountryCode.ToCountry() != dto.Branches[i].Address.Country {
				t.Errorf("expected Branch Address Country %v, got %v", dto.Branches[i].Address.Country, branch.Address.CountryCode.ToCountry())
				return false
			}
		} else if branch.Address != nil || dto.Branches[i].Address != nil {
			t.Errorf("expected Branch Address %v, got %v", dto.Branches[i].Address, branch.Address)
			return false
		}
		if branch.PriceRange != nil && dto.Branches[i].PriceRange != nil {
			if branch.PriceRange.MinPrice != dto.Branches[i].PriceRange.MinPrice {
				t.Errorf("expected Branch Price Range Min Price %v, got %v", dto.Branches[i].PriceRange.MinPrice, branch.PriceRange.MinPrice)
				return false
			}
			if branch.PriceRange.MaxPrice != dto.Branches[i].PriceRange.MaxPrice {
				t.Errorf("expected Branch Price Range Max Price %v, got %v", dto.Branches[i].PriceRange.MaxPrice, branch.PriceRange.MaxPrice)
				return false
			}
		} else if branch.PriceRange != nil || dto.Branches[i].PriceRange != nil {
			t.Errorf("expected Branch Price Range %v, got %v", dto.Branches[i].PriceRange, branch.PriceRange)
			return false
		}
		if len(branch.BranchCategoryRelations) != len(dto.Branches[i].Categories) {
			t.Errorf("expected Branch Category Relations %v, got %v", len(dto.Branches[i].Categories), len(branch.BranchCategoryRelations))
			return false
		}
		for j, category := range branch.BranchCategoryRelations {
			if category.Category.CategoryCode.ToCategory() != dto.Branches[i].Categories[j].Category {
				t.Errorf("expected Branch Category %v, got %v", dto.Branches[i].Categories[j].Category, category.Category.CategoryCode.ToCategory())
				return false
			}
		}
		if len(branch.Tables) != len(dto.Branches[i].Tables) {
			t.Errorf("expected Branch Tables %v, got %v", len(dto.Branches[i].Tables), len(branch.Tables))
			return false
		}
		for j, table := range branch.Tables {
			if table.Capacity != dto.Branches[i].Tables[j].Capacity {
				t.Errorf("expected Branch Table Capacity %v, got %v", dto.Branches[i].Tables[j].Capacity, table.Capacity)
				return false
			}
		}
		if len(branch.Availables) != len(dto.Branches[i].Availables) {
			t.Errorf("expected Branch Availables %v, got %v", len(dto.Branches[i].Availables), len(branch.Availables))
			return false
		}
		for j, available := range branch.Availables {
			if available.Weekday != dto.Branches[i].Availables[j].Weekday {
				t.Errorf("expected Branch Available Weekday %v, got %v", dto.Branches[i].Availables[j].Weekday, available.Weekday)
				return false
			}
			if available.StartTime != dto.Branches[i].Availables[j].StartTime {
				t.Errorf("expected Branch Available Start Time %v, got %v", dto.Branches[i].Availables[j].StartTime, available.StartTime)
				return false
			}
			if available.EndTime != dto.Branches[i].Availables[j].EndTime {
				t.Errorf("expected Branch Available End Time %v, got %v", dto.Branches[i].Availables[j].EndTime, available.EndTime)
				return false
			}
		}
	}
	return true
}
