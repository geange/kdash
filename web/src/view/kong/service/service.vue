<template>
  <div>
    <div v-if="tableVisible">
      <div class="button-box clearflex">
        <el-button size="mini" type="primary" icon="el-icon-plus" @click="add">ADD</el-button>
      </div>
      <el-table :data="tableData">
        <el-table-column label="ID" min-width="250">
          <template slot-scope="scope">
            <el-button type="text" @click="openServiceInstance(scope.row.id)">{{ scope.row.id }}</el-button>
          </template>
        </el-table-column>
        <el-table-column label="Name" min-width="100" prop="name" />
        <el-table-column label="Host" min-width="100" prop="host" />
        <el-table-column label="Port" min-width="100" prop="port" />
        <el-table-column label="Tags" prop="tags">
          <template slot-scope="props">
            <el-tag v-for="tag in props.row.tags" :key="tag" size="mini">{{ tag }}</el-tag>
          </template>
        </el-table-column>
      </el-table>
      <Dialog :visible="dialogFormVisible" @closeHandler="closeDialog" @updateHandler="listServices" />
      <el-pagination
        :current-page="page"
        :page-size="pageSize"
        :page-sizes="[10, 30, 50, 100]"
        :style="{float:'right',padding:'20px'}"
        :total="total"
        layout="total, sizes, prev, pager, next, jumper"
        @current-change="handleCurrentChange"
        @size-change="handleSizeChange"
      />
    </div>
    <div v-if="tabVisible">
      <el-page-header title="back" content="Service" @back="goBack" />
      <Tab :sid="sid" style="margin-top: 20px" />
    </div>
  </div>
</template>

<script>
// 获取列表内容封装在mixins内部  getTableData方法 初始化已封装完成
const path = process.env.VUE_APP_BASE_API
import {
  getUserList,
  setUserAuthorities,
  register,
  deleteUser
} from '@/api/user'

import {
  createService,
  updateService,
  deleteService,
  getService,
  listService,
  listAllService
} from '@/api/kongService'

import { getAuthorityList } from '@/api/authority'
import infoList from '@/mixins/infoList'
import { mapGetters } from 'vuex'

import Tab from '@/components/kong/service/tab'
import Dialog from '@/components/kong/service/dialog'

export default {
  name: 'Api',
  components: { Tab, Dialog },
  mixins: [infoList],
  data() {
    return {
      dialogFormVisible: false,
      tableVisible: true,
      tabVisible: false,
      sid: '',
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
      tableData: [],
      listApi: getUserList,
      createServiceApi: createService,
      updateServiceApi: updateService,
      deleteServiceApi: deleteService,
      getServiceApi: getService,
      listServiceApi: listService,
      listAllServiceApi: listAllService,
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
      }
    }
  },
  computed: {
    ...mapGetters('user', ['token'])
  },
  async mounted() {
    await this.listServices()
  },
  methods: {
    goBack() {
      console.log('go back')
      this.sid = ''
      this.tableVisible = true
      this.tabVisible = false
    },
    openServiceInstance(id) {
      this.sid = id
      this.tableVisible = false
      this.tabVisible = true
    },
    async listServices() {
      const params = this.list_opts.next
      this.listAllServiceApi(params).then((rs) => {
        this.tableData = rs.data.list
        if (rs.data.opts === undefined) {
          this.list_opts.is_end = true
        } else {
          this.list_opts.pre = this.list_opts.next
          this.list_opts.next = rs.data.opts
        }
      }).catch(err => {
        console.log(err)
      })
    },
    add() {
      console.log('add new route')
      this.dialogFormVisible = true
    },
    closeDialog() {
      this.dialogFormVisible = false
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
</style>
