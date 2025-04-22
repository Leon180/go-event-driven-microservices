package featuresdtos

type GetRestaurantBranchRequest struct {
	BranchID string `json:"branch_id" binding:"required"`
}
