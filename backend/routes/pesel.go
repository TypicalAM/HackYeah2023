package routes

import (
	"fmt"
	"net/http"

	"github.com/TypicalAM/hackyeah/prescription"
	"github.com/TypicalAM/hackyeah/validators"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type PeselInput struct {
	Pesel string `json:"pesel"`
	Code  string `json:"code"`
}

type PeselOutput struct {
	UUID  string              `json:"uuid"`
	Drugs []prescription.Drug `json:"drugs"`
}

func (c *Controller) Pesel(e echo.Context) error {
	var input PeselInput
	if err := e.Bind(&input); err != nil {
		return err
	}

	if valid := validators.Pesel(input.Pesel); !valid {
		return e.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid PESEL number"})
	}

	if valid := validators.PeselCode(input.Code); !valid {
		return e.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid code"})
	}

	drugs, err := c.prepository.GetDrugsForPesel(input.Pesel, fmt.Sprint(input.Code))
	if err != nil {
		return err
	}

	return e.JSON(http.StatusOK, PeselOutput{
		Drugs: *drugs,
		UUID:  uuid.New().String(),
	})
}
