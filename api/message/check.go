package message

import (
	"fmt"

	npool "github.com/NpoolPlatform/message/npool/g11n/mgr/v1/message"

	"github.com/google/uuid"
)

func validate(info *npool.MessageReq) error {
	if info.ID != nil {
		if _, err := uuid.Parse(info.GetID()); err != nil {
			return err
		}
	}

	if _, err := uuid.Parse(info.GetAppID()); err != nil {
		return err
	}
	if _, err := uuid.Parse(info.GetLangID()); err != nil {
		return err
	}
	if info.GetMessageID() == "" {
		return fmt.Errorf("messageid is invalid")
	}
	if info.GetMessage() == "" {
		return fmt.Errorf("message is invalid")
	}

	return nil
}

func duplicate(infos []*npool.MessageReq) error {
	keys := map[string]struct{}{}

	for _, info := range infos {
		if err := validate(info); err != nil {
			return err
		}

		key := fmt.Sprintf("%v", info.Message)
		if _, ok := keys[key]; ok {
			return fmt.Errorf("duplicate message")
		}

		keys[key] = struct{}{}
	}

	return nil
}
