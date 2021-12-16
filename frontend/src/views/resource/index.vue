<template>
  <a-spin :tip="'加载中'" :spinning="loadLoading" class="spin-style">
    <div class="tableStyle">
      <no-data-table v-if="list.data.length === 0">
        <a-table-column
            title="资源ID"
            data-index="ID"
            :width="150"
        />
        <a-table-column title="PeerId" data-index="peerId" />
        <a-table-column
            title="机器名"
            data-index="creator"
        />
        <a-table-column title="系统类型" data-index="vmType" />
        <a-table-column title="操作" :width="100"></a-table-column>
      </no-data-table>
      <p-table
          v-show="list.data.length != 0"
          :row-key="t => t.ID"
          :dataList="list"
          @get-data="getList"
          ref="table-page"
      >
        <a-table-column
            title="资源ID"
            data-index="ID"
            :width="150"
        />
        <a-table-column title="PeerId" data-index="peerId" />
        <a-table-column
            title="机器名"
            :ellipsis="true"
            data-index="creator"/>

        <a-table-column title="操作" :width="100">
          <template #default="{ record }">
            <a style="color: #4850FF" @click="link(record)">连接</a>
          </template>
        </a-table-column>
      </p-table>
    </div>
  </a-spin>
  <tip-modal :title="'消息提醒'" :is-show-cancel="true" ref="loginRef" :visible="visible" :title-big-style="false" :close="close">
    <div class="tip-content">
      <span class="font-style">您未配置公钥，请前往"设置"配置公钥</span>
    </div>
    <div class="tip-button">
      <a-button class="ok-btn" @click="goSetting">前往配置</a-button>
    </div>
  </tip-modal>
  <link-modal :title="'连接'" :is-show-cancel="true" ref="linkRef" :visible="linkVisible" :title-big-style="false" :close="linkClose" :loading="loading">
    <div style="margin-top: 32px">
      <a-form ref="linkState" :model="linkForm" :rules="linkRules">
        <a-form-item name="peerId">
          <a-input
              v-model:value="linkForm.peerId"
              type="text"
              autocomplete="off"
          />
        </a-form-item>
        <a-form-item name="port">
          <a-input
              v-model:value="linkForm.port"
              type="text"
              autocomplete="off"
              placeholder="请输入端口号"
          />
        </a-form-item>
        <a-button class="ok-btn" block @click="toLink">
          连接
        </a-button>
      </a-form>
    </div>
  </link-modal>
  <link-tip :title="'初始化配置'" :is-show-cancel="true" :visible="linkTipVisible" :title-big-style="false" :close="linkTipClose" :loading="tipLoading">
    <div style="margin-top: 32px;margin-bottom: 20px">
      <span>您未初始化配置，请初始化配置</span>
    </div>
    <a-button class="ok-btn" block @click="setting">
      初始化配置
    </a-button>
  </link-tip>
</template>

<script>
import {getCurrentInstance, onMounted, reactive, ref, toRefs} from "vue";
import NoDataTable from "../../components/table/NoDataTable";
import PTable from "../../components/table/PTable";
import {timeToDay} from "../../utils/util";
import TipModal from "../../components/model/index"
import {useRouter} from "vue-router";
import LinkModal from "../../components/model/index"
import LinkTip from "../../components/model/index"

export default {
  name: "index",
  components: {
    NoDataTable,
    PTable,
    TipModal,
    LinkModal,
    LinkTip
  },
  setup() {
    const router = useRouter();
    const linkState = ref();
    const { proxy } = getCurrentInstance();
    const state = reactive({
      loadLoading:false,
      tipLoading: false,
      loading: false,
      visible: false,
      linkVisible: false,
      linkTipVisible: false,
      linkForm: {
        port:"",
        peerId: ""
      },
      list: {
        total: 0,
        data: [],
        current: 1
      },
    })
    const linkRules = {

    }
    //p2p连接
    const toLink = () => {
      linkState.value.validate().then(() => {
        state.loading = true
        window.backend.P2p.Link(parseInt(state.linkForm.port),state.linkForm.peerId).then(() => {
          state.loading = false;
          proxy.$message.success("连接成功")
          state.linkVisible = false
        }).catch((err) => {
          proxy.$message.error(err)
          state.loading = false
        })
      }).catch(() => {
        return false
      })
    }
    const getList = () => {

    }
    const link = (record) => {
      //查询是否P2p配置
      window.backend.P2p.IsP2PSetting().then(res => {
        if (res) {
          state.linkForm.peerId = record.peerId
          state.linkVisible = true
        }else {
          //初始化配置
          state.linkTipVisible = true
        }
      })
    }
    onMounted(() => {
      isSettingPublicKey();
    })
    const getResourceList = () => {
      state.loadLoading = true
      window.backend.Resource.GetResources().then(res => {
        state.list.data = res
        state.loadLoading = false
      }).catch(() => {
        state.loadLoading = false
      })
    }
    const close = () => {
      state.visible = false
    }
    const linkClose = () => {
      state.linkVisible = false
    }
    const linkTipClose = () => {
      state.linkTipVisible = false;
    }
    const isSettingPublicKey = () => {
      window.backend.Account.IsAccountSetting().then(res => {
        if (res) {
          getResourceList()
        }else {
          state.visible = true
        }
      })
    }
    //p2p初始化配置
    const setting = () => {
      state.tipLoading = true
      window.backend.Setting.InitP2pSetting().then(() => {
        proxy.$message.success("配置成功")
        state.tipLoading = false
        state.linkTipVisible = false
      }).catch(() => {
        proxy.$message.error("配置失败")
        state.tipLoading = false
      })
    }
    const goSetting = () => {
      router.push("/setting")
    }
    return {
      ...toRefs(state),
      getList,
      link,
      getResourceList,
      timeToDay,
      close,
      goSetting,
      isSettingPublicKey,
      linkClose,
      linkState,
      linkRules,
      toLink,
      linkTipClose,
      setting
    }
  }
}
</script>

<style lang="scss" scoped>
.tableStyle {
  padding: 12px;
  background: white;
  border-radius: 8px;
  height: 100%;
}
.tip-content {
  margin-top: 20px;
  .font-style {
    color: #58637B;
    font-size: 14px;
    line-height: 17px;
  }
}
.tip-button {
  margin-top: 20px;
}
</style>
