package campaign

type Service interface {
	GetCampaigns(UserID int) ([]Campaign, error)
}
type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetCampaigns(UserID int) ([]Campaign, error) {
	if UserID != 0 {
		campaign, err := s.repository.FindByUserID(UserID)
		if err != nil {
			return campaign, err
		}
		return campaign, nil
	}

	campaign, err := s.repository.FindAll()
	if err != nil {
		return campaign, err
	}

	return campaign, nil
}
