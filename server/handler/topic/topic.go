package topic

import (
	"github.com/zqhhh/gomall/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func List(c *gin.Context) {
	q := query{}
	if err := c.BindQuery(&q); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}
	res := Response{
		TopicList: make([]topic, 0),
	}

	res.Count = model.CountTopic()
	if res.Count == 0 {
		c.JSON(http.StatusOK, res)
		return
	}
	offset := 0
	if q.Page > 1 {
		offset = (q.Page - 1) * q.Size
	}
	topList, _ := model.GetTopic([]string{"id", "title", "price_info", "scene_pic_url", "subtitle"}, map[string]interface{}{}, q.Size, offset)
	for _, t := range topList {
		res.TopicList = append(res.TopicList, topic{
			ID:          t.ID,
			Title:       t.Title,
			PriceInfo:   t.PriceInfo,
			ScenePicUrl: t.ScenePicUrl,
			Subtitle:    t.Subtitle,
		})
	}
	if len(topList) < q.Size || q.Size == 0 {
		res.HasNextPage = false
	} else {
		res.HasNextPage = true
	}
	res.PageSize = q.Size
	res.CurrentPage = q.Page
	c.JSON(http.StatusOK, res)
}
