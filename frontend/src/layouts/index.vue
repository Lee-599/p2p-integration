<template>
  <div class="home">
    <div class="home-container">
      <a-layout class="layout-container">
        <a-layout-sider width="142px" class="menu-left-style">
          <div class="logo">
            <img src="@/assets/logo.png" />
          </div>
          <div class="custom-menu">
            <Menu />
          </div>
        </a-layout-sider>
        <a-layout>
          <a-layout-header class="header">
            <Header ref="pcHeader"> </Header>
          </a-layout-header>
          <a-layout-content class="body-content">
            <router-view></router-view>
          </a-layout-content>
        </a-layout>
      </a-layout>
    </div>
  </div>
</template>

<script>
import Header from "./Header";
import Menu from "./Menu";
import {onMounted} from "vue";
import {useRouter} from "vue-router";
export default {
  name: "MobileLayout",
  components: {
    Header,
    Menu,
  },
  setup() {
    const router = useRouter()
    onMounted(() => {
      isExitAccount()
    })
    const isExitAccount = () => {
      window.backend.Account.IsAccount().then(res => {
        if (!res) {
          router.push("/login")
        }
      })
    }
  }
};
</script>

<style lang="scss" scoped>
//.home-container {
//  height: 100vh;
//}
.body-content {
  //height: calc(100vh - 64px);
  overflow-x: hidden;
  overflow-y: scroll;
  padding: 12px;
  background: #f6f6f9;
}
.header {
  height: 64px;
  background-color: #ffffff;
}
.logo {
  cursor: pointer;
  color: #ffffff;
  height: 20px;
  margin: 32px 27px 32px 27px;
  img {
    height: 100%;
  }
}
.custom-menu {
  height: calc(100vh - 84px);
  overflow-y: scroll;
}
.menu-left-style {
  width: 142px;
  background: #202946;
}
.ant-layout-header {
  padding: 0;
}
</style>
