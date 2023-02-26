package utils

import "github.com/Manjit2003/bot-framework/types"

func GetListParams(btnText string, sectionTitle string, data ...string) *types.ListData {

	rows := make([]types.ListRow, 0)

	for _, d := range data {
		rows = append(rows, types.ListRow{
			Id:    d,
			Title: d,
		})
	}

	listData := &types.ListData{
		ButtonText: btnText,
		Sections: []types.ListSection{
			{
				Title: sectionTitle,
				Rows:  rows,
			},
		},
	}

	return listData

}

func GetButtonParams(buttons ...string) []types.Button {

	btns := make([]types.Button, 0)

	for _, b := range buttons {
		btns = append(btns, types.Button{
			Type: "reply",
			Reply: types.ButtonContent{
				Id:    b,
				Title: b,
			},
		})
	}

	return btns
}
