package validatesdtos

import (
	"time"

	customizeerrors "github.com/Leon180/go-event-driven-microservices/internal/pkg/customize_errors"
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/enums"
	"github.com/Leon180/go-event-driven-microservices/internal/services/restaurants/internal/restaurants/dtos"
)

func ValidateRestaurant(restaurant *dtos.Restaurant) error {
	if restaurant == nil {
		return nil
	}
	if restaurant.Name == "" {
		return customizeerrors.RestaurantNameEmptyError
	}
	for _, branch := range restaurant.Branches {
		if err := ValidateBranch(&branch); err != nil {
			return err
		}
	}
	return nil
}

func ValidateBranch(branch *dtos.Branch) error {
	if branch == nil {
		return nil
	}
	if branch.Name == "" {
		return customizeerrors.BranchNameEmptyError
	}
	if branch.Address == nil {
		return customizeerrors.AddressEmptyError
	}
	if err := ValidateAddress(branch.Address); err != nil {
		return err
	}
	if err := ValidatePriceRange(branch.PriceRange); err != nil {
		return err
	}
	for _, category := range branch.Categories {
		if err := ValidateCategory(&category); err != nil {
			return err
		}
	}
	for _, table := range branch.Tables {
		if err := ValidateTable(&table); err != nil {
			return err
		}
	}
	for _, available := range branch.Availables {
		if err := ValidateAvailable(&available); err != nil {
			return err
		}
	}
	return nil
}

func ValidateAddress(address *dtos.Address) error {
	if address == nil {
		return nil
	}
	if !address.City.IsValid() {
		return customizeerrors.InvalidCityError
	}
	if !address.Country.IsValid() {
		return customizeerrors.InvalidCountryError
	}
	return nil
}

func ValidatePriceRange(priceRange *dtos.PriceRange) error {
	if priceRange == nil {
		return nil
	}
	if priceRange.MinPrice < 0 {
		return customizeerrors.PriceRangeInvalidError
	}
	if priceRange.MaxPrice < priceRange.MinPrice {
		return customizeerrors.PriceRangeInvalidError
	}
	return nil
}

func ValidateCategory(category *dtos.Category) error {
	if category == nil {
		return nil
	}
	if !category.Category.IsValid() {
		return customizeerrors.InvalidCategoryError
	}
	return nil
}

func ValidateTable(table *dtos.Table) error {
	if table == nil {
		return nil
	}
	if table.Capacity <= 0 {
		return customizeerrors.TableCapacityInvalidError
	}
	return nil
}

func ValidateAvailable(available *dtos.Available) error {
	if available == nil {
		return nil
	}
	if available.Weekday < time.Sunday || available.Weekday > time.Saturday {
		return customizeerrors.TableAvailableTimeInvalidError
	}
	_, err := time.Parse(enums.TimeFormatClockOnly.ToString(), available.StartTime)
	if err != nil {
		return customizeerrors.TableAvailableTimeInvalidError
	}
	_, err = time.Parse(enums.TimeFormatClockOnly.ToString(), available.EndTime)
	if err != nil {
		return customizeerrors.TableAvailableTimeInvalidError
	}
	return nil
}
