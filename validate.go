package confuse

func (s *Service) validate(result any) error {
	if s.ShouldValidate && s.Validator != nil {
		return s.Validator.Struct(result)
	}

	return nil
}
