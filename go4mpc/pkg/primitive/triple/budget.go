package triple

// type Budget struct {
// 	B1  int
// 	I8  int
// 	I16 int
// 	I32 int
// 	I64 int
// }

// func NewBudget() *Budget {
// 	return new(Budget)
// }

// func (ct *Budget) Mul(t pd.ValType, n int) {
// 	switch t {
// 	case pd.Int8:
// 		ct.I8 += n
// 	case pd.Int16:
// 		ct.I16 += n
// 	case pd.Int32:
// 		ct.I32 += n
// 	case pd.Int64:
// 		ct.I64 += n
// 	default:
// 		log.Panicln(t, n)
// 	}
// }

// func (ct *Budget) And(n int) {
// 	ct.B1 += n
// }
// func (ct *Budget) Or(n int) {
// 	ct.B1 += n
// }
// func (ct *Budget) Eq(t pd.ValType, n int) {
// 	switch t {
// 	case pd.Bool:
// 	case pd.Int8:
// 		ct.B1 += 7 * n
// 	case pd.Int16:
// 		ct.B1 += 15 * n
// 	case pd.Int32:
// 		ct.B1 += 31 * n
// 	case pd.Int64:
// 		ct.B1 += 63 * n
// 	default:
// 		log.Panicln(t, n)
// 	}
// }
// func (ct *Budget) Mux(t pd.ValType, n int) {
// 	switch t {
// 	case pd.Bool:
// 		ct.B1 += 2 * n
// 	case pd.Int8:
// 		ct.I8 += 3 * n
// 	case pd.Int16:
// 		ct.I16 += 3 * n
// 	case pd.Int32:
// 		ct.I32 += 3 * n
// 	case pd.Int64:
// 		ct.I64 += 3 * n
// 	default:
// 		log.Panicln(t, n)
// 	}
// }
