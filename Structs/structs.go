package structs

type User struct {
	Username    string `JSON:"username"`
	FirstName   string `JSON:"first_name"`
	LastName    string `JSON:"last_name"`
	Email       string `JSON:"email"`
	Password    string `JSON:"password"`
	DateOfBirth string `JSON:"date_of_birth"`
	AboutMe     string `JSON:"about_me"`
	HaveImage   bool   `JSON:"have_image"`
}
