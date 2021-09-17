<template>
  <div>
    <el-dialog :before-close="closeDialog" :visible.sync="visible" title="客户">
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
        <el-form-item label="Tags">
          <el-tag v-for="tag in upstream.tags" :key="tag" closable :disable-transitions="false" @close="handleTagsClose(tag)">
            {{ tag }}
          </el-tag>
          <el-input v-if="inputTagsVisible" ref="saveTagsInput" v-model="inputTagValue" class="input-new-tag" size="small" @keyup.enter.native="handleTagsInputConfirm" @blur="handleTagsInputConfirm" />
          <el-button v-else class="button-new-tag" size="small" @click="showTagsInput">+ New Tag</el-button>
        </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click="closeDialog">Cancel</el-button>
        <el-button type="primary" @click="onSubmit">Create</el-button>
      </div>
    </el-dialog>
  </div>
</template>

<script>

import {
  createUpstream
} from '@/api/kongUpstream'

export default {
  props: {
    visible: {
      require: true,
      type: Boolean,
      default: false
    }
  },
  data() {
    return {
      inputTagsVisible: false,
      inputTagValue: '',
      upstream: {
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
  methods: {
    closeDialog() {
      this.$emit('closeHandler')
    },
    enterDialog() {
      console.log('enter')
    },
    async onSubmit() {
      const item = this.upstream

      const data = {
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

      createUpstream(data).then((rs) => {
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
  .el-tag + .el-tag {
    margin-left: 10px;
  }
  .button-new-tag {
    margin-left: 10px;
    height: 32px;
    /* line-height: 30px; */
    padding-top: 0;
    padding-bottom: 0;
  }
  .input-new-tag {
    width: 120px;
    margin-left: 10px;
    vertical-align: bottom;
  }
</style>
