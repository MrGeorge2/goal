package list

type IQueriable[T any] interface {
	Where(T) IQueriable[T]
}
