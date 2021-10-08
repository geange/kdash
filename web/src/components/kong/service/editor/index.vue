<template>
  <el-form ref="service" :model="service" label-width="auto">
    <el-form-item label="Name">
      <el-input v-model="service.name" />
      <p>The Service name.</p>
    </el-form-item>
    <el-form-item label="Protocol">
      <el-input v-model="service.protocol" />
      <p>The protocol used to communicate with the upstream. Accepted values are: <font>"grpc", "grpcs", "http", "https", "tcp", "tls", "udp"</font>. Default: <font>"http"</font>.</p>
    </el-form-item>
    <el-form-item label="Host">
      <el-input v-model="service.host" />
      <p>The host of the upstream server.</p>
    </el-form-item>
    <el-form-item label="Port">
      <el-input v-model="service.port" type="number" />
      <p>The upstream server port. Default: <font>80</font>.</p>
    </el-form-item>
    <el-form-item label="Path">
      <el-input v-model="service.path" />
      <p>The path to be used in requests to the upstream server.</p>
    </el-form-item>
    <el-form-item label="Connect Time">
      <el-input v-model="service.connect_timeout" type="number" />
      <p>The timeout in milliseconds for establishing a connection to the upstream server. Default: <font>60000</font>.</p>
    </el-form-item>
    <el-form-item label="WriteTimeout">
      <el-input v-model="service.write_timeout" type="number" />
      <p>The timeout in milliseconds between two successive write operations for transmitting a request to the upstream server. Default: <font>60000</font>.</p>
    </el-form-item>
    <el-form-item label="ReadTimeout">
      <el-input v-model="service.read_timeout" type="number" />
      <p>The timeout in milliseconds between two successive read operations for transmitting a request to the upstream server. Default: <font>60000</font>.</p>
    </el-form-item>
    <el-form-item label="Retries">
      <el-input v-model.number="service.retries" type="number" />
      <p>The number of retries to execute upon failure to proxy. Default: <font>5</font>.</p>
    </el-form-item>
    <!-- <el-form-item label="Tags">
      <el-tag v-for="tag in service.tags" :key="tag" closable :disable-transitions="false" @close="handleTagsClose(tag)">
        {{ tag }}
      </el-tag>
      <el-input v-if="inputTagsVisible" ref="saveTagsInput" v-model="inputTagValue" class="input-new-tag" size="small" @keyup.enter.native="handleTagsInputConfirm" @blur="handleTagsInputConfirm" />
      <el-button v-else class="button-new-tag" size="small" @click="showTagsInput">+ New Tag</el-button>
      <p>An optional set of strings associated with the Service for grouping and filtering.</p>
    </el-form-item> -->

    <el-form-item label="Tags">
      <div>
        <el-tag v-for="tag in service.tags" :key="tag" closable :disable-transitions="false" @close="removeTags('tags', tag)">
          {{ tag }}
        </el-tag>
      </div>
      <el-input v-model="inputTagValue" @keyup.enter.native="addTags('tags')" />
      <p>An optional set of strings associated with the Service for grouping and filtering.</p>
    </el-form-item>
    <el-form-item label="Client Certificate">
      <el-input v-model="service.client_certificate.id" />
      <p>Certificate to be used as client certificate while TLS handshaking to the upstream server. With form-encoded, the notation is <font>client_certificate.id=${client_certificate id}</font>. With JSON, use <font>"client_certificate":{"id":"${client_certificate id}"}</font>.</p>
    </el-form-item>
    <el-form-item label="TLSVerify">
      <el-switch v-model="service.tls_verify" />
      <p>Whether to enable verification of upstream server TLS certificate. If set to <font>null</font>, then the Nginx default is respected.</p>
    </el-form-item>
    <el-form-item label="TLS Verify Depth">
      <el-input v-model="service.tls_verify_depth" type="number" />
      <p>Maximum depth of chain while verifying Upstream server’s TLS certificate. If set to <font>null</font>, then the Nginx default is respected. Default: <font>null</font>.</p>
    </el-form-item>
    <el-form-item label="CACertificates">
      <div>
        <el-tag v-for="tag in service.ca_certificates" :key="tag" closable :disable-transitions="false" @close="removeTags('ca_certificates', tag)">
          {{ tag }}
        </el-tag>
      </div>
      <el-input v-model="inputCertsValue" @keyup.enter.native="addTags('ca_certificates')" />
      <p>Array of CA Certificate object UUIDs that are used to build the trust store while verifying upstream server’s TLS certificate. If set to null when Nginx default is respected. If default CA list in Nginx are not specified and TLS verification is enabled, then handshake with upstream server will always fail (because no CA are trusted). With form-encoded, the notation is ca_certificates[]=4e3ad2e4-0bc4-4638-8e34-c84a417ba39b&ca_certificates[]=51e77dc2-8f3e-4afa-9d0e-0e3bbbcfd515. With JSON, use an Array.</p>
    </el-form-item>

    <!-- <el-form-item label="CACertificates">
      <el-tag v-for="cert in service.ca_certificates" :key="cert" closable :disable-transitions="false" @close="handleCertsClose(cert)">
        {{ cert }}
      </el-tag>
      <el-input v-if="inputCertsVisible" ref="saveCertsTagInput" v-model="inputCertsValue" class="input-new-tag" size="small" @keyup.enter.native="handleCertsInputConfirm" @blur="handleCertsInputConfirm" />
      <el-button v-else class="button-new-tag" size="small" @click="showCertsInput">+ New Tag</el-button>
      <p>Array of CA Certificate object UUIDs that are used to build the trust store while verifying upstream server’s TLS certificate. If set to null when Nginx default is respected. If default CA list in Nginx are not specified and TLS verification is enabled, then handshake with upstream server will always fail (because no CA are trusted). With form-encoded, the notation is ca_certificates[]=4e3ad2e4-0bc4-4638-8e34-c84a417ba39b&ca_certificates[]=51e77dc2-8f3e-4afa-9d0e-0e3bbbcfd515. With JSON, use an Array.</p>
    </el-form-item> -->
    <el-form-item label="URL">
      <el-input v-model="service.url" />
      <p>Shorthand attribute to set <font>protocol</font>, <font>host</font>, <font>port</font> and <font>path</font> at once. This attribute is write-only (the Admin API never returns the URL).</p>
    </el-form-item>
    <el-form-item>
      <el-button type="primary" @click="onSubmit">Update</el-button>
      <el-button>Cancel</el-button>
    </el-form-item>
  </el-form>
</template>

<script>

import {
  getService,
  updateService
} from '@/api/kongService'

export default {
  name: 'ServiceEditor',
  props: {
    sid: {
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
      service: {
        client_certificate: {
          id: ''
        },
        connect_timeout: 0,
        created_at: 0,
        host: '',
        id: '',
        name: '',
        path: '',
        port: 0,
        protocol: '',
        read_timeout: 0,
        retries: 0,
        updated_at: 0,
        url: '',
        write_timeout: 0,
        tags: [],
        tls_verify: false,
        tls_verify_depth: 0,
        ca_certificates: []
      },
      inputTagsVisible: false,
      inputTagValue: '',
      inputCertsVisible: false,
      inputCertsValue: ''
    }
  },
  async mounted() {
    if (this.op === 'edit') {
      await this.getService(this.sid)
    }
  },
  methods: {
    async getService(id) {
      const params = { id: id }
      getService(params).then((rs) => {
        const v = rs.data
        this.service.id = v.id
        v.client_certificate !== undefined && (this.service.client_certificate = v.client_certificate)
        v.connect_timeout !== undefined && (this.service.connect_timeout = v.connect_timeout)
        v.created_at !== undefined && (this.service.created_at = v.created_at)
        v.host !== undefined && (this.service.host = v.host)
        v.id !== undefined && (this.service.id = v.id)
        v.name !== undefined && (this.service.name = v.name)
        v.path !== undefined && (this.service.path = v.path)
        v.port !== undefined && (this.service.port = v.port)
        v.protocol !== undefined && (this.service.protocol = v.protocol)
        v.read_timeout !== undefined && (this.service.read_timeout = v.read_timeout)
        v.retries !== undefined && (this.service.retries = v.retries)
        v.updated_at !== undefined && (this.service.updated_at = v.updated_at)
        v.url !== undefined && (this.service.url = v.url)
        v.write_timeout !== undefined && (this.service.write_timeout = v.write_timeout)
        v.tags !== undefined && (this.service.tags = v.tags)
        v.tls_verify !== undefined && (this.service.tls_verify = v.tls_verify)
        v.tls_verify_depth !== undefined && (this.service.tls_verify_depth = v.tls_verify_depth)
        v.ca_certificates !== undefined && (this.service.ca_certificates = v.ca_certificates)
      }).catch(err => {
        this.$message({ type: 'error', message: err })
      })
    },
    async onSubmit() {
      const data = {
        id: this.service.id,
        host: this.service.host,
        name: this.service.name
      }
      console.log(this.service.retries)
      this.service.client_certificate.id !== '' && (data.client_certificate.id = this.service.client_certificate.id)
      this.service.path !== '' && (data.path = this.service.path)
      this.service.port !== 0 && (data.port = this.service.port)
      this.service.protocol !== '' && (data.protocol = this.service.protocol)
      this.service.read_timeout !== 0 && (data.read_timeout = this.service.read_timeout)
      this.service.retries !== 0 && (data.retries = this.service.retries)
      this.service.url !== '' && (data.url = this.service.url)
      this.service.write_timeout !== 0 && (data.write_timeout = this.service.write_timeout)
      this.service.tags !== [] && (data.tags = this.service.tags)
      this.service.tls_verify !== false && (data.tls_verify = this.service.tls_verify)
      this.service.tls_verify_depth !== 0 && (data.tls_verify_depth = this.service.tls_verify_depth)
      this.service.ca_certificates !== [] && (data.ca_certificates = this.service.ca_certificates)
      console.log(data)
      updateService(data).then((rs) => {
      }).catch(err => {
        this.$message({ type: 'error', message: err })
      })
      await this.getService(this.sid)
    },
    addTags(addType) {
      switch (addType) {
        case 'tags':
          if (this.service.tags === undefined) {
            this.service.tags = []
          }
          this.service.tags.push(this.inputTagValue)
          this.inputTagValue = ''
          break
        case 'ca_certificates':
          if (this.service.ca_certificates === undefined) {
            this.service.ca_certificates = []
          }
          this.service.ca_certificates.push(this.inputCertsValue)
          this.inputCertsValue = ''
          break
      }
    },
    removeTags(rmType, tag) {
      switch (rmType) {
        case 'tags':
          this.service.tags.splice(this.service.tags.indexOf(tag), 1)
          break
        case 'ca_certificates':
          this.service.ca_certificates.splice(this.service.ca_certificates.indexOf(tag), 1)
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
  .button-new-tag {
    margin-left: 10px;
    height: 32px;
    line-height: 30px;
    padding-top: 0;
    padding-bottom: 0;
  }
  .input-new-tag {
    width: 120px;
    margin-left: 10px;
    vertical-align: bottom;
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
