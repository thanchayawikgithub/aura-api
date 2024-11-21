package handler

import "context"

func (s *Service) ExportUsers(ctx context.Context) error {
	return s.ExportUser.ExportUsers(ctx)
}
