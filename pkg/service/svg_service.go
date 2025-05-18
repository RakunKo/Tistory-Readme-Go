package service

import (
	"fmt"
	"github.com/RakunKo/Tistory-Readme-Go/internal/model"
)

func BuildSvg(channel model.Channel) string {
	blogName := channel.Title
	profileImg := channel.Image.URL
	userName := channel.ManagingEditor

	return fmt.Sprintf(`<svg xmlns="http://www.w3.org/2000/svg" width="180" height="62">
  <foreignObject width="100%%" height="100%%">
    <div xmlns="http://www.w3.org/1999/xhtml" style="background-color: #ff6f61; border-radius: 8px; padding: 10px; box-shadow: 0 1px 2px rgba(0, 0, 0, 0.08); display: flex; align-items: center; font-family: Arial, sans-serif; width: 100%%; box-sizing: border-box; height: 100%%;">
      <div style="width: 36px; height: 36px; border-radius: 50%%; overflow: hidden; margin-right: 10px;">
        <img src="%s" style="width: 36px; height: 36px; object-fit: cover;" alt="profile"/>
      </div>
      <div style="flex: 1; min-width: 0; text-align: left;">
        <div style="font-size: 12px; font-weight: bold; color: #ffffff; margin-bottom: 4px; overflow: hidden; text-overflow: ellipsis; white-space: nowrap;">%s</div>
        <div style="font-size: 10px; color: #eeeeee; overflow: hidden; text-overflow: ellipsis; white-space: nowrap;">%s</div>
      </div>
    </div>
  </foreignObject>
</svg>`, profileImg, userName, blogName)
}
