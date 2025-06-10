package internal

type Items []Customer

type DataProvider struct {
	Items      Items            `json:"items"`
	Pagination PaginationParams `json:"pagination"`
}

// NewDataProvider creates a new DataProvider instance
func NewDataProvider(items Items, pagination PaginationParams) DataProvider {
	return DataProvider{items, pagination}
}

// GetItems returns the items in the DataProvider
func (dp *DataProvider) GetItems() Items {
	return dp.Items
}

// GetPagination returns the pagination parameters in the DataProvider
func (dp *DataProvider) GetPagination() PaginationParams {
	return dp.Pagination
}
