package service

import (
	"fmt"
	"github.com/RakunKo/Tistory-Readme-Go/pkg/model"
)

func BuildSvg(channel model.Channel, color string) string {
	blogName := channel.Title
	userName := channel.ManagingEditor

	return fmt.Sprintf(`<svg xmlns="http://www.w3.org/2000/svg" width="180" height="62">
  <foreignObject width="100%%" height="100%%">
    <div xmlns="http://www.w3.org/1999/xhtml" 
         style="background-color: #%s; border-radius: 8px; padding: 10px; box-shadow: 0 1px 2px rgba(0, 0, 0, 0.08);
                display: flex; align-items: center; font-family: Arial, sans-serif; width: 100%%; box-sizing: border-box; height: 100%%;">
      
      <div style="width: 36px; height: 36px; border-radius: 50%%; overflow: hidden; margin-right: 10px; display: flex; align-items: center; justify-content: center;">
        <!-- 티스토리 로고 SVG 직접 삽입 -->
        <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 459 459" width="36" height="36">
          <title>티스토리 로고</title>
          <g>
            <path fill="#eeeeee" d="M229.5,0C102.75,0,0,102.75,0,229.5S102.75,459,229.5,459,459,356.25,459,229.5,356.25,0,229.5,0ZM130.21,191.45a39.57,39.57,0,1,1,39.56-39.57A39.58,39.58,0,0,1,130.21,191.45ZM229.5,390a39.56,39.56,0,1,1,39.56-39.56A39.56,39.56,0,0,1,229.5,390Zm0-99.29a39.56,39.56,0,1,1,39.56-39.56A39.56,39.56,0,0,1,229.5,290.74Zm0-99.29a39.57,39.57,0,1,1,39.56-39.57A39.57,39.57,0,0,1,229.5,191.45Zm99.29,0a39.57,39.57,0,1,1,39.57-39.57A39.57,39.57,0,0,1,328.79,191.45Z"/>
          </g>
        </svg>
      </div>
      
      <div style="flex: 1; min-width: 0; text-align: left;">
        <div style="font-size: 12px; font-weight: bold; color: #ffffff; margin-bottom: 4px; overflow: hidden; text-overflow: ellipsis; white-space: nowrap;">
          %s
        </div>
        <div style="font-size: 10px; color: #eeeeee; overflow: hidden; text-overflow: ellipsis; white-space: nowrap;">
          %s
        </div>
      </div>
    </div>
  </foreignObject>
</svg>`, color, userName, blogName)
}
