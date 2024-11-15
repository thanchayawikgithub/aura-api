package exception

type ValidateError struct {
	Message string
}

func (ve *ValidateError) Error() string {
	return ve.Message
}
