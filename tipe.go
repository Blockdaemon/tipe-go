package tipe

type CreatedBy struct {
	Email     string `json:"email"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

// Fields
type BooleanField struct {
	Data  map[string]interface{} `json:"data"`
	ID    string                 `json:"id"`
	List  bool                   `json:"list"`
	Name  string                 `json:"name"`
	Type  string                 `json:"type"`
	Value bool                   `json:"value"`
}

type ButtonField struct {
	Data  map[string]interface{} `json:"data"`
	ID    string                 `json:"id"`
	List  bool                   `json:"list"`
	Name  string                 `json:"name"`
	Type  string                 `json:"type"`
	Value string                 `json:"value"`
}

type HTMLField struct {
	Data  map[string]interface{} `json:"data"`
	ID    string                 `json:"id"`
	List  bool                   `json:"list"`
	Name  string                 `json:"name"`
	Type  string                 `json:"type"`
	Value interface{}            `json:"value"`
}

func (f HTMLField) String() string {
	value, ok := f.Value.(string)
	if !ok {
		return ""
	}

	return value
}

func (f HTMLField) StringSlice() []string {
	if f.List {
		values, ok := f.Value.([]interface{})
		if !ok {
			return nil
		}

		strs := []string{}
		for _, v := range values {
			str, ok := v.(string)
			if !ok {
				return nil
			}

			strs = append(strs, str)
		}

		return strs
	}

	return nil
}

type TextField struct {
	Data  map[string]interface{} `json:"data"`
	ID    string                 `json:"id"`
	List  bool                   `json:"list"`
	Name  string                 `json:"name"`
	Type  string                 `json:"type"`
	Value string                 `json:"value"`
}

type Template struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
