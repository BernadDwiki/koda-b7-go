package minitask4

import "fmt"

type Education struct {
	Name  string
	Major string
}

type User struct {
	Name        string
	Image       string
	Email       string
	Age         int
	PhoneNumber string
	IsMarried   bool
	Education   []Education
}

func DefaultUser() User {
	return User{
		Name:        "Dwiki",
		Image:       "kdsjgakdgakd",
		Email:       "dwiki.adicitra6@gmail.com",
		Age:         22,
		PhoneNumber: "081228669806",
		IsMarried:   false,
		Education: []Education{{
			Name:  "SMA Negeri 7 Yogyakarta",
			Major: "Science",
		}, {
			Name:  "Sanata Dharma University",
			Major: "Informatics",
		}},
	}
}

func (u User) Print() {
	fmt.Println("Profil pengguna:")
	fmt.Println("Nama       :", u.Name)
	fmt.Println("Image      :", u.Image)
	fmt.Println("Email      :", u.Email)
	fmt.Println("Umur       :", u.Age)
	fmt.Println("No. Telepon:", u.PhoneNumber)
	fmt.Println("Menikah    :", u.IsMarried)
	for i, edu := range u.Education {
		fmt.Printf("Pendidikan %d: %s (%s)\n", i+1, edu.Name, edu.Major)
	}
}
