package region

import (
	"github.com/zqhhh/gomall/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListAll(c *gin.Context) {
	ack := Ack{
		Id:    0,
		Name:  "",
		Child: []Ack{},
	}
	data := model.GetAllRegions()
	for _, d := range data {
		if d.ParentID == 0 {
			ack.Id = d.ID
			ack.Name = d.Name
		} else if d.ParentID == ack.Id {
			ack.Child = append(ack.Child, Ack{
				Id:    d.ID,
				Name:  d.Name,
				Child: []Ack{},
			})
		}
	}

	for i, c := range ack.Child {
		for _, d := range data {
			if d.ParentID == c.Id {
				ack.Child[i].Child = append(ack.Child[i].Child, Ack{
					Id:    d.ID,
					Name:  d.Name,
					Child: []Ack{},
				})
			}
		}
	}
	for i, _ := range ack.Child {
		for j, c2 := range ack.Child[i].Child {
			for _, d := range data {
				if d.ParentID == c2.Id {
					ack.Child[i].Child[j].Child = append(ack.Child[i].Child[j].Child, Ack{
						Id:    d.ID,
						Name:  d.Name,
						Child: []Ack{},
					})
				}
			}
		}
	}
	c.JSON(http.StatusOK, ack.Child)
}
