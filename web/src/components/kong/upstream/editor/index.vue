<template>
  <el-form ref="upstream" :model="upstream" label-width="140px" label-position="left">
    <el-form-item label="Name">
      <el-input v-model="upstream.name" />
    </el-form-item>
    <el-form-item label="HostHeader">
      <el-input v-model="upstream.host_header" />
    </el-form-item>
    <el-form-item label="ClientCertificate">
      <el-input v-model="upstream.client_certificate.id" />
    </el-form-item>
    <el-form-item label="Algorithm">
      <el-input v-model="upstream.algorithm" />
    </el-form-item>
    <el-form-item label="Slots">
      <el-input v-model="upstream.slots" />
    </el-form-item>
    <el-form-item label="HashOn">
      <el-input v-model="upstream.hash_on" />
    </el-form-item>
    <el-form-item label="HashFallback">
      <el-input v-model="upstream.hash_fallback" />
    </el-form-item>
    <el-form-item label="HashOnHeader">
      <el-input v-model="upstream.hash_on_header" />
    </el-form-item>
    <el-form-item label="HashFallbackHeader">
      <el-input v-model="upstream.hash_fallback_header" />
    </el-form-item>
    <el-form-item label="HashOnCookie">
      <el-input v-model="upstream.hash_on_cookie" />
    </el-form-item>
    <el-form-item label="HashOnCookiePath">
      <el-input v-model="upstream.hash_on_cookie_path" />
    </el-form-item>
    <el-form-item>
      <el-button type="primary" @click="onSubmit">立即创建</el-button>
      <el-button>取消</el-button>
    </el-form-item>
  </el-form>
</template>

<script>
export default {
  name: 'UpstreamTable',
  data() {
    return {
      upstream: {
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
    }
  },
  created() {
    console.log('create')
  },
  methods: {
    onSubmit() {
    }
  }
}
</script>

<style>
  .el-input {
    width: 300px;
  }
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
