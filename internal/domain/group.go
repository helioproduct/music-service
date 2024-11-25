package domain

type Group struct {
	ID    int
	Name  string
	Songs []*Song
}
