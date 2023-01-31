package core_crud

// Error message is return from gorm package
const (
	RecordNotFound = "record not found"
)

type UpdateOrInsert[E any] struct {
	NewData     *E
	ReplaceData *E
}
