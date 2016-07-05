package lastfm

import (
	"testing"
)

// TODO: more coverage?

func TestGetRecentTracks(T *testing.T) {
	T.Parallel()
	lfm := testMock(New("4c563adf68bc357a4570d3e7986f6481"))
	tracks, err := lfm.GetRecentTracks("Kovensky", 1)

	if testExpect(T, "error", nil, err) {
		testExpect(T, "scrobble count", 39679, tracks.Total)
		testExpect(T, "now playing track", &tracks.Tracks[0], tracks.NowPlaying)
		testExpect(T, "first track's loved status to be", true, tracks.Tracks[0].Loved)
	}
}

func TestCompareTaste(T *testing.T) {
	T.Parallel()
	lfm := testMock(New("4c563adf68bc357a4570d3e7986f6481"))
	taste, err := lfm.CompareTaste("Kovensky", "D4RK-PH0ENIX")

	if testExpect(T, "error", nil, err) {
		testExpect(T, "artist count", 5, len(taste.Artists))
		testExpect(T, "top similar artist", "DIR EN GREY", taste.Artists[0])
	}
}

func TestGetUserNeighbours(T *testing.T) {
	T.Parallel()
	lfm := testMock(New("4c563adf68bc357a4570d3e7986f6481"))
	n, err := lfm.GetUserNeighbours("Kovensky", 1)

	if testExpect(T, "error", nil, err) {
		testExpect(T, "neighbour count", 1, len(n))
		testExpect(T, "neighbour", "AT_Field", n[0].Name)
	}
}

func TestGetUserTopArtists(T *testing.T) {
	T.Parallel()
	lfm := testMock(New("4c563adf68bc357a4570d3e7986f6481"))
	t, err := lfm.GetUserTopArtists("Kovensky", Overall, 1)

	if testExpect(T, "error", nil, err) {
		testExpect(T, "user", "Kovensky", t.User)
		testExpect(T, "period", Overall, t.Period)
		if testExpect(T, "artist count", 1, len(t.Artists)) {
			testExpect(T, "top artist", "CROW'SCLAW", t.Artists[0].Name)
		}
	}
}
