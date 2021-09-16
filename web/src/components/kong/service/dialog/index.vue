<template>
  <div>
    <el-dialog :before-close="closeDialog" :visible.sync="visible" title="客户">
      <el-form ref="service" :model="service" label-width="120px" label-position="left">
        <el-form-item label="Name">
          <el-input v-model="service.name" style="width: 300px" />
        </el-form-item>
        <el-form-item label="Path">
          <el-input v-model="service.path" style="width: 300px" />
        </el-form-item>
        <el-form-item label="Host">
          <el-input v-model="service.host" style="width: 300px" />
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
  createService
} from '@/api/kongService'

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
      inputCertsVisible: false,
      inputCertsValue: '',
      service: {
        client_certificate: {
          id: ''
        },
        connect_timeout: 0,
        created_at: 0,
        host: '',
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
      const data = {
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
      createService(data).then((rs) => {
        console.log('create service')
        this.$emit('closeHandler')
        this.$emit('updateHandler')
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
