<template>
      <v-app-bar
        color="primary"
        prominent
      >
        <v-app-bar-nav-icon variant="text" @click.stop="drawer = !drawer"></v-app-bar-nav-icon>

        <v-toolbar-title>Vue Client</v-toolbar-title>

        <v-spacer></v-spacer>

        <!-- TODO: Switch between light and dark with these two buttons -->
        <v-btn
          variant="text"
          icon="mdi-weather-night"
        ></v-btn>

        <v-btn
          variant="text"
          icon="mdi-weather-sunny"
        ></v-btn>

        <v-btn variant="text" icon="mdi-dots-vertical"></v-btn>
      </v-app-bar>

      <v-navigation-drawer
        v-model="drawer"
        permanent
      >
        <v-list>
          <v-list-item
            v-for="item in items"
            :key="item.title"
            :to="item.value"
          >{{ item.title }}</v-list-item>
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

<script lang="ts">
  export default {
    data: () => ({
      drawer: true,
      group: null,
      items: [
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
        }
      ],
    }),

    watch: {
      group () {
        this.drawer = false
      },
    },
  }
</script>
