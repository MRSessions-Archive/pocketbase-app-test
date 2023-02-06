<template>
  <!-- login form with email password and app name fields -->
  <v-container>
    <!-- <v-row align="center" justify="center" style="min-height: 93vh;"> -->
      <v-row class="pt-12" justify="center">
      <v-col cols="12" sm="12" md="8">
        <v-card>
          <v-card-title>
            <span class="headline">Setup</span>
          </v-card-title>
          <v-card-text>
            <v-form ref="form" v-model="valid" lazy-validation>
              <v-text-field v-model="email" label="Email" required></v-text-field>
              <v-text-field v-model="password" :append-icon="showPassword ? 'mdi-eye' : 'mdi-eye-off'"
                :type="showPassword ? 'text' : 'password'" label="Password" required
                @click:append="showPassword = !showPassword"></v-text-field>
              <v-text-field v-model="appName" label="App Name" required></v-text-field>
            </v-form>
          </v-card-text>
          <v-card-actions>
            <v-spacer></v-spacer>
            <v-btn color="blue darken-1" text @click="setup">Setup</v-btn>
          </v-card-actions>
        </v-card>
      </v-col>
    </v-row>
  </v-container>
</template>

<script setup lang="ts">
import { ref } from 'vue';
// import { useRoute } from 'vue-router';
// import { useStore } from 'vuex';
// import { useToast } from 'vue-toastification';
// // import { useI18n } from 'vue-i18n';
// import { useVuelidate } from '@vuelidate/core'
// import { required, email, minLength } from '@vuelidate/validators';

// const store = useStore();
// const route = useRoute();
// const toast = useToast();
// const { t } = useI18n();

const valid = ref(false);
const email = ref('');
const password = ref('');
const appName = ref('');
const showPassword = ref(false);

// const emailRules = ref([
//   (v: string) => !!v || t('validation.required'),
// //   (v: string) => email(v) || t('validation.email'),
// ]);
// const passwordRules = ref([
//   (v: string) => !!v || t('validation.required'),
//   (v: string) => minLength(8)(v) || t('validation.min', { min: 8 }),
// ]);
// const appNameRules = ref([
//   (v: string) => !!v || t('validation.required'),
// ]);

// const { $v } = useVuelidate({
//   email,
//   password,
//   appName,
// }, { $autoDirty: true });

function setup(){
  console.log('setup')
}





</script>