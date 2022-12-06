package country

import (
	"fmt"

	npool "github.com/NpoolPlatform/message/npool/g11n/mgr/v1/country"

	"github.com/google/uuid"
)

func validate(info *npool.CountryReq) error {
	if info.ID != nil {
		if _, err := uuid.Parse(info.GetID()); err != nil {
			return err
		}
	}

	if info.GetCountry() == "" {
		return fmt.Errorf("country is invalid")
	}
	if info.GetFlag() == "" {
		return fmt.Errorf("flag is invalid")
	}
	if info.GetCode() == "" {
		return fmt.Errorf("code is invalid")
	}
	if info.GetShort() == "" {
		return fmt.Errorf("short is invalid")
	}

	return nil
}

func duplicate(infos []*npool.CountryReq) error {
	keys := map[string]struct{}{}

	for _, info := range infos {
		if err := validate(info); err != nil {
			return err
		}

		key := fmt.Sprintf("%v", info.Country)
		if _, ok := keys[key]; ok {
			return fmt.Errorf("duplicate country")
		}

		keys[key] = struct{}{}
	}

	return nil
}
