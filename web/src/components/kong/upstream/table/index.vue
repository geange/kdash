<template>
  <el-table :data="routes" style="width: 100%">
    <el-table-column label="ID">
      <template slot-scope="item">
        <el-button type="text">{{ item.row.id }}</el-button>
      </template>
    </el-table-column>
    <el-table-column label="Name" prop="name" />
    <el-table-column label="Tags" prop="tags">
      <template slot-scope="props">
        <el-tag v-for="tag in props.row.tags" :key="tag" size="mini">{{ tag }}</el-tag>
      </template>
    </el-table-column>
  </el-table>
</template>

<script>
export default {
  name: 'UpstreamTable',
  data() {
    return {
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
