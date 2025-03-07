package repostgresespostgres

import (
	"context"
	"time"

	"github.com/Leon180/go-event-driven-microservices/internal/pkg/enums"
	customizegorm "github.com/Leon180/go-event-driven-microservices/internal/pkg/gorm"
	contextloggers "github.com/Leon180/go-event-driven-microservices/internal/pkg/utilities/context_loggers"
	"github.com/Leon180/go-event-driven-microservices/internal/services/restaurants/internal/restaurants/aggregates"
	"github.com/Leon180/go-event-driven-microservices/internal/services/restaurants/internal/restaurants/dtos"
	"github.com/Leon180/go-event-driven-microservices/internal/services/restaurants/internal/restaurants/entities"
	"github.com/Leon180/go-event-driven-microservices/internal/services/restaurants/internal/restaurants/repositories"
	"github.com/samber/lo"
	"gorm.io/gorm"
)

func NewSearchRestaurantsFullInfo(
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

func (impl *SearchRestaurantsFullInfoImpl) SearchRestaurantsFullInfo(ctx context.Context, search *dtos.SearchRestaurants) (aggregates.Restaurants, error) {
	if search == nil {
		return nil, nil
	}
	sql := impl.buildBaseQuery(ctx).
		Scopes(
			impl.applyNameFilter(search.NameFilter, search.NamePreciseSearch),
			impl.applyDescriptionFilter(search.DescriptionFilter),
			impl.applyCityFilter(search.CityFilter),
			impl.applyCountryFilter(search.CountryFilter),
			impl.applyPriceRangeFilter(search.MinPriceFilter, search.MaxPriceFilter),
			impl.applyCategoryFilter(search.CategoryFilter),
			impl.applyTableAvailabilityFilter(search.TableAvailableWeek, search.TableAvailableStartTime, search.TableAvailableEndTime),
			impl.applyPagination(search.Pagination),
			impl.applyOrdering(search.OrderBy),
		)
	var restaurants []aggregates.Restaurant
	if err := sql.Find(&restaurants).Error; err != nil {
		impl.contextLogger.WithContextInfo(ctx, enums.ContextKeyTraceID).Error("failed to search restaurants full info", err)
		return nil, err
	}
	return restaurants, nil
}

func (impl *SearchRestaurantsFullInfoImpl) buildBaseQuery(ctx context.Context) *gorm.DB {
	return impl.db.WithContext(ctx).
		Preload("Branches").
		Preload("Branches.Address").
		Preload("Branches.PriceRange").
		Preload("Branches.BranchCategoryRelations").
		Preload("Branches.BranchCategoryRelations.Category").
		Preload("Branches.Tables").
		Preload("Branches.Availables")
}

func (impl *SearchRestaurantsFullInfoImpl) applyNameFilter(name *string, preciseSearch bool) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if name != nil {
			if preciseSearch {
				return db.Where("name = ?", *name)
			}
			return db.Where("name LIKE ?", "%"+*name+"%")
		}
		return db
	}
}

func (impl *SearchRestaurantsFullInfoImpl) applyDescriptionFilter(desc *string) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if desc != nil {
			return db.Where("description LIKE ?", "%"+*desc+"%")
		}
		return db
	}
}

func (impl *SearchRestaurantsFullInfoImpl) applyCityFilter(cities []enums.City) func(*gorm.DB) *gorm.DB {
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

func (impl *SearchRestaurantsFullInfoImpl) applyCountryFilter(countries []enums.Country) func(*gorm.DB) *gorm.DB {
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

func (impl *SearchRestaurantsFullInfoImpl) applyPriceRangeFilter(min, max *int) func(*gorm.DB) *gorm.DB {
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

func (impl *SearchRestaurantsFullInfoImpl) applyCategoryFilter(categories []enums.Category) func(*gorm.DB) *gorm.DB {
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

func (impl *SearchRestaurantsFullInfoImpl) applyTableAvailabilityFilter(availableWeek []time.Weekday, availableStartTime *string, availableEndTime *string) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if len(availableWeek) == 0 && availableStartTime == nil && availableEndTime == nil {
			return db
		}

		query := db.Where(`EXISTS (
            SELECT 1 FROM branch
            JOIN available ON branch.id = available.branch_id
            WHERE branch.restaurant_id = restaurant.id`)

		if len(availableWeek) > 0 {
			query = query.Where("available.weekday IN (?)", availableWeek)
		}
		if availableStartTime != nil {
			query = query.Where("available.start_time >= ?", *availableStartTime)
		}
		if availableEndTime != nil {
			query = query.Where("available.end_time <= ?", *availableEndTime)
		}

		return query.Where(")")
	}
}

func (impl *SearchRestaurantsFullInfoImpl) applyPagination(pagination *customizegorm.Pagination) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if pagination != nil {
			return db.Limit(pagination.PageSize).
				Offset((pagination.Page - 1) * pagination.PageSize)
		}
		return db
	}
}

func (impl *SearchRestaurantsFullInfoImpl) applyOrdering(orderBy []customizegorm.OrderBy) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		for _, ob := range orderBy {
			db = db.Order(ob.ToSort())
		}
		return db
	}
}

func NewReadRestaurants(
	db *gorm.DB,
	contextLogger contextloggers.ContextLogger,
) repositories.ReadRestaurants {
	return &ReadRestaurantsImpl{
		db:            db,
		contextLogger: contextLogger,
	}
}

type ReadRestaurantsImpl struct {
	db            *gorm.DB
	contextLogger contextloggers.ContextLogger
}

func (impl *ReadRestaurantsImpl) ReadRestaurantFullInfo(ctx context.Context, id string) (*aggregates.Restaurant, error) {
	if id == "" {
		return nil, nil
	}
	var restaurant aggregates.Restaurant
	if err := impl.db.WithContext(ctx).
		Preload("Branches").
		Preload("Branches.Address").
		Preload("Branches.PriceRange").
		Preload("Branches.BranchCategoryRelations").
		Preload("Branches.BranchCategoryRelations.Category").
		Preload("Branches.Tables").
		Preload("Branches.Availables").
		Where("id = ?", id).Limit(1).Find(&restaurant).Error; err != nil {
		impl.contextLogger.WithContextInfo(ctx, enums.ContextKeyTraceID).Error("failed to read restaurant full info", err)
		return nil, err
	}
	return &restaurant, nil
}

func (impl *ReadRestaurantsImpl) ReadRestaurant(ctx context.Context, id string) (*entities.Restaurant, error) {
	if id == "" {
		return nil, nil
	}
	var restaurant entities.Restaurant
	if err := impl.db.WithContext(ctx).Where("id = ?", id).Limit(1).Find(&restaurant).Error; err != nil {
		impl.contextLogger.WithContextInfo(ctx, enums.ContextKeyTraceID).Error("failed to read restaurant", err)
		return nil, err
	}
	return &restaurant, nil
}

// Update Restaurants Impl
func NewUpdateRestaurants(
	db *gorm.DB,
	contextLogger contextloggers.ContextLogger,
) repositories.UpdateRestaurants {
	return &UpdateRestaurantsImpl{
		db:            db,
		contextLogger: contextLogger,
	}
}

type UpdateRestaurantsImpl struct {
	db            *gorm.DB
	contextLogger contextloggers.ContextLogger
}

func (impl *UpdateRestaurantsImpl) CreateRestaurants(ctx context.Context, restaurants entities.Restaurants) error {
	if len(restaurants) == 0 {
		return nil
	}
	if err := impl.db.WithContext(ctx).Create(&restaurants).Error; err != nil {
		impl.contextLogger.WithContextInfo(ctx, enums.ContextKeyTraceID).Error("failed to create restaurants", err)
		return err
	}
	return nil
}

func (impl *UpdateRestaurantsImpl) UpdateRestaurant(ctx context.Context, updateRestaurant *entities.UpdateRestaurant) error {
	if updateRestaurant == nil || updateRestaurant.ID == "" {
		return nil
	}
	if err := impl.db.WithContext(ctx).Model(&entities.Restaurant{}).Where("id = ?", updateRestaurant.ID).Updates(updateRestaurant.ToUpdateMap()).Error; err != nil {
		impl.contextLogger.WithContextInfo(ctx, enums.ContextKeyTraceID).Error("failed to update restaurant", err)
		return err
	}
	return nil
}

func (impl *UpdateRestaurantsImpl) DeleteRestaurants(ctx context.Context, ids []string) error {
	if len(ids) == 0 {
		return nil
	}
	if err := impl.db.WithContext(ctx).Where("id IN (?)", ids).Delete(&entities.Restaurant{}).Error; err != nil {
		impl.contextLogger.WithContextInfo(ctx, enums.ContextKeyTraceID).Error("failed to delete restaurants", err)
		return err
	}
	return nil
}

func (impl *UpdateRestaurantsImpl) CreateBranches(ctx context.Context, branches entities.Branches) error {
	if len(branches) == 0 {
		return nil
	}
	if err := impl.db.WithContext(ctx).Create(&branches).Error; err != nil {
		impl.contextLogger.WithContextInfo(ctx, enums.ContextKeyTraceID).Error("failed to create branches", err)
		return err
	}
	return nil
}

func (impl *UpdateRestaurantsImpl) UpdateBranch(ctx context.Context, updateBranch *entities.UpdateBranch) error {
	if updateBranch == nil || updateBranch.ID == "" {
		return nil
	}
	if err := impl.db.WithContext(ctx).Model(&entities.Branch{}).Where("id = ?", updateBranch.ID).Updates(updateBranch.ToUpdateMap()).Error; err != nil {
		impl.contextLogger.WithContextInfo(ctx, enums.ContextKeyTraceID).Error("failed to update branch", err)
		return err
	}
	return nil
}

func (impl *UpdateRestaurantsImpl) DeleteBranches(ctx context.Context, ids []string) error {
	if len(ids) == 0 {
		return nil
	}
	if err := impl.db.WithContext(ctx).Where("id IN (?)", ids).Delete(&entities.Branch{}).Error; err != nil {
		impl.contextLogger.WithContextInfo(ctx, enums.ContextKeyTraceID).Error("failed to delete branches", err)
		return err
	}
	return nil
}

func (impl *UpdateRestaurantsImpl) CreateAddresses(ctx context.Context, addresses entities.Addresses) error {
	if len(addresses) == 0 {
		return nil
	}
	if err := impl.db.WithContext(ctx).Create(&addresses).Error; err != nil {
		impl.contextLogger.WithContextInfo(ctx, enums.ContextKeyTraceID).Error("failed to create addresses", err)
		return err
	}
	return nil
}

func (impl *UpdateRestaurantsImpl) UpdateAddress(ctx context.Context, updateAddress *entities.UpdateAddress) error {
	if updateAddress == nil || updateAddress.ID == "" {
		return nil
	}
	if err := impl.db.WithContext(ctx).Model(&entities.Address{}).Where("id = ?", updateAddress.ID).Updates(updateAddress.ToUpdateMap()).Error; err != nil {
		impl.contextLogger.WithContextInfo(ctx, enums.ContextKeyTraceID).Error("failed to update address", err)
		return err
	}
	return nil
}

func (impl *UpdateRestaurantsImpl) DeleteAddresses(ctx context.Context, ids []string) error {
	if len(ids) == 0 {
		return nil
	}
	if err := impl.db.WithContext(ctx).Where("id IN (?)", ids).Delete(&entities.Address{}).Error; err != nil {
		impl.contextLogger.WithContextInfo(ctx, enums.ContextKeyTraceID).Error("failed to delete addresses", err)
		return err
	}
	return nil
}

func (impl *UpdateRestaurantsImpl) CreatePriceRanges(ctx context.Context, priceRanges entities.PriceRanges) error {
	if len(priceRanges) == 0 {
		return nil
	}
	if err := impl.db.WithContext(ctx).Create(&priceRanges).Error; err != nil {
		impl.contextLogger.WithContextInfo(ctx, enums.ContextKeyTraceID).Error("failed to create price ranges", err)
		return err
	}
	return nil
}

func (impl *UpdateRestaurantsImpl) UpdatePriceRange(ctx context.Context, updatePriceRange *entities.UpdatePriceRange) error {
	if updatePriceRange == nil || updatePriceRange.ID == "" {
		return nil
	}
	if err := impl.db.WithContext(ctx).Model(&entities.PriceRange{}).Where("id = ?", updatePriceRange.ID).Updates(updatePriceRange.ToUpdateMap()).Error; err != nil {
		impl.contextLogger.WithContextInfo(ctx, enums.ContextKeyTraceID).Error("failed to update price range", err)
		return err
	}
	return nil
}

func (impl *UpdateRestaurantsImpl) DeletePriceRanges(ctx context.Context, ids []string) error {
	if len(ids) == 0 {
		return nil
	}
	if err := impl.db.WithContext(ctx).Where("id IN (?)", ids).Delete(&entities.PriceRange{}).Error; err != nil {
		impl.contextLogger.WithContextInfo(ctx, enums.ContextKeyTraceID).Error("failed to delete price ranges", err)
		return err
	}
	return nil
}

func (impl *UpdateRestaurantsImpl) CreateBranchCategoryRelations(ctx context.Context, branchCategoryRelations entities.BranchCategoryRelations) error {
	if len(branchCategoryRelations) == 0 {
		return nil
	}
	if err := impl.db.WithContext(ctx).Create(&branchCategoryRelations).Error; err != nil {
		impl.contextLogger.WithContextInfo(ctx, enums.ContextKeyTraceID).Error("failed to create branch category relations", err)
		return err
	}
	return nil
}

func (impl *UpdateRestaurantsImpl) UpdateBranchCategoryRelation(ctx context.Context, updateBranchCategoryRelation *entities.UpdateBranchCategoryRelation) error {
	if updateBranchCategoryRelation == nil || updateBranchCategoryRelation.ID == "" {
		return nil
	}
	if err := impl.db.WithContext(ctx).Model(&entities.BranchCategoryRelation{}).Where("id = ?", updateBranchCategoryRelation.ID).Updates(updateBranchCategoryRelation.ToUpdateMap()).Error; err != nil {
		impl.contextLogger.WithContextInfo(ctx, enums.ContextKeyTraceID).Error("failed to update branch category relation", err)
		return err
	}
	return nil
}

func (impl *UpdateRestaurantsImpl) DeleteBranchCategoryRelations(ctx context.Context, ids []string) error {
	if len(ids) == 0 {
		return nil
	}
	if err := impl.db.WithContext(ctx).Where("id IN (?)", ids).Delete(&entities.BranchCategoryRelation{}).Error; err != nil {
		impl.contextLogger.WithContextInfo(ctx, enums.ContextKeyTraceID).Error("failed to delete branch category relations", err)
		return err
	}
	return nil
}

func (impl *UpdateRestaurantsImpl) CreateTables(ctx context.Context, tables entities.Tables) error {
	if len(tables) == 0 {
		return nil
	}
	if err := impl.db.WithContext(ctx).Create(&tables).Error; err != nil {
		impl.contextLogger.WithContextInfo(ctx, enums.ContextKeyTraceID).Error("failed to create tables", err)
		return err
	}
	return nil
}

func (impl *UpdateRestaurantsImpl) UpdateTable(ctx context.Context, updateTable *entities.UpdateTable) error {
	if updateTable == nil || updateTable.ID == "" {
		return nil
	}
	if err := impl.db.WithContext(ctx).Model(&entities.Table{}).Where("id = ?", updateTable.ID).Updates(updateTable.ToUpdateMap()).Error; err != nil {
		impl.contextLogger.WithContextInfo(ctx, enums.ContextKeyTraceID).Error("failed to update table", err)
		return err
	}
	return nil
}

func (impl *UpdateRestaurantsImpl) DeleteTables(ctx context.Context, ids []string) error {
	if len(ids) == 0 {
		return nil
	}
	if err := impl.db.WithContext(ctx).Where("id IN (?)", ids).Delete(&entities.Table{}).Error; err != nil {
		impl.contextLogger.WithContextInfo(ctx, enums.ContextKeyTraceID).Error("failed to delete tables", err)
		return err
	}
	return nil
}

func (impl *UpdateRestaurantsImpl) CreateAvailables(ctx context.Context, availables entities.Availables) error {
	if len(availables) == 0 {
		return nil
	}
	if err := impl.db.WithContext(ctx).Create(&availables).Error; err != nil {
		impl.contextLogger.WithContextInfo(ctx, enums.ContextKeyTraceID).Error("failed to create availables", err)
		return err
	}
	return nil
}

func (impl *UpdateRestaurantsImpl) UpdateAvailable(ctx context.Context, updateAvailable *entities.UpdateAvailable) error {
	if updateAvailable == nil || updateAvailable.ID == "" {
		return nil
	}
	if err := impl.db.WithContext(ctx).Model(&entities.Available{}).Where("id = ?", updateAvailable.ID).Updates(updateAvailable.ToUpdateMap()).Error; err != nil {
		impl.contextLogger.WithContextInfo(ctx, enums.ContextKeyTraceID).Error("failed to update available", err)
		return err
	}
	return nil
}

func (impl *UpdateRestaurantsImpl) DeleteAvailables(ctx context.Context, ids []string) error {
	if len(ids) == 0 {
		return nil
	}
	if err := impl.db.WithContext(ctx).Where("id IN (?)", ids).Delete(&entities.Available{}).Error; err != nil {
		impl.contextLogger.WithContextInfo(ctx, enums.ContextKeyTraceID).Error("failed to delete availables", err)
		return err
	}
	return nil
}

func NewUpdateRestaurantsWithTransaction(
	db *gorm.DB,
	contextLogger contextloggers.ContextLogger,
) customizegorm.Transactor[repositories.UpdateRestaurantsWithTransaction] {
	return &UpdateRestaurantsWithTransactionImpl{
		TransactorImpl: TransactorImpl{
			db:            db,
			contextLogger: contextLogger,
		},
	}
}

type UpdateRestaurantsWithTransactionImpl struct {
	TransactorImpl
}

func (impl *UpdateRestaurantsWithTransactionImpl) BeginTx(ctx context.Context) (repositories.UpdateRestaurantsWithTransaction, error) {
	tx := impl.db.WithContext(ctx).Begin()
	return &UpdateRestaurantsTransactionImpl{
		TransactionImpl: TransactionImpl{
			Db:            tx,
			ContextLogger: impl.contextLogger,
		},
	}, nil
}

type UpdateRestaurantsTransactionImpl struct {
	TransactionImpl
}

func (impl *UpdateRestaurantsTransactionImpl) CreateRestaurants(ctx context.Context, restaurants entities.Restaurants) error {
	return NewUpdateRestaurants(impl.Db, impl.ContextLogger).CreateRestaurants(ctx, restaurants)
}

func (impl *UpdateRestaurantsTransactionImpl) UpdateRestaurant(ctx context.Context, updateRestaurant *entities.UpdateRestaurant) error {
	return NewUpdateRestaurants(impl.Db, impl.ContextLogger).UpdateRestaurant(ctx, updateRestaurant)
}

func (impl *UpdateRestaurantsTransactionImpl) DeleteRestaurants(ctx context.Context, ids []string) error {
	return NewUpdateRestaurants(impl.Db, impl.ContextLogger).DeleteRestaurants(ctx, ids)
}

func (impl *UpdateRestaurantsTransactionImpl) CreateBranches(ctx context.Context, branches entities.Branches) error {
	return NewUpdateRestaurants(impl.Db, impl.ContextLogger).CreateBranches(ctx, branches)
}

func (impl *UpdateRestaurantsTransactionImpl) UpdateBranch(ctx context.Context, updateBranch *entities.UpdateBranch) error {
	return NewUpdateRestaurants(impl.Db, impl.ContextLogger).UpdateBranch(ctx, updateBranch)
}

func (impl *UpdateRestaurantsTransactionImpl) DeleteBranches(ctx context.Context, ids []string) error {
	return NewUpdateRestaurants(impl.Db, impl.ContextLogger).DeleteBranches(ctx, ids)
}

func (impl *UpdateRestaurantsTransactionImpl) CreateAddresses(ctx context.Context, addresses entities.Addresses) error {
	return NewUpdateRestaurants(impl.Db, impl.ContextLogger).CreateAddresses(ctx, addresses)
}

func (impl *UpdateRestaurantsTransactionImpl) UpdateAddress(ctx context.Context, updateAddress *entities.UpdateAddress) error {
	return NewUpdateRestaurants(impl.Db, impl.ContextLogger).UpdateAddress(ctx, updateAddress)
}

func (impl *UpdateRestaurantsTransactionImpl) DeleteAddresses(ctx context.Context, ids []string) error {
	return NewUpdateRestaurants(impl.Db, impl.ContextLogger).DeleteAddresses(ctx, ids)
}

func (impl *UpdateRestaurantsTransactionImpl) CreatePriceRanges(ctx context.Context, priceRanges entities.PriceRanges) error {
	return NewUpdateRestaurants(impl.Db, impl.ContextLogger).CreatePriceRanges(ctx, priceRanges)
}

func (impl *UpdateRestaurantsTransactionImpl) UpdatePriceRange(ctx context.Context, updatePriceRange *entities.UpdatePriceRange) error {
	return NewUpdateRestaurants(impl.Db, impl.ContextLogger).UpdatePriceRange(ctx, updatePriceRange)
}

func (impl *UpdateRestaurantsTransactionImpl) DeletePriceRanges(ctx context.Context, ids []string) error {
	return NewUpdateRestaurants(impl.Db, impl.ContextLogger).DeletePriceRanges(ctx, ids)
}

func (impl *UpdateRestaurantsTransactionImpl) CreateBranchCategoryRelations(ctx context.Context, branchCategoryRelations entities.BranchCategoryRelations) error {
	return NewUpdateRestaurants(impl.Db, impl.ContextLogger).CreateBranchCategoryRelations(ctx, branchCategoryRelations)
}

func (impl *UpdateRestaurantsTransactionImpl) UpdateBranchCategoryRelation(ctx context.Context, updateBranchCategoryRelation *entities.UpdateBranchCategoryRelation) error {
	return NewUpdateRestaurants(impl.Db, impl.ContextLogger).UpdateBranchCategoryRelation(ctx, updateBranchCategoryRelation)
}

func (impl *UpdateRestaurantsTransactionImpl) DeleteBranchCategoryRelations(ctx context.Context, ids []string) error {
	return NewUpdateRestaurants(impl.Db, impl.ContextLogger).DeleteBranchCategoryRelations(ctx, ids)
}

func (impl *UpdateRestaurantsTransactionImpl) CreateTables(ctx context.Context, tables entities.Tables) error {
	return NewUpdateRestaurants(impl.Db, impl.ContextLogger).CreateTables(ctx, tables)
}

func (impl *UpdateRestaurantsTransactionImpl) UpdateTable(ctx context.Context, updateTable *entities.UpdateTable) error {
	return NewUpdateRestaurants(impl.Db, impl.ContextLogger).UpdateTable(ctx, updateTable)
}

func (impl *UpdateRestaurantsTransactionImpl) DeleteTables(ctx context.Context, ids []string) error {
	return NewUpdateRestaurants(impl.Db, impl.ContextLogger).DeleteTables(ctx, ids)
}

func (impl *UpdateRestaurantsTransactionImpl) CreateAvailables(ctx context.Context, availables entities.Availables) error {
	return NewUpdateRestaurants(impl.Db, impl.ContextLogger).CreateAvailables(ctx, availables)
}

func (impl *UpdateRestaurantsTransactionImpl) UpdateAvailable(ctx context.Context, updateAvailable *entities.UpdateAvailable) error {
	return NewUpdateRestaurants(impl.Db, impl.ContextLogger).UpdateAvailable(ctx, updateAvailable)
}

func (impl *UpdateRestaurantsTransactionImpl) DeleteAvailables(ctx context.Context, ids []string) error {
	return NewUpdateRestaurants(impl.Db, impl.ContextLogger).DeleteAvailables(ctx, ids)
}
