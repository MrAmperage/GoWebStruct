package DataCore

import "github.com/jackc/pgtype"

type TemplatedList struct {
	Template pgtype.JSONB
	DataList []any
}
