package helper

import "database/sql"

func CommitOrRollback(tx *sql.Tx) {
	err := recover()
	if err != nil {
		errR := tx.Rollback()
		PanicError(errR)
		panic(err)
	} else {
		errC := tx.Commit()
		PanicError(errC)
	}
}