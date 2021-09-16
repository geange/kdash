<template>
  <el-form ref="service" :model="service" label-width="120px" label-position="left">
    <el-form-item label="ID">
      <el-tag>{{ service.id }}</el-tag>
    </el-form-item>
    <el-form-item label="Name">
      <el-input v-model="service.name" style="width: 300px" />
    </el-form-item>
    <el-form-item label="Path">
      <el-input v-model="service.path" style="width: 300px" />
    </el-form-item>
    <el-form-item label="Port">
      <el-input v-model="service.port" style="width: 300px" type="number" />
    </el-form-item>
    <el-form-item label="Protocol">
      <el-input v-model="service.protocol" style="width: 300px" />
    </el-form-item>
    <el-form-item label="ReadTimeout">
      <el-input v-model="service.read_timeout" style="width: 300px" type="number" />
    </el-form-item>
    <el-form-item label="Retries">
      <el-input v-model="service.retries" style="width: 300px" type="number" />
    </el-form-item>
    <el-form-item label="URL">
      <el-input v-model="service.url" style="width: 300px" />
    </el-form-item>
    <el-form-item label="WriteTimeout">
      <el-input v-model="service.write_timeout" style="width: 300px" type="number" />
    </el-form-item>
    <el-form-item label="Tags">
      <el-tag v-for="tag in service.tags" :key="tag" closable :disable-transitions="false" @close="handleTagsClose(tag)">
        {{ tag }}
      </el-tag>
      <el-input v-if="inputTagsVisible" ref="saveTagsInput" v-model="inputTagValue" class="input-new-tag" size="small" @keyup.enter.native="handleTagsInputConfirm" @blur="handleTagsInputConfirm" />
      <el-button v-else class="button-new-tag" size="small" @click="showTagsInput">+ New Tag</el-button>
    </el-form-item>
    <el-form-item label="TLSVerify">
      <el-switch v-model="service.tls_verify" />
    </el-form-item>
    <el-form-item label="TLSVerifyDepth">
      <el-input v-model="service.tls_verify_depth" style="width: 300px" type="number" />
    </el-form-item>
    <el-form-item label="CACertificates">
      <el-tag v-for="cert in service.ca_certificates" :key="cert" closable :disable-transitions="false" @close="handleCertsClose(cert)">
        {{ cert }}
      </el-tag>
      <el-input v-if="inputCertsVisible" ref="saveCertsTagInput" v-model="inputCertsValue" class="input-new-tag" size="small" @keyup.enter.native="handleCertsInputConfirm" @blur="handleCertsInputConfirm" />
      <el-button v-else class="button-new-tag" size="small" @click="showCertsInput">+ New Tag</el-button>
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
    await this.getService(this.sid)
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
    handleTagsClose(tag) {
      this.service.tags.splice(this.service.tags.indexOf(tag), 1)
    },
    showTagsInput() {
      this.inputTagsVisible = true
      this.$nextTick(_ => {
        this.$refs.saveTagsInput.$refs.input.focus()
      })
    },
    handleTagsInputConfirm() {
      if (this.inputTagValue) {
        if (this.service.tags === undefined) {
          this.service.tags = []
        }
        this.service.tags.push(this.inputTagValue)
      }
      this.inputTagsVisible = false
      this.inputTagValue = ''
    },
    handleCertsClose(tag) {
      this.service.ca_certificates.splice(this.service.ca_certificates.indexOf(tag), 1)
    },
    showCertsInput() {
      this.inputCertsVisible = true
      this.$nextTick(_ => {
        this.$refs.saveCertsTagInput.$refs.input.focus()
      })
    },
    handleCertsInputConfirm() {
      const inputValue = this.inputCertsValue
      if (inputValue) {
        if (this.service.ca_certificates === undefined) {
          this.service.ca_certificates = []
        }
        this.service.ca_certificates.push(inputValue)
      }
      this.inputCertsVisible = false
      this.inputCertsValue = ''
    },
    async onSubmit() {
      const data = {
        id: this.service.id,
        host: this.service.host,
        name: this.service.name
      }
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
      updateService(data).then((rs) => {
      }).catch(err => {
        this.$message({ type: 'error', message: err })
      })
      await this.getService(this.sid)
    }
  }
}
</script>

<style scoped>
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
    margin-top: 15px;
    width: 100%;
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
</style>
