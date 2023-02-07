<template>
  <!-- login form with email password and confirm password fields -->
  <v-container>
    <!-- <v-row align="center" justify="center" style="min-height: 93vh;"> -->
    <v-row class="pt-12" justify="center">
      <v-col cols="12" sm="12" md="8">
        <v-card variant="tonal">
          <v-card-title>
            <span class="headline">Setup</span>
          </v-card-title>
          <v-card-text class="pb-0">
            <v-form ref="form" v-model="valid" lazy-validation>
              <v-text-field
              v-model="initStore.user.email"
              label="E-mail"
              required
              :rules="[rules.required, rules.email]"></v-text-field>

              <v-text-field
              v-model="initStore.user.password"
              :append-icon="showPassword ? 'mdi-eye' : 'mdi-eye-off'"
              :type="showPassword ? 'text' : 'password'" label="Password" hint="At least 10 characters" counter
              required
              @click:append="showPassword = !showPassword"
              :rules="[rules.min]"></v-text-field>

              <v-text-field
              v-model="initStore.user.passwordConfirm"
              :append-icon="showConfirmPassword ? 'mdi-eye' : 'mdi-eye-off'"
              :type="showConfirmPassword ? 'text' : 'password'" label="Confirm Password" hint="At least 10 characters" counter
              required
              @click:append="showConfirmPassword = !showConfirmPassword"
              :rules="[rules.min, rules.passwordMatch]"></v-text-field>

              <!-- <v-checkbox v-model="showPassword" :label="`Show password`" class="mb-0"></v-checkbox> -->
            </v-form>
          </v-card-text>
          <div class="pb-6 px-6">
            <v-btn size="x-large" variant="tonal" :disabled="!valid" color="success" @click="setup" block>
              Setup
            </v-btn>
          </div>

        </v-card>
      </v-col>
    </v-row>
  </v-container>

</template>

<script setup lang="ts">
import { ref } from 'vue';
import { useInitStore } from '@/stores/init';
import router from '@/router';

const initStore = useInitStore()

// TODO: Check if setup from init checkIsSetup Action
// run this call in views/Setup.vue
if(initStore.isSetup) {
  router.push('/')
}

const rules = {
  required: (v: string) => !!v || 'Required.',
  min: (v: string) => v.length >= 10 || 'At least 10 characters',
  passwordMatch: () => (initStore.user.passwordConfirm === initStore.user.password) || 'Passwords must match',
  email: (v: string) => /^\w+([.-]?\w+)*@\w+([.-]?\w+)*(\.\w{2,3})+$/.test(v) || 'E-mail must be valid',
};

const valid = ref(false);
const showPassword = ref(false);
const showConfirmPassword = ref(false);



async function setup() {
  try {
    const registerSuccess = await initStore.register()
    if(registerSuccess){
      router.push('/')
    }
  } catch (error) {
    if (error instanceof Error) {
      router.push('/n/error?msg=' + error.message)
    }
  }
}





</script>
