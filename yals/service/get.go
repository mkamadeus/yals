package service

func (s yalsService) GetURLFromAlias(alias string) (string, error) {
	url, err := s.Repository.FetchURLFromAlias(alias)
	if err != nil {
		return "", err
	}
	return url, nil
}
