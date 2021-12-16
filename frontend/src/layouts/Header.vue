<template>
  <div class="header">
    <div>
      <a-button class="logout" @click="logout">登出</a-button>
    </div>
  </div>
</template>

<script>
import {getCurrentInstance, reactive, toRefs} from "vue";
import {useRouter} from "vue-router";
export default {
  name: "MobileHeader",
  components: {
  },
  setup() {
    const router = useRouter();
    const { proxy } = getCurrentInstance();
    const state = reactive({
      headerInfo: {}
    });
    const logout = () => {
      window.backend.Login.Logout().then(() => {
        router.push("/login")
      }).catch((err) => {
        proxy.$message.error(err)
      })
    }
    return {
      ...toRefs(state),
      logout
    };
  },
};
</script>

<style lang="scss">
.header {
  display: flex;
  align-items: center;
  justify-content: flex-end;
  .logout {
    font-weight: 600;
    font-size: 12px;
    line-height: 17px;
    color: #4850FF;
    padding: 0px;
    margin-right: 23px;
  }
}
</style>
