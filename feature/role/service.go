package role

import "clean-architect/repository"

type RoleService interface {
	AllRole() ([]RoleRes, error)
}

type roleService struct {
	roleRepo repository.RoleRepository
}

func NewRoleService(r repository.RoleRepository) *roleService {
	return &roleService{
		roleRepo: r,
	}
}

func (s *roleService) AllRole() ([]RoleRes, error) {

	roles, err := s.roleRepo.FindAll()
	if err != nil {
		return nil, err
	}

	// manual map dto
	var res []RoleRes = make([]RoleRes, 0)
	for _, role := range roles {
		item := RoleRes{
			Code: role.Code,
			Name: role.Name,
			Desc: role.Desc,
		}
		res = append(res, item)
	}

	return res, nil
}
