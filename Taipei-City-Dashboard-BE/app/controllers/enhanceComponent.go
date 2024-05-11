package controllers

import (
	"TaipeiCityDashboardBE/app/models"
	"TaipeiCityDashboardBE/app/util"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
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
	Question  string     `json:"question" form:"question"`
	Scores    []TagScore `json:"scores"`
	Threshold int        `json:"threshold"`
}

type RequestUpdateDashboardComponents struct {
	Tags           []string `json:"tags" form:"tags"`
	DashboardIndex string   `json:"dashboard_index" form:"dashboard_index"`
}

func SetDashboardComponents(c *gin.Context) {
	// Get all query parameters from context
	var query RequestUpdateDashboardComponents
	if err := c.BindQuery(&query); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": err.Error()})
		return
	}

	if query.DashboardIndex == "" {
		query.DashboardIndex = "51a92f2df637"
	}
	// select components by contain anyone of tags
	// sql
	sql := "SELECT * FROM components WHERE "
	for i, tag := range query.Tags {
		if i == 0 {
			sql += "tags LIKE '%" + tag + "%' "
		} else {
			sql += "OR tags LIKE '%" + tag + "%' "
		}
	}

	// get components
	components := []models.Component{}
	err := models.DBManager.Raw(sql).Scan(&components).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": err.Error()})
		return
	}

	// find dashboard
	dashboard := models.Dashboard{}
	err = models.DBManager.Where("index = ?", query.DashboardIndex).First(&dashboard).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": err.Error()})
		return
	}

	componentsIDs := pq.Int64Array{}
	for _, component := range components {
		componentsIDs = append(componentsIDs, component.ID)
	}

	// get dashboard_groups by dashboard id
	dashboardGroups := []models.DashboardGroup{}
	err = models.DBManager.Where("dashboard_id = ?", dashboard.ID).Find(&dashboardGroups).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": err.Error()})
		return
	}

	var groups []int
	for _, dashboardGroup := range dashboardGroups {
		groups = append(groups, dashboardGroup.GroupID)
	}

	newDashboard, err := models.UpdateDashboard(dashboard.Index, dashboard.Name, dashboard.Icon, componentsIDs, groups)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "data": newDashboard})
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
	/*
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
	*/
	tags := []string{
		"[創綠]綠電發電用電佔比量", "再生能源", "發電量", "能源結構", "環境保護", "碳中和", "能源轉型", "綠能政策", "發電效率", "減碳目標", "數據分析", "用電", "創綠",
		"[永續]台北市各區即時空氣品質", "空氣污染", "PM2.5濃度", "環境監測", "健康風險", "污染來源", "敏感區域", "揚塵管控", "空品預警", "區域差異", "數據視覺化", "空氣品質", "永續",
		"[永續]台北市各區歷年空氣品質", "空氣污染", "PM2.5濃度", "環境監測", "健康風險", "污染來源", "敏感區域", "揚塵管控", "空品預警", "區域差異", "數據趨勢", "空氣品質", "永續",
		"[減廢]循環杯租借分佈", "環保意識", "減廢政策", "資源再利用", "永續發展", "循環經濟", "社區參與", "公共設施", "地理分佈", "市政建設", "綠色生活", "環保餐具", "減廢",
		"[減廢]資源站可回收項目", "資源回收", "廢棄物處理", "環保意識", "循環經濟", "垃圾分類", "資源再利用", "減廢政策", "回收率", "環境教育", "永續發展", "環保餐具", "減廢",
		"[節電]輔導節電統計", "節能減碳", "能源效率", "用電行為", "節電措施", "用電數據", "家戶用電", "補助政策", "目標管理", "統計分析", "政策評估", "節能減碳", "節能",
		"[綠運]電動巴士淨零貢獻量", "淨零", "電動公車", "綠色交通", "電動車輛", "可再生能源", "空氣污染", "溫室氣體", "能源轉型", "環境保護", "永續發展", "氣候行動", "創新技術", "公車", "市政府", "節能減碳",
		"[除碳]人均減碳", "碳排放", "能源消耗", "環保政策", "節能減碳", "溫室氣體", "氣候變遷", "環境永續", "二氧化碳當量", "減碳管理", "數據分析", "節能減碳", "除碳",
		"[除碳]台北市歷年排碳與人口數", "台北市", "碳排放", "人口變化", "氣候影響", "能源使用", "城市規劃", "減排政策", "碳足跡", "環境數據", "數據分析", "長期趨勢", "節能減碳",
	}

	// random score for each tag, use random number for now
	var tagScores []TagScore
	for _, tag := range tags {
		score := rand.Intn(20)
		// spec list
		if strings.Contains(question.Question, "公車") {

		}
		tagScores = append(tagScores, TagScore{TagName: tag, Score: score})
	}
	res := ResponseTagScores{Question: question.Question, Scores: tagScores, Threshold: rand.Intn(20) + 5}
	c.JSON(http.StatusOK, gin.H{"status": "success", "data": res})
	//c.JSON(http.StatusOK, gin.H{"status": "success", "data": tagScores})
}

func GetAllComponentText(c *gin.Context) {
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

func GetAllComponentTags(c *gin.Context) {
	query := componentQuery{
		PageSize: 1000,
		PageNum:  1,
	}
	components, _, _, err := models.GetAllComponents(query.PageSize, query.PageNum, query.Sort, query.Order, query.FilterBy, query.FilterMode, query.FilterValue, query.SearchByIndex, query.SearchByName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "data": components})
}
