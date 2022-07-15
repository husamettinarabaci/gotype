package slice

//
// TODO
//
func Remove[t []T, T any](slice []T, s int) []T {
	return append(slice[:s], slice[s+1:]...)
}
