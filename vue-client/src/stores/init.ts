import { defineStore } from 'pinia'
import { inject, ref, reactive } from 'vue'
import { useLocalStorage } from '@vueuse/core'
import { pocketBaseSymbol } from '@/symbols/injectionSymbols'

export const useInitStore = defineStore('init', () => {
  const pb = inject(pocketBaseSymbol)

  const isSetup = ref(useLocalStorage("initialSetup", false))

  const user = reactive({
    email: "",
    password: "",
    passwordConfirm: ""
  })

  // Actions
  // done for initial PocketBase setup
  const checkIsSetup = async () => {
    if (!isSetup.value){
      try {
        const initCheck = await pb?.send("/api/init-check", {})
        isSetup.value = initCheck.isSetup
      } catch (e) {
        isSetup.value = false
        throw new Error("Something whent wrong with processing the request. " +
        "Please make sure that PocketBase is running and that the API is accessible.");
      }
    }
  }

  const register = async () => {
    try {
      const register = await pb?.send("/api/admins", {
        method: "POST",
        headers: {
          "Content-Type": "application/json"
        },
        body: user
      })
      console.log(register.email)
      if(register.email === user.email){
        isSetup.value = true
        return true
      }
    } catch (e) {
      throw new Error("Something whent wrong with processing the request. " +
      "Please make sure that PocketBase is running and that the API is accessible.");
    }
  }

  return { isSetup, user, checkIsSetup, register }
})




// curl --request POST \
//     --url http://127.0.0.1:8090/api/admins \
//     --header 'Content-Type: application/json' \
//     --data '{"email": "test@email.com", "password": "P@ssw0rd***", "passwordConfirm": "P@ssw0rd***"}'
