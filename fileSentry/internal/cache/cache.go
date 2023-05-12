package cache

var mapper map[string]string

func Init() {
	mapper = map[string]string{}
}

func Set(key string, value string) {
	mapper[key] = value
}

func Get(key string) string {
	return mapper[key]
}
