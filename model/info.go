package model

func AddInfo(param *Info) (*Info, error) {

	err := DB.Create(param).Error
	return param, err

}

func ListInfo(param *Info) ([]*Info, error) {
	var infos = make([]*Info, 0)

	err := DB.Where("(from_user_id=? and to_user_id = ? ) or (from_user_id=? and to_user_id = ? )", param.FromUserId, param.ToUserId, param.ToUserId, param.FromUserId).Order("created_at desc").Offset((param.PageNum - 1) * param.PageSize).Limit(param.PageSize).Find(&infos).Error

	if err != nil {
		return nil, err
	}

	for _, v := range infos {
		if u, err := GetUserById(v.FromUserId); err == nil {
			v.FromUser = u
		}
		if u, err := GetUserById(v.ToUserId); err == nil {
			v.ToUser = u
		}
	}

	return infos, nil
}
