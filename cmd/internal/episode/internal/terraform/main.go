package terraform

type TerraformEpisode struct{}

func (e *TerraformEpisode) EpisodeSetup() error {
	return nil
}

func (e *TerraformEpisode) EpisodeRun() error {
	return nil
}

func (e *TerraformEpisode) EpisodeCleanup() error {
	return nil
}
