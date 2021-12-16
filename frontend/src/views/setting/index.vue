<template>
  <div class="setting-container">
    <a-form ref="settingState" :model="settingForm" :rules="settingRules">
      <a-form-item name="publicKey">
        <a-textarea v-model:value="settingForm.publicKey" placeholder="请输入公钥" :rows="5"/>
      </a-form-item>
    </a-form>
    <div class="node-address-style" v-if="settingData.peerId != ''">
      <span class="font-style">节点地址：</span>
      <span>{{"/ip4/127.0.0.1/tcp/" + settingData.port + "/p2p/" + settingData.peerId}}</span>
    </div>
    <div class="save-button-style">
      <a-button class="setting-button ok-btn" @click="setting">保存</a-button>
    </div>
  </div>
</template>

<script>
import {getCurrentInstance, onMounted, reactive, ref, toRefs} from "vue";

export default {
  name: "index",
  setup() {
    const { proxy } = getCurrentInstance();
    const settingState = ref()
    const state = reactive({
      settingForm: {
        publicKey: ""
      },
      settingData: {
        peerId: "",
        port: ""
      }
    })
    const settingRules = {
      publicKey: [
        {
          required: true,
          message: "请输入公钥",
          trigger: "blur",
        },
      ],
      privateKey: [
        {
          required: true,
          message: "请输入私钥",
          trigger: "blur",
        },
      ]
    }
    onMounted(() => {
      getSetting()
    })
    const getSetting = () => {
      window.backend.Setting.GetSetting().then(res => {
        state.settingForm.publicKey = res.PublicKey;
        state.settingData.peerId = res.PeerId;
        state.settingData.port = res.Port;
      })
    }
    const setting = () => {
      settingState.value.validate().then(() => {
        window.backend.Setting.Setting(state.settingForm.publicKey).then(res => {
          if (res) {
            proxy.$message.success("配置成功")
          }
        }).catch((err) => {
          proxy.$message.error(err)
        })
      }).catch(() => {
        return false
      })
    }
    return {
      ...toRefs(state),
      settingState,
      settingRules,
      setting
    }
  }
}
</script>

<style lang="scss" scoped>
.setting-container {
  padding: 12px;
  background: white;
  height: 100%;
  border-radius: 8px;
  .setting-button {
    width: 500px;
  }
  .save-button-style {
    display: flex;
    justify-content: center;
    margin-top: 20px;
  }
  .node-address-style {
    .font-style {
      font-size: 12px;
      font-weight: 600;
    }
  }
}
</style>
