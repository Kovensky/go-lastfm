package lastfm

import (
	"testing"
)

func TestGetTrackTopTags(T *testing.T) {
	T.Parallel()
	lfm := testMock(New("4c563adf68bc357a4570d3e7986f6481"))
	topTags, err := lfm.GetTrackTopTags(
		Track{MBID: "48fa1cab-5250-4767-bbdf-14e0ef563d11"}, false)

	if testExpect(T, "error", nil, err) {
		testExpect(T, "track name", "One More Time", topTags.Track)
		testExpect(T, "top tag", "electronic", topTags.Tags[0].Name)
		testExpect(T, "top tag count", 100, topTags.Tags[0].Count)
	}
}

func TestGetArtistTopTags(T *testing.T) {
	T.Parallel()
	lfm := testMock(New("4c563adf68bc357a4570d3e7986f6481"))
	topTags, err := lfm.GetArtistTopTags(
		Artist{Name: "Daft Punk"}, false)

	if testExpect(T, "error", nil, err) {
		testExpect(T, "track name", "", topTags.Track)
		testExpect(T, "top tag", "electronic", topTags.Tags[0].Name)
		testExpect(T, "top tag count", 100, topTags.Tags[0].Count)
	}
}
