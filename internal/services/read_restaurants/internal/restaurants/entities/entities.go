package entities

type Entity interface {
	*struct{}
}

type UpdateEntity[M Entity] interface {
	RemoveUnchangedFields(entity M) UpdateEntity[M]
	ToUpdateMap() map[string]any
}
