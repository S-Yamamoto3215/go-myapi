package apperrors

type ErrCode string

const (
	Unknown ErrCode = "UOOO"

	InsertDataFailed ErrCode = "S001"
	GetDataFailed    ErrCode = "S002"
	NAData           ErrCode = "S003"
)
