package storage

import "time"

const (
	DefaulLimit  = 10
	DefaulOffset = 0
)

type SongFilter struct {
	ReleaseDate *time.Time
	Lyrics      string
	Link        string
	GroupName   string
	Limit       int
	Offset      int
}

type SongFilterBuilder struct {
	filter *SongFilter
}

func NewSongFilter() *SongFilterBuilder {
	return &SongFilterBuilder{
		filter: &SongFilter{
			Limit:  DefaulLimit,
			Offset: DefaulOffset,
		},
	}
}

func (b *SongFilterBuilder) SetReleaseDate(date time.Time) *SongFilterBuilder {
	b.filter.ReleaseDate = &date
	return b
}

func (b *SongFilterBuilder) SetLyrics(lyrics string) *SongFilterBuilder {
	b.filter.Lyrics = lyrics
	return b
}

func (b *SongFilterBuilder) SetLink(link string) *SongFilterBuilder {
	b.filter.Link = link
	return b
}

func (b *SongFilterBuilder) SetGroupName(groupName string) *SongFilterBuilder {
	b.filter.GroupName = groupName
	return b
}

func (b *SongFilterBuilder) SetLimit(limit int) *SongFilterBuilder {
	b.filter.Limit = limit
	return b
}

func (b *SongFilterBuilder) SetOffset(offset int) *SongFilterBuilder {
	b.filter.Offset = offset
	return b
}

func (b *SongFilterBuilder) Build() *SongFilter {
	return b.filter
}
