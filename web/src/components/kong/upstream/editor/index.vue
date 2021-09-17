<template>
  <el-form ref="upstream" :model="upstream" label-width="140px" label-position="left">
    <el-form-item label="ID">
      <el-input v-model="upstream.id" />
    </el-form-item>
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
    <el-form-item label="Tags">
      <el-tag v-for="tag in upstream.tags" :key="tag" closable :disable-transitions="false" @close="handleTagsClose(tag)">
        {{ tag }}
      </el-tag>
      <el-input v-if="inputTagsVisible" ref="saveTagsInput" v-model="inputTagValue" class="input-new-tag" size="small" @keyup.enter.native="handleTagsInputConfirm" @blur="handleTagsInputConfirm" />
      <el-button v-else class="button-new-tag" size="small" @click="showTagsInput">+ New Tag</el-button>
    </el-form-item>
    <el-form-item>
      <el-button type="primary" @click="onSubmit">Update</el-button>
      <el-button>Cancel</el-button>
    </el-form-item>
  </el-form>
</template>

<script>
import {
  getUpstream,
  updateUpstream
} from '@/api/kongUpstream'

export default {
  name: 'Editor',
  props: {
    id: {
      type: String,
      default: ''
    }
  },
  data() {
    return {
      inputTagsVisible: false,
      inputTagValue: '',
      upstream: {
        id: '',
        name: '',
        host_header: '',
        client_certificate: {
          id: ''
        },
        algorithm: '',
        slots: 0,
        healthchecks: {
          threshold: 0,
          active: {
            concurrency: 0,
            healthy: {
              http_statuses: [],
              interval: 0,
              successes: 0
            },
            http_path: '',
            https_sni: '',
            https_verify_certificate: false,
            type: '',
            timeout: 0,
            unhealthy: {
              http_failures: 0,
              http_statuses: [],
              tcp_failures: 0,
              timeouts: 0,
              interval: 0
            }
          },
          passive: {
            healthy: {
              http_statuses: [],
              interval: 0,
              successes: 0
            },
            type: '',
            unhealthy: {
              http_failures: 0,
              http_statuses: [],
              tcp_failures: 0,
              timeouts: 0,
              interval: 0
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
        tags: []
      }
    }
  },
  created() {
  },
  async mounted() {
    await this.getUpstreamData(this.id)
  },
  methods: {
    async getUpstreamData(id) {
      const params = { id: id }
      getUpstream(params).then((rs) => {
        const item = rs.data
        this.upstream.id = item.id
        this.upstream.name = item.name
        this.upstream.host_header = item.host_header

        item.client_certificate !== undefined && (this.upstream.client_certificate.id = item.client_certificate.id)
        item.algorithm !== undefined && (this.upstream.algorithm = item.algorithm)
        item.slots !== undefined && (this.upstream.slots = item.slots)

        item.hash_on !== '' && (this.upstream.hash_on = item.hash_on)
        item.hash_fallback !== '' && (this.upstream.hash_fallback = item.hash_fallback)
        item.hash_on_header !== '' && (this.upstream.hash_on_header = item.hash_on_header)
        item.hash_fallback_header !== '' && (this.upstream.hash_fallback_header = item.hash_fallback_header)
        item.hash_on_cookie !== '' && (this.upstream.hash_on_cookie = item.hash_on_cookie)
        item.hash_on_cookie_path !== '' && (this.upstream.hash_on_cookie_path = item.hash_on_cookie_path)
        item.tags !== [] && (this.upstream.tags = item.tags)
      }).catch(err => {
        this.$message({ type: 'error', message: err })
      })
    },
    async onSubmit() {
      const item = this.upstream

      const data = {
        id: item.id,
        name: item.name
      }

      item.host_header !== '' && (data.host_header = item.host_header)
      item.algorithm !== '' && (data.algorithm = item.algorithm)
      item.client_certificate.id !== '' && (data.client_certificate.id = item.client_certificate.id)
      item.hash_on !== '' && (data.hash_on = item.hash_on)
      item.hash_fallback !== '' && (data.hash_fallback = item.hash_fallback)
      item.hash_on_header !== '' && (data.hash_on_header = item.hash_on_header)
      item.hash_fallback_header !== '' && (data.hash_fallback_header = item.hash_fallback_header)
      item.hash_on_cookie !== '' && (data.hash_on_cookie = item.hash_on_cookie)
      item.hash_on_cookie_path !== '' && (data.hash_on_cookie_path = item.hash_on_cookie_path)
      item.tags !== [] && (data.tags = item.tags)

      updateUpstream(data).then((rs) => {
        console.log('create upstream')
        this.$emit('closeHandler')
        this.$emit('updateHandler')
      }).catch(err => {
        this.$message({ type: 'error', message: err })
      })
    },
    handleTagsClose(tag) {
      this.upstream.tags.splice(this.upstream.tags.indexOf(tag), 1)
    },
    showTagsInput() {
      this.inputTagsVisible = true
      this.$nextTick(_ => {
        this.$refs.saveTagsInput.$refs.input.focus()
      })
    },
    handleTagsInputConfirm() {
      if (this.inputTagValue) {
        if (this.upstream.tags === undefined) {
          this.upstream.tags = []
        }
        this.upstream.tags.push(this.inputTagValue)
      }
      this.inputTagsVisible = false
      this.inputTagValue = ''
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
