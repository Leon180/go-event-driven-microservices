package aggregates

import (
	"time"

	customizeerrors "github.com/Leon180/go-event-driven-microservices/internal/pkg/customize_errors"
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/enums"
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/uuid"
	"github.com/Leon180/go-event-driven-microservices/internal/services/write_restaurants/internal/restaurants/dtos"
	"github.com/Leon180/go-event-driven-microservices/internal/services/write_restaurants/internal/restaurants/entities"
	validatesdtos "github.com/Leon180/go-event-driven-microservices/internal/services/write_restaurants/internal/restaurants/validates/dtos"
	"github.com/samber/lo"
)

type RestaurantDTOAggregateBuilder interface {
	// save restaurant and related entities(branch, address, price range, category relation, table, table available) to the aggregate
	SaveRestaurant(restaurant *dtos.Restaurant) error

	// save branch and related entities(address, price range, category relation, table, table available) to the aggregate
	SaveBranch(restaurantID string, branch *dtos.Branch) error

	// save address to the aggregate
	SaveAddress(branchID string, address *dtos.Address) error

	// save price range to the aggregate
	SavePriceRange(branchID string, priceRange *dtos.PriceRange) error

	// save category relation to the aggregate
	SaveCategoryRelation(branchID string, category *dtos.Category) error

	// save table and related entities(table available) to the aggregate
	SaveTable(branchID string, table *dtos.Table) error

	// save available to the aggregate
	SaveAvailable(branchID string, available *dtos.Available) error

	// set all edit type code to none
	SetAllEditTypeCodeToNone()

	// get restaurant aggregates
	GetAggregates() []Restaurant

	// get pending edit entities during update
	GetEditEntities() []RestaurantEditEntities
}

func NewRestaurantDTOAggregateBuilder(uuidGenerator uuid.UUIDGenerator) RestaurantDTOAggregateBuilder {
	return &restaurantDTOAggregateImpl{
		restaurants:   make([]Restaurant, 0),
		uuidGenerator: uuidGenerator,
	}
}

type restaurantDTOAggregateImpl struct {
	restaurants []Restaurant

	// dependencies
	uuidGenerator uuid.UUIDGenerator
}

type RestaurantEntities entities.Restaurants

func (r RestaurantEntities) ToAggregates() []Restaurant {
	return lo.Map(r, func(restaurant entities.Restaurant, _ int) Restaurant {
		re := RestaurantEntity(restaurant)
		return *re.ToAggregate()
	})
}

type RestaurantEntity entities.Restaurant

func (r *RestaurantEntity) ToAggregate() *Restaurant {
	e := entities.Restaurant(*r)
	return &Restaurant{
		Restaurant:   e,
		editTypeCode: enums.EditTypeCodeNone,
	}
}

type RestaurantEditEntities struct {
	CreateEntities *RestaurantCreateEntities
	UpdateEntities *RestaurantUpdateEntities
	DeleteEntities *RestaurantDeleteEntities
}

func newRestaurantEditEntities() RestaurantEditEntities {
	return RestaurantEditEntities{
		CreateEntities: &RestaurantCreateEntities{},
		UpdateEntities: &RestaurantUpdateEntities{},
		DeleteEntities: &RestaurantDeleteEntities{},
	}
}

type RestaurantCreateEntities struct {
	Restaurants             []entities.Restaurant
	Branches                []entities.Branch
	Addresses               []entities.Address
	PriceRanges             []entities.PriceRange
	BranchCategoryRelations []entities.BranchCategoryRelation
	Tables                  []entities.Table
	Availables              []entities.Available
}

type RestaurantUpdateEntities struct {
	Restaurants             []entities.UpdateRestaurant
	Branches                []entities.UpdateBranch
	Addresses               []entities.UpdateAddress
	PriceRanges             []entities.UpdatePriceRange
	BranchCategoryRelations []entities.UpdateBranchCategoryRelation
	Tables                  []entities.UpdateTable
	Availables              []entities.UpdateAvailable
}

type RestaurantDeleteEntities struct {
	Restaurants             []entities.Restaurant
	Branches                []entities.Branch
	Addresses               []entities.Address
	PriceRanges             []entities.PriceRange
	BranchCategoryRelations []entities.BranchCategoryRelation
	Tables                  []entities.Table
	Availables              []entities.Available
}

type Restaurants []Restaurant

func (r Restaurants) ToDTO() []dtos.Restaurant {
	return lo.Map(r, func(restaurant Restaurant, _ int) dtos.Restaurant {
		return *restaurant.ToDTO()
	})
}

type Restaurant struct {
	entities.Restaurant
	editTypeCode enums.EditTypeCode         `gorm:"-"`
	update       *entities.UpdateRestaurant `gorm:"-"`
	Branches     []Branch                   `gorm:"foreignKey:RestaurantID;references:ID" comment:"Branches"`
}

func (r *Restaurant) TableName() string {
	return "restaurant"
}

func (r *Restaurant) ToDTO() *dtos.Restaurant {
	return &dtos.Restaurant{
		ID:          &r.ID,
		Name:        r.Name,
		Description: r.Description,
		Branches: lo.Map(r.Branches, func(b Branch, _ int) dtos.Branch {
			return *b.ToDTO()
		}),
	}
}

type Branch struct {
	entities.Branch
	editTypeCode            enums.EditTypeCode       `gorm:"-"`
	update                  *entities.UpdateBranch   `gorm:"-"`
	Address                 *Address                 `gorm:"foreignKey:BranchID;references:ID" comment:"Address"`
	PriceRange              *PriceRange              `gorm:"foreignKey:BranchID;references:ID" comment:"Price Range"`
	BranchCategoryRelations []BranchCategoryRelation `gorm:"foreignKey:BranchID;references:ID" comment:"Branch Category Relation"`
	Tables                  []Table                  `gorm:"foreignKey:BranchID;references:ID" comment:"Tables"`
	Availables              []Available              `gorm:"foreignKey:BranchID;references:ID" comment:"Availables"`
}

func (b *Branch) TableName() string {
	return "branch"
}

func (b *Branch) ToDTO() *dtos.Branch {
	return &dtos.Branch{
		ID:          &b.ID,
		Name:        b.Name,
		Description: b.Description,
		Address:     b.Address.ToDTO(),
		PriceRange:  b.PriceRange.ToDTO(),
		Categories: lo.Map(b.BranchCategoryRelations, func(bcr BranchCategoryRelation, _ int) dtos.Category {
			return *bcr.Category.ToDTO()
		}),
		Tables: lo.Map(b.Tables, func(t Table, _ int) dtos.Table {
			return *t.ToDTO()
		}),
		Availables: lo.Map(b.Availables, func(a Available, _ int) dtos.Available {
			return *a.ToDTO()
		}),
	}
}

type Address struct {
	entities.Address
	editTypeCode enums.EditTypeCode      `gorm:"-"`
	update       *entities.UpdateAddress `gorm:"-"`
}

func (a *Address) TableName() string {
	return "address"
}

func (a *Address) ToDTO() *dtos.Address {
	return &dtos.Address{
		ID:         &a.ID,
		Street:     a.Street,
		City:       a.CityCode.ToCity(),
		PostalCode: a.PostalCode,
		Country:    a.CountryCode.ToCountry(),
	}
}

type PriceRange struct {
	entities.PriceRange
	editTypeCode enums.EditTypeCode         `gorm:"-"`
	update       *entities.UpdatePriceRange `gorm:"-"`
}

func (p *PriceRange) TableName() string {
	return "price_range"
}

func (p *PriceRange) ToDTO() *dtos.PriceRange {
	return &dtos.PriceRange{
		ID:       &p.ID,
		MinPrice: p.MinPrice,
		MaxPrice: p.MaxPrice,
	}
}

type BranchCategoryRelation struct {
	entities.BranchCategoryRelation
	editTypeCode enums.EditTypeCode                     `gorm:"-"`
	update       *entities.UpdateBranchCategoryRelation `gorm:"-"`
	Category     Category                               `gorm:"foreignKey:ID;references:CategoryID" comment:"Category"`
}

func (bcr *BranchCategoryRelation) TableName() string {
	return "branch_category_relation"
}

type Category entities.Category

func (c *Category) TableName() string {
	return "category"
}

func (c *Category) ToDTO() *dtos.Category {
	return &dtos.Category{
		ID:       c.ID,
		Category: c.CategoryCode.ToCategory(),
		CommonCQRSHistoryModel: dtos.CommonCQRSHistoryModel{
			ActiveStatus: c.ActiveStatus,
			CreatedAt:    c.CreatedAt,
			UpdatedAt:    c.UpdatedAt,
		},
	}
}

type Table struct {
	entities.Table
	editTypeCode enums.EditTypeCode    `gorm:"-"`
	update       *entities.UpdateTable `gorm:"-"`
}

func (t *Table) TableName() string {
	return "table"
}

func (t *Table) ToDTO() *dtos.Table {
	return &dtos.Table{
		ID:       &t.ID,
		Capacity: t.Capacity,
	}
}

type Available struct {
	entities.Available
	editTypeCode enums.EditTypeCode        `gorm:"-"`
	update       *entities.UpdateAvailable `gorm:"-"`
}

func (a *Available) TableName() string {
	return "available"
}

func (a *Available) ToDTO() *dtos.Available {
	return &dtos.Available{
		ID:        &a.ID,
		Weekday:   a.Weekday,
		StartTime: a.StartTime,
		EndTime:   a.EndTime,
	}
}

// restaurant aggregate methods
func (r *restaurantDTOAggregateImpl) SaveRestaurant(restaurant *dtos.Restaurant) error {
	if restaurant == nil {
		return nil
	}
	if err := validatesdtos.ValidateRestaurant(restaurant); err != nil {
		return err
	}
	for i := range r.restaurants {
		if restaurant.ID != nil && r.restaurants[i].ID == *restaurant.ID {
			return r.updateRestaurant(&r.restaurants[i], restaurant)
		}
		// same name but different id, duplicate restaurant
		if r.restaurants[i].Name == restaurant.Name {
			return customizeerrors.RestaurantAlreadyExistsError
		}
	}
	return r.addRestaurant(restaurant)
}

func (r *restaurantDTOAggregateImpl) addRestaurant(restaurant *dtos.Restaurant) error {
	if restaurant == nil {
		return nil
	}
	var restaurantID string
	if restaurant.ID == nil {
		restaurantID = r.uuidGenerator.GenerateUUID()
		restaurant.ID = &restaurantID
	} else {
		restaurantID = *restaurant.ID
	}
	timeNow := time.Now()
	r.restaurants = append(r.restaurants, Restaurant{
		Restaurant: entities.Restaurant{
			ID:          restaurantID,
			Name:        restaurant.Name,
			Description: restaurant.Description,
			CommonCQRSHistoryModel: entities.CommonCQRSHistoryModel{
				ActiveStatus: true,
				CreatedAt:    timeNow,
				UpdatedAt:    timeNow,
			},
		},
		editTypeCode: enums.EditTypeCodeCreate,
	})
	for i := range restaurant.Branches {
		err := r.addBranch(&r.restaurants[len(r.restaurants)-1], &restaurant.Branches[i])
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *restaurantDTOAggregateImpl) updateRestaurant(updated *Restaurant, update *dtos.Restaurant) error {
	if updated == nil || update == nil {
		return nil
	}
	ori := updated.Restaurant
	updated.Name = update.Name
	updated.Description = update.Description
	updated.CommonCQRSHistoryModel.UpdatedAt = time.Now()
	updated.update = update.ToUpdateRestaurant().RemoveUnchangedFields(ori)
	if updated.update != nil {
		updated.editTypeCode = enums.EditTypeCodeUpdate
	}

	deletedMap := make(map[string]struct{}) // id -> deleted
	for i := range update.Branches {
		err := r.saveBranch(updated, &update.Branches[i])
		if err != nil {
			return err
		}
		if update.Branches[i].ID != nil {
			deletedMap[*update.Branches[i].ID] = struct{}{}
		}
	}
	for i := range updated.Branches {
		if _, ok := deletedMap[updated.Branches[i].ID]; !ok {
			err := r.deleteBranch(&updated.Branches[i])
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// branch aggregate methods
func (r *restaurantDTOAggregateImpl) SaveBranch(restaurantID string, branch *dtos.Branch) error {
	if err := validatesdtos.ValidateBranch(branch); err != nil {
		return err
	}
	for i := range r.restaurants {
		if r.restaurants[i].ID == restaurantID {
			return r.saveBranch(&r.restaurants[i], branch)
		}
	}
	return customizeerrors.RestaurantNotFoundError
}

func (r *restaurantDTOAggregateImpl) saveBranch(updated *Restaurant, update *dtos.Branch) error {
	if updated == nil || update == nil {
		return nil
	}
	for i := range updated.Branches {
		if update.ID != nil && updated.Branches[i].ID == *update.ID {
			return r.updateBranch(&updated.Branches[i], update)
		}
		// same name but different id, duplicate branch
		if updated.Branches[i].Name == update.Name {
			return customizeerrors.BranchAlreadyExistsError
		}
	}
	return r.addBranch(updated, update)
}

func (r *restaurantDTOAggregateImpl) addBranch(restaurant *Restaurant, branch *dtos.Branch) error {
	if restaurant == nil || branch == nil {
		return nil
	}
	var branchID string
	if branch.ID == nil {
		branchID = r.uuidGenerator.GenerateUUID()
		branch.ID = &branchID
	} else {
		branchID = *branch.ID
	}
	timeNow := time.Now()
	restaurant.Branches = append(restaurant.Branches, Branch{
		Branch: entities.Branch{
			ID:           branchID,
			RestaurantID: restaurant.ID,
			Name:         branch.Name,
			Description:  branch.Description,
			CommonCQRSHistoryModel: entities.CommonCQRSHistoryModel{
				ActiveStatus: true,
				CreatedAt:    timeNow,
				UpdatedAt:    timeNow,
			},
		},
		editTypeCode: enums.EditTypeCodeCreate,
	})

	err := r.addAddress(&restaurant.Branches[len(restaurant.Branches)-1], branch.Address)
	if err != nil {
		return err
	}
	err = r.addPriceRange(&restaurant.Branches[len(restaurant.Branches)-1], branch.PriceRange)
	if err != nil {
		return err
	}
	for i := range branch.Categories {
		err := r.addCategoryRelation(&restaurant.Branches[len(restaurant.Branches)-1], &branch.Categories[i])
		if err != nil {
			return err
		}
	}
	for i := range branch.Tables {
		err := r.addTable(&restaurant.Branches[len(restaurant.Branches)-1], &branch.Tables[i])
		if err != nil {
			return err
		}
	}
	for i := range branch.Availables {
		err := r.addAvailable(&restaurant.Branches[len(restaurant.Branches)-1], &branch.Availables[i])
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *restaurantDTOAggregateImpl) deleteBranch(branch *Branch) error {
	if branch == nil {
		return nil
	}
	branch.ActiveStatus = false
	branch.CommonCQRSHistoryModel.UpdatedAt = time.Now()
	branch.editTypeCode = enums.EditTypeCodeDelete

	err := r.deleteAddress(branch.Address)
	if err != nil {
		return err
	}
	err = r.deletePriceRange(branch.PriceRange)
	if err != nil {
		return err
	}
	for i := range branch.BranchCategoryRelations {
		err := r.deleteCategoryRelation(&branch.BranchCategoryRelations[i])
		if err != nil {
			return err
		}
	}
	for i := range branch.Tables {
		err := r.deleteTable(&branch.Tables[i])
		if err != nil {
			return err
		}
	}
	for i := range branch.Availables {
		err := r.deleteAvailable(&branch.Availables[i])
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *restaurantDTOAggregateImpl) updateBranch(updated *Branch, update *dtos.Branch) error {
	if updated == nil || update == nil {
		return nil
	}
	ori := updated.Branch
	updated.Name = update.Name
	updated.Description = update.Description
	updated.CommonCQRSHistoryModel.UpdatedAt = time.Now()
	updated.update = update.ToUpdateBranch().RemoveUnchangedFields(ori)
	if updated.update != nil {
		updated.editTypeCode = enums.EditTypeCodeUpdate
	}

	err := r.saveAddress(updated, update.Address)
	if err != nil {
		return err
	}
	err = r.savePriceRange(updated, update.PriceRange)
	if err != nil {
		return err
	}
	deletedMap := make(map[string]struct{}) // id -> deleted
	for i := range update.Categories {
		err = r.saveCategoryRelation(updated, &update.Categories[i])
		if err != nil {
			return err
		}
		deletedMap[update.Categories[i].ID] = struct{}{}
	}
	for i := range updated.BranchCategoryRelations {
		if _, ok := deletedMap[updated.BranchCategoryRelations[i].Category.ID]; !ok {
			err = r.deleteCategoryRelation(&updated.BranchCategoryRelations[i])
			if err != nil {
				return err
			}
		}
	}
	deletedMap = make(map[string]struct{})
	for i := range update.Tables {
		err = r.saveTable(updated, &update.Tables[i])
		if err != nil {
			return err
		}
		if update.Tables[i].ID != nil {
			deletedMap[*update.Tables[i].ID] = struct{}{}
		}
	}
	for i := range updated.Tables {
		if _, ok := deletedMap[updated.Tables[i].ID]; !ok {
			err = r.deleteTable(&updated.Tables[i])
			if err != nil {
				return err
			}
		}
	}
	deletedMap = make(map[string]struct{})
	for i := range update.Availables {
		err = r.saveAvailable(updated, &update.Availables[i])
		if err != nil {
			return err
		}
		if update.Availables[i].ID != nil {
			deletedMap[*update.Availables[i].ID] = struct{}{}
		}
	}
	for i := range updated.Availables {
		if _, ok := deletedMap[updated.Availables[i].ID]; !ok {
			err = r.deleteAvailable(&updated.Availables[i])
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// address aggregate methods
func (r *restaurantDTOAggregateImpl) SaveAddress(branchID string, address *dtos.Address) error {
	if err := validatesdtos.ValidateAddress(address); err != nil {
		return err
	}
	for i := range r.restaurants {
		for j := range r.restaurants[i].Branches {
			if r.restaurants[i].Branches[j].ID == branchID {
				return r.saveAddress(&r.restaurants[i].Branches[j], address)
			}
		}
	}
	return customizeerrors.BranchNotFoundError
}

func (r *restaurantDTOAggregateImpl) saveAddress(updated *Branch, update *dtos.Address) error {
	if updated == nil || update == nil {
		return nil
	}
	if updated.Address != nil {
		ori := updated.Address.Address
		updated.Address.Street = update.Street
		updated.Address.CityCode = update.City.ToCityCode()
		updated.Address.PostalCode = update.PostalCode
		updated.Address.CountryCode = update.Country.ToCountryCode()
		updated.Address.CommonCQRSHistoryModel.UpdatedAt = time.Now()
		updated.Address.update = update.ToUpdateAddress().RemoveUnchangedFields(ori)
		if updated.Address.update != nil {
			updated.Address.editTypeCode = enums.EditTypeCodeUpdate
		}
		return nil
	}
	return r.addAddress(updated, update)
}

func (r *restaurantDTOAggregateImpl) addAddress(updated *Branch, update *dtos.Address) error {
	if updated == nil || update == nil {
		return nil
	}
	var addressID string
	if update.ID == nil {
		addressID = r.uuidGenerator.GenerateUUID()
		update.ID = &addressID
	} else {
		addressID = *update.ID
	}
	timeNow := time.Now()
	updated.Address = &Address{
		Address: entities.Address{
			ID:          addressID,
			BranchID:    updated.ID,
			Street:      update.Street,
			CityCode:    update.City.ToCityCode(),
			PostalCode:  update.PostalCode,
			CountryCode: update.Country.ToCountryCode(),
			CommonCQRSHistoryModel: entities.CommonCQRSHistoryModel{
				ActiveStatus: true,
				CreatedAt:    timeNow,
				UpdatedAt:    timeNow,
			},
		},
		editTypeCode: enums.EditTypeCodeCreate,
	}
	return nil
}

func (r *restaurantDTOAggregateImpl) deleteAddress(address *Address) error {
	if address == nil {
		return nil
	}
	address.ActiveStatus = false
	address.CommonCQRSHistoryModel.UpdatedAt = time.Now()
	address.editTypeCode = enums.EditTypeCodeDelete
	return nil
}

// price range aggregate methods
func (r *restaurantDTOAggregateImpl) SavePriceRange(branchID string, priceRange *dtos.PriceRange) error {
	if err := validatesdtos.ValidatePriceRange(priceRange); err != nil {
		return err
	}
	for i := range r.restaurants {
		for j := range r.restaurants[i].Branches {
			if r.restaurants[i].Branches[j].ID == branchID {
				return r.savePriceRange(&r.restaurants[i].Branches[j], priceRange)
			}
		}
	}
	return customizeerrors.BranchNotFoundError
}

func (r *restaurantDTOAggregateImpl) savePriceRange(updated *Branch, update *dtos.PriceRange) error {
	if updated == nil || update == nil {
		return nil
	}
	if updated.PriceRange != nil {
		ori := updated.PriceRange.PriceRange
		updated.PriceRange.MinPrice = update.MinPrice
		updated.PriceRange.MaxPrice = update.MaxPrice
		updated.PriceRange.CommonCQRSHistoryModel.UpdatedAt = time.Now()
		updated.PriceRange.update = update.ToUpdatePriceRange().RemoveUnchangedFields(ori)
		if updated.PriceRange.update != nil {
			updated.PriceRange.editTypeCode = enums.EditTypeCodeUpdate
		}
		return nil
	}
	return r.addPriceRange(updated, update)
}

func (r *restaurantDTOAggregateImpl) addPriceRange(updated *Branch, update *dtos.PriceRange) error {
	if updated == nil || update == nil {
		return nil
	}
	var priceRangeID string
	if update.ID == nil {
		priceRangeID = r.uuidGenerator.GenerateUUID()
		update.ID = &priceRangeID
	} else {
		priceRangeID = *update.ID
	}
	timeNow := time.Now()
	updated.PriceRange = &PriceRange{
		PriceRange: entities.PriceRange{
			ID:       priceRangeID,
			BranchID: updated.ID,
			MinPrice: update.MinPrice,
			MaxPrice: update.MaxPrice,
			CommonCQRSHistoryModel: entities.CommonCQRSHistoryModel{
				ActiveStatus: true,
				CreatedAt:    timeNow,
				UpdatedAt:    timeNow,
			},
		},
		editTypeCode: enums.EditTypeCodeCreate,
	}
	return nil
}

func (r *restaurantDTOAggregateImpl) deletePriceRange(priceRange *PriceRange) error {
	if priceRange == nil {
		return nil
	}
	priceRange.ActiveStatus = false
	priceRange.CommonCQRSHistoryModel.UpdatedAt = time.Now()
	priceRange.editTypeCode = enums.EditTypeCodeDelete
	return nil
}

// category aggregate methods
func (r *restaurantDTOAggregateImpl) SaveCategoryRelation(branchID string, category *dtos.Category) error {
	if err := validatesdtos.ValidateCategory(category); err != nil {
		return err
	}
	for i := range r.restaurants {
		for j := range r.restaurants[i].Branches {
			if r.restaurants[i].Branches[j].ID == branchID {
				return r.saveCategoryRelation(&r.restaurants[i].Branches[j], category)
			}
		}
	}
	return customizeerrors.BranchNotFoundError
}

func (r *restaurantDTOAggregateImpl) saveCategoryRelation(updated *Branch, update *dtos.Category) error {
	if updated == nil || update == nil {
		return nil
	}
	if update.ID == "" {
		return customizeerrors.CategoryIDEmptyError
	}
	for i := range updated.BranchCategoryRelations {
		if updated.BranchCategoryRelations[i].Category.ID == update.ID {
			// category already exists, do nothing
			return nil
		}
	}
	return r.addCategoryRelation(updated, update)
}

func (r *restaurantDTOAggregateImpl) addCategoryRelation(branch *Branch, category *dtos.Category) error {
	if branch == nil || category == nil {
		return nil
	}
	timeNow := time.Now()
	branch.BranchCategoryRelations = append(branch.BranchCategoryRelations, BranchCategoryRelation{
		BranchCategoryRelation: entities.BranchCategoryRelation{
			ID:         r.uuidGenerator.GenerateUUID(),
			BranchID:   branch.ID,
			CategoryID: category.ID,
			CommonCQRSHistoryModel: entities.CommonCQRSHistoryModel{
				ActiveStatus: true,
				CreatedAt:    timeNow,
				UpdatedAt:    timeNow,
			},
		},
		Category: Category{
			ID:           category.ID,
			CategoryCode: category.Category.ToCategoryCode(),
			CommonCQRSHistoryModel: entities.CommonCQRSHistoryModel{
				ActiveStatus: category.ActiveStatus,
				CreatedAt:    category.CreatedAt,
				UpdatedAt:    category.UpdatedAt,
			},
		},
		editTypeCode: enums.EditTypeCodeCreate,
	})

	return nil
}

func (r *restaurantDTOAggregateImpl) deleteCategoryRelation(category *BranchCategoryRelation) error {
	if category == nil {
		return nil
	}
	category.ActiveStatus = false
	category.CommonCQRSHistoryModel.UpdatedAt = time.Now()
	category.editTypeCode = enums.EditTypeCodeDelete
	return nil
}

// table aggregate methods
func (r *restaurantDTOAggregateImpl) SaveTable(branchID string, table *dtos.Table) error {
	if err := validatesdtos.ValidateTable(table); err != nil {
		return err
	}
	for i := range r.restaurants {
		for j := range r.restaurants[i].Branches {
			if r.restaurants[i].Branches[j].ID == branchID {
				return r.saveTable(&r.restaurants[i].Branches[j], table)
			}
		}
	}
	return customizeerrors.BranchNotFoundError
}

func (r *restaurantDTOAggregateImpl) saveTable(updated *Branch, update *dtos.Table) error {
	if updated == nil || update == nil {
		return nil
	}
	for i := range updated.Tables {
		if update.ID != nil && updated.Tables[i].ID == *update.ID {
			return r.updateTable(&updated.Tables[i], update)
		}
	}
	return r.addTable(updated, update)
}

func (r *restaurantDTOAggregateImpl) addTable(branch *Branch, table *dtos.Table) error {
	if branch == nil || table == nil {
		return nil
	}
	var tableID string
	if table.ID == nil {
		tableID = r.uuidGenerator.GenerateUUID()
		table.ID = &tableID
	} else {
		tableID = *table.ID
	}
	timeNow := time.Now()
	branch.Tables = append(branch.Tables, Table{
		Table: entities.Table{
			ID:       tableID,
			BranchID: branch.ID,
			Capacity: table.Capacity,
			CommonCQRSHistoryModel: entities.CommonCQRSHistoryModel{
				ActiveStatus: true,
				CreatedAt:    timeNow,
				UpdatedAt:    timeNow,
			},
		},
		editTypeCode: enums.EditTypeCodeCreate,
	})
	return nil
}

func (r *restaurantDTOAggregateImpl) deleteTable(table *Table) error {
	if table == nil {
		return nil
	}
	table.ActiveStatus = false
	table.CommonCQRSHistoryModel.UpdatedAt = time.Now()
	table.editTypeCode = enums.EditTypeCodeDelete
	return nil
}

func (r *restaurantDTOAggregateImpl) updateTable(updated *Table, update *dtos.Table) error {
	if updated == nil || update == nil {
		return nil
	}
	ori := updated.Table
	updated.Capacity = update.Capacity
	updated.CommonCQRSHistoryModel.UpdatedAt = time.Now()
	updated.update = update.ToUpdateTable().RemoveUnchangedFields(ori)
	if updated.update != nil {
		updated.editTypeCode = enums.EditTypeCodeUpdate
	}
	return nil
}

// table available aggregate methods
func (r *restaurantDTOAggregateImpl) SaveAvailable(branchID string, available *dtos.Available) error {
	if err := validatesdtos.ValidateAvailable(available); err != nil {
		return err
	}
	for i := range r.restaurants {
		for j := range r.restaurants[i].Branches {
			if r.restaurants[i].Branches[j].ID == branchID {
				return r.saveAvailable(&r.restaurants[i].Branches[j], available)
			}
		}
	}
	return customizeerrors.BranchNotFoundError
}

func (r *restaurantDTOAggregateImpl) saveAvailable(updated *Branch, update *dtos.Available) error {
	if updated == nil || update == nil {
		return nil
	}
	for i := range updated.Availables {
		if update.ID != nil && updated.Availables[i].ID == *update.ID {
			return r.updateAvailable(&updated.Availables[i], update)
		}
	}
	return r.addAvailable(updated, update)
}

func (r *restaurantDTOAggregateImpl) addAvailable(branch *Branch, available *dtos.Available) error {
	if branch == nil || available == nil {
		return nil
	}
	var availableID string
	if available.ID == nil {
		availableID = r.uuidGenerator.GenerateUUID()
		available.ID = &availableID
	} else {
		availableID = *available.ID
	}

	timeNow := time.Now()
	branch.Availables = append(branch.Availables, Available{
		Available: entities.Available{
			ID:        availableID,
			BranchID:  branch.ID,
			Weekday:   available.Weekday,
			StartTime: available.StartTime,
			EndTime:   available.EndTime,
			CommonCQRSHistoryModel: entities.CommonCQRSHistoryModel{
				ActiveStatus: true,
				CreatedAt:    timeNow,
				UpdatedAt:    timeNow,
			},
		},
		editTypeCode: enums.EditTypeCodeCreate,
	})

	return nil
}

func (r *restaurantDTOAggregateImpl) updateAvailable(updated *Available, update *dtos.Available) error {
	if updated == nil || update == nil {
		return nil
	}
	ori := updated.Available
	updated.Weekday = update.Weekday
	updated.StartTime = update.StartTime
	updated.EndTime = update.EndTime
	updated.CommonCQRSHistoryModel.UpdatedAt = time.Now()
	updated.update = update.ToUpdateAvailable().RemoveUnchangedFields(ori)
	if updated.update != nil {
		updated.editTypeCode = enums.EditTypeCodeUpdate
	}
	return nil
}

func (r *restaurantDTOAggregateImpl) deleteAvailable(Available *Available) error {
	if Available == nil {
		return nil
	}
	Available.ActiveStatus = false
	Available.CommonCQRSHistoryModel.UpdatedAt = time.Now()
	Available.editTypeCode = enums.EditTypeCodeDelete
	return nil
}

func (r *restaurantDTOAggregateImpl) SetAllEditTypeCodeToNone() {
	for i := range r.restaurants {
		r.restaurants[i].editTypeCode = enums.EditTypeCodeNone
		for j := range r.restaurants[i].Branches {
			r.restaurants[i].Branches[j].editTypeCode = enums.EditTypeCodeNone
			if r.restaurants[i].Branches[j].Address != nil {
				r.restaurants[i].Branches[j].Address.editTypeCode = enums.EditTypeCodeNone
			}
			if r.restaurants[i].Branches[j].PriceRange != nil {
				r.restaurants[i].Branches[j].PriceRange.editTypeCode = enums.EditTypeCodeNone
			}
			for k := range r.restaurants[i].Branches[j].BranchCategoryRelations {
				r.restaurants[i].Branches[j].BranchCategoryRelations[k].editTypeCode = enums.EditTypeCodeNone
			}
			for k := range r.restaurants[i].Branches[j].Tables {
				r.restaurants[i].Branches[j].Tables[k].editTypeCode = enums.EditTypeCodeNone
			}
			for k := range r.restaurants[i].Branches[j].Availables {
				r.restaurants[i].Branches[j].Availables[k].editTypeCode = enums.EditTypeCodeNone
			}
		}
	}
}

func (r *restaurantDTOAggregateImpl) GetAggregates() []Restaurant {
	return r.restaurants
}

func (r *restaurantDTOAggregateImpl) GetEditEntities() []RestaurantEditEntities {
	res := make([]RestaurantEditEntities, len(r.restaurants))
	for i, restaurant := range r.restaurants {
		// Initialize container
		res[i] = newRestaurantEditEntities()
		createEntities, updateEntities, deleteEntities := res[i].CreateEntities, res[i].UpdateEntities, res[i].DeleteEntities

		// Handle restaurant
		appendEntityByType(
			&createEntities.Restaurants,
			&updateEntities.Restaurants,
			&deleteEntities.Restaurants,
			restaurant.Restaurant,
			restaurant.update,
			restaurant.editTypeCode,
		)

		for _, branch := range restaurant.Branches {
			// Handle branch
			appendEntityByType(
				&createEntities.Branches,
				&updateEntities.Branches,
				&deleteEntities.Branches,
				branch.Branch,
				branch.update,
				branch.editTypeCode,
			)

			// Handle address
			if branch.Address != nil {
				appendEntityByType(
					&createEntities.Addresses,
					&updateEntities.Addresses,
					&deleteEntities.Addresses,
					branch.Address.Address,
					branch.Address.update,
					branch.Address.editTypeCode,
				)
			}

			// Handle price range
			if branch.PriceRange != nil {
				appendEntityByType(
					&createEntities.PriceRanges,
					&updateEntities.PriceRanges,
					&deleteEntities.PriceRanges,
					branch.PriceRange.PriceRange,
					branch.PriceRange.update,
					branch.PriceRange.editTypeCode,
				)
			}

			// Handle category relations
			for _, relation := range branch.BranchCategoryRelations {
				appendEntityByType(
					&createEntities.BranchCategoryRelations,
					&updateEntities.BranchCategoryRelations,
					&deleteEntities.BranchCategoryRelations,
					relation.BranchCategoryRelation,
					relation.update,
					relation.editTypeCode,
				)
			}

			// Handle table
			for _, table := range branch.Tables {
				appendEntityByType(
					&createEntities.Tables,
					&updateEntities.Tables,
					&deleteEntities.Tables,
					table.Table,
					table.update,
					table.editTypeCode,
				)
			}

			// Handle availabilities
			for _, available := range branch.Availables {
				appendEntityByType(
					&createEntities.Availables,
					&updateEntities.Availables,
					&deleteEntities.Availables,
					available.Available,
					available.update,
					available.editTypeCode,
				)
			}
		}
	}
	return res
}
