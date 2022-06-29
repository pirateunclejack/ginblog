<template>
  <div>
    <h3>User List</h3>
    <a-card>
      <a-row :gutter="20">
        <a-col :span="6">
          <a-input-search
            v-model="queryParam.username"
            placeholder="input username to search"
            enter-button
            allowClear
            @search="getUserList"
          />
        </a-col>
        <a-col :span="4">
          <a-button type="primary" @click="addUserVisible = true">Add</a-button>
        </a-col>
      </a-row>
      <a-table
        rowKey="username"
        :columns="columns"
        :pagination="pagination"
        :dataSource="userlist"
        bordered
        @change="handleTableChange"
      >
        <span slot="role" slot-scope="role">{{
          role === 1 ? 'Administrator' : 'Subscriber'
        }}</span>
        <template slot="action" slot-scope="text, data">
          <div class="actionSlot">
            <a-button
              type="primary"
              icon="edit"
              style="margin-right: 15px"
              @click="editUser(data.ID)"
              >Edit</a-button
            >
            <a-button
              type="danger"
              icon="delete"
              style="margin-right: 15px"
              @click="deleteUser(data.ID)"
              >Delete</a-button
            >
            <a-button type="info" icon="info" @click="ChangePassword(data.ID)"
              >Change password</a-button
            >
          </div>
        </template>
      </a-table>
    </a-card>
    <!-- Add user -->
    <a-modal
      closable
      title="Add user"
      :visible="addUserVisible"
      width="60%"
      @ok="addUserOk"
      @cancel="addUserCancel"
      destroyOnClose
    >
      <a-form-model :model="newUser" :rules="addUserRules" ref="addUserRef">
        <a-form-model-item label="username" prop="username">
          <a-input v-model="newUser.username"></a-input>
        </a-form-model-item>
        <a-form-model-item has-feedback label="password" prop="password">
          <a-input-password v-model="newUser.password"></a-input-password>
        </a-form-model-item>
        <a-form-model-item
          has-feedback
          label="confirm password"
          prop="checkpass"
        >
          <a-input-password v-model="newUser.checkpass"></a-input-password>
        </a-form-model-item>
      </a-form-model>
    </a-modal>
    <!-- Edit user -->
    <a-modal
      closable
      title="Edit user"
      :visible="editUserVisible"
      width="60%"
      @ok="editUserOk"
      @cancel="editUserCancel"
    >
      <a-form-model :model="userInfo" :rules="editUserRules" ref="editUserRef">
        <a-form-model-item label="Username" prop="username"
          ><a-input v-model="userInfo.username"></a-input
        ></a-form-model-item>
        <a-form-model-item label="If is an administrator">
          <a-switch
            :checked="IsAdmin"
            checked-children="Yes"
            un-checked-children="No"
            @change="adminChange"
        /></a-form-model-item>
      </a-form-model>
    </a-modal>
    <!-- Change password -->
    <a-modal
      closable
      title="Change password"
      :visible="changePasswordVisible"
      width="60%"
      @ok="changePasswordOk"
      @cancel="changePasswordCancel"
      destroyOnClose
    >
      <a-form-model
        :model="changePassword"
        :rules="changePasswordRules"
        ref="changePasswordRef"
      >
        <a-form-model-item has-feedback label="Password" prop="password">
          <a-input-password
            v-model="changePassword.password"
          ></a-input-password>
        </a-form-model-item>
        <a-form-model-item
          has-feedback
          label="Confirm password"
          prop="checkpass"
        >
          <a-input-password
            v-model="changePassword.checkpass"
          ></a-input-password>
        </a-form-model-item>
      </a-form-model>
    </a-modal>
  </div>
</template>

<script>
const columns = [
  {
    title: 'ID',
    dataIndex: 'ID',
    width: '10%',
    key: 'id',
    align: 'center'
  },
  {
    title: 'Username',
    dataIndex: 'username',
    width: '20%',
    key: 'username',
    align: 'center'
  },
  {
    title: 'Role',
    dataIndex: 'role',
    width: '20%',
    key: 'role',
    scopedSlots: { customRender: 'role' },
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
      userlist: [],
      columns,
      queryParam: {
        username: '',
        pagesize: 2,
        pagenum: 1
      },
      userInfo: {
        id: 0,
        username: '',
        password: '',
        checkPass: '',
        role: 2
      },
      newUser: {
        username: '',
        password: '',
        role: 2,
        checkPass: ''
      },
      changePassword: {
        id: 0,
        password: '',
        checkPass: ''
      },
      visible: false,
      addUserRules: {
        username: [
          {
            validator: (rule, value, callback) => {
              if (this.newUser.username === '') {
                callback(new Error('Please input username'))
              }
              if (
                [...this.newUser.username].length < 4 ||
                [...this.newUser.username].length > 12
              ) {
                callback(
                  new Error('The length of username should >=4 and <=12')
                )
              } else {
                callback()
              }
            },
            trigger: 'blur'
          }
        ],
        password: [
          {
            validator: (rule, value, callback) => {
              if (this.newUser.password === '') {
                callback(new Error('Please input password'))
              }
              if (
                [...this.newUser.password].length < 6 ||
                [...this.newUser.password].length > 20
              ) {
                callback(
                  new Error('The length of password should >=6 and <=20')
                )
              } else {
                callback()
              }
            },
            trigger: 'blur'
          }
        ],
        checkpass: [
          {
            validator: (rule, value, callback) => {
              if (this.newUser.checkpass === '') {
                callback(new Error('Please input password'))
              }
              if (this.newUser.password !== this.newUser.checkpass) {
                callback(new Error('Password mismatch, please input again'))
              } else {
                callback()
              }
            },
            trigger: 'blur'
          }
        ]
      },
      editUserRules: {
        username: [
          {
            validator: (rule, value, callback) => {
              if (this.userInfo.username === '') {
                callback(new Error('Please input username'))
              }
              if (
                [...this.userInfo.username].length < 4 ||
                [...this.userInfo.username].length > 12
              ) {
                callback(
                  new Error('The length of username should >=4 and <=12')
                )
              } else {
                callback()
              }
            },
            trigger: 'blur'
          }
        ],
        password: [
          {
            validator: (rule, value, callback) => {
              if (this.userInfo.password === '') {
                callback(new Error('Please input password'))
              }
              if (
                [...this.userInfo.password].length < 6 ||
                [...this.userInfo.password].length > 20
              ) {
                callback(
                  new Error('The length of password should >=6 and <=20')
                )
              } else {
                callback()
              }
            },
            trigger: 'blur'
          }
        ],
        checkpass: [
          {
            validator: (rule, value, callback) => {
              if (this.userInfo.checkpass === '') {
                callback(new Error('Please input password'))
              }
              if (this.userInfo.password !== this.userInfo.checkpass) {
                callback(new Error('Password mismatch, please input again'))
              } else {
                callback()
              }
            },
            trigger: 'blur'
          }
        ]
      },
      changePasswordRules: {
        password: [
          {
            validator: (rule, value, callback) => {
              if (this.changePassword.password === '') {
                callback(new Error('Please input password'))
              }
              if (
                [...this.changePassword.password].length < 6 ||
                [...this.changePassword.password].length > 20
              ) {
                callback(
                  new Error('The length of password should >=6 and <=20')
                )
              } else {
                callback()
              }
            },
            trigger: 'blur'
          }
        ],
        checkpass: [
          {
            validator: (rule, value, callback) => {
              if (this.changePassword.checkpass === '') {
                callback(new Error('Please input password'))
              }
              if (
                this.changePassword.password !== this.changePassword.checkpass
              ) {
                callback(new Error('Password mismatch, please input again'))
              } else {
                callback()
              }
            },
            trigger: 'blur'
          }
        ]
      },
      addUserVisible: false,
      editUserVisible: false,
      changePasswordVisible: false
    }
  },
  created() {
    this.getUserList()
  },
  computed: {
    IsAdmin: function () {
      if (this.userInfo.role === 1) {
        return true
      } else {
        return false
      }
    }
  },
  methods: {
    async getUserList() {
      const { data: res } = await this.$http.get('admin/users', {
        params: {
          username: this.queryParam.username,
          pagesize: this.queryParam.pagesize,
          pagenum: this.queryParam.pagenum
        }
      })
      if (res.status !== 200) {
        return this.$message.error(res.message)
      }
      this.userlist = res.data
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
      this.getUserList()
    },
    deleteUser(id) {
      this.$confirm({
        title: 'Do you want to delete this user?',
        content:
          'Are you sure you want to delete this user? Once it is deleted , it will be impossible to recover it!',
        onOk: async () => {
          const { data: res } = await this.$http.delete(`user/${id}`)
          if (res.status !== 200) return this.$message.error(res.message)
          this.$message.success('Delete success')
          this.getUserList()
        },
        onCancel: () => {
          this.$message.info('Cancel delete')
        }
      })
    },
    // Add user
    addUserOk() {
      this.$refs.addUserRef.validate(async (valid) => {
        if (!valid) {
          return this.$message.error('Invalid parameters, please input again')
        }
        const { data: res } = await this.$http.post('user/add', {
          username: this.newUser.username,
          password: this.newUser.password,
          role: this.newUser.role
        })
        if (res.status !== 200) return this.$message.error(res.message)
        this.$refs.addUserRef.resetFields()
        this.addUserVisible = false
        this.$message.success('Add user success')
        this.getUserList()
      })
    },
    addUserCancel() {
      this.$refs.addUserRef.resetFields()
      this.addUserVisible = false
      this.$message.info('Add user canceled')
    },
    adminChange(checked) {
      if (checked) {
        this.userInfo.role = 1
      } else {
        this.userInfo.role = 2
      }
    },
    // Edit user
    async editUser(id) {
      this.editUserVisible = true
      const { data: res } = await this.$http.get(`user/${id}`)
      this.userInfo = res.data
      this.userInfo.id = id
    },
    editUserOk() {
      this.$refs.editUserRef.validate(async (valid) => {
        if (!valid) {
          return this.$message.error('Invalid parameters , please input again')
        }
        const { data: res } = await this.$http.put(`user/${this.userInfo.id}`, {
          username: this.userInfo.username,
          role: this.userInfo.role
        })
        if (res.status !== 200) {
          console.log(res)
          return this.$message.error(res.message)
        }
        console.log(this.userInfo.role)
        console.log(this.userInfo.username)
        console.log(res)
        this.editUserVisible = false
        this.$message.success('Edit user success')
        this.getUserList()
      })
    },
    editUserCancel() {
      this.$refs.editUserRef.resetFields()
      this.editUserVisible = false
      this.$message.info('Edit canceled')
    },
    // Change password
    async ChangePassword(id) {
      this.changePasswordVisible = true
      // const { data: res } = await this.$http.get(`user/${id}`)
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
        this.getUserList()
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
