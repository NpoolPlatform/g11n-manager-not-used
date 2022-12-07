package appcountry

import (
	"fmt"

	npool "github.com/NpoolPlatform/message/npool/g11n/mgr/v1/appcountry"

	"github.com/google/uuid"
)

func validate(info *npool.CountryReq) error {
	if info.ID != nil {
		if _, err := uuid.Parse(info.GetID()); err != nil {
			return err
		}
	}

	if _, err := uuid.Parse(info.GetAppID()); err != nil {
		return err
	}
	if _, err := uuid.Parse(info.GetCountryID()); err != nil {
		return err
	}

	return nil
}

func duplicate(infos []*npool.CountryReq) error {
	keys := map[string]struct{}{}
	apps := map[string]struct{}{}

	for _, info := range infos {
		if err := validate(info); err != nil {
			return err
		}

		key := fmt.Sprintf("%v:%v", info.AppID, info.CountryID)
		if _, ok := keys[key]; ok {
			return fmt.Errorf("duplicate appcountry")
		}

		keys[key] = struct{}{}
		apps[info.GetAppID()] = struct{}{}
	}

	if len(apps) > 1 {
		return fmt.Errorf("invalid apps")
	}

	return nil
}

func Validate(info *npool.CountryReq) error {
	return validate(info)
}

func Duplicate(infos []*npool.CountryReq) error {
	return duplicate(infos)
}
