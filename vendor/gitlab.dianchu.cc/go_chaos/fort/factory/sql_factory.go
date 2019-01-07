package factory

//SQLQuery 查询sql语句信息
type SQLInfo struct {
	Sql             string
	Args            []interface{}
	TableName       string
	PrimaryKeyValue interface{}
}

//SQLAtom 原子操作
type SQLAtom struct {
	Id      string
	CmdList []*SQLInfo
}

type SQLFactory struct {
	SchemaFactory
}

func NewSQLFactory() *SQLFactory {
	var f = new(SQLFactory)
	f.ModelSpec = make(map[string]*Schema)
	return f
}
