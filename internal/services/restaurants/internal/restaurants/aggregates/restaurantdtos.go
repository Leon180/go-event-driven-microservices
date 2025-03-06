package aggregates

import (
	"time"

	customizeerrors "github.com/Leon180/go-event-driven-microservices/internal/pkg/customize_errors"
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/enums"
	"github.com/Leon180/go-event-driven-microservices/internal/pkg/uuid"
	"github.com/Leon180/go-event-driven-microservices/internal/services/restaurants/internal/restaurants/dtos"
	"github.com/Leon180/go-event-driven-microservices/internal/services/restaurants/internal/restaurants/entities"
)

type RestaurantDTOAggregate interface {
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

	// save table available to the aggregate
	SaveTableAvailable(tableID string, tableAvailable *dtos.TableAvailable) error

	// get restaurant aggregates
	GetAggregates() []Restaurant

	// get pending edit entities during update
	GetEditEntities() []RestaurantEditEntities
}

func NewRestaurantDTOAggregate(uuidGenerator uuid.UUIDGenerator) RestaurantDTOAggregate {
	return &RestaurantDTOAggregateImpl{
		uuidGenerator: uuidGenerator,
	}
}

type RestaurantDTOAggregateImpl struct {
	restaurants []Restaurant

	// dependencies
	uuidGenerator uuid.UUIDGenerator
}

type RestaurantEntities struct {
	Restaurants       entities.Restaurant
	Branches          []entities.Branch
	Addresses         map[string]entities.Address                  // branch id -> address
	PriceRanges       map[string]entities.PriceRange               // branch id -> price range
	CategoryRelations map[string][]entities.BranchCategoryRelation // branch id -> category relation
	Category          []entities.Category
	Tables            map[string][]entities.Table          // branch id -> table
	TableAvailables   map[string][]entities.TableAvailable // table id -> table available
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
	Restaurants       []entities.Restaurant
	Branches          []entities.Branch
	Addresses         []entities.Address
	PriceRanges       []entities.PriceRange
	CategoryRelations []entities.BranchCategoryRelation
	Tables            []entities.Table
	TableAvailables   []entities.TableAvailable
}

type RestaurantUpdateEntities struct {
	Restaurants       []entities.Restaurant
	Branches          []entities.Branch
	Addresses         []entities.Address
	PriceRanges       []entities.PriceRange
	CategoryRelations []entities.BranchCategoryRelation
	Tables            []entities.Table
	TableAvailables   []entities.TableAvailable
}

type RestaurantDeleteEntities struct {
	Restaurants       []entities.Restaurant
	Branches          []entities.Branch
	Addresses         []entities.Address
	PriceRanges       []entities.PriceRange
	CategoryRelations []entities.BranchCategoryRelation
	Tables            []entities.Table
	TableAvailables   []entities.TableAvailable
}

type Restaurant struct {
	entities.Restaurant
	editTypeCode enums.EditTypeCode

	// Relations
	Branches []Branch `gorm:"foreignKey:RestaurantID;references:ID" comment:"Branches"`
}

func (r *Restaurant) TableName() string {
	return "restaurant"
}

type Branch struct {
	entities.Branch
	editTypeCode enums.EditTypeCode

	// Relations
	Address                 *Address                 `gorm:"foreignKey:BranchID;references:ID" comment:"Address"`
	PriceRange              *PriceRange              `gorm:"foreignKey:BranchID;references:ID" comment:"Price Range"`
	BranchCategoryRelations []BranchCategoryRelation `gorm:"foreignKey:BranchID;references:ID" comment:"Branch Category Relation"`
	Tables                  []Table                  `gorm:"foreignKey:BranchID;references:ID" comment:"Tables"`
}

func (b *Branch) TableName() string {
	return "branch"
}

type Address struct {
	entities.Address
	editTypeCode enums.EditTypeCode
}

type PriceRange struct {
	entities.PriceRange
	editTypeCode enums.EditTypeCode
}

type BranchCategoryRelation struct {
	entities.BranchCategoryRelation
	editTypeCode enums.EditTypeCode

	// Relations
	Category Category `gorm:"foreignKey:ID;references:CategoryID" comment:"Category"`
}

func (bcr *BranchCategoryRelation) TableName() string {
	return "branch_category_relation"
}

type Category entities.Category

type Table struct {
	entities.Table
	editTypeCode enums.EditTypeCode

	// Relations
	TableAvailables []TableAvailable `gorm:"foreignKey:TableID;references:ID" comment:"Table Availables"`
}

func (t *Table) TableName() string {
	return "table"
}

type TableAvailable struct {
	entities.TableAvailable
	editTypeCode enums.EditTypeCode
}

// restaurant aggregate methods
func (r *RestaurantDTOAggregateImpl) SaveRestaurant(restaurant *dtos.Restaurant) error {
	if restaurant == nil {
		return nil
	}
	if restaurant.Name == "" {
		return customizeerrors.RestaurantNameEmptyError
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

func (r *RestaurantDTOAggregateImpl) addRestaurant(restaurant *dtos.Restaurant) error {
	if restaurant == nil {
		return nil
	}
	var restaurantID string
	if restaurant.ID == nil {
		restaurantID = r.uuidGenerator.GenerateUUID()
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

func (r *RestaurantDTOAggregateImpl) updateRestaurant(updated *Restaurant, update *dtos.Restaurant) error {
	if updated == nil || update == nil {
		return nil
	}
	updated.Name = update.Name
	updated.Description = update.Description
	updated.CommonCQRSHistoryModel.UpdatedAt = time.Now()
	updated.editTypeCode = enums.EditTypeCodeUpdate

	deletedMap := make(map[string]struct{}) // id -> deleted
	for i := range update.Branches {
		err := r.saveBranch(updated, &update.Branches[i])
		if err != nil {
			return err
		}
		deletedMap[*update.Branches[i].ID] = struct{}{}
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
func (r *RestaurantDTOAggregateImpl) SaveBranch(restaurantID string, branch *dtos.Branch) error {
	for i := range r.restaurants {
		if r.restaurants[i].ID == restaurantID {
			return r.saveBranch(&r.restaurants[i], branch)
		}
	}
	return customizeerrors.RestaurantNotFoundError
}

func (r *RestaurantDTOAggregateImpl) saveBranch(updated *Restaurant, update *dtos.Branch) error {
	if updated == nil || update == nil {
		return nil
	}
	if update.Name == "" {
		return customizeerrors.BranchNameEmptyError
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

func (r *RestaurantDTOAggregateImpl) addBranch(restaurant *Restaurant, branch *dtos.Branch) error {
	if restaurant == nil || branch == nil {
		return nil
	}
	var branchID string
	if branch.ID == nil {
		branchID = r.uuidGenerator.GenerateUUID()
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

	return nil
}

func (r *RestaurantDTOAggregateImpl) deleteBranch(branch *Branch) error {
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
	return nil
}

func (r *RestaurantDTOAggregateImpl) updateBranch(updated *Branch, update *dtos.Branch) error {
	if updated == nil || update == nil {
		return nil
	}
	updated.Name = update.Name
	updated.Description = update.Description
	updated.CommonCQRSHistoryModel.UpdatedAt = time.Now()
	updated.editTypeCode = enums.EditTypeCodeUpdate

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
		deletedMap[*update.Tables[i].ID] = struct{}{}
	}
	for i := range updated.Tables {
		if _, ok := deletedMap[updated.Tables[i].ID]; !ok {
			err = r.deleteTable(&updated.Tables[i])
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// address aggregate methods
func (r *RestaurantDTOAggregateImpl) SaveAddress(branchID string, address *dtos.Address) error {
	for i := range r.restaurants {
		for j := range r.restaurants[i].Branches {
			if r.restaurants[i].Branches[j].ID == branchID {
				return r.saveAddress(&r.restaurants[i].Branches[j], address)
			}
		}
	}
	return customizeerrors.BranchNotFoundError
}

func (r *RestaurantDTOAggregateImpl) saveAddress(updated *Branch, update *dtos.Address) error {
	if updated == nil || update == nil {
		return nil
	}
	if update.Country == "" || update.City == "" || update.PostalCode == "" || update.Street == "" {
		return customizeerrors.AddressEmptyError
	}
	if updated.Address != nil {
		updated.Address.Street = update.Street
		updated.Address.CityCode = update.City.ToCityCode()
		updated.Address.PostalCode = update.PostalCode
		updated.Address.CountryCode = update.Country.ToCountryCode()
		updated.Address.CommonCQRSHistoryModel.UpdatedAt = time.Now()
		updated.Address.editTypeCode = enums.EditTypeCodeUpdate
		return nil
	}
	return r.addAddress(updated, update)
}

func (r *RestaurantDTOAggregateImpl) addAddress(updated *Branch, update *dtos.Address) error {
	if updated == nil || update == nil {
		return nil
	}
	var addressID string
	if update.ID == nil {
		addressID = r.uuidGenerator.GenerateUUID()
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

func (r *RestaurantDTOAggregateImpl) deleteAddress(address *Address) error {
	if address == nil {
		return nil
	}
	address.ActiveStatus = false
	address.CommonCQRSHistoryModel.UpdatedAt = time.Now()
	address.editTypeCode = enums.EditTypeCodeDelete
	return nil
}

// price range aggregate methods
func (r *RestaurantDTOAggregateImpl) SavePriceRange(branchID string, priceRange *dtos.PriceRange) error {
	for i := range r.restaurants {
		for j := range r.restaurants[i].Branches {
			if r.restaurants[i].Branches[j].ID == branchID {
				return r.savePriceRange(&r.restaurants[i].Branches[j], priceRange)
			}
		}
	}
	return customizeerrors.BranchNotFoundError
}

func (r *RestaurantDTOAggregateImpl) savePriceRange(updated *Branch, update *dtos.PriceRange) error {
	if updated == nil || update == nil {
		return nil
	}
	if update.MinPrice == 0 || update.MaxPrice <= update.MinPrice {
		return customizeerrors.PriceRangeInvalidError
	}
	if updated.PriceRange != nil {
		updated.PriceRange.MinPrice = update.MinPrice
		updated.PriceRange.MaxPrice = update.MaxPrice
		updated.PriceRange.CommonCQRSHistoryModel.UpdatedAt = time.Now()
		return nil
	}
	return r.addPriceRange(updated, update)
}

func (r *RestaurantDTOAggregateImpl) addPriceRange(updated *Branch, update *dtos.PriceRange) error {
	if updated == nil || update == nil {
		return nil
	}
	var priceRangeID string
	if update.ID == nil {
		priceRangeID = r.uuidGenerator.GenerateUUID()
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

func (r *RestaurantDTOAggregateImpl) deletePriceRange(priceRange *PriceRange) error {
	if priceRange == nil {
		return nil
	}
	priceRange.ActiveStatus = false
	priceRange.CommonCQRSHistoryModel.UpdatedAt = time.Now()
	priceRange.editTypeCode = enums.EditTypeCodeDelete
	return nil
}

// category aggregate methods
func (r *RestaurantDTOAggregateImpl) SaveCategoryRelation(branchID string, category *dtos.Category) error {
	for i := range r.restaurants {
		for j := range r.restaurants[i].Branches {
			if r.restaurants[i].Branches[j].ID == branchID {
				return r.saveCategoryRelation(&r.restaurants[i].Branches[j], category)
			}
		}
	}
	return customizeerrors.BranchNotFoundError
}

func (r *RestaurantDTOAggregateImpl) saveCategoryRelation(updated *Branch, update *dtos.Category) error {
	if updated == nil || update == nil {
		return nil
	}
	if update.Category == "" {
		return customizeerrors.CategoryNameEmptyError
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

func (r *RestaurantDTOAggregateImpl) addCategoryRelation(branch *Branch, category *dtos.Category) error {
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

func (r *RestaurantDTOAggregateImpl) deleteCategoryRelation(category *BranchCategoryRelation) error {
	if category == nil {
		return nil
	}
	category.ActiveStatus = false
	category.CommonCQRSHistoryModel.UpdatedAt = time.Now()
	category.editTypeCode = enums.EditTypeCodeDelete
	return nil
}

// table aggregate methods
func (r *RestaurantDTOAggregateImpl) SaveTable(branchID string, table *dtos.Table) error {
	for i := range r.restaurants {
		for j := range r.restaurants[i].Branches {
			if r.restaurants[i].Branches[j].ID == branchID {
				return r.saveTable(&r.restaurants[i].Branches[j], table)
			}
		}
	}
	return customizeerrors.BranchNotFoundError
}

func (r *RestaurantDTOAggregateImpl) saveTable(updated *Branch, update *dtos.Table) error {
	if updated == nil || update == nil {
		return nil
	}
	if update.Capacity <= 0 {
		return customizeerrors.TableCapacityInvalidError
	}
	for i := range updated.Tables {
		if update.ID != nil && updated.Tables[i].ID == *update.ID {
			return r.updateTable(&updated.Tables[i], update)
		}
	}
	return r.addTable(updated, update)
}

func (r *RestaurantDTOAggregateImpl) addTable(branch *Branch, table *dtos.Table) error {
	if branch == nil || table == nil {
		return nil
	}
	var tableID string
	if table.ID == nil {
		tableID = r.uuidGenerator.GenerateUUID()
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

	for i := range table.TableAvailables {
		err := r.addTableAvailable(&branch.Tables[len(branch.Tables)-1], &table.TableAvailables[i])
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *RestaurantDTOAggregateImpl) deleteTable(table *Table) error {
	if table == nil {
		return nil
	}
	table.ActiveStatus = false
	table.CommonCQRSHistoryModel.UpdatedAt = time.Now()
	table.editTypeCode = enums.EditTypeCodeDelete
	return nil
}

func (r *RestaurantDTOAggregateImpl) updateTable(updated *Table, update *dtos.Table) error {
	if updated == nil || update == nil {
		return nil
	}
	updated.Capacity = update.Capacity
	updated.CommonCQRSHistoryModel.UpdatedAt = time.Now()
	updated.editTypeCode = enums.EditTypeCodeUpdate

	deletedMap := make(map[string]struct{}) // id -> deleted
	for i := range update.TableAvailables {
		err := r.saveTableAvailable(updated, &update.TableAvailables[i])
		if err != nil {
			return err
		}
		deletedMap[*update.TableAvailables[i].ID] = struct{}{}
	}
	for i := range updated.TableAvailables {
		if _, ok := deletedMap[updated.TableAvailables[i].ID]; !ok {
			err := r.deleteTableAvailable(&updated.TableAvailables[i])
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// table available aggregate methods
func (r *RestaurantDTOAggregateImpl) SaveTableAvailable(tableID string, tableAvailable *dtos.TableAvailable) error {
	for i := range r.restaurants {
		for j := range r.restaurants[i].Branches {
			for k := range r.restaurants[i].Branches[j].Tables {
				if r.restaurants[i].Branches[j].Tables[k].ID == tableID {
					return r.saveTableAvailable(&r.restaurants[i].Branches[j].Tables[k], tableAvailable)
				}
			}
		}
	}
	return customizeerrors.TableNotFoundError
}

func (r *RestaurantDTOAggregateImpl) saveTableAvailable(updated *Table, update *dtos.TableAvailable) error {
	if updated == nil || update == nil {
		return nil
	}
	if update.Weekday < 0 || update.Weekday > 6 {
		return customizeerrors.TableAvailableTimeInvalidError
	}
	_, err := time.Parse(enums.TimeFormatClockOnly.ToString(), update.StartTime)
	if err != nil {
		return customizeerrors.TableAvailableTimeInvalidError
	}
	_, err = time.Parse(enums.TimeFormatClockOnly.ToString(), update.EndTime)
	if err != nil {
		return customizeerrors.TableAvailableTimeInvalidError
	}
	for i := range updated.TableAvailables {
		if update.ID != nil && updated.TableAvailables[i].ID == *update.ID {
			return r.updateTableAvailable(&updated.TableAvailables[i], update)
		}
	}
	return r.addTableAvailable(updated, update)
}

func (r *RestaurantDTOAggregateImpl) addTableAvailable(table *Table, tableAvailable *dtos.TableAvailable) error {
	if tableAvailable == nil || table == nil {
		return nil
	}
	var tableAvailableID string
	if tableAvailable.ID == nil {
		tableAvailableID = r.uuidGenerator.GenerateUUID()
	} else {
		tableAvailableID = *tableAvailable.ID
	}

	timeNow := time.Now()
	table.TableAvailables = append(table.TableAvailables, TableAvailable{
		TableAvailable: entities.TableAvailable{
			ID:        tableAvailableID,
			TableID:   table.ID,
			Weekday:   tableAvailable.Weekday,
			StartTime: tableAvailable.StartTime,
			EndTime:   tableAvailable.EndTime,
			Booked:    tableAvailable.Booked,
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

func (r *RestaurantDTOAggregateImpl) updateTableAvailable(updated *TableAvailable, update *dtos.TableAvailable) error {
	if updated == nil || update == nil {
		return nil
	}
	updated.Weekday = update.Weekday
	updated.StartTime = update.StartTime
	updated.EndTime = update.EndTime
	updated.Booked = update.Booked
	updated.CommonCQRSHistoryModel.UpdatedAt = time.Now()
	updated.editTypeCode = enums.EditTypeCodeUpdate
	return nil
}

func (r *RestaurantDTOAggregateImpl) deleteTableAvailable(tableAvailable *TableAvailable) error {
	if tableAvailable == nil {
		return nil
	}
	tableAvailable.ActiveStatus = false
	tableAvailable.CommonCQRSHistoryModel.UpdatedAt = time.Now()
	tableAvailable.editTypeCode = enums.EditTypeCodeDelete
	return nil
}

func (r *RestaurantDTOAggregateImpl) GetAggregates() []Restaurant {
	return r.restaurants
}

func (r *RestaurantDTOAggregateImpl) GetEditEntities() []RestaurantEditEntities {
	res := make([]RestaurantEditEntities, len(r.restaurants))
	for i, restaurant := range r.restaurants {
		// Initialize container
		res[i] = newRestaurantEditEntities()
		createEntities, updateEntities, deleteEntities := res[i].CreateEntities, res[i].UpdateEntities, res[i].DeleteEntities

		// Handle restaurant
		appendEntityByType(&createEntities.Restaurants, &updateEntities.Restaurants, &deleteEntities.Restaurants, restaurant.Restaurant, restaurant.editTypeCode)

		for _, branch := range restaurant.Branches {
			// Handle branch
			appendEntityByType(&createEntities.Branches, &updateEntities.Branches, &deleteEntities.Branches, branch.Branch, branch.editTypeCode)

			// Handle address
			if branch.Address != nil {
				appendEntityByType(&createEntities.Addresses, &updateEntities.Addresses, &deleteEntities.Addresses, branch.Address.Address, branch.Address.editTypeCode)
			}

			// Handle price range
			if branch.PriceRange != nil {
				appendEntityByType(&createEntities.PriceRanges, &updateEntities.PriceRanges, &deleteEntities.PriceRanges, branch.PriceRange.PriceRange, branch.PriceRange.editTypeCode)
			}

			// Handle category relations
			for _, relation := range branch.BranchCategoryRelations {
				appendEntityByType(&createEntities.CategoryRelations, &updateEntities.CategoryRelations, &deleteEntities.CategoryRelations, relation.BranchCategoryRelation, relation.editTypeCode)
			}

			for _, table := range branch.Tables {
				// Handle table
				appendEntityByType(&createEntities.Tables, &updateEntities.Tables, &deleteEntities.Tables, table.Table, table.editTypeCode)

				// Handle table availabilities
				for _, available := range table.TableAvailables {
					appendEntityByType(&createEntities.TableAvailables, &updateEntities.TableAvailables, &deleteEntities.TableAvailables, available.TableAvailable, available.editTypeCode)
				}
			}
		}
	}
	return res
}

func appendEntityByType[T any](create *[]T, update *[]T, delete *[]T, entity T, editTypeCode enums.EditTypeCode) {
	switch editTypeCode {
	case enums.EditTypeCodeCreate:
		*create = append(*create, entity)
	case enums.EditTypeCodeUpdate:
		*update = append(*update, entity)
	case enums.EditTypeCodeDelete:
		*delete = append(*delete, entity)
	}
}
