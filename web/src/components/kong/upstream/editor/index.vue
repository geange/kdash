<template>
  <el-form ref="upstream" :model="upstream" label-width="auto">
    <el-form-item label="Name">
      <el-input v-model="upstream.name" />
      <p>This is a hostname, which must be equal to the <font>host</font> of a Service.</p>
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
      <p>Which load balancing algorithm to use. Accepted values are: <font>consistent-hashing</font>, <font>least-connections</font>, <font>round-robin</font>. Default: <font>round-robin</font>.</p>
    </el-form-item>
    <el-form-item label="Slots">
      <el-input v-model="upstream.slots" type="number" />
      <p>The number of slots in the load balancer algorithm. If <font>algorithm</font> is set to <font>round-robin</font>, this setting determines the maximum number of slots. If <font>algorithm</font> is set to <font>consistent-hashing</font>, this setting determines the actual number of slots in the algorithm. Accepts an integer in the range <font>10-65536</font>. Default: <font>10000</font>.</p>
    </el-form-item>
    <el-form-item label="Hash On">
      <el-select v-model="upstream.hash_on">
        <el-option v-for="item in hash_on_options" :key="item.value" :label="item.label" :value="item.value" />
      </el-select>
      <p>What to use as hashing input. Using <font>none</font> results in a weighted-round-robin scheme with no hashing. Accepted values are: <font>none</font>, <font>consumer</font>, <font>ip</font>, <font>header</font>, <font>cookie</font>. Default: <font>none</font>.</p>
    </el-form-item>
    <el-form-item label="Hash Fallback">
      <el-select v-model="upstream.hash_fallback">
        <el-option v-for="item in hash_fallback_options" :key="item.value" :label="item.label" :value="item.value" />
      </el-select>
      <p>What to use as hashing input if the primary <font>hash_on</font> does not return a hash (eg. header is missing, or no Consumer identified). Not available if <font>hash_on</font> is set to <font>cookie</font>. Accepted values are: <font>none</font>, <font>consumer</font>, <font>ip</font>, <font>header</font>, <font>cookie</font>. Default: <font>none</font>.</p>
    </el-form-item>
    <el-form-item label="Hash On Header">
      <el-input v-model="upstream.hash_on_header" />
      <p>The header name to take the value from as hash input. Only required when <font>hash_on</font> is set to <font>header</font>.</p>
    </el-form-item>
    <el-form-item label="Hash Fallback Header">
      <el-input v-model="upstream.hash_fallback_header" />
      <p>The header name to take the value from as hash input. Only required when <font>hash_fallback</font> is set to <font>header</font>.</p>
    </el-form-item>
    <el-form-item label="Hash On Cookie">
      <el-input v-model="upstream.hash_on_cookie" />
      <p>The cookie name to take the value from as hash input. Only required when <font>hash_on</font> or <font>hash_fallback</font> is set to <font>cookie</font>. If the specified cookie is not in the request, Kong will generate a value and set the cookie in the response.</p>
    </el-form-item>
    <el-form-item label="Hash On Cookie Path">
      <el-input v-model="upstream.hash_on_cookie_path" />
      <p>The cookie path to set in the response headers. Only required when <font>hash_on</font> or <font>hash_fallback</font> is set to cookie. Default: <font>"/"</font>.</p>
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
          <el-form label-width="180px" label-position="right">
            <el-form-item label="Concurrency">
              <el-input v-model="upstream.healthchecks.active.concurrency" />
              <p>Number of targets to check concurrently in active health checks. Default: <font>10</font>.</p>
            </el-form-item>
            <el-form-item label="Http Path">
              <el-input v-model="upstream.healthchecks.active.http_path" />
              <p>Path to use in GET HTTP request to run as a probe on active health checks. Default: <font>"/"</font>.</p>
            </el-form-item>
            <el-form-item label="Https SNI">
              <el-input v-model="upstream.healthchecks.active.https_sni" />
              <p>The hostname to use as an SNI (Server Name Identification) when performing active health checks using HTTPS. This is particularly useful when Targets are configured using IPs, so that the target host’s certificate can be verified with the proper SNI.</p>
            </el-form-item>
            <el-form-item label="Https Verify Certificate">
              <el-switch v-model="upstream.healthchecks.active.https_verify_certificate" />
              <p>Whether to check the validity of the SSL certificate of the remote host when performing active health checks using HTTPS. Default: <font>true</font>.</p>
            </el-form-item>
            <el-form-item label="Type">
              <el-select v-model="upstream.healthchecks.active.type">
                <el-option v-for="item in net_type_options" :key="item.value" :label="item.label" :value="item.value" />
              </el-select>
              <p>Whether to perform active health checks using HTTP or HTTPS, or just attempt a TCP connection. Accepted values are: <font>tcp</font>, <font>http</font>, <font>https</font>, <font>grpc</font>, <font>grpcs</font>. Default: <font>http</font>.</p>
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
              <p>An array of HTTP statuses to consider a success, indicating healthiness, when returned by a probe in active health checks. Default: <font>[200, 302]</font>. With form-encoded, the notation is http_statuses[]=200&http_statuses[]=302. With JSON, use an Array.</p>
            </el-form-item>
            <el-form-item label="Healthy Interval">
              <el-input v-model="upstream.healthchecks.active.healthy.interval" />
              <p>Interval between active health checks for healthy targets (in seconds). A value of zero indicates that active probes for healthy targets should not be performed. Default: <font>0</font>.</p>
            </el-form-item>
            <el-form-item label="Healthy Successes">
              <el-input v-model="upstream.healthchecks.active.healthy.successes" />
              <p>Number of successes in active probes (as defined by healthchecks.active.healthy.http_statuses) to consider a target healthy. Default: <font>0</font>.</p>
            </el-form-item>
            <el-form-item label="Unhealthy Http Failures">
              <el-input v-model="upstream.healthchecks.active.unhealthy.http_failures" />
              <p>Number of HTTP failures in active probes (as defined by healthchecks.active.unhealthy.http_statuses) to consider a target unhealthy. Default: <font>0</font>.</p>
            </el-form-item>
            <el-form-item label="Unhealthy Http Statuses">
              <div>
                <el-tag v-for="status in upstream.healthchecks.active.unhealthy.http_statuses" :key="status" closable :disable-transitions="false" @close="removeTags('active.unhealthy.http_statuses', status)">
                  {{ status }}
                </el-tag>
              </div>
              <el-input v-model="activeUnhealthyStatus" @keyup.enter.native="addTags('active.unhealthy.http_statuses')" />
              <p>An array of HTTP statuses to consider a failure, indicating unhealthiness, when returned by a probe in active health checks. Default: <font>[429, 404, 500, 501, 502, 503, 504, 505]</font>. With form-encoded, the notation is http_statuses[]=429&http_statuses[]=404. With JSON, use an Array.</p>
            </el-form-item>
            <el-form-item label="Unhealthy TCP Failures">
              <el-input v-model="upstream.healthchecks.active.unhealthy.tcp_failures" />
              <p>Number of TCP failures in active probes to consider a target unhealthy. Default: <font>0</font>.</p>
            </el-form-item>
            <el-form-item label="Unhealthy Timeouts">
              <el-input v-model="upstream.healthchecks.active.unhealthy.timeouts" />
              <p>Number of timeouts in active probes to consider a target unhealthy. Default: <font>0</font>.</p>
            </el-form-item>
            <el-form-item label="Unhealthy Interval">
              <el-input v-model="upstream.healthchecks.active.unhealthy.interval" />
              <p>Interval between active health checks for unhealthy targets (in seconds). A value of zero indicates that active probes for unhealthy targets should not be performed. Default: <font>0</font>.</p>
            </el-form-item>
          </el-form>
        </el-collapse-item>
        <el-collapse-item title="Passive" name="passic">
          <el-form label-width="180px" label-position="right">
            <el-form-item label="Unhealthy Timeouts">
              <el-input v-model="upstream.healthchecks.passive.unhealthy.timeouts" />
              <p>Number of timeouts in proxied traffic to consider a target unhealthy, as observed by passive health checks. Default: <font>0</font>.</p>
            </el-form-item>
            <el-form-item label="Unhealthy TCP Failures">
              <el-input v-model="upstream.healthchecks.passive.unhealthy.tcp_failures" />
              <p>Number of TCP failures in proxied traffic to consider a target unhealthy, as observed by passive health checks. Default: <font>0</font>.</p>
            </el-form-item>
            <el-form-item label="Unhealthy TCP Statuses">
              <el-input v-model="upstream.healthchecks.passive.unhealthy.tcp_failures" />
              <p>An array of HTTP statuses which represent unhealthiness when produced by proxied traffic, as observed by passive health checks. Default: <font>[429, 500, 503]</font>. With form-encoded, the notation is http_statuses[]=429&http_statuses[]=500. With JSON, use an Array.</p>
            </el-form-item>
            <el-form-item label="Unhealthy http Failures">
              <el-input v-model="upstream.healthchecks.passive.unhealthy.http_failures" />
              <p>Number of HTTP failures in proxied traffic (as defined by healthchecks.passive.unhealthy.http_statuses) to consider a target unhealthy, as observed by passive health checks. Default: 0.</p>
            </el-form-item>
            <el-form-item label="Type">
              <el-input v-model="upstream.healthchecks.passive.type" />
              <p>Whether to perform passive health checks interpreting HTTP/HTTPS statuses, or just check for TCP connection success. In passive checks, http and https options are equivalent. Accepted values are: <font>"tcp", "http", "https", "grpc", "grpcs"</font>. Default: <font>"http"</font>.</p>
            </el-form-item>
            <el-form-item label="Healthy http Statuses">
              <!-- <el-input v-model="upstream.healthchecks.passive.healthy.http_statuses" /> -->
              <div>
                <el-tag v-for="status in upstream.healthchecks.passive.healthy.http_statuses" :key="status" closable :disable-transitions="false" @close="removeTags('passive.healthy.http_statuses', status)">
                  {{ status }}
                </el-tag>
              </div>
              <el-input v-model="activeUnhealthyStatus" @keyup.enter.native="addTags('passive.healthy.http_statuses')" />
              <p>An array of HTTP statuses which represent healthiness when produced by proxied traffic, as observed by passive health checks. Default: <font>[200, 201, 202, 203, 204, 205, 206, 207, 208, 226, 300, 301, 302, 303, 304, 305, 306, 307, 308]</font> With form-encoded, the notation is http_statuses[]=200&http_statuses[]=201. With JSON, use an Array.</p>
            </el-form-item>
            <el-form-item label="Healthy Successes">
              <el-input v-model="upstream.healthchecks.passive.unhealthy.http_failures" />
              <p>Number of successes in proxied traffic (as defined by <font>healthchecks.passive.healthy.http_statuses</font>) to consider a target healthy, as observed by passive health checks. Default: 0.</p>
            </el-form-item>
          </el-form>
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

        const passivePtr = ptr.healthchecks.passive
        const passive = item.healthchecks.passive
        passive.type !== undefined && (passivePtr.type = passive.type)

        passive.healthy.http_statuses !== undefined && (passivePtr.healthy.http_statuses = passive.healthy.http_statuses)
        passive.healthy.interval !== undefined && (passivePtr.healthy.interval = passive.healthy.interval)
        passive.healthy.successes !== undefined && (passivePtr.healthy.successes = passive.healthy.successes)

        passive.unhealthy.http_failures !== undefined && (passivePtr.unhealthy.http_failures = passive.unhealthy.http_failures)
        passive.unhealthy.http_statuses !== undefined && (passivePtr.unhealthy.http_statuses = passive.unhealthy.http_statuses)
        passive.unhealthy.tcp_failures !== undefined && (passivePtr.unhealthy.tcp_failures = passive.unhealthy.tcp_failures)
        passive.unhealthy.timeouts !== undefined && (passivePtr.unhealthy.timeouts = passive.unhealthy.timeouts)
        passive.unhealthy.interval !== undefined && (passivePtr.unhealthy.interval = passive.unhealthy.interval)
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
        this.$refs.saveTagsInput.$refs.input.font()
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

<style scoped>
  .el-input {
    width: 300px;
  }
  .el-tag + .el-tag {
    margin-left: 10px;
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
  p {
    font-size: 13px;
    color: #8b95aa;
    line-height: 1.5;
    line-break: strict;
    margin-right: 20px;
  }
  font {
    font-size: 13px;
    font-weight: bold;
    color: #fa6863;
  }
</style>
