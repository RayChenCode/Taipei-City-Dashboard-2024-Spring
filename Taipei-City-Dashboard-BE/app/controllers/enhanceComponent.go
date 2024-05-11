package controllers

import (
	"TaipeiCityDashboardBE/app/models"
	"TaipeiCityDashboardBE/app/util"
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
)

type TagScore struct {
	TagName string `json:"tag_name"`
	Score   int    `json:"score"`
}

type RequestTagScores struct {
	Question string `json:"question" form:"question"`
}

type ResponseTagScores struct {
	Question string     `json:"question" form:"question"`
	Scores   []TagScore `json:"scores"`
}

func GetTagScores(c *gin.Context) {
	question := RequestTagScores{}
	if err := c.BindQuery(&question); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": err.Error()})
		return
	}
	/*
		主題,TAG1,TAG2,TAG3,TAG4,TAG5,TAG6,TAG7,TAG8,TAG9,TAG10
		人均減碳,碳排放,能源消耗,環保政策,節能減碳,溫室氣體,氣候變遷,環境永續,二氧化碳當量,減碳管理,能源效率
		刑事統計,犯罪率,治安維護,罪案類型,地區熱點,破案率,警力資源,預防犯罪,社區安全,資料分析,政策規劃
		台北市各區即時空氣品質,空氣污染,PM2.5濃度,環境監測,健康風險,污染來源,敏感區域,揚塵管控,空品預警,民生品質,永續城市
		抽水站狀態,防災應變,水利設施,抽水站運作,水位監控,排水系統,防汛整備,易淹水區域,雨災風險,基礎建設,緊急應變
		社福人口,低收入戶,身心障礙,弱勢關懷,社會資源,生活補助,福利政策,人口統計,經濟狀況,社會公平,風險評估
		社福機構,社會服務,機構分布,弱勢族群,服務資源,社區照顧,長期照護,兒少保護,婦女權益,專業輔導,民間組織
		綠電發電量,再生能源,發電量,能源結構,環境保護,碳中和,能源轉型,綠能政策,發電效率,減碳目標,科技創新
		資源站可回收項目,資源回收,廢棄物處理,環保意識,循環經濟,垃圾分類,資源再利用,減廢政策,回收率,環境教育,永續發展
		閒置市有財產,公有地,閒置資產,資產管理,土地利用,都市規劃,地政,土地徵收,興辦事業,場所再利用,市地重劃
	*/

	tags := []string{
		"人均減碳", "碳排放", "能源消耗", "環保政策", "節能減碳", "溫室氣體", "氣候變遷", "環境永續", "二氧化碳當量", "減碳管理", "能源效率",
		"刑事統計", "犯罪率", "治安維護", "罪案類型", "地區熱點", "破案率", "警力資源", "預防犯罪", "社區安全", "資料分析", "政策規劃",
		"台北市各區即時空氣品質", "空氣污染", "PM2.5濃度", "環境監測", "健康風險", "污染來源", "敏感區域", "揚塵管控", "空品預警", "民生品質", "永續城市",
		"抽水站狀態", "防災應變", "水利設施", "抽水站運作", "水位監控", "排水系統", "防汛整備", "易淹水區域", "雨災風險", "基礎建設", "緊急應變",
		"社福人口", "低收入戶", "身心障礙", "弱勢關懷", "社會資源", "生活補助", "福利政策", "人口統計", "經濟狀況", "社會公平", "風險評估",
		"社福機構", "社會服務", "機構分布", "弱勢族群", "服務資源", "社區照顧", "長期照護", "兒少保護", "婦女權益", "專業輔導", "民間組織",
		"綠電發電量", "再生能源", "發電量", "能源結構", "環境保護", "碳中和", "能源轉型", "綠能政策", "發電效率", "減碳目標", "科技創新",
		"資源站可回收項目", "資源回收", "廢棄物處理", "環保意識", "循環經濟", "垃圾分類", "資源再利用", "減廢政策", "回收率", "環境教育", "永續發展",
		"閒置市有財產", "公有地", "閒置資產", "資產管理", "土地利用", "都市規劃", "地政", "土地徵收", "興辦事業", "場所再利用", "市地重劃",
	}

	// random score for each tag, use random number for now
	var tagScores []TagScore
	for _, tag := range tags {
		score := rand.Intn(20)
		tagScores = append(tagScores, TagScore{TagName: tag, Score: score})
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "data": ResponseTagScores{Question: question.Question, Scores: tagScores}})
	//c.JSON(http.StatusOK, gin.H{"status": "success", "data": tagScores})
}

func GetAllComponentTag(c *gin.Context) {
	query := componentQuery{
		PageSize: 1000,
		PageNum:  1,
	}
	components, _, _, err := models.GetAllComponents(query.PageSize, query.PageNum, query.Sort, query.Order, query.FilterBy, query.FilterMode, query.FilterValue, query.SearchByIndex, query.SearchByName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": err.Error()})
		return
	}

	componentTags := make(map[string]string)
	// Return the components
	for _, component := range components {
		queryType, queryString, err := models.GetComponentChartDataQuery(int(component.ID))
		if err != nil {
			//c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": err.Error()})
			print(err)
			continue
		}

		timeFrom, timeTo := util.GetTime(c)
		tags := make(map[string]string)
		//components name, source, short_desc, long_desc, use_case
		tags["name"] = component.Name
		tags["source"] = component.Source
		tags["short_desc"] = component.ShortDesc
		tags["long_desc"] = component.LongDesc
		tags["use_case"] = component.UseCase
		//tags = append(tags, []string{component.Name, component.Source, component.ShortDesc, component.LongDesc, component.UseCase}...)

		if queryType == "two_d" {
			chartData, err := models.GetTwoDimensionalData(&queryString, timeFrom, timeTo)
			if err != nil {
				print(err)
				continue
			}

			for _, datum := range chartData {
				for _, data := range datum.Data {
					//tags = append(tags, data.Xaxis)
					tags[data.Xaxis] = data.Xaxis
				}
			}

		} else if queryType == "three_d" || queryType == "percent" {
			chartData, categories, err := models.GetThreeDimensionalData(&queryString, timeFrom, timeTo)
			if err != nil {
				print(err)
				continue
			}

			for _, datum := range chartData {
				//tags = append(tags, datum.Name)
				//tags = append(tags, datum.Icon)
				tags[datum.Name] = datum.Name
				tags[datum.Icon] = datum.Icon
			}

			//tags = append(tags, categories...)
			for _, category := range categories {
				tags[category] = category
			}

			//c.JSON(http.StatusOK, gin.H{"status": "success", "data": chartData, "categories": categories})
		} else if queryType == "time" {
			chartData, err := models.GetTimeSeriesData(&queryString, timeFrom, timeTo)
			if err != nil {
				print(err)
				continue
			}

			for _, datum := range chartData {
				//tags = append(tags, datum.Name)
				tags[datum.Name] = datum.Name
			}
			//c.JSON(http.StatusOK, gin.H{"status": "success", "data": chartData})
		} else if queryType == "map_legend" {
			chartData, err := models.GetMapLegendData(&queryString, timeFrom, timeTo)
			if err != nil {
				print(err)
				continue
			}

			for _, datum := range chartData {
				//tags = append(tags, datum.Name)
				tags[datum.Name] = datum.Name
			}
		}

		var tagList []string
		for key := range tags {
			value := tags[key]
			// check if key is number
			_, err := strconv.Atoi(key)
			if key != "" && err != nil {
				// if key 最後是 "區"，不加入tagList
				if !strings.HasSuffix(key, "區") {
					tagList = append(tagList, key)
				}
			}

			// check if value is number
			_, err = strconv.Atoi(value)
			if value != "" && err != nil {
				if !strings.HasSuffix(value, "區") {
					tagList = append(tagList, value)
				}
			}
		}

		tagList = append(tagList, component.Name)
		// join tagList with comma
		componentTags[component.Name] = strings.Join(tagList, ",")
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "data": componentTags})
}
