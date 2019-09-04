package dbprovider

func (mgr *manager) Delete(id uint) (bool, error) {
	// check if sensortype with id exists
	sensorType, err := mgr.GetById(id)
	if err != nil {
		return false, err
	}

	// delete sensor type
	err = mgr.GetDb().Unscoped().Delete(sensorType).Error
	if err != nil {
		return false, err
	}
	return true, nil
}
