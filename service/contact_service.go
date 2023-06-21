package service

import (
	"sort"

	"github.com/handrixn/contacts-api/model"
	"github.com/handrixn/contacts-api/repository"
)

type ContactService interface {
	GetAllContacts(filterParams map[string]string, sortField string, sortOrder string, page int, pageSize int) ([]model.Contact, error)
	GetContactByID(id string) (*model.Contact, error)
	CreateContact(contact *model.Contact) error
	UpdateContact(contact *model.Contact) error
	DeleteContact(id string) error
}

type ContactServiceImpl struct {
	contactRepo repository.ContactRepository
}

func NewContactService(contactRepo repository.ContactRepository) *ContactServiceImpl {
	return &ContactServiceImpl{
		contactRepo: contactRepo,
	}
}

func (s *ContactServiceImpl) GetAllContacts(filterParams map[string]string, sortField string, sortOrder string, page int, pageSize int) ([]model.Contact, error) {
	contacts, err := s.contactRepo.GetAll()
	if err != nil {
		return nil, err
	}

	// Apply filters
	filteredContacts := filterContacts(contacts, filterParams)

	// Apply sorting
	sortedContacts := sortContacts(filteredContacts, sortField, sortOrder)

	// Apply pagination
	paginatedContacts := paginateContacts(sortedContacts, page, pageSize)

	return paginatedContacts, nil
}

func (s *ContactServiceImpl) GetContactByID(id string) (*model.Contact, error) {
	return s.contactRepo.GetByID(id)
}

func (s *ContactServiceImpl) CreateContact(contact *model.Contact) error {
	return s.contactRepo.Create(contact)
}

func (s *ContactServiceImpl) UpdateContact(contact *model.Contact) error {
	return s.contactRepo.Update(contact)
}

func (s *ContactServiceImpl) DeleteContact(id string) error {
	return s.contactRepo.Delete(id)
}

func filterContacts(contacts []model.Contact, filterParams map[string]string) []model.Contact {
	if len(filterParams) == 0 {
		return contacts
	}

	filteredContacts := make([]model.Contact, 0)
	for _, contact := range contacts {
		if filterMatch(contact, filterParams) {
			filteredContacts = append(filteredContacts, contact)
		}
	}

	return filteredContacts
}

func filterMatch(contact model.Contact, filterParams map[string]string) bool {
	for key, value := range filterParams {
		switch key {
		case "name":
			if contact.Name != value {
				return false
			}
		case "gender":
			if contact.Gender != value {
				return false
			}
		case "phone":
			if contact.Phone != value {
				return false
			}
		case "email":
			if contact.Email != value {
				return false
			}
		}
	}

	return true
}

func sortContacts(contacts []model.Contact, sortField string, sortOrder string) []model.Contact {
	if sortField == "" {
		return contacts
	}

	sort.Slice(contacts, func(i, j int) bool {
		switch sortField {
		case "name":
			if sortOrder == "desc" {
				return contacts[i].Name > contacts[j].Name
			}
			return contacts[i].Name < contacts[j].Name
		case "gender":
			if sortOrder == "desc" {
				return contacts[i].Gender > contacts[j].Gender
			}
			return contacts[i].Gender < contacts[j].Gender
		case "phone":
			if sortOrder == "desc" {
				return contacts[i].Phone > contacts[j].Phone
			}
			return contacts[i].Phone < contacts[j].Phone
		case "email":
			if sortOrder == "desc" {
				return contacts[i].Email > contacts[j].Email
			}
			return contacts[i].Email < contacts[j].Email
		default:
			return i < j
		}
	})

	return contacts
}

func paginateContacts(contacts []model.Contact, page int, pageSize int) []model.Contact {
	if page < 1 || pageSize < 1 {
		return contacts
	}

	startIndex := (page - 1) * pageSize
	endIndex := startIndex + pageSize

	if startIndex >= len(contacts) {
		return []model.Contact{}
	}

	if endIndex > len(contacts) {
		endIndex = len(contacts)
	}

	return contacts[startIndex:endIndex]
}
