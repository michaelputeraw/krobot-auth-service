package constant

type FieldName map[string]string

var Fields = map[string]FieldName{
	"fullName": {
		LANG_DEFAULT: "full name",
		LANG_ID:      "nama lengkap",
	},
	"gender": {
		LANG_DEFAULT: "gender",
		LANG_ID:      "jenis kelamin",
	},
	"email": {
		LANG_DEFAULT: "email address",
		LANG_ID:      "alamat surel",
	},
	"password": {
		LANG_DEFAULT: "password",
		LANG_ID:      "kata sandi",
	},
}
