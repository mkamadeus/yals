package service

import "fmt"

func (s yalsService) SetURLToAlias(alias, url string, force bool) error {
	if !force {
		used, err := s.Repository.IsAliasUsed(alias)
		if err != nil {
			return err
		}
		if used {
			return fmt.Errorf("alias is already used")
		}
	}

	err := s.Repository.SetURLToAlias(alias, url)
	if err != nil {
		return err
	}
	return nil
}
