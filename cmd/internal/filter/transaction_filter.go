package filter

type TransactionFilter struct {
	addresses []string
}

func NewTransactionFilter(addresses []string) *TransactionFilter {
	return &TransactionFilter{
		addresses: addresses,
	}
}

func (filter *TransactionFilter) ContainsAddress(address string) bool {
	for _, filterAddress := range filter.addresses {
		if address == filterAddress {
			return true
		}
	}

	return false
}
