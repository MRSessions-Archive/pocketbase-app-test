<script setup lang="ts">
import { ref } from 'vue';
import { useRoute } from 'vue-router';
// let query = useRoute().query.msg
const query = ref(useRoute().query.msg)
</script>

<!-- Error Page -->
<template>
  <!-- <v-main> -->
    <v-container>
      <v-row>
        <v-col cols="12">
          <h1 class="display-1">Error</h1>
          <h2 class="display-2">{{ query }}</h2>
        </v-col>
      </v-row>
    </v-container>
  <!-- </v-main> -->
</template>