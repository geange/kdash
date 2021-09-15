<template>
  <div>
    <el-dialog :before-close="closeDialog" :visible.sync="visible" title="客户">
      <el-form :inline="true" :model="form" label-width="80px">
        <el-form-item label="客户名">
          <el-input v-model="form.customerName" autocomplete="off" />
        </el-form-item>
        <el-form-item label="客户电话">
          <el-input v-model="form.customerPhoneData" autocomplete="off" />
        </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click="closeDialog">Cancel</el-button>
        <el-button type="primary" @click="enterDialog">Create</el-button>
      </div>
    </el-dialog>
  </div>
</template>

<script>
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
      form: {
        customerName: '',
        customerPhoneData: ''
      }
    }
  },
  methods: {
    closeDialog() {
      this.$emit('closeHandler')
    },
    enterDialog() {
      console.log('enter')
    }
  }
}
</script>

<style>
</style>
