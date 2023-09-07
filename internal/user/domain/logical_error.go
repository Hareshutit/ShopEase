package domain

import (
	"fmt"
)

type PassNonComporableErr struct{}

func (e PassNonComporableErr) Error() string {
	return fmt.Sprintf("Пароли не совпадают")
}
