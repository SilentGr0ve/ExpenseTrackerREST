package web_repository

import (
	"fmt"
	"os"

	core_errors "github.com/SilentGr0ve/ExpenseTrackerREST/internal/core/errors"
)

func (r *WebRepository) GetFile(filepath string) ([]byte, error) {
	file, err := os.ReadFile(filepath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, fmt.Errorf(
				"file: %s: %w",
				filepath,
				core_errors.ErrNotFound,
			)
		}

		return nil, fmt.Errorf(
			"get file: %s: %w",
			filepath,
			err,
		)

	}

	return file, nil
}
