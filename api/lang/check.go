package lang

import (
	"fmt"

	npool "github.com/NpoolPlatform/message/npool/g11n/mgr/v1/lang"

	"github.com/google/uuid"
)

func validate(info *npool.LangReq) error {
	if info.ID != nil {
		if _, err := uuid.Parse(info.GetID()); err != nil {
			return err
		}
	}

	if info.GetLang() == "" {
		return fmt.Errorf("lang is invalid")
	}
	if info.GetLogo() == "" {
		return fmt.Errorf("logo is invalid")
	}
	if info.GetName() == "" {
		return fmt.Errorf("name is invalid")
	}
	if info.GetShort() == "" {
		return fmt.Errorf("short is invalid")
	}

	return nil
}

func duplicate(infos []*npool.LangReq) error {
	keys := map[string]struct{}{}

	for _, info := range infos {
		if err := validate(info); err != nil {
			return err
		}

		key := fmt.Sprintf("%v", info.Lang)
		if _, ok := keys[key]; ok {
			return fmt.Errorf("duplicate lang")
		}

		keys[key] = struct{}{}
	}

	return nil
}
