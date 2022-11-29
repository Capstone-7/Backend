package helpers

type ValidationErrors struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func MsgForTag(tag string) string {
	switch tag {
	case "required":
		return "Kolom ini harus diisi"
	case "email":
		return "Email tidak valid"
	case "min":
		return "Kolom ini setidaknya berisi [PARAM] karakter"
	case "max":
		return "Kolom ini maksimal berisi [PARAM] karakter"
	case "alphanum":
		return "Kolom ini harus berisi alphanumeric"
	case "containsany":
		return "Kolom ini harus berisi setidaknya satu karakter spesial, satu huruf kapital, satu huruf kecil, dan satu angka"
	case "alpha":
		return "Kolom ini harus berisi alphabetic"
	case "uppercase":
		return "Kolom ini setidaknya berisi 1 huruf kapital"
	case "lowercase":
		return "Kolom ini setidaknya berisi 1 huruf kapital"
	case "alphanumunicode":
		return "Kolom ini harus berisi alphanumeric and unicode"
	case "eqfield":
		return "Kolom ini harus sama dengan [PARAM]"
	case "len":
		return "Kolom ini harus berisi [PARAM] karakter"
	case "gte":
		return "Kolom ini harus lebih besar atau sama dengan [PARAM]"
	case "url":
		return "Kolom ini harus berisi URL valid"
	default:
		return "Kolom tidak valid " + tag
	}
}