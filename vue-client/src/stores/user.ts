import { defineStore } from 'pinia'
import { inject, ref, reactive } from 'vue'
import { useLocalStorage } from '@vueuse/core'
import { pocketBaseSymbol } from '@/symbols/injectionSymbols'

export const useUserStore = defineStore('user', () => {
  const pb = inject(pocketBaseSymbol)

  const user = reactive({
    email: "",
    password: ""
  })

  // Actions

  return { user }
})
