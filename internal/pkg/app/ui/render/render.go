package render

import (
	"fyne.io/fyne/v2/widget"
	"go-qn2management/internal/pkg/app/config"
	"go-qn2management/internal/pkg/app/i18n"
	"go-qn2management/internal/pkg/app/repository"
	"go-qn2management/internal/pkg/app/service"
	"log"
	"sort"
	"strconv"
)

// Config contains all UI configs, and distribute service for components working together
type Config struct {
	Toolbar       *widget.Toolbar
	SessionWidget *widget.Table
}

type Render interface {
	GetRenderConfig() Config
	SetSessionWidget(sessionWidget *widget.Table)
	GetSessionProps() (map[string][]*repository.SessionItem, [][]interface{})
}

type render struct {
	service service.Service
}

type Item struct {
	sessionItem repository.SessionItem
}

func New(service service.Service) *render {
	return &render{
		service: service,
	}
}

var DrawConfig Config

func (r *render) GetRenderConfig() Config {
	return DrawConfig
}

func (r *render) SetSessionWidget(widgetTable *widget.Table) {
	DrawConfig.SessionWidget = widgetTable
}

func (r *render) GetSessionProps() (map[string][]*repository.SessionItem, [][]interface{}) {
	var slice [][]interface{}

	// Get sessions from mongo
	sessions, err := r.service.GetAllSessions()
	if err != nil {
		log.Println(err)
		return nil, nil
	}
	sort.Slice(sessions, func(i, j int) bool {
		return sessions[i].CreatedAt.After(sessions[j].CreatedAt)
	})

	// Get items for session from mongo
	items, err := r.service.GetAllItems()
	if err != nil {
		log.Println(err)
		return nil, nil
	}

	sessionMap := make(map[string][]*repository.SessionItem)

	slice = append(slice, []interface{}{
		"ID1", // origin by session id
		"ID2", // origin by item id
		i18n.Messages["delete_session"][config.GlobalConfig.Language],
		i18n.Messages["session_name"][config.GlobalConfig.Language],
		i18n.Messages["session_key"][config.GlobalConfig.Language],
		i18n.Messages["title"][config.GlobalConfig.Language],
		i18n.Messages["youtube_extension"][config.GlobalConfig.Language],
		i18n.Messages["add_item"][config.GlobalConfig.Language],
		i18n.Messages["delete_item"][config.GlobalConfig.Language],
	},
	)

	for _, session := range sessions {
		sessionMap[session.ID] = []*repository.SessionItem{}

		for _, item := range items {
			if session.ID == item.SessionID {
				sessionMap[session.ID] = append(sessionMap[session.ID], item)
			}
		}
	}

	indexSession := 0
	for _, session := range sessions {
		indexSession++
		if len(sessionMap[session.ID]) > 0 {
			for index, item := range sessionMap[session.ID] {
				var currentRow []interface{}

				currentRow = append(currentRow, session.ID)
				currentRow = append(currentRow, item.ID)
				currentRow = append(currentRow, "") // delete session button
				currentRow = append(currentRow, session.SessionName)
				currentRow = append(currentRow, session.SessionKey)
				currentRow = append(currentRow, item.Title)
				currentRow = append(currentRow, item.Extension)

				if index == 0 {
					currentRow = append(currentRow, true)                       // add item button
					currentRow = append(currentRow, strconv.Itoa(indexSession)) // delete item button (use with index for session)
				} else {
					currentRow = append(currentRow, false)
					currentRow = append(currentRow, strconv.Itoa(indexSession))
				}
				slice = append(slice, currentRow)
			}
		} else {
			var currentRow []interface{}
			currentRow = append(currentRow, session.ID)
			currentRow = append(currentRow, "") // null item id
			currentRow = append(currentRow, "") // delete session button
			currentRow = append(currentRow, session.SessionName)
			currentRow = append(currentRow, session.SessionKey)
			currentRow = append(currentRow, "")
			currentRow = append(currentRow, "")
			currentRow = append(currentRow, true)                       // add item button
			currentRow = append(currentRow, strconv.Itoa(indexSession)) // delete item button (use with index for session)
			slice = append(slice, currentRow)
		}
	}

	return sessionMap, slice
}
