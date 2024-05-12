<script setup>
import { reactive, ref, onMounted, nextTick } from 'vue'
import { useAssistant } from '../store/assistantStore'
import { useDialogStore } from '../store/dialogStore'
import MoreInfo from '../components/dialogs/MoreInfo.vue'
import { useRouter } from 'vue-router'
//import { axios } from "axios";
const router = useRouter();
const assistant = useAssistant()
const dialogStore = useDialogStore()
const emit = defineEmits(['open'])
const dashboard = ref('');
const handleLineClick = (e) => {
console.log('isSelect',isSelect.value)
	let tagsString = ''
	if(isSelect.value.length > 0){
		isSelect.value.forEach((el)=>{
			console.log('el',el)
			tagsString = tagsString + '&tags=' + el
		})
	}
	fetch(`/api/dev/enhance/setDashboardComponents?dashboard_index=${dashboard.value}${tagsString}`)
		.then(function (response) {
			return response.json();
		})
		.then(function (myJson) {
			staffMemberArr.value = myJson.staffMemberArr
			if (result.Msg.account) {
				staffNum.value = staffMemberArr.value.find((item, index) => {
					if (item === result.Msg.account) {
						return true
					}
				})
				if (staffNum.value) {
					isStaff.value = true
				} else {
					alert('無權限')
					cm.closeWindow()
				}
			} else {
				alert('無權限')
				cm.closeWindow()
			}
		});
	router.push('/mapview?index=283f362e9566')
}


// route.push('/cityGPT')
// const cc =  {
// 	"id": 151,
// 	"index": "TestChart3",
// 	"map_config": [
// 		{
// 			"index": "TestChart3",
// 			"type": "fill",
// 			"paint": {
// 				"fill-color": [
// 					"match",
// 					["get", "category"],
// 					"#d1a26c",
// 					"#b9734a",
// 					"#9f4333",
// 					"#800026"
// 				],
// 				"fill-opacity": 0.8
// 			},
// 			"property": [
// 				{
// 					"key": "category",
// 					"name": "總體需求"
// 				}
// 			],
// 			"title": "潛在需求"
// 		}
// 	],
// 	"chart_config": {
// 		"types": ["TestChart3"],
// 		"color": ["#d1a26c", "#b9734a", "#9f4333", "#800026"],
// 		"unit": "件"
// 	},
// 	"chart_data": [
// 		{
// 			"name": "",
// 			"data": [
// 				{
// 					"x": "2023/01/01",
// 					"y": 300,
// 					"data": [1, 2, 3, 4]
// 				},
// 				{
// 					"x": "2023/04/01",
// 					"y": 496,
// 					"data": [1, 3, 2, 4]
// 				},
// 				{
// 					"x": "2023/07/01",
// 					"y": 457,
// 					"data": [2, 3, 1, 4]
// 				},
// 				{
// 					"x": "2023/10/01",
// 					"y": 403,
// 					"data": [3, 4, 2, 1]
// 				}
// 			]
// 		}
// 	],
// 	"name": "綠化路線新增量與累計量",
// 	"source": "鴻海",
// 	"time_from": "2023-11-11T00:00:00+08:00",
// 	"time_to": null,
// 	"short_desc": "",
// 	"long_desc": "測試組件的說明",
// 	"use_case": "測試組件的情境",
// 	"tags": [],
// 	"links": [
// 		"https://tuic.gov.taipei/youbike",
// 		"https://github.com/tpe-doit/YouBike-Optimization"
// 	],
// 	"contributors": ["tuic"]
// }

// dialogStore.showMoreInfo(cc)
////
// let clickedElId = e.target.id
// if (clickedElId === 'temp_testing_div2') {
//   emit('open', true)
// }
var isCheck = ref(false);
let isSelect = ref([])

onMounted(async () => {
})

const getDisabled = () => {
	
}

</script>
<template>
	<div class="conversation">
		<div v-for="(text, idx) in assistant.textScript" :key="idx + 'text'">
			<div class="conversation-wrapper" :class="{ res: idx % 2 === 1 }">
				<div class="conversation-contnet">
					<div v-for="(response, index) in text.response" :key="index">
						<input v-model="isSelect" type="checkbox" id="scales" name="scales" :checked="response.score > 50"
							@click="getDisabled()" @change="onChangee" :value="response.tag_name"/>
						<label for="scales">標籤: {{ response.tag_name }} 分數:{{ response.score }}</label>
					</div>
					<div v-html="text.content" @click="handleLineClick">
					</div>

				</div>
				<span>{{ text.time }}</span>

			</div>
		</div>
	</div>
	<MoreInfo />
</template>
<style scoped lang="scss">
.conversation {
	display: flex;
	flex-direction: column;
	gap: 10px;

	&-wrapper {
		display: flex;
		flex-direction: column;
		align-items: flex-end;


		&.res {
			align-items: flex-start;
			// height: 100vh;
			//overflow: unset;

			.conversation-contnet {
				background-color: #3c3c47;
				border-radius: 8px 8px 8px 0px;
				max-width: 312px;
			}
		}


		>span {
			font-size: 14px;
			font-weight: 400;
			color: rgba(255, 255, 255, 0.5);
			margin-top: 4px;
		}
	}

	&-contnet {
		padding: 12px;
		background-color: #4C37CD;
		max-width: 264px;
		border-radius: 8px 8px 0px 8px;
	}
}
</style>
