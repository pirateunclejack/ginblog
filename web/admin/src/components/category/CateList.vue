<template>
  <div>
    <h3>Category List</h3>
    <a-card>
      <a-row :gutter="20">
        <a-col :span="4">
          <a-button type="primary" @click="addCateVisible = true"
            >Add category</a-button
          >
        </a-col>
      </a-row>
      <a-table
        rowKey="id"
        :columns="columns"
        :pagination="pagination"
        :dataSource="Catelist"
        bordered
        @change="handleTableChange"
      >
        <template slot="action" slot-scope="text, data">
          <div class="actionSlot">
            <a-button
              type="primary"
              icon="edit"
              style="margin-right: 15px"
              @click="editCate(data.id)"
              >Edit</a-button
            >
            <a-button
              type="danger"
              icon="delete"
              style="margin-right: 15px"
              @click="deleteCate(data.id)"
              >Delete</a-button
            >
          </div>
        </template>
      </a-table>
    </a-card>
    <!-- Add Category -->
    <a-modal
      closable
      title="Add Category"
      :visible="addCateVisible"
      width="60%"
      @ok="addCateOk"
      @cancel="addCateCancel"
      destroyOnClose
    >
      <a-form-model :model="newCate" :rules="addCateRules" ref="addCateRef">
        <a-form-model-item label="Catename" prop="name">
          <a-input v-model="newCate.name"></a-input>
        </a-form-model-item>
      </a-form-model>
    </a-modal>
    <!-- Edit Category -->
    <a-modal
      closable
      title="Edit Category"
      :visible="editCateVisible"
      width="60%"
      @ok="editCateOk"
      @cancel="editCateCancel"
    >
      <a-form-model :model="CateInfo" :rules="editCateRules" ref="editCateRef">
        <a-form-model-item label="Catename" prop="name"
          ><a-input v-model="CateInfo.name"></a-input
        ></a-form-model-item>
      </a-form-model>
    </a-modal>
  </div>
</template>

<script>
const columns = [
  {
    title: 'ID',
    dataIndex: 'id',
    width: '10%',
    key: 'id',
    align: 'center'
  },
  {
    title: 'Catename',
    dataIndex: 'name',
    width: '20%',
    key: 'name',
    align: 'center'
  },
  {
    title: 'Action',
    dataIndex: 'action',
    width: '30%',
    key: 'action',
    scopedSlots: { customRender: 'action' },
    align: 'center'
  }
]
export default {
  data() {
    return {
      pagination: {
        pageSizeOptions: ['2', '4', '8'],
        pageSize: 2,
        total: 0,
        showSizeChanger: true,
        showTotal: (total) => `Total ${total}`
      },
      Catelist: [],
      columns,
      queryParam: {
        pagesize: 5,
        pagenum: 1
      },
      CateInfo: {
        id: 0,
        name: ''
      },
      newCate: {
        name: ''
      },
      editVisible: false,
      addCateRules: {
        name: [
          {
            validator: (rule, value, callback) => {
              if (this.newCate.name === '') {
                callback(new Error('Please input Category name'))
              } else {
                callback()
              }
            },
            trigger: 'blur'
          }
        ]
      },
      editCateRules: {
        name: [
          {
            validator: (rule, value, callback) => {
              if (this.CateInfo.name === '') {
                callback(new Error('Please input Category name'))
              } else {
                callback()
              }
            },
            trigger: 'blur'
          }
        ]
      },
      addCateVisible: false,
      editCateVisible: false
    }
  },
  created() {
    this.getCateList()
  },
  computed: {
    IsAdmin: function () {
      if (this.CateInfo.role === 1) {
        return true
      } else {
        return false
      }
    }
  },
  methods: {
    // Get category list
    async getCateList() {
      const { data: res } = await this.$http.get('category', {
        params: {
          pagesize: this.queryParam.pagesize,
          pagenum: this.queryParam.pagenum
        }
      })
      if (res.status !== 200) {
        return this.$message.error(res.message)
      }
      this.Catelist = res.data
      this.pagination.total = res.total
    },
    handleTableChange(pagination, filters, sorter) {
      const pager = { ...this.pagination }
      pager.current = pagination.current
      pager.pageSize = pagination.pageSize
      this.queryParam.pagesize = pagination.pageSize
      this.queryParam.pagenum = pagination.current

      if (pagination.pageSize !== this.pagination.pageSize) {
        this.queryParam.pagenum = 1
        pager.current = 1
      }
      this.pagination = pager
      this.getCateList()
    },
    deleteCate(id) {
      this.$confirm({
        title: 'Do you want to delete this Cate?',
        content:
          'Are you sure you want to delete this Cate? Once it is deleted , it will be impossible to recover it!',
        onOk: async () => {
          const { data: res } = await this.$http.delete(`category/${id}`)
          if (res.status !== 200) return this.$message.error(res.message)
          this.$message.success('Delete success')
          this.getCateList()
        },
        onCancel: () => {
          this.$message.info('Cancel delete')
        }
      })
    },
    // Add Cate
    addCateOk() {
      this.$refs.addCateRef.validate(async (valid) => {
        if (!valid) {
          return this.$message.error('Invalid parameters, please input again')
        }
        const { data: res } = await this.$http.post('category/add', {
          name: this.newCate.name,
          password: this.newCate.password,
          role: this.newCate.role
        })
        if (res.status !== 200) return this.$message.error(res.message)
        this.$refs.addCateRef.resetFields()
        this.addCateVisible = false
        this.$message.success('Add Category success')
        await this.getCateList()
      })
    },
    addCateCancel() {
      this.$refs.addCateRef.resetFields()
      this.addCateVisible = false
      this.$message.info('Add Category canceled')
    },
    // Edit Category
    async editCate(id) {
      this.editCateVisible = true
      const { data: res } = await this.$http.get(`category/${id}`)
      this.CateInfo = res.data
      this.CateInfo.id = id
    },
    editCateOk() {
      this.$refs.editCateRef.validate(async (valid) => {
        if (!valid) {
          return this.$message.error('Invalid parameters , please input again')
        }
        const { data: res } = await this.$http.put(
          `category/${this.CateInfo.id}`,
          {
            name: this.CateInfo.name
          }
        )
        if (res.status !== 200) {
          return this.$message.error(res.message)
        }
        this.editCateVisible = false
        this.$message.success('Edit Category success')
        this.getCateList()
      })
    },
    editCateCancel() {
      this.$refs.editCateRef.resetFields()
      this.editCateVisible = false
      this.$message.info('Edit canceled')
    },
    // Change password
    async ChangePassword(id) {
      this.changePasswordVisible = true
      this.changePassword.id = id
    },
    changePasswordOk() {
      this.$refs.changePasswordRef.validate(async (valid) => {
        if (!valid) {
          return this.$message.error('Invalid parameters, please input again')
        }
        const { data: res } = await this.$http.put(
          `admin/changepw/${this.changePassword.id}`,
          {
            password: this.changePassword.password
          }
        )
        if (res.status !== 200) return this.$message.error(res.message)
        this.changePasswordVisible = false
        this.$message.success('Change password success')
        this.getCateList()
      })
    },
    changePasswordCancel() {
      this.$refs.changePasswordRef.resetFields()
      this.changePasswordVisible = false
      this.$message.info('Change password canceled')
    }
  }
}
</script>

<style scoped>
.actionSlot {
  display: flex;
  justify-content: center;
}
</style>
