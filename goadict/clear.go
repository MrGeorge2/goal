package goadict

func (d *Goadict[K, T]) Clear() {
	*d = Goadict[K, T]{}
}
