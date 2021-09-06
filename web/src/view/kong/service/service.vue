<template>
  <div>
    <div v-if="tableVisible">
      <div class="button-box clearflex">
        <el-button size="mini" type="primary" icon="el-icon-plus" @click="addUser">添加Service</el-button>
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
export default {
  name: 'Api',
  components: { Tab },
  mixins: [infoList],
  data() {
    return {
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
  async created() {
    // const svcs = this.list
    //
    await this.listServices()
    this.setAuthorityIds()
    const res = await getAuthorityList({ page: 1, pageSize: 999 })
    this.setOptions(res.data.list)
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
    setAuthorityIds() {
      this.tableData && this.tableData.forEach((user) => {
        const authorityIds = user.authorities && user.authorities.map(i => {
          return i.authorityId
        })
        this.$set(user, 'authorityIds', authorityIds)
      })
    },
    openHeaderChange() {
      this.$refs.chooseImg.open()
    },
    setOptions(authData) {
      this.authOptions = []
      this.setAuthorityOptions(authData, this.authOptions)
    },
    setAuthorityOptions(AuthorityData, optionsData) {
      AuthorityData &&
        AuthorityData.map(item => {
          if (item.children && item.children.length) {
            const option = {
              authorityId: item.authorityId,
              authorityName: item.authorityName,
              children: []
            }
            this.setAuthorityOptions(item.children, option.children)
            optionsData.push(option)
          } else {
            const option = {
              authorityId: item.authorityId,
              authorityName: item.authorityName
            }
            optionsData.push(option)
          }
        })
    },
    async deleteUser(row) {
      const res = await deleteUser({ id: row.ID })
      if (res.code === 0) {
        this.$message.success('删除成功')
        await this.getTableData()
        this.setAuthorityIds()
        row.visible = false
      }
    },
    async enterAddUserDialog() {
      this.userInfo.authorityId = this.userInfo.authorityIds[0]
      this.$refs.userForm.validate(async valid => {
        if (valid) {
          const res = await register(this.userInfo)
          if (res.code === 0) {
            this.$message({ type: 'success', message: '创建成功' })
          }
          await this.getTableData()
          this.setAuthorityIds()
          this.closeAddUserDialog()
        }
      })
    },
    closeAddUserDialog() {
      this.$refs.userForm.resetFields()
      this.addUserDialog = false
    },
    addUser() {
      this.addUserDialog = true
    },
    async changeAuthority(row, flag) {
      if (flag) {
        return
      }
      this.$nextTick(async() => {
        const res = await setUserAuthorities({
          ID: row.ID,
          authorityIds: row.authorityIds
        })
        if (res.code === 0) {
          this.$message({ type: 'success', message: '角色设置成功' })
        }
      })
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
