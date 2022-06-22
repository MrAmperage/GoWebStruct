package DataCore

import "github.com/jackc/pgtype"

type TemplatedList[Data any] struct {
	Template pgtype.JSONB
	Data     []Data
}
