package episode

type Episode interface {
	SetupEpisode() error
	RunEpisode(proceed bool) error
	CleanupEpisode() error
}
