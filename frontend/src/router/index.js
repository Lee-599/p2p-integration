import {createMemoryHistory, createRouter} from 'vue-router'
import Layout from "../layouts/index"
import Resource from "../views/resource/index"
//import Setting from "../views/setting/index"
import Status from "../views/status/index"
import Share from "../views/share"
import Login from "../views/login"

const routes = [
  {
    path: '/',
    name: 'Home',
    showLevelFlag: false,
    isShow: false,
    component: Layout,
    redirect: "/resource",
    children: [
      {
        path: '/resource',
        name: 'resource',
        showLevelFlag: false,
        isShow: true,
        meta: {
          title: "资源"
        },
        // You can only use pre-loading to add routes, not the on-demand loading method.
        component: Resource
      },
      // {
      //   path: '/setting',
      //   name: 'setting',
      //   showLevelFlag: false,
      //   isShow: true,
      //   meta: {
      //     title: "设置"
      //   },
      //   // You can only use pre-loading to add routes, not the on-demand loading method.
      //   component: Setting
      // },
      {
        path: '/status',
        name: 'status',
        showLevelFlag: false,
        isShow: true,
        meta: {
          title: "状态"
        },
        // You can only use pre-loading to add routes, not the on-demand loading method.
        component: Status
      },
      {
        path: '/share',
        name: 'share',
        showLevelFlag: false,
        isShow: true,
        meta: {
          title: "共享"
        },
        // You can only use pre-loading to add routes, not the on-demand loading method.
        component: Share
      }
    ]
  },
  {
    path: "/login",
    name: "login",
    component: Login
  }
]

const router = createRouter({
  mode: "abstract",
  history: createMemoryHistory(),
  routes
})

export default router
