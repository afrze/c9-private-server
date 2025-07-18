package db

import (
	"context"
	"database/sql"
)

// CreateAccount returns (dup, accNo, err)
// dup   → true  if RETURN code == 2
// accNo → the SCOPE_IDENTITY() value from the proc
func CreateAccount(ctx context.Context, db *sql.DB, accId, pwd string) (dup bool, accNo int64, err error) {
	var ret int32

	_, err = db.ExecContext(ctx,
		`EXEC @ret = [C9Unity].[Auth].[UspCreateAccount]
		     @pAccId    = @accId,
		     @pPassword = @pwd,
		     @pAccNo    = @outAccNo OUTPUT`,
		sql.Named("accId", accId),
		sql.Named("pwd", pwd),
		sql.Named("outAccNo", sql.Out{Dest: &accNo}),
		sql.Named("ret", sql.Out{Dest: &ret}),
	)
	if err != nil {
		return false, 0, err
	}

	return ret == 2, accNo, nil
}
