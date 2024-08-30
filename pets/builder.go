package pets

import "errors"

type PetInterface interface {
	Build() (*Pet, error)
	SetSpecies(s string) *Pet
	SetBreed(b string) *Pet
	SetMinWeight(mw int) *Pet
	SetMaxWeight(mw int) *Pet
	SetWeight(w int) *Pet
	SetDescription(d string) *Pet
	SetLifespan(l string) *Pet
	SetGeographicOrigin(g string) *Pet
	SetColor(c string) *Pet
	SetAge(a int) *Pet
	SetAgeEstimated(ae bool) *Pet
}

func NewPetBuilder() PetInterface {
	return &Pet{}
}

func (p *Pet) SetSpecies(s string) *Pet {
	p.Species = s
	return p
}

func (p *Pet) SetBreed(b string) *Pet {
	p.Breed = b
	return p
}

func (p *Pet) SetMinWeight(mw int) *Pet {
	p.MinWeight = mw
	return p
}

func (p *Pet) SetMaxWeight(mw int) *Pet {
	p.MaxWeight = mw
	return p
}

func (p *Pet) SetWeight(w int) *Pet {
	p.Weight = w
	return p
}

func (p *Pet) SetDescription(d string) *Pet {
	p.Description = d
	return p
}

func (p *Pet) SetLifespan(l string) *Pet {
	p.Lifespan = l
	return p
}

func (p *Pet) SetGeographicOrigin(g string) *Pet {
	p.GeographicOrigin = g
	return p
}

func (p *Pet) SetColor(c string) *Pet {
	p.Color = c
	return p
}

func (p *Pet) SetAge(a int) *Pet {
	p.Age = a
	return p
}

func (p *Pet) SetAgeEstimated(ae bool) *Pet {
	p.AgeEstimated = ae
	return p
}

func (p *Pet) Build() (*Pet, error) {
	if p.MinWeight > p.MaxWeight {
		return nil, errors.New("Minimum weight cannot be greater than maximum weight")
	}

	p.AverageWeight = (p.MinWeight + p.MaxWeight) / 2

	return p, nil
}
