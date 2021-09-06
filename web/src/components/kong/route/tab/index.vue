<template>
  <el-tabs v-model="activeName" @tab-click="handleClick">
    <el-tab-pane label="Route" name="route">
      <RouteEditor />
    </el-tab-pane>
    <el-tab-pane label="Plugins" name="plugins">
      <PluginTable />
    </el-tab-pane>
  </el-tabs>
</template>

<script>
// 获取列表内容封装在mixins内部  getTableData方法 初始化已封装完成
const path = process.env.VUE_APP_BASE_API
import {
  getUserList
} from '@/api/user'

import {
  createRoute,
  createRouteInService,
  updateRoute,
  deleteRoute,
  getRoute,
  listRoute,
  listAllRoute,
  listAllRouteInService
} from '@/api/kongRoute'

import infoList from '@/mixins/infoList'
import { mapGetters } from 'vuex'
import RouteEditor from '@/components/kong/route/editor'
import PluginTable from '@/components/kong/plugin/table'
export default {
  name: 'RouteTab',
  components: { RouteEditor, PluginTable },
  mixins: [infoList],
  props: {
    sid: {
      type: String,
      default: ''
    }
  },
  data() {
    return {
      activeName: 'route',
      rs: {},
      service: {},
      plugins: [],
      tableData: [],
      listApi: getUserList,
      createRouteApi: createRoute,
      createRouteInServiceApi: createRouteInService,
      updateRouteApi: updateRoute,
      deleteRouteApi: deleteRoute,
      getRouteApi: getRoute,
      listRouteApi: listRoute,
      listAllRouteApi: listAllRoute,
      listAllRouteInServiceApi: listAllRouteInService,
      path: path,
      authOptions: [],
      addUserDialog: false
    }
  },
  computed: {
    ...mapGetters('user', ['token'])
  },
  async created() {
  },
  methods: {
    handleClick(tab, event) {
      console.log(tab, event)
    }
  }
}
</script>

<style lang="scss">
.demo-table-expand {
  font-size: 0;
}
.demo-table-expand label {
  width: 120px;
  color: #99a9bf;
}
.demo-table-expand .el-form-item {
  margin-right: 0;
  margin-bottom: 0;
  width: 50%;
}
</style>
