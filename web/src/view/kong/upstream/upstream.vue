<template>
  <div>
    <div v-if="tableVisible">
      <div class="button-box clearflex">
        <el-button size="mini" type="primary" icon="el-icon-plus" @click="add">ADD</el-button>
      </div>
      <el-table :data="upstreams" style="width: 100%">
        <el-table-column label="ID">
          <template slot-scope="item">
            <el-button type="text" @click="openTab(item.row.id)">{{ item.row.id }}</el-button>
          </template>
        </el-table-column>
        <el-table-column label="Name" prop="name" />
        <el-table-column label="Tags" prop="tags">
          <template slot-scope="props">
            <el-tag v-for="tag in props.row.tags" :key="tag" size="mini">{{ tag }}</el-tag>
          </template>
        </el-table-column>
      </el-table>
    </div>
    <Dialog :visible="dialogFormVisible" @closeHandler="closeDialog" @updateHandler="getTableData" />
    <div v-if="tabVisible">
      <el-page-header title="back" content="Upstream" @back="goBack" />
      <Tab :id="id" style="margin-top: 20px" />
    </div>
  </div>
</template>

<script>

import {
  listAllUpstream
} from '@/api/kongUpstream'

import Tab from '@/components/kong/upstream/tab'
import Dialog from '@/components/kong/upstream/dialog'

export default {
  name: 'Upstream',
  components: { Tab, Dialog },
  data() {
    return {
      dialogFormVisible: false,
      id: '',
      tableVisible: true,
      tabVisible: false,
      upstreams: [
        {
          id: 'xxxxxxxxx',
          name: 'xxxx',
          host_header: 'XHHH',
          client_certificate: {
            id: '45666dd'
          },
          algorithm: 'sss',
          slots: 0,
          healthchecks: {
            threshold: 1.8,
            active: {
              concurrency: 0,
              healthy: {
                http_statuses: [200, 201],
                interval: 10,
                successes: 6
              },
              http_path: '',
              https_sni: '',
              https_verify_certificate: false,
              type: 'no',
              timeout: 10,
              unhealthy: {
                http_failures: 5,
                http_statuses: [404, 500],
                tcp_failures: 10,
                timeouts: 10,
                interval: 5
              }
            },
            passive: {
              healthy: {
                http_statuses: [200, 201],
                interval: 10,
                successes: 6
              },
              type: 'no',
              unhealthy: {
                http_failures: 5,
                http_statuses: [404, 500],
                tcp_failures: 10,
                timeouts: 10,
                interval: 5
              }
            }
          },
          created_at: 0,
          hash_on: '',
          hash_fallback: '',
          hash_on_header: '',
          hash_fallback_header: '',
          hash_on_cookie: '',
          hash_on_cookie_path: '',
          tags: ['a', 'b']
        }
      ]
    }
  },
  created() {
    console.log('create')
  },
  async mounted() {
    await this.getTableData()
  },
  methods: {
    goBack() {
      this.tableVisible = true
      this.tabVisible = false
    },
    openTab(id) {
      this.id = id
      this.tableVisible = false
      this.tabVisible = true
    },
    add() {
      console.log('add new route')
      this.dialogFormVisible = true
    },
    closeDialog() {
      this.dialogFormVisible = false
    },
    async getTableData() {
      listAllUpstream().then((rs) => {
        if (rs.data == null) {
          this.upstreams = []
        } else {
          this.upstreams = rs.data.list
        }
      }).catch(err => {
        console.log(err)
      })
    }
  }
}
</script>

<style>
  .demo-table-expand {
    font-size: 0;
  }
  .demo-table-expand label {
    width: 100px;
    color: #99a9bf;
  }
  .demo-table-expand .el-form-item {
    margin-right: 0;
    margin-bottom: 0;
    width: 50%;
  }
  .el-tag + .el-tag {
    margin-left: 10px;
  }
</style>
