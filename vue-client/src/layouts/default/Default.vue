<template>
  <v-app theme="dark">
    <default-app-bar />
    <default-view />
  </v-app>
</template>

<script lang="ts" setup>
import DefaultView from './View.vue'
import DefaultAppBar from './AppBar.vue'
import PocketBase from 'pocketbase'
import router from '@/router';

const pb = new PocketBase('http://127.0.0.1:8090')

// let blah = await pb.send("/api/init-check", {})

try {
  let initCheck = await pb.send("/api/init-check", {})
  if(!initCheck.isSetup) {
    router.push('/setup')
  }
} catch(e) {
  router.push('/error?msg=Something whent wrong with processing the request. ' +
  'Please make sure that PocketBase is running and that the API is accessible.')
  // router.push('/error?msg=blah')
}

// console.log(blah)
</script>
