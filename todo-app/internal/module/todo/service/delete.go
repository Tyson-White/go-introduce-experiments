package todo

func (s *TodoService) DeleteTodo(todoId int) error {
	return s.repository.DeleteTodo(todoId)
}
