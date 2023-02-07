<template>
  <v-app theme="dark">
    <default-app-bar />
    <default-view />
  </v-app>
</template>

<script lang="ts" setup>
import DefaultView from './View.vue'
import DefaultAppBar from './AppBar.vue'
import { useInitStore } from '@/stores/init';
import router from '@/router';

const store = useInitStore()

try {
  await store.checkIsSetup()
  if (!store.isSetup) {
    router.push('/n/setup')
  }
} catch (error) {
  if (error instanceof Error) {
    router.push('/n/error?msg=' + error.message)
  }
}

console.log(store.isSetup)
// import PocketBase from 'pocketbase'
// import router from '@/router';

// const pb = new PocketBase('http://127.0.0.1:8090')

// try {
//   let initCheck = await pb.send("/api/init-check", {})
//   if(!initCheck.isSetup) {
//     router.push('/setup')
//   }
// } catch(e) {
//   router.push('/n/error?msg=Something whent wrong with processing the request. ' +
//   'Please make sure that PocketBase is running and that the API is accessible.')
// }
</script>
