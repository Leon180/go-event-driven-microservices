package repostgresespostgres

import (
	"context"

	"github.com/Leon180/go-event-driven-microservices/internal/pkg/enums"
	contextloggers "github.com/Leon180/go-event-driven-microservices/internal/pkg/utilities/context_loggers"
	"github.com/Leon180/go-event-driven-microservices/internal/services/restaurants/internal/restaurants/aggregates"
	"github.com/Leon180/go-event-driven-microservices/internal/services/restaurants/internal/restaurants/dtos"
	"github.com/Leon180/go-event-driven-microservices/internal/services/restaurants/internal/restaurants/entities"
	"github.com/Leon180/go-event-driven-microservices/internal/services/restaurants/internal/restaurants/repositories"
	"github.com/samber/lo"
	"gorm.io/gorm"
)

func NewSearchRestaurantsFullInfoRepository(
	db *gorm.DB,
	contextLogger contextloggers.ContextLogger,
) repositories.SearchRestaurantsFullInfo {
	return &SearchRestaurantsFullInfoImpl{
		db:            db,
		contextLogger: contextLogger,
	}
}

type SearchRestaurantsFullInfoImpl struct {
	db            *gorm.DB
	contextLogger contextloggers.ContextLogger
}

func (r *SearchRestaurantsFullInfoImpl) SearchRestaurantsFullInfo(ctx context.Context, search dtos.SearchRestaurants) ([]aggregates.Restaurant, error) {
	sql := r.buildBaseQuery(ctx).
		Scopes(
			r.applyNameFilter(search.NameFilter),
			r.applyDescriptionFilter(search.DescriptionFilter),
			r.applyCityFilter(search.CityFilter),
			r.applyCountryFilter(search.CountryFilter),
			r.applyPriceRangeFilter(search.MinPriceFilter, search.MaxPriceFilter),
			r.applyCategoryFilter(search.CategoryFilter),
			r.applyTableAvailabilityFilter(search),
			r.applyPagination(search.Pagination),
			r.applyOrdering(search.OrderBy),
		)

	var restaurants []aggregates.Restaurant
	if err := sql.Find(&restaurants).Error; err != nil {
		r.contextLogger.WithContextInfo(ctx, enums.ContextKeyTraceID).Error("failed to search restaurants full info", err)
		return nil, err
	}
	return restaurants, nil
}

func (r *SearchRestaurantsFullInfoImpl) buildBaseQuery(ctx context.Context) *gorm.DB {
	return r.db.WithContext(ctx).
		Preload("Branches").
		Preload("Branches.Address").
		Preload("Branches.PriceRange").
		Preload("Branches.BranchCategoryRelations").
		Preload("Branches.BranchCategoryRelations.Category").
		Preload("Branches.Tables").
		Preload("Branches.Availables")
}

func (r *SearchRestaurantsFullInfoImpl) applyNameFilter(name *string) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if name != nil {
			return db.Where("name LIKE ?", "%"+*name+"%")
		}
		return db
	}
}

func (r *SearchRestaurantsFullInfoImpl) applyDescriptionFilter(desc *string) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if desc != nil {
			return db.Where("description LIKE ?", "%"+*desc+"%")
		}
		return db
	}
}

func (r *SearchRestaurantsFullInfoImpl) applyCityFilter(cities []enums.City) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if len(cities) > 0 {
			cityCodes := lo.Map(cities, func(city enums.City, _ int) int {
				return int(city.ToCityCode())
			})
			return db.Where(`EXISTS (
                SELECT 1 FROM branch
                JOIN address ON branch.id = address.branch_id
                WHERE branch.restaurant_id = restaurant.id
                AND address.city_code IN (?)
            )`, cityCodes)
		}
		return db
	}
}

func (r *SearchRestaurantsFullInfoImpl) applyCountryFilter(countries []enums.Country) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if len(countries) > 0 {
			countryCodes := lo.Map(countries, func(country enums.Country, _ int) int {
				return int(country.ToCountryCode())
			})
			return db.Where(`EXISTS (
                SELECT 1 FROM branch
                JOIN address ON branch.id = address.branch_id
                WHERE branch.restaurant_id = restaurant.id
                AND address.country_code IN (?)
            )`, countryCodes)
		}
		return db
	}
}

func (r *SearchRestaurantsFullInfoImpl) applyPriceRangeFilter(min, max *int) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if min != nil || max != nil {
			query := db.Where(`EXISTS (
                SELECT 1 FROM branch
                JOIN price_range ON branch.id = price_range.branch_id
                WHERE branch.restaurant_id = restaurant.id`)

			if min != nil {
				query = query.Where("price_range.min_price >= ?", *min)
			}
			if max != nil {
				query = query.Where("price_range.max_price <= ?", *max)
			}
			return query.Where(")")
		}
		return db
	}
}

func (r *SearchRestaurantsFullInfoImpl) applyCategoryFilter(categories []enums.Category) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if len(categories) > 0 {
			categoryCodes := lo.Map(categories, func(category enums.Category, _ int) int {
				return int(category.ToCategoryCode())
			})
			return db.Where(`EXISTS (
                SELECT 1 FROM branch
                JOIN branch_category_relation ON branch.id = branch_category_relation.branch_id
                JOIN category ON branch_category_relation.category_id = category.id
                WHERE branch.restaurant_id = restaurant.id
                AND category.category_code IN (?)
            )`, categoryCodes)
		}
		return db
	}
}

func (r *SearchRestaurantsFullInfoImpl) applyTableAvailabilityFilter(search dtos.SearchRestaurants) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if len(search.TableAvailableWeek) == 0 && search.TableAvailableStartTime == nil && search.TableAvailableEndTime == nil {
			return db
		}

		query := db.Where(`EXISTS (
            SELECT 1 FROM branch
            JOIN available ON branch.id = available.branch_id
            WHERE branch.restaurant_id = restaurant.id`)

		if len(search.TableAvailableWeek) > 0 {
			query = query.Where("available.weekday IN (?)", search.TableAvailableWeek)
		}
		if search.TableAvailableStartTime != nil {
			query = query.Where("available.start_time >= ?", *search.TableAvailableStartTime)
		}
		if search.TableAvailableEndTime != nil {
			query = query.Where("available.end_time <= ?", *search.TableAvailableEndTime)
		}

		return query.Where(")")
	}
}

func (r *SearchRestaurantsFullInfoImpl) applyPagination(pagination *dtos.Pagination) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if pagination != nil {
			return db.Limit(pagination.PageSize).
				Offset((pagination.Page - 1) * pagination.PageSize)
		}
		return db
	}
}

func (r *SearchRestaurantsFullInfoImpl) applyOrdering(orderBy []dtos.OrderBy) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		for _, ob := range orderBy {
			db = db.Order(ob.ToSort())
		}
		return db
	}
}

// Restaurants Impl
func NewRestaurantsRepository(
	db *gorm.DB,
	contextLogger contextloggers.ContextLogger,
) repositories.Restaurants {
	return &RestaurantsImpl{
		db:            db,
		contextLogger: contextLogger,
	}
}

type RestaurantsImpl struct {
	db            *gorm.DB
	contextLogger contextloggers.ContextLogger
}

func (r *RestaurantsImpl) CreateRestaurants(ctx context.Context, restaurants entities.Restaurants) error {
	if len(restaurants) == 0 {
		return nil
	}
	if err := r.db.WithContext(ctx).Create(&restaurants).Error; err != nil {
		r.contextLogger.WithContextInfo(ctx, enums.ContextKeyTraceID).Error("failed to create restaurants", err)
		return err
	}
	return nil
}

func (r *RestaurantsImpl) ReadRestaurantFullInfo(ctx context.Context, id string) (aggregates.Restaurant, error) {
	if id == "" {
		return aggregates.Restaurant{}, nil
	}
	var restaurant aggregates.Restaurant
	if err := r.db.WithContext(ctx).
		Preload("Branches").
		Preload("Branches.Address").
		Preload("Branches.PriceRange").
		Preload("Branches.BranchCategoryRelations").
		Preload("Branches.BranchCategoryRelations.Category").
		Preload("Branches.Tables").
		Preload("Branches.Availables").
		Where("id = ?", id).Limit(1).Find(&restaurant).Error; err != nil {
		r.contextLogger.WithContextInfo(ctx, enums.ContextKeyTraceID).Error("failed to read restaurant full info", err)
		return aggregates.Restaurant{}, err
	}
	return restaurant, nil
}

func (r *RestaurantsImpl) UpdateRestaurant(ctx context.Context, updateRestaurant entities.UpdateRestaurant) error {
	if updateRestaurant.ID == "" {
		return nil
	}
	if err := r.db.WithContext(ctx).Model(&entities.Restaurant{}).Where("id = ?", updateRestaurant.ID).Updates(updateRestaurant.ToUpdateMap()).Error; err != nil {
		r.contextLogger.WithContextInfo(ctx, enums.ContextKeyTraceID).Error("failed to update restaurant", err)
		return err
	}
	return nil
}

func (r *RestaurantsImpl) DeleteRestaurants(ctx context.Context, ids []string) error {
	if len(ids) == 0 {
		return nil
	}
	if err := r.db.WithContext(ctx).Where("id IN (?)", ids).Delete(&entities.Restaurant{}).Error; err != nil {
		r.contextLogger.WithContextInfo(ctx, enums.ContextKeyTraceID).Error("failed to delete restaurants", err)
		return err
	}
	return nil
}

// Branches Impl
func NewBranchesRepository(
	db *gorm.DB,
	contextLogger contextloggers.ContextLogger,
) repositories.Branches {
	return &BranchesImpl{
		db:            db,
		contextLogger: contextLogger,
	}
}

type BranchesImpl struct {
	db            *gorm.DB
	contextLogger contextloggers.ContextLogger
}

func (r *BranchesImpl) CreateBranches(ctx context.Context, branches entities.Branches) error {
	if len(branches) == 0 {
		return nil
	}
	if err := r.db.WithContext(ctx).Create(&branches).Error; err != nil {
		r.contextLogger.WithContextInfo(ctx, enums.ContextKeyTraceID).Error("failed to create branches", err)
		return err
	}
	return nil
}

func (r *BranchesImpl) UpdateBranch(ctx context.Context, updateBranch entities.UpdateBranch) error {
	if updateBranch.ID == "" {
		return nil
	}
	if err := r.db.WithContext(ctx).Model(&entities.Branch{}).Where("id = ?", updateBranch.ID).Updates(updateBranch.ToUpdateMap()).Error; err != nil {
		r.contextLogger.WithContextInfo(ctx, enums.ContextKeyTraceID).Error("failed to update branch", err)
		return err
	}
	return nil
}

func (r *BranchesImpl) DeleteBranches(ctx context.Context, ids []string) error {
	if len(ids) == 0 {
		return nil
	}
	if err := r.db.WithContext(ctx).Where("id IN (?)", ids).Delete(&entities.Branch{}).Error; err != nil {
		r.contextLogger.WithContextInfo(ctx, enums.ContextKeyTraceID).Error("failed to delete branches", err)
		return err
	}
	return nil
}

// Addresses Impl
func NewAddressesRepository(
	db *gorm.DB,
	contextLogger contextloggers.ContextLogger,
) repositories.Addresses {
	return &AddressesImpl{
		db:            db,
		contextLogger: contextLogger,
	}
}

type AddressesImpl struct {
	db            *gorm.DB
	contextLogger contextloggers.ContextLogger
}

func (r *AddressesImpl) CreateAddresses(ctx context.Context, addresses entities.Addresses) error {
	if len(addresses) == 0 {
		return nil
	}
	if err := r.db.WithContext(ctx).Create(&addresses).Error; err != nil {
		r.contextLogger.WithContextInfo(ctx, enums.ContextKeyTraceID).Error("failed to create addresses", err)
		return err
	}
	return nil
}

func (r *AddressesImpl) UpdateAddress(ctx context.Context, updateAddress entities.UpdateAddress) error {
	if updateAddress.ID == "" {
		return nil
	}
	if err := r.db.WithContext(ctx).Model(&entities.Address{}).Where("id = ?", updateAddress.ID).Updates(updateAddress.ToUpdateMap()).Error; err != nil {
		r.contextLogger.WithContextInfo(ctx, enums.ContextKeyTraceID).Error("failed to update address", err)
		return err
	}
	return nil
}

func (r *AddressesImpl) DeleteAddresses(ctx context.Context, ids []string) error {
	if len(ids) == 0 {
		return nil
	}
	if err := r.db.WithContext(ctx).Where("id IN (?)", ids).Delete(&entities.Address{}).Error; err != nil {
		r.contextLogger.WithContextInfo(ctx, enums.ContextKeyTraceID).Error("failed to delete addresses", err)
		return err
	}
	return nil
}

// PriceRanges Impl
func NewPriceRangesRepository(
	db *gorm.DB,
	contextLogger contextloggers.ContextLogger,
) repositories.PriceRanges {
	return &PriceRangesImpl{
		db:            db,
		contextLogger: contextLogger,
	}
}

type PriceRangesImpl struct {
	db            *gorm.DB
	contextLogger contextloggers.ContextLogger
}

func (r *PriceRangesImpl) CreatePriceRanges(ctx context.Context, priceRanges entities.PriceRanges) error {
	if len(priceRanges) == 0 {
		return nil
	}
	if err := r.db.WithContext(ctx).Create(&priceRanges).Error; err != nil {
		r.contextLogger.WithContextInfo(ctx, enums.ContextKeyTraceID).Error("failed to create price ranges", err)
		return err
	}
	return nil
}

func (r *PriceRangesImpl) UpdatePriceRange(ctx context.Context, updatePriceRange entities.UpdatePriceRange) error {
	if updatePriceRange.ID == "" {
		return nil
	}
	if err := r.db.WithContext(ctx).Model(&entities.PriceRange{}).Where("id = ?", updatePriceRange.ID).Updates(updatePriceRange.ToUpdateMap()).Error; err != nil {
		r.contextLogger.WithContextInfo(ctx, enums.ContextKeyTraceID).Error("failed to update price range", err)
		return err
	}
	return nil
}

func (r *PriceRangesImpl) DeletePriceRanges(ctx context.Context, ids []string) error {
	if len(ids) == 0 {
		return nil
	}
	if err := r.db.WithContext(ctx).Where("id IN (?)", ids).Delete(&entities.PriceRange{}).Error; err != nil {
		r.contextLogger.WithContextInfo(ctx, enums.ContextKeyTraceID).Error("failed to delete price ranges", err)
		return err
	}
	return nil
}

// BranchCategoryRelations Impl
func NewBranchCategoryRelationsRepository(
	db *gorm.DB,
	contextLogger contextloggers.ContextLogger,
) repositories.BranchCategoryRelations {
	return &BranchCategoryRelationsImpl{
		db:            db,
		contextLogger: contextLogger,
	}
}

type BranchCategoryRelationsImpl struct {
	db            *gorm.DB
	contextLogger contextloggers.ContextLogger
}

func (r *BranchCategoryRelationsImpl) CreateBranchCategoryRelations(ctx context.Context, branchCategoryRelations entities.BranchCategoryRelations) error {
	if len(branchCategoryRelations) == 0 {
		return nil
	}
	if err := r.db.WithContext(ctx).Create(&branchCategoryRelations).Error; err != nil {
		r.contextLogger.WithContextInfo(ctx, enums.ContextKeyTraceID).Error("failed to create branch category relations", err)
		return err
	}
	return nil
}

func (r *BranchCategoryRelationsImpl) UpdateBranchCategoryRelation(ctx context.Context, updateBranchCategoryRelation entities.UpdateBranchCategoryRelation) error {
	if updateBranchCategoryRelation.ID == "" {
		return nil
	}
	if err := r.db.WithContext(ctx).Model(&entities.BranchCategoryRelation{}).Where("id = ?", updateBranchCategoryRelation.ID).Updates(updateBranchCategoryRelation.ToUpdateMap()).Error; err != nil {
		r.contextLogger.WithContextInfo(ctx, enums.ContextKeyTraceID).Error("failed to update branch category relation", err)
		return err
	}
	return nil
}

func (r *BranchCategoryRelationsImpl) DeleteBranchCategoryRelations(ctx context.Context, ids []string) error {
	if len(ids) == 0 {
		return nil
	}
	if err := r.db.WithContext(ctx).Where("id IN (?)", ids).Delete(&entities.BranchCategoryRelation{}).Error; err != nil {
		r.contextLogger.WithContextInfo(ctx, enums.ContextKeyTraceID).Error("failed to delete branch category relations", err)
		return err
	}
	return nil
}

// Tables Impl
func NewTablesRepository(
	db *gorm.DB,
	contextLogger contextloggers.ContextLogger,
) repositories.Tables {
	return &TablesImpl{
		db:            db,
		contextLogger: contextLogger,
	}
}

type TablesImpl struct {
	db            *gorm.DB
	contextLogger contextloggers.ContextLogger
}

func (r *TablesImpl) CreateTables(ctx context.Context, tables entities.Tables) error {
	if len(tables) == 0 {
		return nil
	}
	if err := r.db.WithContext(ctx).Create(&tables).Error; err != nil {
		r.contextLogger.WithContextInfo(ctx, enums.ContextKeyTraceID).Error("failed to create tables", err)
		return err
	}
	return nil
}

func (r *TablesImpl) UpdateTable(ctx context.Context, updateTable entities.UpdateTable) error {
	if updateTable.ID == "" {
		return nil
	}
	if err := r.db.WithContext(ctx).Model(&entities.Table{}).Where("id = ?", updateTable.ID).Updates(updateTable.ToUpdateMap()).Error; err != nil {
		r.contextLogger.WithContextInfo(ctx, enums.ContextKeyTraceID).Error("failed to update table", err)
		return err
	}
	return nil
}

func (r *TablesImpl) DeleteTables(ctx context.Context, ids []string) error {
	if len(ids) == 0 {
		return nil
	}
	if err := r.db.WithContext(ctx).Where("id IN (?)", ids).Delete(&entities.Table{}).Error; err != nil {
		r.contextLogger.WithContextInfo(ctx, enums.ContextKeyTraceID).Error("failed to delete tables", err)
		return err
	}
	return nil
}

// Availables Impl
func NewAvailablesRepository(
	db *gorm.DB,
	contextLogger contextloggers.ContextLogger,
) repositories.Availables {
	return &AvailablesImpl{
		db:            db,
		contextLogger: contextLogger,
	}
}

type AvailablesImpl struct {
	db            *gorm.DB
	contextLogger contextloggers.ContextLogger
}

func (r *AvailablesImpl) CreateAvailables(ctx context.Context, availables entities.Availables) error {
	if len(availables) == 0 {
		return nil
	}
	if err := r.db.WithContext(ctx).Create(&availables).Error; err != nil {
		r.contextLogger.WithContextInfo(ctx, enums.ContextKeyTraceID).Error("failed to create availables", err)
		return err
	}
	return nil
}

func (r *AvailablesImpl) UpdateAvailable(ctx context.Context, updateAvailable entities.UpdateAvailable) error {
	if updateAvailable.ID == "" {
		return nil
	}
	if err := r.db.WithContext(ctx).Model(&entities.Available{}).Where("id = ?", updateAvailable.ID).Updates(updateAvailable.ToUpdateMap()).Error; err != nil {
		r.contextLogger.WithContextInfo(ctx, enums.ContextKeyTraceID).Error("failed to update available", err)
		return err
	}
	return nil
}

func (r *AvailablesImpl) DeleteAvailables(ctx context.Context, ids []string) error {
	if len(ids) == 0 {
		return nil
	}
	if err := r.db.WithContext(ctx).Where("id IN (?)", ids).Delete(&entities.Available{}).Error; err != nil {
		r.contextLogger.WithContextInfo(ctx, enums.ContextKeyTraceID).Error("failed to delete availables", err)
		return err
	}
	return nil
}
