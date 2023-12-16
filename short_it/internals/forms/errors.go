package forms

type errors map[string][]string

func (e errors) Add(field, msg string) {
	e[field] = append(e[field], msg)
}

func (e errors) Get(field string) string {
	es, ok := e[field]

	if !ok {
		return ""
	}

	return es[0]
}
