<template>
  <v-app-bar color="primary" prominent>
    <v-app-bar-nav-icon variant="text" @click.stop="drawer = !drawer"></v-app-bar-nav-icon>

    <v-toolbar-title>Vue Client</v-toolbar-title>

    <v-spacer></v-spacer>

    <!-- TODO: Switch between light and dark with these two buttons -->
    <v-btn variant="text" icon="mdi-weather-night"></v-btn>

    <v-btn variant="text" icon="mdi-weather-sunny"></v-btn>

    <v-btn variant="text" icon="mdi-dots-vertical"></v-btn>
  </v-app-bar>

  <v-navigation-drawer
  v-model="drawer"
  :location="$vuetify.display.smAndDown ? 'top': 'start'"
  :permanent="$vuetify.display.smAndDown ? false : true">
    <v-list>
      <v-list-item v-for="item in items" :key="item.title" :to="item.value">{{ item.title }}</v-list-item>
    </v-list>

    <!-- TODO: Check if user logged in or out and swap button -->
    <template v-slot:append>
      <div class="pa-2">
        <v-btn block color="error">
          Logout
        </v-btn>
      </div>
    </template>
  </v-navigation-drawer>
</template>

<script setup lang="ts">
import { ref, watch } from 'vue'

const drawer = ref(true)
const items = ref([
  {
    title: 'Home',
    value: '/home',
  },
  {
    title: 'About',
    value: '/about',
  },
  {
    title: 'Settings',
    value: '/settings',
  },
  {
    title: 'blah'
  },
  {
    title: 'blah'
  },
  {
    title: 'blah'
  },
  {
    title: 'blah'
  },
  {
    title: 'blah'
  },
  {
    title: 'blah'
  },
  {
    title: 'blah'
  }
])

watch(drawer, (v) => {
  drawer.value = v
})
</script>
