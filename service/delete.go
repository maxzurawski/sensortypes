package service

func (s *service) Delete(id uint) (bool, error) {
	result, err := s.Mgr.Delete(id)
	if err != nil {
		return result, err
	}
	return result, err
}
