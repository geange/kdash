<template>
  <el-form ref="route" :model="route" label-width="auto">
    <el-form-item label="Name">
      <el-input v-model="route.name" />
    </el-form-item>
    <el-form-item label="Hosts">
      <el-tag v-for="host in route.hosts" :key="host" closable :disable-transitions="false" @close="handleHostsClose(host)">{{ host }}</el-tag>
      <el-input v-if="inputHostsVisible" ref="saveHostsInput" v-model="inputHostValue" class="input-new-tag" size="small" @keyup.enter.native="handleHostsInputConfirm" @blur="handleHostsInputConfirm" />
      <el-button v-else class="button-new-tag" size="small" @click="showHostsInput">+ New Tag</el-button>
    </el-form-item>
    <el-form-item label="Methods">
      <el-tag v-for="method in route.methods" :key="method" closable :disable-transitions="false" @close="handleMethodsClose(method)">{{ method }}</el-tag>
      <el-input v-if="inputMethodsVisible" ref="saveMethodsInput" v-model="inputMethodValue" class="input-new-tag" size="small" @keyup.enter.native="handleMethodsInputConfirm" @blur="handleMethodsInputConfirm" />
      <el-button v-else class="button-new-tag" size="small" @click="showMethodsInput">+ New Tag</el-button>
    </el-form-item>
    <el-form-item label="Paths">
      <el-tag v-for="path in route.paths" :key="path" closable :disable-transitions="false" @close="handlePathsClose(path)">{{ path }}</el-tag>
      <el-input v-if="inputPathsVisible" ref="savePathsInput" v-model="inputPathValue" class="input-new-tag" size="small" @keyup.enter.native="handlePathsInputConfirm" @blur="handlePathsInputConfirm" />
      <el-button v-else class="button-new-tag" size="small" @click="showPathsInput">+ New Tag</el-button>
    </el-form-item>
    <el-form-item label="PathHandling">
      <el-input v-model="route.path_handling" />
    </el-form-item>
    <el-form-item label="PreserveHost">
      <el-switch v-model="route.preserve_host" />
    </el-form-item>
    <el-form-item label="Protocols">
      <el-tag v-for="protocol in route.protocols" :key="protocol" closable :disable-transitions="false" @close="handleProtocolsClose(protocol)">{{ protocol }}</el-tag>
      <el-input v-if="inputProtocolsVisible" ref="saveProtocolsInput" v-model="inputProtocolValue" class="input-new-tag" size="small" @keyup.enter.native="handleProtocolsInputConfirm" @blur="handleProtocolsInputConfirm" />
      <el-button v-else class="button-new-tag" size="small" @click="showProtocolsInput">+ New Tag</el-button>
    </el-form-item>
    <el-form-item label="RegexPriority">
      <el-input v-model="route.regex_priority" type="number" />
    </el-form-item>
    <el-form-item label="Service">
      <el-input v-model="route.service.id" />
    </el-form-item>
    <el-form-item label="StripPath">
      <el-switch v-model="route.strip_path" />
    </el-form-item>
    <el-form-item label="SNIs">
      <el-tag v-for="sni in route.snis" :key="sni" closable :disable-transitions="false" @close="handleSNIsClose(sni)">{{ sni }}</el-tag>
      <el-input v-if="inputSNIsVisible" ref="saveSNIsInput" v-model="inputSNIValue" class="input-new-tag" size="small" @keyup.enter.native="handleSNIsInputConfirm" @blur="handleSNIsInputConfirm" />
      <el-button v-else class="button-new-tag" size="small" @click="showSNIsInput">+ New Tag</el-button>
    </el-form-item>
    <el-form-item label="Sources" />
    <el-form-item label="Destinations" />
    <el-form-item label="Tags">
      <el-tag v-for="tag in route.tags" :key="tag" closable :disable-transitions="false" @close="handleTagsClose(tag)">{{ tag }}</el-tag>
      <el-input v-if="inputTagsVisible" ref="saveTagsInput" v-model="inputTagValue" class="input-new-tag" size="small" @keyup.enter.native="handleTagsInputConfirm" @blur="handleTagsInputConfirm" />
      <el-button v-else class="button-new-tag" size="small" @click="showTagsInput">+ New Tag</el-button>
    </el-form-item>
    <el-form-item label="HTTPSRedirectStatusCode">
      <el-input v-model="route.https_redirect_status_code" type="number" />
    </el-form-item>
    <el-form-item label="RequestBuffering">
      <el-switch v-model="route.request_buffering" />
    </el-form-item>
    <el-form-item label="ResponseBuffering">
      <el-switch v-model="route.response_buffering" />
    </el-form-item>
    <el-form-item>
      <el-button type="primary" @click="onSubmit">立即创建</el-button>
      <el-button>取消</el-button>
    </el-form-item>
  </el-form>
</template>

<script>
export default {
  name: 'RouteEditor',
  data() {
    return {
      inputHostsVisible: false,
      inputHostValue: '',
      inputProtocolsVisible: false,
      inputProtocolValue: '',
      inputPathsVisible: false,
      inputPathValue: '',
      inputMethodsVisible: false,
      inputMethodValue: '',
      inputTagsVisible: false,
      inputTagValue: '',
      inputSNIsVisible: false,
      inputSNIValue: '',
      route: {
        created_at: 0,
        hosts: ['a.cn', 'b.cn'],
        headers: {
          'name': ['a', 'b']
        },
        id: 'sssssssss',
        name: 'ss',
        methods: ['GET'],
        paths: ['/abssssssssssssssssssssssssc9', '/abssssssssssssssssssssssssc'],
        path_handling: 'ss',
        preserve_host: 'sss.cn',
        protocols: ['http', 'https'],
        regex_priority: 0,
        service: {
          id: 'sssssffffff'
        },
        strip_path: false,
        updated_at: 0,
        snis: ['a1', 'a2'],
        sources: [{
          ip: '',
          port: 8000
        }],
        destinations: [],
        tags: ['a', 'b'],
        https_redirect_status_code: 302,
        request_buffering: false,
        response_buffering: false
      }
    }
  },
  created() {
    console.log('create')
  },
  methods: {
    onSubmit() {
    },
    handleHostsClose(key) {
      this.route.hosts.splice(this.route.hosts.indexOf(key), 1)
    },
    showHostsInput() {
      this.inputHostsVisible = true
      this.$nextTick(_ => {
        this.$refs.saveHostsInput.$refs.input.focus()
      })
    },
    handleHostsInputConfirm() {
      if (this.inputHostValue) {
        if (this.route.hosts === undefined) {
          this.route.hosts = []
        }
        this.route.hosts.push(this.inputHostValue)
      }
      this.inputHostsVisible = false
      this.inputHostValue = ''
    },

    handleMethodsClose(key) {
      this.route.methods.splice(this.route.methods.indexOf(key), 1)
    },
    showMethodsInput() {
      this.inputMethodsVisible = true
      this.$nextTick(_ => {
        this.$refs.saveMethodsInput.$refs.input.focus()
      })
    },
    handleMethodsInputConfirm() {
      if (this.inputMethodValue) {
        if (this.route.methods === undefined) {
          this.route.methods = []
        }
        this.route.methods.push(this.inputMethodValue)
      }
      this.inputMethodsVisible = false
      this.inputMethodValue = ''
    },

    handlePathsClose(key) {
      this.route.paths.splice(this.route.paths.indexOf(key), 1)
    },
    showPathsInput() {
      this.inputPathsVisible = true
      this.$nextTick(_ => {
        this.$refs.savePathsInput.$refs.input.focus()
      })
    },
    handlePathsInputConfirm() {
      if (this.inputPathValue) {
        if (this.route.paths === undefined) {
          this.route.paths = []
        }
        this.route.paths.push(this.inputPathValue)
      }
      this.inputPathsVisible = false
      this.inputPathValue = ''
    },

    handleTagsClose(key) {
      this.route.tags.splice(this.route.tags.indexOf(key), 1)
    },
    showTagsInput() {
      this.inputTagsVisible = true
      this.$nextTick(_ => {
        this.$refs.saveTagsInput.$refs.input.focus()
      })
    },
    handleTagsInputConfirm() {
      if (this.inputTagValue) {
        if (this.route.tags === undefined) {
          this.route.tags = []
        }
        this.route.tags.push(this.inputTagValue)
      }
      this.inputTagsVisible = false
      this.inputTagValue = ''
    },

    handleSNIsClose(key) {
      this.route.snis.splice(this.route.snis.indexOf(key), 1)
    },
    showSNIsInput() {
      this.inputSNIsVisible = true
      this.$nextTick(_ => {
        this.$refs.saveSNIsInput.$refs.input.focus()
      })
    },
    handleSNIsInputConfirm() {
      if (this.inputSNIValue) {
        if (this.route.snis === undefined) {
          this.route.snis = []
        }
        this.route.snis.push(this.inputSNIValue)
      }
      this.inputSNIsVisible = false
      this.inputSNIValue = ''
    },

    handleProtocolsClose(key) {
      this.route.protocols.splice(this.route.protocols.indexOf(key), 1)
    },
    showProtocolsInput() {
      this.inputProtocolsVisible = true
      this.$nextTick(_ => {
        this.$refs.saveProtocolsInput.$refs.input.focus()
      })
    },
    handleProtocolsInputConfirm() {
      if (this.inputProtocolValue) {
        if (this.route.protocols === undefined) {
          this.route.protocols = []
        }
        this.route.protocols.push(this.inputProtocolValue)
      }
      this.inputProtocolsVisible = false
      this.inputProtocolValue = ''
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
  .button-new-tag {
    margin-left: 10px;
  }
  .input-new-tag {
    width: 120px;
    margin-left: 10px;
    vertical-align: bottom;
  }
</style>
