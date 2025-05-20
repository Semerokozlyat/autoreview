package database

import "strconv"

// Company represents a company
type Company struct {
	ID      string
	Company string
	Contact string
	Country string
}

// List of companies
type Companies struct {
	Companies []Company
}

// Company data access
var Data Companies

func init() {

	Data = Companies{
		Companies: []Company{
			{
				ID:      "1",
				Company: "Amazon",
				Contact: "Jeff Bezos",
				Country: "United States",
			},
			{
				ID:      "2",
				Company: "Apple",
				Contact: "Tim Cook",
				Country: "United States",
			},
			{
				ID:      "3",
				Company: "Microsoft",
				Contact: "Satya Nadella",
				Country: "United States",
			},
		},
	}
}

func (c *Companies) GetByID(id string) Company {
	var result Company
	for _, i := range c.Companies {
		if i.ID == id {
			result = i
			break
		}
	}
	return result
}

func (c *Companies) Update(company Company) {
	result := []Company{}
	for _, i := range c.Companies {
		if i.ID == company.ID {
			i.Company = company.Company
			i.Contact = company.Contact
			i.Country = company.Country
		}
		result = append(result, i)
	}
	c.Companies = result
}

func (c *Companies) Add(company Company) {
	max := 0
	for _, i := range c.Companies {
		n, _ := strconv.Atoi(i.ID)
		if n > max {
			max = n
		}
	}
	max++
	id := strconv.Itoa(max)

	c.Companies = append(c.Companies, Company{
		ID:      id,
		Company: company.Company,
		Contact: company.Contact,
		Country: company.Country,
	})
}

func (c *Companies) Delete(id string) {
	result := []Company{}
	for _, i := range c.Companies {
		if i.ID != id {
			result = append(result, i)
		}
	}
	c.Companies = result
}
