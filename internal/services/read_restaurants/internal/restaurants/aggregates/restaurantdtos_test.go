package aggregates

import (
	"testing"
	"time"

	enums "github.com/Leon180/go-event-driven-microservices/internal/pkg/enums"
	uuid "github.com/Leon180/go-event-driven-microservices/internal/pkg/uuid"
	"github.com/Leon180/go-event-driven-microservices/internal/services/restaurants/internal/restaurants/dtos"
	"github.com/Leon180/go-event-driven-microservices/internal/services/restaurants/internal/restaurants/entities"
	"github.com/samber/lo"
)

func TestRestaurantDTOAggregate_AddRestaurant_Aggregate(t *testing.T) {
	uuidGenerator := uuid.NewUUIDGenerator()
	restaurantDTOAggregate := NewRestaurantDTOAggregateBuilder(uuidGenerator)
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
					{Capacity: 4},
					{Capacity: 4},
					{Capacity: 4},
					{Capacity: 4},
					{Capacity: 2},
					{Capacity: 2},
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
					{Capacity: 8},
					{Capacity: 8},
					{Capacity: 4},
					{Capacity: 4},
					{Capacity: 4},
					{Capacity: 4},
					{Capacity: 4},
					{Capacity: 4},
					{Capacity: 2},
					{Capacity: 2},
					{Capacity: 2},
					{Capacity: 2},
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
		if !checkAggregates(t, &aggregates[0], restaurant) {
			t.Errorf("expected aggregates to match")
		}
		editEntities := restaurantDTOAggregate.GetEditEntities()
		if !checkEditEntitiesWhileCreate(t, &editEntities[0], restaurant) {
			t.Errorf("expected edit entities to match")
		}
	})
}

func checkAggregates(t *testing.T, aggregate *Restaurant, dto *dtos.Restaurant) bool {
	if aggregate == nil && dto == nil {
		return true
	}
	if aggregate == nil || dto == nil {
		t.Errorf("expected aggregate and dto, got nil")
		return false
	}
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
				t.Errorf(
					"expected Branch Address Street %v, got %v",
					dto.Branches[i].Address.Street,
					branch.Address.Street,
				)
				return false
			}
			if branch.Address.CityCode.ToCity() != dto.Branches[i].Address.City {
				t.Errorf(
					"expected Branch Address City %v, got %v",
					dto.Branches[i].Address.City,
					branch.Address.CityCode.ToCity(),
				)
				return false
			}
			if branch.Address.PostalCode != dto.Branches[i].Address.PostalCode {
				t.Errorf(
					"expected Branch Address Postal Code %v, got %v",
					dto.Branches[i].Address.PostalCode,
					branch.Address.PostalCode,
				)
				return false
			}
			if branch.Address.CountryCode.ToCountry() != dto.Branches[i].Address.Country {
				t.Errorf(
					"expected Branch Address Country %v, got %v",
					dto.Branches[i].Address.Country,
					branch.Address.CountryCode.ToCountry(),
				)
				return false
			}
		} else if branch.Address != nil || dto.Branches[i].Address != nil {
			t.Errorf("expected Branch Address %v, got %v", dto.Branches[i].Address, branch.Address)
			return false
		}
		if branch.PriceRange != nil && dto.Branches[i].PriceRange != nil {
			if branch.PriceRange.MinPrice != dto.Branches[i].PriceRange.MinPrice {
				t.Errorf(
					"expected Branch Price Range Min Price %v, got %v",
					dto.Branches[i].PriceRange.MinPrice,
					branch.PriceRange.MinPrice,
				)
				return false
			}
			if branch.PriceRange.MaxPrice != dto.Branches[i].PriceRange.MaxPrice {
				t.Errorf(
					"expected Branch Price Range Max Price %v, got %v",
					dto.Branches[i].PriceRange.MaxPrice,
					branch.PriceRange.MaxPrice,
				)
				return false
			}
		} else if branch.PriceRange != nil || dto.Branches[i].PriceRange != nil {
			t.Errorf("expected Branch Price Range %v, got %v", dto.Branches[i].PriceRange, branch.PriceRange)
			return false
		}
		if len(branch.BranchCategoryRelations) != len(dto.Branches[i].Categories) {
			t.Errorf(
				"expected Branch Category Relations %v, got %v",
				len(dto.Branches[i].Categories),
				len(branch.BranchCategoryRelations),
			)
			return false
		}
		for j, category := range branch.BranchCategoryRelations {
			if category.Category.CategoryCode.ToCategory() != dto.Branches[i].Categories[j].Category {
				t.Errorf(
					"expected Branch Category %v, got %v",
					dto.Branches[i].Categories[j].Category,
					category.Category.CategoryCode.ToCategory(),
				)
				return false
			}
		}
		if len(branch.Tables) != len(dto.Branches[i].Tables) {
			t.Errorf("expected Branch Tables %v, got %v", len(dto.Branches[i].Tables), len(branch.Tables))
			return false
		}
		for j, table := range branch.Tables {
			if table.Capacity != dto.Branches[i].Tables[j].Capacity {
				t.Errorf(
					"expected Branch Table Capacity %v, got %v",
					dto.Branches[i].Tables[j].Capacity,
					table.Capacity,
				)
				return false
			}
		}
		if len(branch.Availables) != len(dto.Branches[i].Availables) {
			t.Errorf("expected Branch Availables %v, got %v", len(dto.Branches[i].Availables), len(branch.Availables))
			return false
		}
		for j, available := range branch.Availables {
			if available.Weekday != dto.Branches[i].Availables[j].Weekday {
				t.Errorf(
					"expected Branch Available Weekday %v, got %v",
					dto.Branches[i].Availables[j].Weekday,
					available.Weekday,
				)
				return false
			}
			if available.StartTime != dto.Branches[i].Availables[j].StartTime {
				t.Errorf(
					"expected Branch Available Start Time %v, got %v",
					dto.Branches[i].Availables[j].StartTime,
					available.StartTime,
				)
				return false
			}
			if available.EndTime != dto.Branches[i].Availables[j].EndTime {
				t.Errorf(
					"expected Branch Available End Time %v, got %v",
					dto.Branches[i].Availables[j].EndTime,
					available.EndTime,
				)
				return false
			}
		}
	}
	return true
}

func checkEditEntitiesWhileCreate(t *testing.T, editEntities *RestaurantEditEntities, dto *dtos.Restaurant) bool {
	if editEntities == nil && dto == nil {
		return true
	}
	if editEntities == nil || dto == nil {
		t.Errorf("expected edit entities and dto, got nil")
		return false
	}
	if editEntities.CreateEntities == nil {
		t.Errorf("expected create entities, got nil")
		return false
	}
	// restaurant
	if len(editEntities.CreateEntities.Restaurants) != 1 {
		t.Errorf("expected create 1 restaurant, got %v", len(editEntities.CreateEntities.Restaurants))
		return false
	}
	if editEntities.CreateEntities.Restaurants[0].Name != dto.Name {
		t.Errorf("expected Restaurant Name %v, got %v", dto.Name, editEntities.CreateEntities.Restaurants[0].Name)
		return false
	}
	if editEntities.CreateEntities.Restaurants[0].Description != dto.Description {
		t.Errorf(
			"expected Restaurant Description %v, got %v",
			dto.Description,
			editEntities.CreateEntities.Restaurants[0].Description,
		)
		return false
	}
	// branches
	if len(editEntities.CreateEntities.Branches) != len(dto.Branches) {
		t.Errorf("expected create 2 branches, got %v", len(editEntities.CreateEntities.Branches))
		return false
	}
	if len(editEntities.CreateEntities.Addresses) != len(dto.Branches) {
		t.Errorf("expected 1 address, got %v", len(editEntities.CreateEntities.Addresses))
		return false
	}
	if len(editEntities.CreateEntities.PriceRanges) != len(dto.Branches) {
		t.Errorf("expected 1 price range, got %v", len(editEntities.CreateEntities.PriceRanges))
		return false
	}
	for i, branch := range editEntities.CreateEntities.Branches {
		if branch.Name != dto.Branches[i].Name {
			t.Errorf("expected Branch Name %v, got %v", dto.Branches[i].Name, branch.Name)
			return false
		}
		if branch.Description != dto.Branches[i].Description {
			t.Errorf("expected Branch Description %v, got %v", dto.Branches[i].Description, branch.Description)
			return false
		}
		// Address
		address, ok := lo.Find(editEntities.CreateEntities.Addresses, func(address entities.Address) bool {
			return address.BranchID == branch.ID
		})
		if !ok {
			t.Errorf("expected address, got nil")
			return false
		}
		if address.Street != dto.Branches[i].Address.Street {
			t.Errorf("expected Branch Address Street %v, got %v", dto.Branches[i].Address.Street, address.Street)
			return false
		}
		if address.CityCode.ToCity() != dto.Branches[i].Address.City {
			t.Errorf("expected Branch Address City %v, got %v", dto.Branches[i].Address.City, address.CityCode.ToCity())
			return false
		}
		if address.PostalCode != dto.Branches[i].Address.PostalCode {
			t.Errorf(
				"expected Branch Address Postal Code %v, got %v",
				dto.Branches[i].Address.PostalCode,
				address.PostalCode,
			)
			return false
		}
		if address.CountryCode.ToCountry() != dto.Branches[i].Address.Country {
			t.Errorf(
				"expected Branch Address Country %v, got %v",
				dto.Branches[i].Address.Country,
				address.CountryCode.ToCountry(),
			)
			return false
		}
		// Price Range
		priceRange, ok := lo.Find(editEntities.CreateEntities.PriceRanges, func(priceRange entities.PriceRange) bool {
			return priceRange.BranchID == branch.ID
		})
		if !ok {
			t.Errorf("expected price range, got nil")
			return false
		}
		if priceRange.MinPrice != dto.Branches[i].PriceRange.MinPrice {
			t.Errorf(
				"expected Branch Price Range Min Price %v, got %v",
				dto.Branches[i].PriceRange.MinPrice,
				priceRange.MinPrice,
			)
			return false
		}
		if priceRange.MaxPrice != dto.Branches[i].PriceRange.MaxPrice {
			t.Errorf(
				"expected Branch Price Range Max Price %v, got %v",
				dto.Branches[i].PriceRange.MaxPrice,
				priceRange.MaxPrice,
			)
			return false
		}
		// Category Relations
		categoryRelations := lo.UniqBy(
			lo.Filter(
				editEntities.CreateEntities.BranchCategoryRelations,
				func(relation entities.BranchCategoryRelation, _ int) bool {
					return relation.BranchID == branch.ID
				},
			),
			func(relation entities.BranchCategoryRelation) string {
				return relation.CategoryID
			},
		)
		if len(categoryRelations) != len(dto.Branches[i].Categories) {
			t.Errorf(
				"expected Branch Category Relations %v, got %v",
				len(dto.Branches[i].Categories),
				len(categoryRelations),
			)
			return false
		}
		categoryRelationCategoryIDCountMap := make(map[string]int)
		for _, categoryRelation := range categoryRelations {
			categoryRelationCategoryIDCountMap[categoryRelation.CategoryID]++
		}
		for _, category := range dto.Branches[i].Categories {
			categoryRelationCategoryIDCountMap[category.ID]--
		}
		for _, count := range categoryRelationCategoryIDCountMap {
			if count != 0 {
				t.Errorf("expected Branch Category should be same")
				return false
			}
		}
		// Tables
		tableRelations := lo.Filter(editEntities.CreateEntities.Tables, func(table entities.Table, _ int) bool {
			return table.BranchID == branch.ID
		})
		if len(tableRelations) != len(dto.Branches[i].Tables) {
			t.Errorf("expected Branch Tables %v, got %v", len(dto.Branches[i].Tables), len(tableRelations))
			return false
		}
		tableCapacityCountMap := make(map[int]int)
		for _, tableRelation := range tableRelations {
			tableCapacityCountMap[tableRelation.Capacity]++
		}
		for _, table := range dto.Branches[i].Tables {
			tableCapacityCountMap[table.Capacity]--
		}
		for _, count := range tableCapacityCountMap {
			if count != 0 {
				t.Errorf("expected Branch Table Capacity should be same")
				return false
			}
		}
		// Availables
		availableRelations := lo.Filter(
			editEntities.CreateEntities.Availables,
			func(available entities.Available, _ int) bool {
				return available.BranchID == branch.ID
			},
		)
		if len(availableRelations) != len(dto.Branches[i].Availables) {
			t.Errorf("expected Branch Availables %v, got %v", len(dto.Branches[i].Availables), len(availableRelations))
			return false
		}
		availableWeekdayCountMap := make(map[time.Weekday]int)
		for _, available := range availableRelations {
			availableWeekdayCountMap[available.Weekday]++
		}
		for _, available := range dto.Branches[i].Availables {
			availableWeekdayCountMap[available.Weekday]--
		}
		for _, count := range availableWeekdayCountMap {
			if count != 0 {
				t.Errorf("expected Branch Available Weekday should be same")
				return false
			}
		}
	}
	return true
}
