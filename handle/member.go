package handle

import "github.com/sirwarj/go-api-gin/models"

type MemberData struct {
	data []models.Member
}

func NewMember() *MemberData {
	return &MemberData{}
}

func (m *MemberData) AllData() []models.Member {
	return m.data
}

func (m *MemberData) AddData(d models.Member) models.Member {
	m.data = append(m.data, d)
	return d
}
