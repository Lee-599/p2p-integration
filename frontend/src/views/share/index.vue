<template>
  <div>
    <h1>共享本地端口</h1>
    <a-form :label-col="{ span: 5 }" :wrapper-col="{ span: 12 }">
      <a-form-item label="共享的本地端口">
        <a-input-number v-model:value="port" :min="1" :max="65535" v-model:disabled="allow"/>
      </a-form-item>
      <a-form-item label="显示描述信息">
        <a-input v-model:value="hostname" v-model:disabled="allow"/>
      </a-form-item>
      <a-form-item :wrapper-col="{ span: 12, offset: 5 }">
        <a-switch v-model:checked="allow" checked-children="开" un-checked-children="关" @change="onChange"  />
      </a-form-item>
  </a-form>
  </div>
</template>

<script>
import {reactive, toRefs, getCurrentInstance, onMounted} from "vue";

export default {
  name: "share",
  setup(){
    const { proxy } = getCurrentInstance();

    const state = reactive({
      allow: false,
      hostname: '',
      port: 8080,
    })

    onMounted(() => {

      window.backend.P2p.IsSharing().then((result) => {
        proxy.allow = result.IsShare
        if(proxy.allow) {
          proxy.port = result.Port
        }
      })
    })

    const onChange = (checked) => {
      let target = `/ip4/127.0.0.1/tcp/${proxy.port}`
      console.log(checked,target)
      if(checked) {
        //开启p2p端口
        window.backend.P2p.Share(target,proxy.hostname)
            .then(() => {
              proxy.$message.success("打开p2p端口成功")
            })
            .catch(() => {
              proxy.$message.error("打开p2p端口失败")
            })
      }else {
        window.backend.P2p.UnShare().then(() => {
          proxy.$message.success("关闭p2p端口成功")
        }).catch(() => {
          proxy.$message.error("关闭p2p端口失败")
        })
      }
    }

    return {
      ...toRefs(state),
      onChange,
    }
  }
};
</script>