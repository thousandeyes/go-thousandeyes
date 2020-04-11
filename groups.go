package thousandeyes

// AccountGroups - list of account groups
type AccountGroups []AccountGroup

// AccountGroup - an account group
type AccountGroup struct {
	Aid  int    `json:"aid"`
	Name string `json:"name"`
}

// GroupLabels - list of group labels
type GroupLabels []GroupLabel

// GroupLabel - group label
type GroupLabel struct {
	GroupName string
	GroupID   int
	BuiltIn   int
}
