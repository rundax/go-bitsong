package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// query endpoints supported by the Album Querier
const (
	QueryParams = "params"
	QueryAlbums = "albums"
	QueryAlbum  = "album"

	QueryTracks = "tracks"
)

// Params for queries:
// - 'custom/album/album'
type QueryAlbumParams struct {
	AlbumID uint64
}

// creates a new instance of QueryAlbumParams
func NewQueryAlbumParams(albumID uint64) QueryAlbumParams {
	return QueryAlbumParams{
		AlbumID: albumID,
	}
}

// Params for query 'custom/album/albums'
type QueryAlbumsParams struct {
	Owner       sdk.AccAddress
	AlbumStatus AlbumStatus
	Limit       uint64
}

// creates a new instance of QueryAlbumsParams
func NewQueryAlbumsParams(owner sdk.AccAddress, status AlbumStatus, limit uint64) QueryAlbumsParams {
	return QueryAlbumsParams{
		Owner:       owner,
		AlbumStatus: status,
		Limit:       limit,
	}
}
