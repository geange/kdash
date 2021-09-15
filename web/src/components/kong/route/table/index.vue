<template>
  <div>
    <div class="button-box clearflex">
      <el-button size="mini" type="primary" icon="el-icon-plus" @click="add">ADD</el-button>
    </div>
    <el-table :data="routes" style="width: 100%">
      <el-table-column label="ID">
        <template slot-scope="item">
          <el-button type="text">{{ item.row.id }}</el-button>
        </template>
      </el-table-column>
      <el-table-column label="Name" prop="name" />
      <el-table-column label="Paths" prop="paths">
        <template slot-scope="props">
          <el-tag v-for="path in props.row.paths" :key="path" size="mini">{{ path }}</el-tag>
        </template>
      </el-table-column>
      <el-table-column label="Tags" prop="tags">
        <template slot-scope="props">
          <el-tag v-for="tag in props.row.tags" :key="tag" size="mini">{{ tag }}</el-tag>
        </template>
      </el-table-column>
    </el-table>
    <Dialog :visible="dialogFormVisible" @closeHandler="closeDialog" />
  </div>
</template>

<script>
import {
  listAllRouteInService
} from '@/api/kongRoute'

import Dialog from '@/components/kong/route/dialog'

export default {
  name: 'RouteTable',
  components: { Dialog },
  props: {
    sid: {
      type: String,
      default: ''
    },
    rtype: {
      type: String,
      default: ''
    }
  },
  data() {
    return {
      dialogFormVisible: false,
      routes: [
        {
          created_at: 0,
          hosts: ['a.cn', 'b.cn'],
          headers: {
            'name': ['a', 'b']
          },
          id: 'sssssssss',
          name: 'ss',
          methods: ['GET'],
          paths: ['/abssssssssssssssssssssssssc9', '/abssssssssssssssssssssssssc'],
          path_handling: 'ss',
          preserve_host: 'sss.cn',
          protocols: ['http', 'https'],
          regex_priority: 0,
          service: {
            id: 'sssssffffff'
          },
          strip_path: false,
          updated_at: 0,
          snis: ['a1', 'a2'],
          sources: [{
            ip: '',
            port: 8000
          }],
          destinations: [],
          tags: ['a', 'b'],
          https_redirect_status_code: 302,
          request_buffering: false,
          response_buffering: false
        }
      ]
    }
  },
  created() {
  },
  mounted() {
    console.log('create routes', this.sid, this.rtype)
    console.log('======>', this.rtype)
    if (this.rtype === 'service') {
      this.getRoutesByService(this.sid)
    }
  },
  methods: {
    add() {
      console.log('add new route')
      this.dialogFormVisible = true
    },
    closeDialog() {
      this.dialogFormVisible = false
    },
    async getRoutesByService(id) {
      console.log('get route')
      const params = { id: id }
      listAllRouteInService(params).then((rs) => {
        console.log(rs)
        if (rs.msg === 'listRouteError') {
          this.routes = []
        }
      }).catch(err => {
        this.$message({ type: 'error', message: err })
        this.routes = []
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
    width: 90px;
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
