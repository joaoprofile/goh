package sqlg

type Manager struct {
	page   int
	limit  int
	offset int
}

// The Manager will be a helper to help generate SQL DML,
// with INSERT INTO, UPDATE, DELETE, SELECT, ORDER, GROUP, JOIN, Pagination...

// test only
// limit = 10
// offset = limit * (page - 1)
// select count(*) over() as total, columns from tableslimit $1 offset $2
