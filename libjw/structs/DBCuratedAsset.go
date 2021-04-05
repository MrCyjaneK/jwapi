package structs

// DBCuratedAsset -
type DBCuratedAsset struct {
	ListType           int `json:"list_type"`
	SortOrder          int `json:"sort_order"`
	PublicationAssetID int `json:"publication_asset_id"`
}
