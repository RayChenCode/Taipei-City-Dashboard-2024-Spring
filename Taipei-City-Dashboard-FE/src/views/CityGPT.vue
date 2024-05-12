<script setup>
import {ref, computed} from "vue";
import conversation from '../components/conversation.vue'
import {useAssistant} from '../store/assistantStore'

const assistant = useAssistant()
const textVal = ref('')
let data = null;
let threshold=null;
const getresponse =async () => {
 await fetch(`/api/dev/enhance/tagScores?question=${textVal.value}`)
      .then(function (response) {
        return response.json();
      })
      .then(function (myJson) {
       // staffMemberArr.value = myJson.data.scores
        data = myJson.data.scores;
        threshold=myJson.data.threshold;
        console.log('data',data)

     //  if (result.Msg.account) {
     //    staffNum.value = staffMemberArr.value.find((item, index) => {
     //      if (item === result.Msg.account) {
     //        return true
     //      }
     //    })
     //    if (staffNum.value) {
     //      isStaff.value = true
     //    } else {
     //      alert('無權限')
     //      cm.closeWindow()
     //    }
     //  } else {
     //    alert('無權限')
     //    cm.closeWindow()
     //  }
     });
}
const  sendQuestion =async () => {
  let date = new Date()
  ///console.log('4567')
 await getresponse();
  ///if (textVal.value.includes('淨零')) {
  assistant.textScript.push
  ({content: `<span>${textVal.value}</span>`, time: date.getHours() + ':' + date.getMinutes()}, {
        content: `<div style="margin-top: 5px;display: flex; gap: 9px; cursor: pointer; align-items: center;"><span style="color: #9E8EFF; font-weight: 700; font-size: 14px;" >產生圖表</span></div> `, time: date.getHours() + ':' + date.getMinutes(),
        response: data
      })
  // } else {
  //   assistant.textScript.push({ content: `<span>${textVal.value}</span>`, time: date.getHours() + ':' + date.getMinutes() },
  //     { content: '<div style="display: flex; gap:2px;">目前顯示三至四個月前的交通事故統計及熱力圖，資料來源為台北市交通局內部資料</div><div style="margin-top: 5px;display: flex; gap: 9px; cursor: pointer; align-items: center;"><span style="color: #9E8EFF; font-weight: 700; font-size: 14px;" >查看圖表</span></div>', time: date.getHours() + ':' + date.getMinutes() })

  // }
  //let date = new Date()
  //assistant.textScript.push({ content: `<span>${textVal.value}</span>`, time: date.getHours() + ':' + date.getMinutes() }, { content: '<div style="display: flex; gap:2px;">目前顯示三至四個月前的交通事故統計及熱力圖，資料來源為台北市交通局內部資料</div><div style="margin-top: 5px;display: flex; gap: 9px; cursor: pointer; align-items: center;"><span style="color: #9E8EFF; font-weight: 700; font-size: 14px;" >查看圖表</span></div>', time: date.getHours() + ':' + date.getMinutes() })

}
</script>

<template>
  <div class="assistant_content">
    <div class="assistantWrapper" :class="{ checked: checked }">
      <div class="assistantWrapper-header">
        <div>
          <div>
            <h3>小幫手</h3>
            <!-- <span @click="dialogStore.showNotification('info', '開啟使用語音助理功能，尋找圖表資訊更便利')">info</span> -->
          </div>
        </div>
      </div>
      <div class="assistantWrapper-opened">
        <div class="assistantWrapper-content">
          <conversation/>
        </div>
        <div class="assistantWrapper-textOrAudio">
          <!-- <img src="../assets/images/Mic.svg" alt="mic" @click="startTTS()"> -->
          <!-- <div class="assistantWrapper-circle"></div> -->
          <input type="text" v-model="textVal">
          <img src="../assets/images/Vector.svg" alt="send" @click="sendQuestion()">
        </div>
      </div>
      <!-- <div v-else-if="checked" class="componentmapchart-loading">
      <div></div>
    </div> -->
    </div>
  </div>
</template>

<style scoped lang="scss">
.assistant_content {


  .assistantWrapper {
    width: calc(100% - var(--font-m) * 2);
    max-width: calc(100% - var(--font-m) * 2);
    display: flex;
    flex-direction: column;
    justify-content: space-between;
    position: relative;
    padding: var(--font-m);
    border-radius: 5px;
    background-color: var(--color-component-background);
    height: calc(100vh - 90px);
    overflow: scroll;

    &-circle {
      width: 10px;
      height: 10px;
      border-radius: 50%;
      background-color: red;
    }

    &-header {
      display: flex;
      justify-content: space-between;
      align-items: baseline;

      h3 {
        font-size: var(--font-m);
      }

      h4 {
        color: var(--color-complement-text);
        font-size: var(--font-s);
        font-weight: 400;
      }

      div:first-child {
        div {
          display: flex;
          align-items: center;
        }

        span {
          margin-left: 8px;
          color: var(--color-complement-text);
          font-family: var(--font-icon);
          user-select: none;
        }
      }

      &-toggle {
        min-height: 1rem;
        min-width: 2rem;
        margin-top: 4px;
      }
    }

    &-control {
      width: 100%;
      display: flex;
      justify-content: center;
      align-items: center;
      position: absolute;
      top: 4rem;
      left: 0;
      z-index: 10;

      button {
        margin: 0 4px;
        padding: 4px 4px;
        border-radius: 5px;
        background-color: rgb(77, 77, 77);
        opacity: 0.25;
        color: var(--color-complement-text);
        font-size: var(--font-s);
        text-align: center;;
        transition: color 0.2s, opacity 0.2s;

        &:hover {
          opacity: 1;
          color: white;
        }
      }
    }

    &-opened,
    &-loading {
      height: 90%;
      // position: relative;
      // overflow-y: scroll;
      display: flex;
      flex-direction: column;
      justify-content: space-between;
    }

    &-loading {
      display: flex;
      align-items: center;
      justify-content: center;

      div {
        width: 2rem;
        height: 2rem;
        border-radius: 50%;
        border: solid 4px var(--color-border);
        border-top: solid 4px var(--color-highlight);
        animation: spin 0.7s ease-in-out infinite;
      }
    }

    &-content {
      width: 100%;
      height: 85%;
      overflow-y: scroll;
    }

    &-textOrAudio {
      border-radius: 5px;
      width: 100%;
      height: 13%;
      background-color: #1e1e22;
      display: flex;
      align-items: center;
      justify-content: space-between;
      padding: 0 8px 0 0;
      box-sizing: border-box;

      > input {
        width: 95%;
      }

      > img {
        cursor: pointer;
      }
    }
  }

  .checked {
    // max-height: 300px;
    // height: 300px;
    height: 500px;
  }
}
</style>
