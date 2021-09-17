<template>
  <el-tabs v-model="activeName" @tab-click="handleClick">
    <el-tab-pane label="Upstream" name="upstream">
      <UpstreamEditor :id="id" />
    </el-tab-pane>
    <el-tab-pane label="Targets" name="targets">
      <TargetTable />
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
// import ServiceEditor from '@/components/kong/service/editor'
// import RouteTable from '@/components/kong/route/table'
// import PluginTable from '@/components/kong/plugin/table'
import UpstreamEditor from '@/components/kong/upstream/editor'
import TargetTable from '@/components/kong/target/table'
export default {
  name: 'UpstreamTab',
  components: { UpstreamEditor, TargetTable },
  mixins: [infoList],
  props: {
    id: {
      type: String,
      default: ''
    }
  },
  data() {
    return {
      activeName: 'upstream',
      list_opts: {
        pre: {
          offset: 0,
          size: 30,
          tags: []
        },
        next: {
          offset: 0,
          size: 30,
          tags: []
        },
        is_begin: true,
        is_end: false
      },
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
      addUserDialog: false,
      userInfo: {
        username: '',
        password: '',
        nickName: '',
        headerImg: '',
        authorityId: '',
        authorityIds: []
      },
      rules: {
        username: [
          { required: true, message: '请输入用户名', trigger: 'blur' },
          { min: 5, message: '最低5位字符', trigger: 'blur' }
        ],
        password: [
          { required: true, message: '请输入用户密码', trigger: 'blur' },
          { min: 6, message: '最低6位字符', trigger: 'blur' }
        ],
        nickName: [
          { required: true, message: '请输入用户昵称', trigger: 'blur' }
        ],
        authorityId: [
          { required: true, message: '请选择用户角色', trigger: 'blur' }
        ]
      }
    }
  },
  computed: {
    ...mapGetters('user', ['token'])
  },
  async created() {
    console.log(this.id)
  },
  methods: {
    handleClick(tab, event) {
    }
  }
}
</script>

<style lang="scss">
.button-box {
  padding: 10px 20px;
  .el-button {
    float: right;
  }
}

.user-dialog {
  .header-img-box {
  width: 200px;
  height: 200px;
  border: 1px dashed #ccc;
  border-radius: 20px;
  text-align: center;
  line-height: 200px;
  cursor: pointer;
}
  .avatar-uploader .el-upload:hover {
    border-color: #409eff;
  }
  .avatar-uploader-icon {
    border: 1px dashed #d9d9d9 !important;
    border-radius: 6px;
    font-size: 28px;
    color: #8c939d;
    width: 178px;
    height: 178px;
    line-height: 178px;
    text-align: center;
  }
  .avatar {
    width: 178px;
    height: 178px;
    display: block;
  }
}

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
