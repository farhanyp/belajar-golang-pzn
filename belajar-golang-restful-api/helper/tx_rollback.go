package helper

import "database/sql"

func TxRollback(err error, tx *sql.Tx) {
	if err != nil {
		ErrorRollback := tx.Rollback()
		IfError(ErrorRollback)
	}

}