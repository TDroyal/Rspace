<template>
  <navbar />
  <router-view/>
</template>

<script>
// import Navbar  from "@/component/Navbar"
import Navbar from './components/Navbar.vue'
import {useStore} from 'vuex'
export default{
  components: { Navbar },
  setup(){
    const store = useStore()
    // console.log("App.vue:", store.state)
    window.addEventListener("beforeunload",()=>{
      // console.log("App.vue:", JSON.stringify(store.state))
      localStorage.setItem("userStore",JSON.stringify(store.state))
    })
    localStorage.getItem("userStore") && store.replaceState(Object.assign(store.state,JSON.parse(localStorage.getItem("userStore"))));
  }  
}

</script>

<style>

</style>
