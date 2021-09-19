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
      <p class="desc-font">The cookie path to set in the response headers. Only required when <span class="mark-font">hash_on</span> or <span class="mark-font">hash_fallback</span> is set to cookie. Default: <span class="mark-font">"/"</span>.</p>
    </el-form-item>
    <el-form-item label="Tags">
      <el-tag v-for="tag in upstream.tags" :key="tag" closable :disable-transitions="false" @close="handleTagsClose(tag)">
        {{ tag }}
      </el-tag>
      <el-input v-if="inputTagsVisible" ref="saveTagsInput" v-model="inputTagValue" class="input-new-tag" size="small" @keyup.enter.native="handleTagsInputConfirm" @blur="handleTagsInputConfirm" />
      <el-button v-else class="button-new-tag" size="small" @click="showTagsInput">+ New Tag</el-button>
    </el-form-item>
    <el-form-item label="Healthchecks">
      <el-collapse accordion>
        <el-collapse-item title="Active">
          <el-form ref="activeCheck" label-width="180px" label-position="right">
            <el-form-item label="Concurrency">
              <el-input v-model="upstream.healthchecks.active.concurrency" />
              <p class="desc-font">Number of targets to check concurrently in active health checks. Default: <span class="mark-font">10</span>.</p>
            </el-form-item>
            <el-form-item label="Http Path">
              <el-input v-model="upstream.healthchecks.active.http_path" />
              <p class="desc-font">Path to use in GET HTTP request to run as a probe on active health checks. Default: <span class="mark-font">"/"</span>.</p>
            </el-form-item>
            <el-form-item label="Https SNI">
              <el-input v-model="upstream.healthchecks.active.https_sni" />
              <p class="desc-font">The hostname to use as an SNI (Server Name Identification) when performing active health checks using HTTPS. This is particularly useful when Targets are configured using IPs, so that the target host’s certificate can be verified with the proper SNI.</p>
            </el-form-item>
            <el-form-item label="Https Verify Certificate">
              <el-switch v-model="upstream.healthchecks.active.https_verify_certificate" />
              <p class="desc-font">Whether to check the validity of the SSL certificate of the remote host when performing active health checks using HTTPS. Default: <span class="mark-font">true</span>.</p>
            </el-form-item>
            <el-form-item label="Type">
              <el-select v-model="upstream.healthchecks.active.type">
                <el-option v-for="item in net_type_options" :key="item.value" :label="item.label" :value="item.value" />
              </el-select>
              <p class="desc-font">Whether to perform active health checks using HTTP or HTTPS, or just attempt a TCP connection. Accepted values are: <span class="mark-font">tcp</span>, <span class="mark-font">http</span>, <span class="mark-font">https</span>, <span class="mark-font">grpc</span>, <span class="mark-font">grpcs</span>. Default: <span class="mark-font">http</span>.</p>
            </el-form-item>
            <el-form-item label="Timeout">
              <el-input v-model="upstream.healthchecks.active.timeout" />
            </el-form-item>
            <el-form-item label="Healthy Http Statuses">
              <div>
                <el-tag v-for="status in upstream.healthchecks.active.healthy.http_statuses" :key="status" closable :disable-transitions="false" @close="removeTags('active.healthy.http_statuses', status)">
                  {{ status }}
                </el-tag>
              </div>
              <el-input v-model="activeHealthyStatus" @keyup.enter.native="addTags('active.healthy.http_statuses')" />
              <p class="desc-font">An array of HTTP statuses to consider a success, indicating healthiness, when returned by a probe in active health checks. Default: <span class="mark-font">[200, 302]</span>. With form-encoded, the notation is http_statuses[]=200&http_statuses[]=302. With JSON, use an Array.</p>
            </el-form-item>
            <el-form-item label="Healthy Interval">
              <el-input v-model="upstream.healthchecks.active.healthy.interval" />
              <p class="desc-font">Interval between active health checks for healthy targets (in seconds). A value of zero indicates that active probes for healthy targets should not be performed. Default: <span class="mark-font">0</span>.</p>
            </el-form-item>
            <el-form-item label="Healthy Successes">
              <el-input v-model="upstream.healthchecks.active.healthy.successes" />
              <p class="desc-font">Number of successes in active probes (as defined by healthchecks.active.healthy.http_statuses) to consider a target healthy. Default: <span class="mark-font">0</span>.</p>
            </el-form-item>
            <el-form-item label="Unhealthy Http Failures">
              <el-input v-model="upstream.healthchecks.active.unhealthy.http_failures" />
              <p class="desc-font">Number of HTTP failures in active probes (as defined by healthchecks.active.unhealthy.http_statuses) to consider a target unhealthy. Default: <span class="mark-font">0</span>.</p>
            </el-form-item>
            <el-form-item label="Unhealthy Http Statuses">
              <div>
                <el-tag v-for="status in upstream.healthchecks.active.unhealthy.http_statuses" :key="status" closable :disable-transitions="false" @close="removeTags('active.unhealthy.http_statuses', status)">
                  {{ status }}
                </el-tag>
              </div>
              <el-input v-model="activeUnhealthyStatus" @keyup.enter.native="addTags('active.unhealthy.http_statuses')" />
              <p class="desc-font">An array of HTTP statuses to consider a failure, indicating unhealthiness, when returned by a probe in active health checks. Default: <span class="mark-font">[429, 404, 500, 501, 502, 503, 504, 505]</span>. With form-encoded, the notation is http_statuses[]=429&http_statuses[]=404. With JSON, use an Array.</p>
            </el-form-item>
            <el-form-item label="Unhealthy TCP Failures">
              <el-input v-model="upstream.healthchecks.active.unhealthy.tcp_failures" />
              <p class="desc-font">Number of TCP failures in active probes to consider a target unhealthy. Default: <span class="mark-font">0</span>.</p>
            </el-form-item>
            <el-form-item label="Unhealthy Timeouts">
              <el-input v-model="upstream.healthchecks.active.unhealthy.timeouts" />
              <p class="desc-font">Number of timeouts in active probes to consider a target unhealthy. Default: <span class="mark-font">0</span>.</p>
            </el-form-item>
            <el-form-item label="Unhealthy Interval">
              <el-input v-model="upstream.healthchecks.active.unhealthy.interval" />
              <p class="desc-font">Interval between active health checks for unhealthy targets (in seconds). A value of zero indicates that active probes for unhealthy targets should not be performed. Default: <span class="mark-font">0</span>.</p>
            </el-form-item>
          </el-form>
        </el-collapse-item>
        <el-collapse-item title="Passive" name="passic">
          <el-form ref="passive" :model="upstream.healthchecks.passive" label-width="auto" />
        </el-collapse-item>
      </el-collapse>
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
      activeHealthyStatus: '',
      activeUnhealthyStatus: '',
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

      net_type_options: [
        {
          value: 'http',
          label: ''
        },
        {
          value: 'https',
          label: 'https'
        },
        {
          value: 'tcp',
          label: 'tcp'
        },

        {
          value: 'grpc',
          label: 'grpc'
        },
        {
          value: 'grpcs',
          label: 'grpcs'
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

      inputActiveHealStatusVisible: false,
      inputActiveHealStatusValue: '',

      upstream: {
        id: '',
        name: '',
        host_header: '',
        client_certificate: {
          id: ''
        },
        algorithm: 'round-robin',
        slots: 10000,
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
        hash_on: 'none',
        hash_fallback: 'none',
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
        const ptr = this.upstream
        const item = rs.data
        ptr.id = item.id
        ptr.name = item.name
        ptr.host_header = item.host_header

        item.client_certificate !== undefined && (ptr.client_certificate.id = item.client_certificate.id)
        item.algorithm !== undefined && (ptr.algorithm = item.algorithm)
        item.slots !== undefined && (ptr.slots = item.slots)

        item.hash_on !== undefined && (ptr.hash_on = item.hash_on)
        item.hash_fallback !== undefined && (ptr.hash_fallback = item.hash_fallback)
        item.hash_on_header !== undefined && (ptr.hash_on_header = item.hash_on_header)
        item.hash_fallback_header !== undefined && (ptr.hash_fallback_header = item.hash_fallback_header)
        item.hash_on_cookie !== undefined && (ptr.hash_on_cookie = item.hash_on_cookie)
        item.hash_on_cookie_path !== undefined && (ptr.hash_on_cookie_path = item.hash_on_cookie_path)
        item.tags !== undefined && (ptr.tags = item.tags)
        item.healthchecks.threshold !== undefined && (ptr.healthchecks.threshold = item.healthchecks.threshold)

        const activePtr = ptr.healthchecks.active
        const active = item.healthchecks.active
        active.concurrency !== undefined && (activePtr.concurrency = active.concurrency)
        active.http_path !== undefined && (activePtr.http_path = active.http_path)
        active.https_sni !== undefined && (activePtr.https_sni = active.https_sni)
        active.https_verify_certificate !== undefined && (activePtr.https_verify_certificate = active.https_verify_certificate)
        active.type !== undefined && (activePtr.type = active.type)
        active.timeout !== undefined && (activePtr.timeout = active.timeout)

        active.healthy.http_statuses !== undefined && (activePtr.healthy.http_statuses = active.healthy.http_statuses)
        active.healthy.interval !== undefined && (activePtr.healthy.interval = active.healthy.interval)
        active.healthy.successes !== undefined && (activePtr.healthy.successes = active.healthy.successes)

        active.unhealthy.http_failures !== undefined && (activePtr.unhealthy.http_failures = active.unhealthy.http_failures)
        active.unhealthy.http_statuses !== undefined && (activePtr.unhealthy.http_statuses = active.unhealthy.http_statuses)
        active.unhealthy.tcp_failures !== undefined && (activePtr.unhealthy.tcp_failures = active.unhealthy.tcp_failures)
        active.unhealthy.timeouts !== undefined && (activePtr.unhealthy.timeouts = active.unhealthy.timeouts)
        active.unhealthy.interval !== undefined && (activePtr.unhealthy.interval = active.unhealthy.interval)
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
    },
    addTags(addType) {
      switch (addType) {
        case 'active.healthy.http_statuses':
          if (this.upstream.healthchecks.active.healthy.http_statuses === undefined) {
            this.upstream.healthchecks.active.healthy.http_statuses = []
          }
          this.upstream.healthchecks.active.healthy.http_statuses.push(this.activeHealthyStatus)
          this.activeHealthyStatus = ''
          break
        case 'active.unhealthy.http_statuses':
          if (this.upstream.healthchecks.active.unhealthy.http_statuses === undefined) {
            this.upstream.healthchecks.active.unhealthy.http_statuses = []
          }
          this.upstream.healthchecks.active.unhealthy.http_statuses.push(this.activeUnhealthyStatus)
          this.activeUnhealthyStatus = ''
          break
      }
    },
    removeTags(rmType, tag) {
      switch (rmType) {
        case 'active.healthy.http_statuses':
          this.upstream.healthchecks.active.healthy.http_statuses.splice(this.upstream.healthchecks.active.healthy.http_statuses.indexOf(tag), 1)
          break
        case 'active.unhealthy.http_statuses':
          this.upstream.healthchecks.active.unhealthy.http_statuses.splice(this.upstream.healthchecks.active.unhealthy.http_statuses.indexOf(tag), 1)
          break
      }
    }
  }
}
</script>

<style>
  .el-input {
    width: 300px;
  }
  .el-tag + .el-tag {
    margin-left: 10px;
  }
  .span {
    font-size: 13px;
    color: #fa6863;
  }
  .el-form-item__label {
    color: #7e879a;
  }
  .el-collapse-item__header {
    color: #a1adc5;
    height: 60px;
  }

  .el-form-item__content {
    padding-right: 30px;
  }
</style>
