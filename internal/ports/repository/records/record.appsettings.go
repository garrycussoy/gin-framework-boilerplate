package records

type Appsettings struct {
	Id    string `db:"id"`
	Key   string `db:"key"`
	Value string `db:"value"`
}
