package database

type User struct {
	ID       string `db:"id"`
	FullName string `db:"full_name"`
	Gender   string `db:"gender"`
	Email    string `db:"email"`
	Password string `db:"password"`
	DateColumn
}
