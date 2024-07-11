package pets

import (
	"dpatterns/models"
	"errors"
	"fmt"
)

type AnimalInterface interface {
	Show() string
}

type DogFromFactory struct {
	Pet *models.Dog
}

func (dff *DogFromFactory) Show() string {
	return fmt.Sprintf("This animal is a %s.", dff.Pet.Breed.Breed)
}

type CatFromFactory struct {
	Pet *models.Cat
}

func (cff *CatFromFactory) Show() string {
	return fmt.Sprintf("This animal is a %s.", cff.Pet.Breed.Breed)
}

type PetFactoryInterface interface {
	newPet() AnimalInterface
}

type DogAbstractFactory struct{}

func (daf *DogAbstractFactory) newPet() AnimalInterface {
	return &DogFromFactory{
		Pet: &models.Dog{},
	}
}

type CatAbstractFactory struct{}

func (caf *CatAbstractFactory) newPet() AnimalInterface {
	return &CatFromFactory{
		Pet: &models.Cat{},
	}
}

func NewPetFromAbstractFactory(species string) (AnimalInterface, error) {
	switch species {
	case "dog":
		var daf DogAbstractFactory
		dog := daf.newPet()
		return dog, nil

	case "cat":
		var caf CatAbstractFactory
		cat := caf.newPet()
		return cat, nil

	default:
		return nil, errors.New("unknown species")
	}
}
