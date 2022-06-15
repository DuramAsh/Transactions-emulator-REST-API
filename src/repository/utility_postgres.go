package repository

func (r *TransactionPostgres) TransactionExists(transactionId int) bool {
	query := "SELECT count(*) FROM transactions WHERE id = $1"
	var amount int
	err := r.db.Get(&amount, query, transactionId)
	if err != nil {
		return false
	}
	if amount != 0 {
		return true
	}
	return false
}
