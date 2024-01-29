package terraform

type TerraformEpisode struct{}

func (e *TerraformEpisode) SetupEpisode() error {
	return nil
}

func (e *TerraformEpisode) RunEpisode() error {
	return nil
}

func (e *TerraformEpisode) CleanupEpisode() error {
	return nil
}
