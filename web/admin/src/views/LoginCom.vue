<template>
  <div class="container">
    <div class="loginBox">
      <a-form-model :rules="rules" :model="formdata" class="loginForm">
        <a-form-model-item prop="username">
          <a-input v-model="formdata.username" placeholder="Username">
            <a-icon
              slot="prefix"
              type="user"
              style="color: rgba(0, 0, 0, 0.25)"
            ></a-icon>
          </a-input>
        </a-form-model-item>
        <a-form-model-item prop="password">
          <a-input
            v-model="formdata.password"
            placeholder="Password"
            type="password"
          >
            <a-icon
              slot="prefix"
              type="lock"
              style="color: rgba(0, 0, 0, 0.25)"
            ></a-icon>
          </a-input>
        </a-form-model-item>
        <a-form-model-item class="loginBtn">
          <a-button type="primary" style="margin: 10px">Login</a-button>
          <a-button type="info">Cancel</a-button>
        </a-form-model-item>
      </a-form-model>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      formdata: {
        username: '',
        password: ''
      },
      rules: {
        username: [
          { required: true, message: 'Please input username', trigger: 'blur' },
          {
            min: 4,
            max: 12,
            message: 'Length should be 4 to 12',
            trigger: 'blur'
          }
        ],
        password: [
          { required: true, message: 'Please input password', trigger: 'blur' },
          {
            min: 4,
            max: 12,
            message: 'Length should be 6 to 20',
            trigger: 'blur'
          }
        ]
      }
    }
  }
}
</script>

<style scoped>
.container {
  height: 100%;
  background-color: black;
}

.loginBox {
  width: 450px;
  height: 300px;
  background-color: white;
  position: absolute;
  top: 50%;
  left: 70%;
  transform: translate(-50%, -50%);
  border-radius: 9px;
}

.loginForm {
  width: 100%;
  position: absolute;
  bottom: 10%;
  padding: 0 20px;
  box-sizing: border-box;
}

.loginBtn {
  display: flex;
  justify-content: flex-end;
}
</style>
