<template>
  <div class="home">
    <a-table
        bordered
        :dataSource="dataSource"
        :columns="columns"
        :pagination="false"
    >
      <template #operation="{ record }">
          <UseResource v-if="record.status===0" :id="record.ID">使用</UseResource>
          <a v-if="record.status===1" @click="deleteUse(record.ID)">移除</a>
      </template>
    </a-table>
  </div>
</template>

<script>
import {onMounted, ref} from "vue";
import * as Wails from '@wailsapp/runtime';
import UseResource from "@/components/UseResource.vue";
import {message} from "ant-design-vue";
export default {
  name: "Home",
  components: {
    UseResource
  },
  setup() {
    let dataSource = ref([])
    onMounted(() =>{
      Wails.Events.On("resources",t=>{
        dataSource.value = t.map(t => {
          t.key = t.ID
          return t
        })
      })
    })
    function deleteUse(id){
      window.backend.WailsApi.DeleteUseResource(id).then(() => {
        message.success('操作成功');
      });
    }
    return {
      dataSource,
      columns: [
        {
          title: '资源ID',
          dataIndex: 'ID',
          key: 'ID',
        },
        {
          title: '资源节点ID',
          dataIndex: 'peerId',
          key: 'peerId',
        },
        {
          title: 'cpu核',
          dataIndex: 'cpu',
          key: 'cpu',
        },
        {
          title: '内存',
          dataIndex: 'memory',
          key: 'memory',
        },
        {
          title: '系统类型',
          dataIndex: 'vmType',
          key: 'vmType',
        },
        {
          title: '系统镜像',
          dataIndex: 'systemImage',
          key: 'systemImage',
        },
        {
          title: '创建人',
          dataIndex: 'creator',
          key: 'creator',
        },
        {
          title: '使用人',
          dataIndex: 'user',
          key: 'user',
        },
        {
          title: '系统状态',
          dataIndex: 'status',
          key: 'status',
        },
        {
          title: '过期时间',
          dataIndex: 'expireTime',
          key: 'expireTime',
        },
        {
          title: '操作',
          dataIndex: 'operation',
          key: 'operation',
          slots: {
            customRender: 'operation',
          },
        },
      ],
      deleteUse
    };
  },
};
</script>
