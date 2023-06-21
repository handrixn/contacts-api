package repository

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/handrixn/contacts-api/model"
)

type ContactRepository interface {
	GetAll() ([]model.Contact, error)
	GetByID(id string) (*model.Contact, error)
	Create(contact *model.Contact) error
	Update(contact *model.Contact) error
	Delete(id string) error
}

type FileContactRepository struct {
	dataFile string
}

func NewFileContactRepository(dataFile string) *FileContactRepository {
	return &FileContactRepository{
		dataFile: dataFile,
	}
}

func (r *FileContactRepository) GetAll() ([]model.Contact, error) {
	contacts := []model.Contact{}

	data, err := r.readFile()
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &contacts)
	if err != nil {
		return nil, err
	}

	return contacts, nil
}

func (r *FileContactRepository) GetByID(id string) (*model.Contact, error) {
	contacts, err := r.GetAll()
	if err != nil {
		return nil, err
	}

	for _, contact := range contacts {
		if contact.ID == id {
			return &contact, nil
		}
	}

	return nil, fmt.Errorf("contact not found")
}

func (r *FileContactRepository) Create(contact *model.Contact) error {
	contacts, err := r.GetAll()
	if err != nil {
		return err
	}

	contact.ID = uuid.New().String()
	contact.CreatedAt = time.Now()
	contact.UpdatedAt = time.Now()

	contacts = append(contacts, *contact)

	err = r.saveFile(contacts)
	if err != nil {
		return err
	}

	return nil
}

func (r *FileContactRepository) Update(contact *model.Contact) error {
	contacts, err := r.GetAll()
	if err != nil {
		return err
	}

	for i, c := range contacts {
		if c.ID == contact.ID {
			contact.CreatedAt = c.CreatedAt
			contact.UpdatedAt = time.Now()
			contacts[i] = *contact

			err = r.saveFile(contacts)
			if err != nil {
				fmt.Println(err)
				return err
			}

			return nil
		}
	}

	return fmt.Errorf("contact not found")
}

func (r *FileContactRepository) Delete(id string) error {
	contacts, err := r.GetAll()
	if err != nil {
		return err
	}

	for i, contact := range contacts {
		if contact.ID == id {
			contacts = append(contacts[:i], contacts[i+1:]...)

			err = r.saveFile(contacts)
			if err != nil {
				return err
			}

			return nil
		}
	}

	return fmt.Errorf("contact not found")
}

func (r *FileContactRepository) readFile() ([]byte, error) {
	file, err := os.Open(r.dataFile)
	if err != nil {
		if os.IsNotExist(err) {
			// Create an empty file if it doesn't exist
			emptyData, _ := json.Marshal([]model.Contact{})
			err = ioutil.WriteFile(r.dataFile, emptyData, 0644)
			if err != nil {
				return nil, fmt.Errorf("failed to create data file: %v", err)
			}

			return emptyData, nil
		}
		return nil, fmt.Errorf("failed to open data file: %v", err)
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("failed to read data file: %v", err)
	}

	return data, nil
}

func (r *FileContactRepository) saveFile(contacts []model.Contact) error {
	file, err := os.Create(r.dataFile)
	if err != nil {
		return fmt.Errorf("failed to create data file: %v", err)
	}
	defer file.Close()

	data, err := json.Marshal(contacts)
	if err != nil {
		return fmt.Errorf("failed to encode data file: %v", err)
	}

	_, err = file.Write(data)
	if err != nil {
		return fmt.Errorf("failed to write data file: %v", err)
	}

	return nil
}
