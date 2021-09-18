<template>
  <el-form ref="upstream" :model="upstream" label-width="auto">
    <el-form-item label="Name">
      <el-input v-model="upstream.name" />
      <p class="desc-font">This is a hostname, which must be equal to the <span class="mark-font">host</span> of a Service.</p>
    </el-form-item>
    <el-form-item label="Host Header (optional)">
      <el-input v-model="upstream.host_header" />
    </el-form-item>
    <el-form-item label="Client Certificate">
      <el-input v-model="upstream.client_certificate.id" />
    </el-form-item>
    <el-form-item label="Algorithm">
      <el-select v-model="upstream.algorithm">
        <el-option v-for="item in algorithm_options" :key="item.value" :label="item.label" :value="item.value" />
      </el-select>
      <p class="desc-font">Which load balancing algorithm to use. Accepted values are: <span class="mark-font">consistent-hashing</span>, <span class="mark-font">least-connections</span>, <span class="mark-font">round-robin</span>. Default: <span class="mark-font">round-robin</span>.</p>
    </el-form-item>
    <el-form-item label="Slots">
      <el-input v-model="upstream.slots" type="number" />
      <p class="desc-font">The number of slots in the load balancer algorithm. If <span class="mark-font">algorithm</span> is set to <span class="mark-font">round-robin</span>, this setting determines the maximum number of slots. If <span class="mark-font">algorithm</span> is set to <span class="mark-font">consistent-hashing</span>, this setting determines the actual number of slots in the algorithm. Accepts an integer in the range <span class="mark-font">10-65536</span>. Default: <span class="mark-font">10000</span>.</p>
    </el-form-item>
    <el-form-item label="Hash On">
      <el-select v-model="upstream.hash_on">
        <el-option v-for="item in hash_on_options" :key="item.value" :label="item.label" :value="item.value" />
      </el-select>
      <p class="desc-font">What to use as hashing input. Using <span class="mark-font">none</span> results in a weighted-round-robin scheme with no hashing. Accepted values are: <span class="mark-font">none</span>, <span class="mark-font">consumer</span>, <span class="mark-font">ip</span>, <span class="mark-font">header</span>, <span class="mark-font">cookie</span>. Default: <span class="mark-font">none</span>.</p>
    </el-form-item>
    <el-form-item label="Hash Fallback">
      <el-select v-model="upstream.hash_fallback">
        <el-option v-for="item in hash_fallback_options" :key="item.value" :label="item.label" :value="item.value" />
      </el-select>
      <p class="desc-font">What to use as hashing input if the primary <span class="mark-font">hash_on</span> does not return a hash (eg. header is missing, or no Consumer identified). Not available if <span class="mark-font">hash_on</span> is set to <span class="mark-font">cookie</span>. Accepted values are: <span class="mark-font">none</span>, <span class="mark-font">consumer</span>, <span class="mark-font">ip</span>, <span class="mark-font">header</span>, <span class="mark-font">cookie</span>. Default: <span class="mark-font">none</span>.</p>
    </el-form-item>
    <el-form-item label="Hash On Header">
      <el-input v-model="upstream.hash_on_header" />
      <p class="desc-font">The header name to take the value from as hash input. Only required when <span class="mark-font">hash_on</span> is set to <span class="mark-font">header</span>.</p>
    </el-form-item>
    <el-form-item label="Hash Fallback Header">
      <el-input v-model="upstream.hash_fallback_header" />
      <p class="desc-font">The header name to take the value from as hash input. Only required when <span class="mark-font">hash_fallback</span> is set to <span class="mark-font">header</span>.</p>
    </el-form-item>
    <el-form-item label="Hash On Cookie">
      <el-input v-model="upstream.hash_on_cookie" />
      <p class="desc-font">The cookie name to take the value from as hash input. Only required when <span class="mark-font">hash_on</span> or <span class="mark-font">hash_fallback</span> is set to <span class="mark-font">cookie</span>. If the specified cookie is not in the request, Kong will generate a value and set the cookie in the response.</p>
    </el-form-item>
    <el-form-item label="Hash On Cookie Path">
      <el-input v-model="upstream.hash_on_cookie_path" />
      <p class="desc-font">The cookie path to set in the response headers. Only required when <span class="desc-font">hash_on</span> or <span class="desc-font">hash_fallback</span> is set to cookie. Default: <span class="mark-font">"/"</span>.</p>
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
    },
    op: {
      type: String,
      default: ''
    }
  },
  data() {
    return {
      hash_on_options: [
        {
          value: 'none',
          label: 'none'
        },
        {
          value: 'consumer',
          label: 'consumer'
        },
        {
          value: 'ip',
          label: 'ip'
        },
        {
          value: 'header',
          label: 'header'
        },
        {
          value: 'cookie',
          label: 'cookie'
        }
      ],

      hash_fallback_options: [
        {
          value: 'none',
          label: 'none'
        },
        {
          value: 'consumer',
          label: 'consumer'
        },
        {
          value: 'ip',
          label: 'ip'
        },
        {
          value: 'header',
          label: 'header'
        },
        {
          value: 'cookie',
          label: 'cookie'
        }
      ],

      algorithm_options: [
        {
          value: 'round-robin',
          label: 'round-robin'
        },
        {
          value: 'least-connections',
          label: 'least-connections'
        },
        {
          value: 'consistent-hashing',
          label: 'consistent-hashing'
        }
      ],

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
    console.log(this.op)
    if (this.op === 'edit') {
      await this.getUpstreamData(this.id)
    }
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
        name: item.name
      }

      // 编辑的话需要设置id
      if (this.op === 'edit') {
        data.id = item.id
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
  .span {
    font-size: 13px;
    color: #fa6863;
  }
</style>
